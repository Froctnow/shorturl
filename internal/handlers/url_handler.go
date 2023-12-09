package handlers

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"regexp"
	"shorturl/internal/services"
	"strings"
)

const RegexpURL = "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"

type URLHandler struct {
	urlService *services.URLService
}

func NewURLHandler(urlService *services.URLService) *URLHandler {
	urlHandler := &URLHandler{urlService: urlService}

	return urlHandler
}

func (uh *URLHandler) HandleRequest() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			uh.handleIndexPost(res, req)
			return
		case http.MethodGet:
			uh.handleIndexGet(res, req)
			return
		default:
			http.Error(res, "Method isn't allowed. Only POST, GET", http.StatusMethodNotAllowed)
			return
		}
	}
}

func (uh *URLHandler) handleIndexPost(res http.ResponseWriter, req *http.Request) {
	headerContentType := req.Header.Get("Content-Type")

	if !strings.Contains(headerContentType, "text/plain") {
		http.Error(res, MessageErrorIncorrectContentType, http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Something went wrong", http.StatusInternalServerError)
		return
	}

	url := string(body)

	isMatched, err := regexp.MatchString(RegexpURL, url)

	if !isMatched {
		http.Error(res, MessageErrorIncorrectURL, http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(res, "Something went wrong", http.StatusInternalServerError)
		return
	}

	result := uh.urlService.CreateShortURL(url)

	res.Header().Add("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(result))
	if err != nil {
		http.Error(res, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func (uh *URLHandler) handleIndexGet(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/")

	_, err := uuid.Parse(id)

	if err != nil {
		http.Error(res, MessageErrorIncorrectID, http.StatusBadRequest)
		return
	}

	headerContentType := req.Header.Get("Content-Type")

	if !strings.Contains(headerContentType, "text/plain") {
		http.Error(res, MessageErrorIncorrectContentType, http.StatusBadRequest)
		return
	}

	url, err := uh.urlService.GetURL(id)

	if err != nil {
		http.Error(res, MessageErrorShortURLNotFound, http.StatusNotFound)
		return
	}

	res.Header().Add("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
