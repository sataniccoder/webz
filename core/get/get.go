package get

import (
	"fmt"
	"net/http"
	"os"
)

var wd string

func Set_wd(new_wd string) {
	wd = new_wd
}

func Get_handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	//_, err := os.ReadDir(wd)

	if path == "" {

		if _, err := os.Stat(wd + "/index.html"); err == nil {
			http.ServeFile(w, r, wd+"/index.html")
			fmt.Println("[200] index.html")
		} else {
			fmt.Fprintf(w, "GO-WEB works! please add a index.html file or use the '--list' option to change this page")
			fmt.Println("[404] index.html not found!")

		}
	} else {
		// search current working dir or 'wd' var too check if the file exits
		if _, err := os.Stat(wd + "/" + path); err == nil {
			http.ServeFile(w, r, wd+"/"+path)
			fmt.Println("[200] /" + path)
		} else {
			fmt.Fprintln(w, "[404] File Not Found!")
			fmt.Println("[404] file " + path + " not found!")
		}
	}
}
