package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/uuid"

	"github.com/ysomad/uniplay/internal/pkg/demoparser"
)

type httpError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func writeHttpError(w http.ResponseWriter, code int, err error) error {
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(httpError{
		Msg:  err.Error(),
		Code: http.StatusText(code),
	})
}

type uploadDemoRes struct {
	DemoJobID uuid.UUID `json:"demo_job_id"`
}

type demoHandlerV1 struct{}

func (h *demoHandlerV1) uploadDemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseMultipartForm(10 >> 20); err != nil {
		writeHttpError(w, http.StatusBadRequest, err)
		return
	}

	f, fh, err := r.FormFile("demo")
	if err != nil {
		return
	}
	defer f.Close()

	slog.Info("got demo file", "header", fh)

	parser, err := demoparser.New(f, fh)
	if err != nil {
		writeHttpError(w, http.StatusBadRequest, err)
		return
	}
	defer parser.Close()

	if err := parser.Parse(); err != nil {
		writeHttpError(w, http.StatusBadRequest, err)
		return
	}

	if err := json.NewEncoder(w).Encode(uploadDemoRes{
		DemoJobID: uuid.New(), // TODO: return REAL demo job id
	}); err != nil {
		writeHttpError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	})))

	demoHandlerV1 := &demoHandlerV1{}

	http.HandleFunc("/v1/demos", demoHandlerV1.uploadDemo)
	http.ListenAndServe(":8080", nil)
}
