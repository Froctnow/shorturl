package handlers

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"regexp"
	"shorturl/internal/services"
	"strings"
)

const RegexpURL = "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/|\\/|\\/\\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"

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
		http.Error(res, "Incorrect Content-Type. Only text/plain allowed", http.StatusBadRequest)
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
		http.Error(res, "Incorrect URL", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(res, "Something went wrong", http.StatusInternalServerError)
		return
	}

	result := uh.urlService.CreateShortUrl(url)

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
		http.Error(res, "Incorrect id", http.StatusBadRequest)
		return
	}

	url, err := uh.urlService.GetURL(id)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Add("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
