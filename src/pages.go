package src

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
)

//ici on définit les pages statics!!!
func pages(w http.ResponseWriter, r *http.Request) {
	content, _ := getDefaultVar(r)
	lang := content["lang"].(string)
	//url for static pages switch
	ctmpl := "notfound"
	switch content["currentpage"] {
	case "/tos":
		ctmpl = "tos"
	case "/faq":
		ctmpl = "faq"
	case "/":
		ctmpl = "home"
	}
	Tmpl := template.Must(template.ParseFiles(
		TEMPLATE+"/base.html",
		TEMPLATE_I18N+"/"+ctmpl+"."+lang+".html"))
	Tmpl.Execute(w, content)
}

/*Genere les données de login*/
func getLoginData(r *http.Request) (*user.User, string, *appengine.Context) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	url := ""
	if u == nil {
		url, _ = user.LoginURL(c, "/")
	} else {
		url, _ = user.LogoutURL(c, "/")
	}
	return u, url, &c
}

/*Recupere la lang, le user, le context, l'url, etc*/
func getDefaultVar(r *http.Request) (map[string]interface{}, *appengine.Context) {
	lang := getHeaderLang(r)
	page := r.URL.Path
	user, url, context := getLoginData(r)
	isAdmin := ""
	if user != nil && user.Admin {
		isAdmin = "true"
	}

	content := map[string]interface{}{
		"lang":        lang,
		"i18n":        i18n[lang],
		"user":        user,
		"urlLogin":    url,
		"currentpage": page,
		"isAdmin":     isAdmin,
	}
	return content, context
}
