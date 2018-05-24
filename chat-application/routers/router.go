package routers

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"net/http"
)

func New() *http.ServeMux {
	router := http.NewServeMux()
	loginFs := http.FileServer(http.Dir("public/login"))
	homeFs  := http.FileServer(http.Dir("public/home"))
	router.Handle("/home/",http.StripPrefix("/home/",homeFs))
	//router.Handle("/home", &templates.TemplateHandler{FileName: "index.html", Once: &sync.Once{}})
	router.Handle("/login/", http.StripPrefix("/login/", loginFs))
	router.HandleFunc("/auth/callback/google", googleCallBackHandler)
	router.HandleFunc("/auth/login/google", googleLoginHandler)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/login/", http.StatusFound)
	})
	return router
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
		return
	}
	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get begin auth url %s: %s", provider, err), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, loginUrl, http.StatusTemporaryRedirect)
	return
}

func googleCallBackHandler(w http.ResponseWriter, r *http.Request) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
		return
	}
	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get credential %s: %s", provider, err), http.StatusInternalServerError)
		return
	}
	user, err := provider.GetUser(creds)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get user %s: %s", provider, err), http.StatusInternalServerError)
		return
	}
	authCookieValue := objx.New(map[string]interface{}{
		"name":  user.Name(),
		"email": user.Email(),
	}).MustBase64()
	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/",
	})
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)

}
