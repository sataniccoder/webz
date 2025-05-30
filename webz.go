package main

// version 0.2
import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"webz/core/get"
	"webz/core/list"
	"webz/core/log"
	"webz/core/password"
	"webz/core/update"
)

var wd string
var port string
var list_mode bool
var silent bool
var log_path string
var log_stat bool
var version float64
var passwd string
var need_passwd bool

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
		} else if data == "-h" || data == "--help" {
			help_menu()
		} else if data == "--silent" {
			silent = true
		} else if data == "-l" {
			if num+1 >= len(args) {
				fmt.Println("[error] can't read -p verable, please make sure it has an argument")
				os.Exit(0)
			}
			if strings.Contains(args[num+1], "webz.log") {
				log_path = args[num+1]
			} else {
				log_path = args[num+1] + "/webz.log"
			}

			log_stat = true
		} else if data == "--update" {
			if update.Check_version(version) {
				update.Update()
			}
		} else if data == "--force-update" {
			update.Update()
		} else if data == "--passwd" {
			if num+1 >= len(args) {
				fmt.Println("[error] can't read -p verable, please make sure it has an argument")
				os.Exit(0)
			}

			passwd = args[num+1]
			need_passwd = true
		}
	}
}

func help_menu() {
	fmt.Println(`
	 WEBZ Usage

====== [ argument needed ] ======
-p {port}      : port to use (default 8080)
-d {dir}       : directory to use (defult is current working dir)
-l {dir}       : set log file path, if empty it won't be used (EG: /home/usr/logs/)
--passwd {pas} : set password for site (doesn't matter on the mode)

======== [ no argument ] ========
-h             : display help menu
--list         : enter list mode allowing for file viewing and downloading
--silent       : silent mode only allows fetal error's to be shown in output
--update       : update webz program
--force-update : force update webz even if current version is the latest

Example: webz -p 8081 -d /home/user/templates
	`)
	os.Exit(0)
}

func main() {
	// update check
	version = 0.2
	update.Check_version(version)

	// set defualt args
	port = "8080"
	wd, _ = os.Getwd()
	list_mode = false
	silent = false
	log_path = ""
	log_stat = false
	need_passwd = false
	passwd = "None"

	// arg handler
	arg_parser(os.Args[1:])

	// log config
	log.Set_config(log_path, log_stat, silent)

	fmt.Println(`
==================== [ GO-WEB ] ====================
            [+] a simple HTTP server [+] 
			            V0.1
WD     : `+wd+`
URL    : http://127.0.0.1:`+port+`
LOG    : `+log_path+`
SILENT :`, silent, `
LIST   :`, list_mode, `


=====================================================
			
		`)

	// setup login stuff if needed
	if need_passwd {
		password.Config(passwd)
		http.HandleFunc("/auth", password.Auth_handler)
		http.HandleFunc("/auth/login", password.Login)
	}

	if !list_mode {
		// config for get module
		get.Set_config(wd, need_passwd)

		// start get module
		http.HandleFunc("/", get.Get_handle)
		http.ListenAndServe(":"+port, nil)
	} else {
		// config setip for list mode
		list.Set_config(wd)

		// start list module
		http.HandleFunc("/", list.List_handle)
		http.ListenAndServe(":"+port, nil)
	}
}
