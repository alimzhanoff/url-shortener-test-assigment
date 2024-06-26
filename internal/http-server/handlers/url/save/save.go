package save

import (
	"errors"
	resp "github.com/alimzhanoff/url-shortener-test-assigment/internal/lib/api/response"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/lib/logger/sl"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/lib/random"
	"github.com/alimzhanoff/url-shortener-test-assigment/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
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
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to devode request"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		//validate
		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLenght)
		}

		id, err := urlSaver.SaveURl(req.URL, alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLExists) {
				log.Info("url already exists", slog.String("url", req.URL))
				render.JSON(w, r, resp.Error("url already exists"))

				return
			}
			log.Error("failed to add url", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to add url"))
			return
		}
		render.JSON(w, r, Response{resp.OK(), id, alias})
		return
	}
}
