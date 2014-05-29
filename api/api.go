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
	User     *models.User
	Database *mgo.Session
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
	r.Middleware((*Context).RequestLogger)
	r.Middleware((*Context).MongoDatabase)

	// Setup general routes
	r.Get("/", (*Context).Index)

	// Setup Docker registry API routes as per
	// http://docs.docker.io/reference/api/registry_index_spec/
	api := r.Subrouter(APIContext{}, "/v1")
	api.Middleware(ContentTypeJSON)
	api.Middleware((*APIContext).Authenticate)
	//api.Middleware((*APIContext).MultipartUpload)

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

	api.Get("/images/:image_id/checksum", (*APIContext).ImageChecksumGet)
	api.Put("/images/:image_id/checksum", (*APIContext).ImageChecksumPut)
	api.Delete("/images/:image_id/checksum", (*APIContext).ImageChecksumDelete)

	api.Get("/images/:image_id/ancestry", (*APIContext).AncestryGet)
	api.Put("/images/:image_id/ancestry", (*APIContext).AncestryPut)
	api.Delete("/images/:image_id/ancestry", (*APIContext).AncestryDelete)

	api.Put("/repositories/:owner/:repository", (*APIContext).RepositoryPut)
	api.Put("/repositories/:owner/:repository/auth", (*APIContext).RepositoryAuth)
	api.Delete("/repositories/:owner/:repository", (*APIContext).RepositoryDelete)

	api.Get("/repositories/:owner/:repository/images", (*APIContext).RepositoryImageGet)
	api.Put("/repositories/:owner/:repository/images", (*APIContext).RepositoryImagePut)

	api.Get("/repositories/:owner/:repository/tags", (*APIContext).TagsGet)
	api.Get("/repositories/:owner/:repository/tags/:tag", (*APIContext).TagGet)
	api.Put("/repositories/:owner/:repository/tags/:tag", (*APIContext).TagPut)
	api.Delete("/repositories/:owner/:repository/tags/:tag", (*APIContext).TagDelete)

	api.Get("/_ping", (*APIContext).Status)

	return r

}

// Index handles incoming GET requests on our root index '/'
func (c *Context) Index(res web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(res, "Hello World")
}
