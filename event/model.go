package event

type Event struct {
	ID         string `json:"id"`
	Timestamp  int64  `json:"timestamp"`
	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`
	Action     string `json:"action"`
}

// retrieve events between specific times
//
// includes from, excludes to
func (e *Event) GetBetween(from, to int64) ([]Event, error) {
	return nil, nil
}

// retrieves events from epoch until now
func (e *Event) GetSince(epoch int64) ([]Event, error) {
	return nil, nil
}

// creates an event
func (e *Event) Create() (*Event, error) {
	return nil, nil
}
