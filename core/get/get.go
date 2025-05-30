package get

import (
	"net/http"
	"os"
	"webz/core/log"
	"webz/core/password"
)

var wd string
var need_passwd bool

func Set_config(new_wd string, new_need_passwd bool) {
	wd = new_wd
	need_passwd = new_need_passwd

}

func Get_handle(w http.ResponseWriter, r *http.Request) {
	// check if password is needed
	if need_passwd {
		if !password.Check_coockie(w, r) {
			// redirect to login page
			http.Redirect(w, r, "/auth", http.StatusFound)
			return
		}
	}

	path := r.URL.Path[1:]
	//_, err := os.ReadDir(wd)

	// check if path is '' that means it's the index
	if path == "" {
		// check for an index.html
		if _, err := os.Stat(wd + "/index.html"); err == nil {
			http.ServeFile(w, r, wd+"/index.html")
			log.Info("[200] index.html")
		} else {
			http.Error(w, "GO-WEB works! please add a index.html file or use the '--list' option to change this page", http.StatusNotFound)
			log.Warning("[404] index.html")

		}

		return
	} else {
		// search current working dir or 'wd' var to check if the file exits within it
		if _, err := os.Stat(wd + "/" + path); err == nil {
			http.ServeFile(w, r, wd+"/"+path)
			log.Info("[200] /" + path)
		} else {
			http.Error(w, "[404] File Not Found!", http.StatusNotFound)
			log.Warning("[404] file " + path + " not found!")
		}

		return
	}
}
