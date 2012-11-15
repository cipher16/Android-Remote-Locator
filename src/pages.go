package src

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
)

//ici on définit les pages !!!
func pages(w http.ResponseWriter, r *http.Request) {
	lang := getHeaderLang(r)
	page := r.URL.Path
	user, url, context := getLoginData(r)
	//url switch
	ctmpl := "templates/i18n/notfound." + lang + ".html"
	switch page {
	case "/gcm":
		GCM(context, w, r)
		return
	case "/phones":
		Phones(w, r, lang, user, context, url)
		return
		//		ctmpl = "templates/i18n/phone." + lang + ".html"
	case "/tos":
		ctmpl = "templates/i18n/tos." + lang + ".html"
	case "/faq":
		ctmpl = "templates/i18n/faq." + lang + ".html"
	case "/":
		ctmpl = "templates/i18n/home." + lang + ".html"
		/*case "/contact":
		ctmpl = "templates/i18n/contact." + lang + ".html"*/
	}

	content := map[string]interface{}{
		"i18n":        i18n[lang],
		"user":        user,
		"urlLogin":    url,
		"currentpage": page,
	}

	Tmpl := template.Must(template.ParseFiles("templates/base.html", ctmpl))
	Tmpl.Execute(w, content)
}

/*Les funcs des pages*/

func getLoginData(r *http.Request) (*user.User, string, appengine.Context) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	url := ""

	if u == nil {
		url, _ = user.LoginURL(c, "/")
	} else {
		url, _ = user.LogoutURL(c, "/")
	}
	return u, url, c
}
