package middleware

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocraft/web"
)

// RequestLogger prints a pretty JSON representation of incoming requests
func RequestLogger(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	b, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		log.Printf("Unable to create JSON representation of incoming request")
		next(res, req)
		return
	}

	fmt.Print(string(b))
	next(res, req)

}
