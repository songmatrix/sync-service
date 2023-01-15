package server

import (
	"net/http"
	"time"

	"github.com/songmatrix/sync-service/event"
)

func getEvents(w http.ResponseWriter, r *http.Request) {
	responseOk(w, []event.Event{{
		ID:         "foo",
		Timestamp:  time.Now().Unix(),
		EntityType: "track",
		EntityID:   "bar",
		Action:     "create",
	}})
}

func createEvents(w http.ResponseWriter, r *http.Request) {
	responseOk[any](w, nil)
}
