package function

import (
	"net/http"
	"os"
)

var fsPath = os.Getenv("ROOT_DIR")
var fs = http.FileServer(http.Dir(fsPath))
var ps = http.StripPrefix("/policies", fs)

//Handle serves any files that located in 'policies' path
func Handle(w http.ResponseWriter, r *http.Request) {
	ps.ServeHTTP(w, r)
}
