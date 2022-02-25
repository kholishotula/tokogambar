package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"image"
	"net/http"

	"github.com/riandyrn/tokogambar/internal/core/getsimilar"

	"github.com/gorilla/mux"
)

type API struct {
	getsimilarService getsimilar.Service
}

func (a *API) Handler() http.Handler {
	r := mux.NewRouter()

	r.Handle("/", a.showPage())
	r.PathPrefix("/images/").Handler(a.listImages())
	r.HandleFunc("/similars", a.checkSimilarImages())

	return r
}

func (a *API) showPage() http.Handler {
	return http.FileServer(http.Dir("./static/web"))
}

func (a *API) listImages() http.Handler {
	return http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images")))
}

func (a *API) checkSimilarImages() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			WriteRespBody(w, NewErrorResp(NewErrMethodNotAllowed()))
			return
		}
		// parse request body
		var rb searchReqBody
		err := json.NewDecoder(r.Body).Decode(&rb)
		if err != nil {
			WriteRespBody(w, NewErrorResp(NewErrBadRequest(err.Error())))
			return
		}
		// validate request body
		err = rb.Validate()
		if err != nil {
			WriteRespBody(w, NewErrorResp(err))
			return
		}

		imgData, err := rb.GetByte()
		if err != nil {
			WriteRespBody(w, NewErrorResp(err))
			return
		}

		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			WriteRespBody(w, NewErrorResp(err))
			return
		}

		similarImages, err := a.getsimilarService.GetSimilarImages(context.Background(), img)
		if err != nil {
			WriteRespBody(w, NewErrorResp(err))
			return
		}

		WriteRespBody(w, NewSuccessResp(similarImages))
		return
	}
}

type APIConfig struct {
	GetSimilarService getsimilar.Service
}

func NewAPI(cfg APIConfig) *API {
	return &API{getsimilarService: cfg.GetSimilarService}
}
