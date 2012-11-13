package src

import (
	"net/http"
	"regexp"
)

/*
	Traductions et autres fonction liées aux langues, par ici ;)
*/
var i18n = map[string](map[string]string){
	"fr": {
		"Home":    "Accueil",
		"Hello":   "Bonjour",
		"Contact": "Contact",
		"Phone":   "Téléphone(s)",
		"FAQ":     "FaQ",
		"TOS":     "Condition d'utilisation",
		"Login":   "Se logguer",
		"Logout":  "Se déconnecter",
	},
	"en": {
		"Home":    "Home",
		"Hello":   "Hi",
		"Phone":   "Phone(s)",
		"FAQ":     "FAQ",
		"TOS":     "ToS",
		"Contact": "Contact",
		"Login":   "Log in",
		"Logout":  "Log out",
	},
}

func getHeaderLang(r *http.Request) string {
	lang := "en"
	lahe := r.Header.Get("accept-language")
	lare := regexp.MustCompile("([a-z]+)-[A-Z]+")

	if lare.Match([]byte(lahe)) {
		lang = lare.FindStringSubmatch(lahe)[1] //full substring is 0 (fr-FR), so 1 = fr 
	}
	switch lang {
	case "fr": //do nothing
	default:
		lang = "en"
	}
	return lang
}
