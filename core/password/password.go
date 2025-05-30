package password

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"
	"webz/core/log"
)

var coockie_database []string
var passwd string
var template string

// config stuff
func Config(new_passwd string) {
	passwd = new_passwd
	template = `	
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Webz - Aauth</title>
</head>
<body>
    <h3>Login Required</h3>
    <form method="post" action="/auth/login">
        Password: <input type="password" placeholder="Password" name="pass">
        <input type="submit">
    </form>
</body>
</html>

	`
}

// coockie storage/deletion stuff
func add_coockie() http.Cookie {
	// create a random string
	coc := rand.Text()

	// add it to the database
	coockie_database = append(coockie_database, coc)

	cookie := &http.Cookie{
		Name:     "verify",
		Value:    coc,
		Path:     "/",
		Expires:  time.Now().Add(2 * time.Hour),
		HttpOnly: true,
	}

	return *cookie
}

func delete_coockie(coockie string) {

}

// auth handle
func Check_coockie(w http.ResponseWriter, r *http.Request) bool {
	// grab the coockie
	coc, err := r.Cookie("verify")

	if err != nil {
		log.Error("[error] couldn't read coockie due too " + err.Error())
		return false
	}

	// check if coockie is within the "db"
	for _, coc_chk := range coockie_database {
		if coc.Value == coc_chk {
			return true
		}
	}

	return false

}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	} else {
		// grab form data
		err := r.ParseForm()
		if err != nil {
			log.Error("[error] couldn't parse login form due too " + err.Error())
			http.Redirect(w, r, "/auth", 200)

			return
		} else {

			pass := r.FormValue("pass")
			if pass == passwd {
				// create & add coockie
				coc := add_coockie()

				http.SetCookie(w, &coc)
				http.Redirect(w, r, "/", http.StatusFound)

			} else {
				http.Redirect(w, r, "/auth", 200)
			}
			return
		}
	}
}

func Auth_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, template)
	return
}
