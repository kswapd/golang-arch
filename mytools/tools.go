package mytools

import (
	"fmt"
	"net/http"

	"github.com/dolthub/vitess/go/vt/log"
)

var (
	// Version is the version of the binary
	staticDir       string = "./mytools/static"
	portalIndexFile string = "index.html"
)

type Data struct {
	Title string
	Body  string
}

func handler(w http.ResponseWriter, r *http.Request) {

	/*tmpl, err := template.ParseFiles(fmt.Sprintf("%s%s", staticPath, portalIndex))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}*/

	// Create a file server handler for the specified directory.
	//staticDir := fmt.Sprintf("%s%s", staticPath, portalIndex)
	fs := http.FileServer(http.Dir(staticDir))

	// Handle requests to the root path by serving files from the static directory.
	http.Handle("/", fs)
}

func RunHtmlView() {
	var port int = 8888
	log.Infof("Starting server on :%d", port)
	fs := http.FileServer(http.Dir(staticDir))
	// Handle requests to the root path by serving files from the static directory.
	http.Handle("/", fs)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
