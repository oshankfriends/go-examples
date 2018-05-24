package templates

import (
	"github.com/stretchr/objx"
	"html/template"
	"net/http"
	"sync"
	"path/filepath"
)

type TemplateHandler struct {
	FileName string
	tmpl     *template.Template
	Once     *sync.Once
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var UserData objx.Map
	if cookie, err := r.Cookie("auth"); err == nil {
		UserData = objx.MustFromBase64(cookie.Value)
	}
	t.Once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("public","home",t.FileName)))
	})
	t.tmpl.Execute(w,UserData)
}
