package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/modules/auth"
	"github.com/ngrink/url-shortener/internal/modules/urls"
)

type WebController struct {
	urlsService urls.IUrlsService
}

func NewWebController(urlsService urls.IUrlsService) *WebController {
	return &WebController{urlsService: urlsService}
}

func (c *WebController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./web/views/layouts/main.html",
		"./web/views/partials/url-table.html",
		"./web/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	userId := auth.GetUserIdFromContext(r)
	urlsList, err := c.urlsService.GetUserUrls(uint64(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	data := struct {
		Urls []urls.Url
	}{
		Urls: urlsList,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func (c *WebController) Url(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./web/views/layouts/main.html",
		"./web/views/partials/url-table.html",
		"./web/views/partials/visits-table.html",
		"./web/views/url.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	vars := mux.Vars(r)
	urlId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url, err := c.urlsService.GetUrl(uint64(urlId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	visits, err := c.urlsService.GetUrlVisits(uint64(urlId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := struct {
		Urls   []urls.Url
		Visits []urls.Visit
	}{
		Urls:   []urls.Url{url},
		Visits: visits,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func (c *WebController) Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./web/views/layouts/main.html",
		"./web/views/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func (c *WebController) Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./web/views/layouts/main.html",
		"./web/views/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}
