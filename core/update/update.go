package update

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"webz/core/log"
)

func Check_version(version float64) bool {
	url := "https://raw.githubusercontent.com/sataniccoder/webz/refs/heads/main/webz.go"

	log.Info("[info] checking version")

	res, err := http.Get(url)

	if err != nil {
		log.Error("[error] couldn't get requested url due too " + err.Error())
		return false
	}

	html, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("[error] couldn't read reqested body due too " + err.Error())
		return false
	}

	html_string := strings.Split(string(html), "\n")[2]

	html_string = strings.Split(html_string, " ")[2]

	version_chk, err := strconv.ParseFloat(html_string, 64)
	if err != nil {
		log.Error("[error] couldn't parse version number " + html_string + " due to " + err.Error())
		return false
	}

	if version < version_chk {
		log.Warning("[warning] please update your webz install via the '--update' command")
		return true
	} else {
		log.Info("[info] webz is upto-date at version " + html_string)
		return false
	}
}

func Update() {
	fmt.Println("Sorry, but this hasn't been fully programmned yet, please check the roadmap to see when it should be added")
	os.Exit(0)

	user_valid := false
	var inp_chk string
	var inp bool

	for !user_valid {
		fmt.Println(`
============ [START OF WARNING] ============

this webz about to self-update
if you compiled webz from source then you should hit CNTRL+C or type 'n'
if you downloaded the pre-compiled version of webz then hit 'y' as it will auto 
download and setup the new webz version

============= [END OF WARNING] =============
		`)

		fmt.Print("Would you like to continue updating [y/n] -> ")
		fmt.Scanln(&inp_chk)

		if inp_chk == "n" {
			inp = false
			user_valid = true
		} else if inp_chk == "y" {
			inp = true
			user_valid = true
		} else {
			fmt.Println(inp_chk + " isn't a valid input")
		}
	}

	if !inp {
		os.Exit(0)
	}

}
