package download

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" || strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func GetDownloadFolderHandler() http.Handler {
	df := "/" + DownloadFolder()
	return http.StripPrefix(df, noDirListing(http.FileServer(http.Dir("."+df))))
}

func SetDownloadRoute(r *mux.Router) {
	df := "/" + DownloadFolder()
	r.PathPrefix(df).Handler(GetDownloadFolderHandler())
}
