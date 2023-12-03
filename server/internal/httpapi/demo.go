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

	"github.com/ysomad/uniplay/server/internal/appctx"
	"github.com/ysomad/uniplay/server/internal/demoparser"
	"github.com/ysomad/uniplay/server/internal/domain"
	"github.com/ysomad/uniplay/server/internal/httpapi/reswriter"
	"github.com/ysomad/uniplay/server/internal/postgres"
)

var errDemoNotUploaded = errors.New("demo not uploaded to storage")

type uploadDemoRes struct {
	DemoID uuid.UUID `json:"demo_id"`
}

type demoV1 struct {
	minio   *minio.Client
	storage postgres.DemoStorage
	bucket  string
}

func NewDemoV1(c *minio.Client, bucket string, s postgres.DemoStorage) *demoV1 {
	return &demoV1{
		minio:   c,
		bucket:  bucket,
		storage: s,
	}
}

const (
	demoMaxSize       = 200 << 20
	demoMemoryMaxSize = 50 << 20
	demoTTL           = time.Hour * 24 * 7
)

var errDemoTooLarge = errors.New("demo file is too large")

func (d *demoV1) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		reswriter.Status(w, http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, demoMaxSize)

	if err := r.ParseMultipartForm(demoMemoryMaxSize); err != nil {
		reswriter.Error(w, http.StatusBadRequest,
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
		reswriter.Error(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()

	uploader, ok := appctx.IdentityID(ctx)
	if !ok {
		reswriter.Status(w, http.StatusUnauthorized)
		return
	}

	now := time.Now()
	filename := demo.Filename()

	// Upload demo only if it its not already in object storage.
	if _, err = d.minio.StatObject(ctx, d.bucket, filename, minio.GetObjectOptions{}); err != nil {
		res, err := d.minio.PutObject(ctx, d.bucket, filename,
			demo, demo.Size,
			minio.PutObjectOptions{
				UserMetadata: map[string]string{"uploader": uploader},
				Expires:      now.Add(demoTTL),
			})
		if err != nil {
			reswriter.Error(w, http.StatusInternalServerError,
				fmt.Errorf("%w, reason: %w", errDemoNotUploaded, err))
			return
		}

		slog.Info("demo uploaded to object storage", "upload_info", res)
	}

	err = d.storage.Save(ctx, domain.Demo{
		ID:         demo.ID,
		Status:     domain.DemoStatusAwaiting,
		IdentityID: uploader,
		UploadedAt: now,
	})
	if err != nil && !errors.Is(err, postgres.ErrDemoAlreadyExists) {
		slog.Error("demo not saved to db", "error", err, "demo_id", demo.ID)
		reswriter.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(uploadDemoRes{DemoID: demo.ID}); err != nil {
		reswriter.Error(w, http.StatusInternalServerError, err)
		return
	}
}
