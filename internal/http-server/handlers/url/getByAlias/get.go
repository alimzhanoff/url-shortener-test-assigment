package getByAlias

import (
	resp "github.com/alimzhanoff/url-shortener-test-assigment/internal/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

const aliasLenght = 6

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}
type Response struct {
	resp.Response
	ID    int64  `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
}
type URLSaver interface {
	SaveURl(urlToSave, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.getByAlias.New"

		alias := chi.URLParam(r, "alias")
		render.JSON(w, r, alias)
		return
	}
}
