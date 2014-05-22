package api

import (
	"fmt"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
	"labix.org/v2/mgo"
)

// Context is inialised for each user by the middleware and passed
// to the URL handlers when it can be inspected/updated.
type Context struct {
	User          *models.User
	Database      *mgo.Session
	AuthWhitelist []string
}

// APIContext is inialised for each API user by the middleware and passed
// to the URL handlers when it can be inspected/updated.
type APIContext struct {
	*Context
}

// Database is used to store a handle to the *mgo.Session
// so that it can be injected by the middleware into the context
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
	api := r.Subrouter(APIContext{}, "/v1")
	api.Middleware(ContentTypeJSON)
	api.Middleware((*APIContext).AuthenticationWhitelist)
	api.Middleware((*APIContext).BasicAuthentication)

	api.Get("/users", (*APIContext).UserLogin)
	api.Post("/users", (*APIContext).UserCreate)
	api.Put("/users/:username", (*APIContext).UserUpdate)

	api.Get("/images/:image_id/layer", (*APIContext).LayerGet)
	api.Put("/images/:image_id/layer", (*APIContext).LayerPut)
	api.Delete("/images/:image_id/layer", (*APIContext).LayerDelete)

	api.Get("/images", (*APIContext).ImageIndex)
	api.Get("/images/:image_id/json", (*APIContext).ImageGet)
	api.Put("/images/:image_id/json", (*APIContext).ImagePut)
	api.Delete("/images/:image_id/json", (*APIContext).ImageDelete)

	api.Get("/images/:image_id/ancestry", (*APIContext).AncestryGet)
	api.Put("/images/:image_id/ancestry", (*APIContext).AncestryPut)
	api.Delete("/images/:image_id/ancestry", (*APIContext).AncestryDelete)

	api.Put("/repositories/:namespace/:repository", (*APIContext).RepositoryPut)
	api.Put("/repositories/:namespace/:repository/auth", (*APIContext).RepositoryAuth)
	api.Delete("/repositories/:namespace/:repository", (*APIContext).RepositoryDelete)

	api.Get("/repositories/:namespace/:repository/tags", (*APIContext).TagsGet)
	api.Get("/repositories/:namespace/:repository/tags/:tag", (*APIContext).TagGet)
	api.Put("/repositories/:namespace/:repository/tags/:tag", (*APIContext).TagPut)
	api.Delete("/repositories/:namespace/:repository/tags/:tag", (*APIContext).TagDelete)

	api.Get("/repositories/:namespace/:repository/images", (*APIContext).RepositoryImageGet)
	api.Put("/repositories/:namespace/:repository/images", (*APIContext).RepositoryImagePut)

	api.Get("/_ping", (*APIContext).Status)

	return r

}

// Index handles incoming GET requests on our root index '/'
func (c *Context) Index(res web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(res, "Hello World")
}
