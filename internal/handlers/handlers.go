package handlers

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/KokoulinM/exchanges-history-app/internal/csv"
	"github.com/KokoulinM/exchanges-history-app/internal/models"
	"github.com/go-chi/chi/v5"
)

type Repository interface {
	Ping(ctx context.Context) error
	UploadFile(ctx context.Context, exchangesHistory []models.ExchangesHistory) error
}

type Handlers struct {
	repo    Repository
	baseURL string
}

// New is the handlers constructor
func New(repo Repository, baseURL string) *Handlers {
	return &Handlers{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (h *Handlers) UploadExchangesHistory(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	fileName := chi.URLParam(r, "file")
	if fileName == "" {
		http.Error(w, "the parameter is missing", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reader := bytes.NewReader(buf.Bytes())

	data, err := csv.Reader(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.repo.UploadFile(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h *Handlers) PingDB(w http.ResponseWriter, r *http.Request) {
	err := h.repo.Ping(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
