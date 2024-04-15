package urls

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/modules/auth"
	"github.com/ngrink/url-shortener/internal/utils"
)

type UrlsController struct {
	service IUrlsService
}

func NewUrlsController(service IUrlsService) *UrlsController {
	return &UrlsController{service: service}
}

func (c *UrlsController) CreateUrl(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r)

	var data CreateUrlDto

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url, err := c.service.CreateUrl(userId, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, url)
}

func (c *UrlsController) GetAllUrls(w http.ResponseWriter, r *http.Request) {
	urls, err := c.service.GetAllUrls()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *UrlsController) GetUserUrls(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["userId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	urls, err := c.service.GetUserUrls(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *UrlsController) GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url, err := c.service.GetUrl(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *UrlsController) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.DeleteUrl(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (c *UrlsController) RedirectByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	url, err := c.service.GetUrlByKey(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	c.service.RegisterVisit(url.ID, r.UserAgent(), r.RemoteAddr)

	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}

func (c *UrlsController) GetUrlVisits(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	visits, err := c.service.GetUrlVisits(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	visitsResponse := make([]VisitResponse, len(visits))
	for i, visit := range visits {
		visitsResponse[i] = VisitResponse{IpAddress: visit.IpAddress, UserAgent: visit.UserAgent}
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"visits": visitsResponse,
	})
}
