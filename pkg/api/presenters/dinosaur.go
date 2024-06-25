package presenters

import (
	"github.com/eemurphy/brontosaurus/pkg/api"
	"github.com/eemurphy/brontosaurus/pkg/api/openapi"
	"github.com/eemurphy/brontosaurus/pkg/util"
)

func ConvertDinosaur(dinosaur openapi.Dinosaur) *api.Dinosaur {
	return &api.Dinosaur{
		Meta: api.Meta{
			ID: util.NilToEmptyString(dinosaur.Id),
		},
		Species: util.NilToEmptyString(dinosaur.Species),
	}
}

func PresentDinosaur(dinosaur *api.Dinosaur) openapi.Dinosaur {
	reference := PresentReference(dinosaur.ID, dinosaur)
	return openapi.Dinosaur{
		Id:        reference.Id,
		Kind:      reference.Kind,
		Href:      reference.Href,
		Species:   openapi.PtrString(dinosaur.Species),
		CreatedAt: openapi.PtrTime(dinosaur.CreatedAt),
		UpdatedAt: openapi.PtrTime(dinosaur.UpdatedAt),
	}
}
