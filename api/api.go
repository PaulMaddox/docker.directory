package api

import (
	"fmt"
	"time"
	"labix.org/v2/mgo"

	"github.com/gocraft/web"
)

type Context struct {
	Id       int64
	Started  time.Time
	Modified time.Time
	Database *mgo.Session
}

type ApiContext struct {
	*Context
}

var Database *mgo.Session

// NewRouter generates a new Mux router that implements
// the various Docker Registry API calls required
func NewRouter(db *mgo.Session) *web.Router {

	Database = db
	r := web.New(Context{})

	// Setup middleware
	r.Middleware(web.LoggerMiddleware)
	r.Middleware(web.ShowErrorsMiddleware)
	r.Middleware((*Context).Version)
	r.Middleware((*Context).MongoDatabase)
	//r.Middleware((*Context).RequestLogger)

	// Setup general routes
	r.Get("/", (*Context).Index)

	// Setup Docker registry API routes as per
	// http://docs.docker.io/reference/api/registry_index_spec/
	api := r.Subrouter(ApiContext{}, "/v1")
	api.Middleware(ContentTypeJson)

	api.Get("/users", (*ApiContext).UserLogin)
	api.Post("/users", (*ApiContext).UserCreate)
	api.Put("/users/:username", (*ApiContext).UserUpdate)

	api.Get("/images/:image_id/layer", (*ApiContext).LayerGet)
	api.Put("/images/:image_id/layer", (*ApiContext).LayerPut)
	api.Delete("/images/:image_id/layer", (*ApiContext).LayerDelete)

	api.Get("/images", (*ApiContext).ImageIndex)
	api.Get("/images/:image_id/json", (*ApiContext).ImageGet)
	api.Put("/images/:image_id/json", (*ApiContext).ImagePut)
	api.Delete("/images/:image_id/json", (*ApiContext).ImageDelete)

	api.Get("/images/:image_id/ancestry", (*ApiContext).AncestryGet)
	api.Put("/images/:image_id/ancestry", (*ApiContext).AncestryPut)
	api.Delete("/images/:image_id/ancestry", (*ApiContext).AncestryDelete)

	api.Put("/repositories/:namespace/:repository", (*ApiContext).RepositoryPut)
	api.Put("/repositories/:namespace/:repository/auth", (*ApiContext).RepositoryAuth)
	api.Delete("/repositories/:namespace/:repository", (*ApiContext).RepositoryDelete)

	api.Get("/repositories/:namespace/:repository/tags", (*ApiContext).TagsGet)
	api.Get("/repositories/:namespace/:repository/tags/:tag", (*ApiContext).TagGet)
	api.Put("/repositories/:namespace/:repository/tags/:tag", (*ApiContext).TagPut)
	api.Delete("/repositories/:namespace/:repository/tags/:tag", (*ApiContext).TagDelete)

	api.Get("/repositories/:namespace/:repository/images", (*ApiContext).RepositoryImageGet)
	api.Put("/repositories/:namespace/:repository/images", (*ApiContext).RepositoryImagePut)

	api.Get("/_ping", (*ApiContext).Status)

	return r

}

// Index handles incoming GET requests on our root index '/'
func (c *Context) Index(res web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(res, "Hello World")
}
