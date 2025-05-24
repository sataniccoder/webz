package main

import (
	"fmt"
	"net/http"
	"os"
	"webz/core/get"
	"webz/core/list"
)

var wd string
var port string
var list_mode bool
var silent bool

func arg_parser(args []string) {
	for num, data := range args {
		// sort args out
		if data == "--list" {
			list_mode = true
		} else if data == "-d" {
			if num+1 >= len(args) {
				fmt.Println("[error] can't read -d verable, please make sure it has an argument")
				os.Exit(0)
			}
			wd = args[num+1]
		} else if data == "-p" {
			if num+1 >= len(args) {
				fmt.Println("[error] can't read -p verable, please make sure it has an argument")
				os.Exit(0)
			}
			port = args[num+1]
		} else if data == "-h" {
			help_menu()
		} else if data == "--silent" {
			silent = true
		} else {
			fmt.Println("[error] no such argument know '" + data + "' please use '-h' to see help menu")
			os.Exit(0)
		}
	}
}

func help_menu() {
	fmt.Println(`
	 WEBZ Usage

====== [ argument needed ] ======
-p {port} : port to use (default 8080)
-d {dir}  : directory to use (defult is current working dir)

======== [ no argument ] ========
-h        : display help menu
--list    : enter list mode allowing for file viewing and downloading
--silent  : silent mode only allows fetal error's to be shown in output

Example: webz -p 8081 -d /home/user/templates
	`)
	os.Exit(0)
}

func main() {
	// set defualt args
	port = "8080"
	wd, _ = os.Getwd()
	list_mode = false
	silent = false

	// arg handler
	arg_parser(os.Args[1:])
	fmt.Println(`
==================== [ GO-WEB ] ====================
            [+] a simple HTTP server [+] 
			            V0.1

URL    : http://127.0.0.1:`+port+`
WD     : `+wd+`
SILENT : `, silent, `
LIST   : `, list_mode, `

=====================================================
			
		`)

	if !list_mode {
		// config for get module
		get.Set_config(wd, silent)

		// start get module
		http.HandleFunc("/", get.Get_handle)
		http.ListenAndServe(":"+port, nil)
	} else {
		list.Set_config(wd, silent)

		http.HandleFunc("/", list.List_handle)
		http.ListenAndServe(":"+port, nil)
	}
}
