package httpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/ysomad/uniplay/server/internal/demoparser"
	"github.com/ysomad/uniplay/server/internal/domain"
	appctx "github.com/ysomad/uniplay/server/internal/kratosctx"
	"github.com/ysomad/uniplay/server/internal/postgres"
)

const (
	demoMaxSize       = 200 << 20
	demoMemoryMaxSize = 50 << 20
	demoTTL           = time.Hour * 24 * 7
)

var (
	errDemoNotUploaded = errors.New("demo not uploaded to storage")
	errDemoTooLarge    = errors.New("demo file is too large")
)

type demoV1 struct {
	minio  *minio.Client
	demo   postgres.DemoStorage
	bucket string
}

type uploadDemoResponse struct {
	DemoID uuid.UUID `json:"demo_id"`
}

func (d *demoV1) upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeStatus(w, http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, demoMaxSize)

	if err := r.ParseMultipartForm(demoMemoryMaxSize); err != nil {
		writerError(w, http.StatusBadRequest,
			fmt.Errorf("%w, must be equal or less than %dMB", errDemoTooLarge, demoMaxSize/1024/1024))
		return
	}

	file, fileHdr, err := r.FormFile("demo")
	if err != nil {
		return
	}
	defer file.Close()

	demo, err := demoparser.NewDemo(file, fileHdr)
	if err != nil {
		writerError(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()
	identityID := appctx.IdentityID(ctx)
	filename := demo.Filename()
	now := time.Now()

	// Upload demo only if it its not already in object storage.
	if _, err = d.minio.StatObject(ctx, d.bucket, filename, minio.GetObjectOptions{}); err != nil {
		res, err := d.minio.PutObject(ctx, d.bucket, filename,
			demo, demo.Size,
			minio.PutObjectOptions{
				UserMetadata: map[string]string{"uploader": identityID},
				Expires:      now.Add(demoTTL),
			})
		if err != nil {
			writerError(w, http.StatusInternalServerError,
				fmt.Errorf("%w, reason: %w", errDemoNotUploaded, err))
			return
		}

		slog.Info("demo uploaded to object storage", "upload_info", res)
	}

	err = d.demo.Save(ctx, domain.Demo{
		UploadedAt: now,
		Status:     domain.DemoStatusAwaiting,
		IdentityID: identityID,
		ID:         demo.ID,
	})
	if err != nil && !errors.Is(err, postgres.ErrDemoAlreadyExists) {
		slog.Error("demo not saved to db", "error", err, "demo_id", demo.ID)
		writerError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(uploadDemoResponse{DemoID: demo.ID}); err != nil {
		writerError(w, http.StatusInternalServerError, err)
		return
	}
}
