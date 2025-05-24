package list

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"webz/core/misc"
)

var wd string
var silent bool
var base_html_top string
var base_html_bottom string

func Set_config(new_wd string, new_slient bool) {
	wd = new_wd
	silent = new_slient

	base_html_top = "<!DOCTYPE html><head><title>WEBZ - list</title></head><body>"
	base_html_bottom = "<p>Powered by WEBZ</p></body>"
}

func List_handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if strings.Contains(path, "..") {
		misc.Print("[warning] detected '..' in requested url", silent)
		fmt.Fprintln(w, "[500] Internal server error")
	}

	if path == "" {
		// base of the 'wd' dir
		files, err := os.ReadDir(wd)
		if err != nil {
			misc.Print("[error] couldn't read '"+wd+"' due too "+err.Error(), silent)
			fmt.Fprintln(w, "[500] Internal server error")
		} else {
			// construct 'list'
			html := base_html_top
			for _, file := range files {
				if file.IsDir() {
					html += `<a href="/dir/` + file.Name() + `">` + file.Name() + " DIR </a></br>"
				} else {
					html += `<a href="/file/` + file.Name() + `">` + file.Name() + " FILE </a></br>"
				}
			}

			html += base_html_bottom

			fmt.Fprint(w, html)
		}
	} else if strings.Contains(path, "dir/") {
		path = strings.Replace(path, "dir/", "", 1)

		files, err := os.ReadDir(wd + "/" + path)
		if err != nil {
			misc.Print("[error] couldn't read '"+wd+"' due too "+err.Error(), silent)
			fmt.Fprintln(w, "[500] Internal server error")
		} else {
			misc.Print("[200] sendind dir "+path, silent)

			// construct 'list'
			html := base_html_top
			for _, file := range files {
				if file.IsDir() {
					html += `<a href="/dir/` + path + "/" + file.Name() + `">` + file.Name() + " DIR </a></br>"
				} else {
					html += `<a href="/file/` + path + "/" + file.Name() + `">` + file.Name() + " FILE </a></br>"
				}
			}

			html += base_html_bottom

			fmt.Fprint(w, html)
		}
	} else if strings.Contains(path, "file/") {
		path = strings.Replace(path, "file/", "", 1)
		if _, err := os.Stat(wd + "/" + path); err == nil {
			http.ServeFile(w, r, wd+"/"+path)
			misc.Print("[200] /"+path, silent)
		} else {
			fmt.Fprintln(w, "[404] File Not Found!")
			misc.Print("[404] file "+path+" not found!", silent)
		}
	}

}
