package presenters

import (
	"github.com/eemurphy/brontosaurus/pkg/api"
	"github.com/eemurphy/brontosaurus/pkg/api/openapi"
	"github.com/eemurphy/brontosaurus/pkg/errors"
)

func ObjectKind(i interface{}) *string {
	result := ""
	switch i.(type) {
	case api.Dinosaur, *api.Dinosaur:
		result = "Dinosaur"
	case errors.ServiceError, *errors.ServiceError:
		result = "Error"
	}

	return openapi.PtrString(result)
}
