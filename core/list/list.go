package list

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"webz/core/log"
)

var wd string
var base_html_top string
var base_html_bottom string

func Set_config(new_wd string) {
	wd = new_wd

	base_html_top = "<!DOCTYPE html><head><title>WEBZ - list</title></head><body>"
	base_html_bottom = "<p>Powered by WEBZ</p></body>"
}

func list_constructer(files []os.DirEntry, path string) string {
	html := base_html_top
	for _, file := range files {
		if file.IsDir() {
			html += `<a href="/dir/` + path + "/" + file.Name() + `">` + file.Name() + " DIR </a></br>"
		} else {
			html += `<a href="/file/` + path + "/" + file.Name() + `">` + file.Name() + " FILE </a></br>"
		}
	}

	html += base_html_bottom

	return html
}

func List_handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if strings.Contains(path, "..") {
		fmt.Fprintln(w, "[500] Internal server error")
		return
	}

	if path == "" {
		// base of the 'wd' dir instead of looking for a 'index.html' we'll server a list of files and dir's within 'wd'
		files, err := os.ReadDir(wd)
		if err != nil {
			log.Error("[error] couldn't read '" + wd + "' due too " + err.Error())
			fmt.Fprintln(w, "[500] Internal server error")
		} else {
			// construct 'list'
			fmt.Fprint(w, list_constructer(files, path))
		}
		return
	} else if strings.Contains(path, "dir/") {
		// we know it's a directory so we can search it and do the same as we did before
		path = strings.Replace(path, "dir/", "", 1)

		files, err := os.ReadDir(wd + "/" + path)
		if err != nil {
			log.Error("[error] couldn't read '" + wd + "' due too " + err.Error())
			http.Error(w, "[500] Internal server error", http.StatusInternalServerError)
		} else {
			log.Info("[200] sendind dir " + path)

			// construct 'list'
			fmt.Fprint(w, list_constructer(files, path))
		}
		return
	} else if strings.Contains(path, "file/") {
		// we know it's a file so we can read and serve the file
		path = strings.Replace(path, "file/", "", 1)
		if _, err := os.Stat(wd + "/" + path); err == nil {
			http.ServeFile(w, r, wd+"/"+path)
			log.Info("[200] /" + path)
		} else {
			http.Error(w, "[404] File Not Found!", http.StatusNotFound)
			log.Info("[404] File Not " + path + " Found!")
		}
		return
	}

}
