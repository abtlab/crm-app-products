package products

import (
	"fmt"
	"net/http"
)

// Index :
func Index(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(writer, "OK")
}
