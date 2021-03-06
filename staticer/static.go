/*
	Staticer is a Middleware used to provide access to static files
*/
package staticer

import (
	"github.com/HairyMezican/TheRack/httper"
	"net/http"
)

type StaticProvider struct {
	prefix       string
	fileLocation string
}

//When getting a new StaticProvider, you must specify the prefix you want incoming requests, and where the files are located
func New(prefix, loc string) *StaticProvider {
	sp := new(StaticProvider)
	sp.prefix = prefix
	sp.fileLocation = loc
	return sp
}

//tests whether a starts with b
func StartsWith(a, b string) bool {
	if len(a) > len(b) && a[:len(b)] == b {
		return true
	}
	return false
}

func (this StaticProvider) Run(vars map[string]interface{}, next func()) {
	r := (httper.V)(vars).GetRequest()

	if StartsWith(r.URL.String(), this.prefix) {
		w := (httper.V)(vars).BlankResponse()
		http.StripPrefix(this.prefix, http.FileServer(http.Dir(this.fileLocation))).ServeHTTP(w, r)
		status, _, _ := w.Results()
		if status != http.StatusNotFound {
			w.Save()
			return
		}
	}
	next()
}
