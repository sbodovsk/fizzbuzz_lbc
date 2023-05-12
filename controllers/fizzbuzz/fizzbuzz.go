package fizzbuzz

import (
	"net/http"

	"github.com/gorilla/mux"
)

// nameSpace defines how to reach this controller by route api
const (
	nameSpace = "fizzbuzz"
)

// Controller is the interface describing the abilities of the fizzbuzz controller.
type Controller interface {
	GetName() string
	GetFizzBuzz(w http.ResponseWriter, r *http.Request)
	Register(r *mux.Router)
}

// controller struct defines what's needed by the controller to be functional
type controller struct {
	name string
}

// NewController creates and returns a new Controller.
func NewController() Controller {
	return &controller{
		name: nameSpace,
	}
}

// Register takes a mux.Router and registers the controllers routes.
func (c controller) Register(r *mux.Router) {
	r.HandleFunc("", c.GetFizzBuzz).Methods(http.MethodGet)
}

// GetName returns the name of the controller.
//
// It is used to register the controller on the mux.Router.
func (c controller) GetName() string {
	return c.name
}
