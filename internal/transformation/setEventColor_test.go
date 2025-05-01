package transformation

import (
	"testing"
	"time"

	"github.com/inovex/CalendarSync/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSetEventColor_Transform(t *testing.T) {
	tests := []struct {
		name     string
		colorID  string
		source   models.Event
		sink     models.Event
		expected models.Event
	}{
		{
			name:    "sets color ID on sink event",
			colorID: "7",
			source:  models.Event{},
			sink:    models.Event{},
			expected: models.Event{
				ColorID: "7",
			},
		},
		{
			name:    "overrides existing color ID",
			colorID: "5",
			source:  models.Event{},
			sink: models.Event{
				ColorID: "9",
			},
			expected: models.Event{
				ColorID: "5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformer := &SetEventColor{
				ColorID: tt.colorID,
			}

			result, err := transformer.Transform(tt.source, tt.sink)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.ColorID, result.ColorID)
		})
	}
}

// TestSetEventColor_ExistingEvent tests the scenario where an existing event
// that was created before the color feature was added now gets a color assigned.
// This simulates what happens when updating existing events after deploying the color feature.
func TestSetEventColor_ExistingEvent(t *testing.T) {
	// Setup existing event (as it would have been before this feature)
	metadata := models.NewEventMetadata("event123", "https://calendar.com/event/123", "source123")
	existingEvent := models.Event{
		ID:          "event123",
		Title:       "Existing meeting",
		Description: "Meeting description",
		Location:    "Conference room",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(1 * time.Hour),
		Metadata:    metadata,
		// No ColorID set (as it would be for events created before this feature)
	}

	// Set ColorID to our existing event via the transformer
	transformer := &SetEventColor{
		ColorID: "4",
	}

	// Transform the event
	result, err := transformer.Transform(models.Event{}, existingEvent)
	
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "4", result.ColorID)
	
	// Verify other properties remained unchanged
	assert.Equal(t, existingEvent.ID, result.ID)
	assert.Equal(t, existingEvent.Title, result.Title)
	assert.Equal(t, existingEvent.Description, result.Description)
	assert.Equal(t, existingEvent.StartTime, result.StartTime)
	assert.Equal(t, existingEvent.EndTime, result.EndTime)
	assert.Equal(t, existingEvent.Metadata, result.Metadata)
}