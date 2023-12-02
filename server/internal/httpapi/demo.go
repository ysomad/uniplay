package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/ysomad/uniplay/internal/demoparser"
	"github.com/ysomad/uniplay/internal/httpapi/writer"
)

type uploadDemoRes struct {
	DemoJobID uuid.UUID `json:"demo_job_id"`
}

type demoV1 struct{}

func NewDemoV1() *demoV1 {
	return &demoV1{}
}

func (d *demoV1) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writer.Status(w, http.StatusMethodNotAllowed)
		return
	}

	// TODO: return specific error
	if err := r.ParseMultipartForm(10 >> 20); err != nil {
		writer.Error(w, http.StatusBadRequest, err)
		return
	}

	// TODO: save demo file to some shared storage some-uuid.dem
	// save job with filename in job db
	// return job id
	// parse demo in background worker
	f, fh, err := r.FormFile("demo")
	if err != nil {
		return
	}
	defer f.Close()

	slog.Info("got demo file", "header", fh)

	parser, err := demoparser.New(f, fh)
	if err != nil {
		writer.Error(w, http.StatusBadRequest, err)
		return
	}
	defer parser.Close()

	if err := parser.Parse(); err != nil {
		writer.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(uploadDemoRes{
		DemoJobID: uuid.New(), // TODO: return REAL demo job id
	}); err != nil {
		writer.Error(w, http.StatusInternalServerError, err)
		return
	}
}
