package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/eemurphy/brontosaurus/cmd/brontosaurus/server"
	"github.com/eemurphy/brontosaurus/pkg/api"
	"github.com/eemurphy/brontosaurus/pkg/controllers"
	"github.com/eemurphy/brontosaurus/pkg/dao"
	"github.com/eemurphy/brontosaurus/pkg/db"
	"github.com/eemurphy/brontosaurus/test"
	. "github.com/onsi/gomega"
)

func TestControllerRacing(t *testing.T) {
	h, _ := test.RegisterIntegration(t)

	account := h.NewRandAccount()
	authCtx := h.NewAuthenticatedContext(account)
	dao := dao.NewEventDao(&h.Env().Database.SessionFactory)

	// The handler filters the events by source id/type/reconciled, and only record
	// the event with create type. Due to the event lock, each create event
	// should be only processed once.
	var proccessedEvent []string
	onUpsert := func(ctx context.Context, id string) error {
		events, err := dao.All(authCtx)
		if err != nil {
			return err
		}

		for _, evt := range events {
			if evt.SourceID != id {
				continue
			}
			if evt.EventType != api.CreateEventType {
				continue
			}
			// the event has been reconciled by others, ignore.
			if evt.ReconciledDate != nil {
				continue
			}
			proccessedEvent = append(proccessedEvent, id)
		}

		return nil
	}

	// Start 3 controllers concurrently
	threads := 3
	for i := 0; i < threads; i++ {
		go func() {
			s := &server.ControllersServer{
				KindControllerManager: controllers.NewKindControllerManager(
					db.NewAdvisoryLockFactory(h.Env().Database.SessionFactory),
					h.Env().Services.Events(),
				),
			}

			s.KindControllerManager.Add(&controllers.ControllerConfig{
				Source: "Dinosaurs",
				Handlers: map[api.EventType][]controllers.ControllerHandlerFunc{
					api.CreateEventType: {onUpsert},
					api.UpdateEventType: {onUpsert},
				},
			})

			s.Start()
		}()
	}

	_, err := h.Factories.NewDinosaurList("bronto", 50)
	Expect(err).NotTo(HaveOccurred())

	// This is to check only two create events is processed. It waits for 5 seconds to ensure all events have been
	// processed by the controllers.
	Eventually(func() error {
		if len(proccessedEvent) != 50 {
			return fmt.Errorf("should have only 2 create events but got %d", len(proccessedEvent))
		}
		return nil
	}, 5*time.Second, 1*time.Second).Should(Succeed())
}
