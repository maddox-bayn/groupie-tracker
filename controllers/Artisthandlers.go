package controllers

import (
	"groupie-tracker/data"
	"groupie-tracker/model"
	"log"
	"net/http"
	"strconv"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "404 | Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := RendersTemplates(w, http.StatusOK, "index.html", data.CombinedData)
	if err != nil {
		return
	}
}

func HandleArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		http.NotFound(w, r)
		log.Println("404 page not found =>", r.URL.Path, "❌")
		return
	}
	QueryParam := r.URL.Query()
	id := QueryParam.Get("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		RendersTemplates(w, http.StatusBadRequest, "error.html", nil)
	}
	var artist model.Artist
	for _, v := range data.CombinedData.Artists {
		if v.ID == Id {
			artist.ID = v.ID
			artist.CreationDate = v.CreationDate
			artist.Image = v.Image
			artist.Members = v.Members
			artist.Name = v.Name
			artist.ty
		}
	}

}
