package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ysomad/uniplay/internal/pkg/demoparser"
)

func uploadDemoHandlerV1(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 >> 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, fh, err := r.FormFile("demo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()

	slog.Info("got demo file", "header", fh)

	parser, err := demoparser.New(f, fh)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer parser.Close()

	if err := parser.Parse(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	})))
	http.HandleFunc("/v1/demos", uploadDemoHandlerV1)
	http.ListenAndServe(":8080", nil)
}
