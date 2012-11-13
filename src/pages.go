package src

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
)

//ici on définit les pages !!!
func pages(w http.ResponseWriter, r *http.Request) {
	user, url := getLoginData(r)
	lang := getHeaderLang(r)
	page := r.URL.Path

	content := map[string]interface{}{
		"i18n":     i18n[lang],
		"user":     user,
		"urlLogin": url,
	}

	ctmpl := "templates/i18n/home." + lang + ".html"
	switch page {
	case "/phones":
		ctmpl = "templates/i18n/phone." + lang + ".html"
	case "/tos":
		ctmpl = "templates/i18n/tos." + lang + ".html"
	case "/faq":
		ctmpl = "templates/i18n/faq." + lang + ".html"
		/*case "/contact":
		ctmpl = "templates/i18n/contact." + lang + ".html"*/
	}

	Tmpl := template.Must(template.ParseFiles("templates/base.html", ctmpl))
	Tmpl.Execute(w, content)
}

/*Les funcs des pages*/

func getLoginData(r *http.Request) (*user.User, string) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	url := ""

	if u == nil {
		url, _ = user.LoginURL(c, "/")
	} else {
		url, _ = user.LogoutURL(c, "/")
	}
	return u, url
}
