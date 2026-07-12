package controllers

import (
	"errors"
	"fmt"
	"groupie-tracker/control_utils"
	"groupie-tracker/data"
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
	idStr := QueryParam.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RendersTemplates(w, http.StatusBadRequest, "error.html", nil)
		return
	}
	artist, err := control_utils.FetchArtist(id)
	if errors.Is(err, control_utils.Err404) {
		err = RendersTemplates(w, http.StatusNotFound, "error.html", artist)
		if err != nil {
			http.Error(w, "Internal server Error", http.StatusInternalServerError)
			fmt.Println("Error parsing template")
		}
		return
	} else if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template")
		return
	}

	err = RendersTemplates(w, http.StatusOK, "artist.html", artist)
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template")
	}
}
