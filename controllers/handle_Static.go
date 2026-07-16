package controllers

import (
	"net/http"
	"strings"
)


// handle request to static directiy
// HandleStatic handler checks if request path has "/" as suffix 
// to specify that the user is requesting for the static dir if true return
// serve static file to the browser
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		RenderError(w, http.StatusForbidden, "403 | forbideen Access to this resource.")
		return
	}
	fileServer := http.FileServer(http.Dir("./static"))
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}
