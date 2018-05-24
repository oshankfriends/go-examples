package authentication

import (
	"github.com/dghubble/sling"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Auth struct {
	Sling *sling.Sling
	Creds Credentials
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}
type AuthQueryParams struct {
	Url string `url:"url"`
}

func NewAuth(baseSling *sling.Sling, creds Credentials) *Auth {
	return &Auth{
		Creds: creds,
		Sling: baseSling.New().Get("/authorization/api/v1/token"),
	}
}

func (auth *Auth) GetToken() (string, error) {
	auth.Sling.QueryStruct(&AuthQueryParams{auth.Creds.Url})
	auth.Sling.SetBasicAuth(auth.Creds.Username, auth.Creds.Password)
	req, err := auth.Sling.Request()
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	token, err := ioutil.ReadAll(resp.Body)
	if err != nil && resp.StatusCode < 300 {
		return "", err
	}
	return url.QueryUnescape(string(token))
}
