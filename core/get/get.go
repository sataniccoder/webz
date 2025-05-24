package get

import (
	"fmt"
	"net/http"
	"os"
	"webz/core/misc"
)

var wd string
var silent bool

func Set_config(new_wd string, new_silent bool) {
	wd = new_wd
	silent = new_silent
}

func Get_handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	//_, err := os.ReadDir(wd)

	if path == "" {

		if _, err := os.Stat(wd + "/index.html"); err == nil {
			http.ServeFile(w, r, wd+"/index.html")
			misc.Print("[200] index.html", silent)
		} else {
			fmt.Fprintf(w, "GO-WEB works! please add a index.html file or use the '--list' option to change this page")
			misc.Print("[404] index.html not found!", silent)

		}
	} else {
		// search current working dir or 'wd' var too check if the file exits
		if _, err := os.Stat(wd + "/" + path); err == nil {
			http.ServeFile(w, r, wd+"/"+path)
			misc.Print("[200] /"+path, silent)
		} else {
			fmt.Fprintln(w, "[404] File Not Found!")
			misc.Print("[404] file "+path+" not found!", silent)
		}
	}
}
