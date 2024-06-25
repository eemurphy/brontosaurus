package presenters

import (
	"fmt"

	"github.com/eemurphy/brontosaurus/pkg/api/openapi"

	"github.com/eemurphy/brontosaurus/pkg/api"
	"github.com/eemurphy/brontosaurus/pkg/errors"
)

const (
	BasePath = "/api/brontosaurus/v1"
)

func ObjectPath(id string, obj interface{}) *string {
	return openapi.PtrString(fmt.Sprintf("%s/%s/%s", BasePath, path(obj), id))
}

func path(i interface{}) string {
	switch i.(type) {
	case api.Dinosaur, *api.Dinosaur:
		return "dinosaurs"
	case errors.ServiceError, *errors.ServiceError:
		return "errors"
	default:
		return ""
	}
}
