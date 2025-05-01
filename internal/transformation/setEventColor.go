package transformation

import (
	"github.com/inovex/CalendarSync/internal/models"
)

// SetEventColor allows setting a specific color for events in Google Calendar.
// This is currently only implemented for Google Calendar. Other calendars will ignore this.
type SetEventColor struct {
	// ColorID is the color identifier used by Google Calendar.
	// Valid values are "1" through "11" (as strings).
	ColorID string
}

func (t *SetEventColor) Name() string {
	return "SetEventColor"
}

func (t *SetEventColor) Transform(source models.Event, sink models.Event) (models.Event, error) {
	sink.ColorID = t.ColorID
	return sink, nil
}