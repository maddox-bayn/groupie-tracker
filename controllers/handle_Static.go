package controllers

import (
	"net/http"
	"strings"
)

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		RenderError(w, http.StatusForbidden, "403 | forbideen Access to this resource.")
		return
	}
	fileServer := http.FileServer(http.Dir("./static"))
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}
