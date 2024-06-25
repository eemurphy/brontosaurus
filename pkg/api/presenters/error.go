package presenters

import (
	"github.com/eemurphy/brontosaurus/pkg/api/openapi"
	"github.com/eemurphy/brontosaurus/pkg/errors"
)

func PresentError(err *errors.ServiceError) openapi.Error {
	return err.AsOpenapiError("")
}
