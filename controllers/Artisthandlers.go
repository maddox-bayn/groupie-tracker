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
		RenderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	err := RendersTemplates(w, http.StatusOK, "index.html", data.CombinedData)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func HandleArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		RenderError(w, http.StatusNotFound, "Page Not Found")
		log.Println("404 page not found =>", r.URL.Path, "❌")
		return
	}
	QueryParam := r.URL.Query()
	idStr := QueryParam.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RenderError(w, http.StatusBadRequest, "404 | Bad reuquest try something else")
		return
	}
	artist, err := control_utils.FetchArtist(id)
	if err != nil {
		if errors.Is(err, control_utils.Err404) {
			RenderError(w, http.StatusNotFound, "Bad Request Artist Not Found")
			return
		} else {
			RenderError(w, http.StatusInternalServerError, "Internal Server Error")
			fmt.Println("Error parsing template")
		}
		return
	}

	err = RendersTemplates(w, http.StatusOK, "artist.html", artist)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println("Error parsing template")
	}
}
