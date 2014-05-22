package api

import (
	"fmt"

	"github.com/gocraft/web"
)

func (c *ApiContext) ImageIndex(res web.ResponseWriter, req *web.Request) {
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) ImageGet(res web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(res, "Not implemented")
}

func (c *ApiContext) ImagePut(res web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(res, "Not implemented")
}

func (c *ApiContext) ImageDelete(res web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(res, "Not implemented")
}
