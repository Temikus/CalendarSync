package transformation

import (
	"testing"

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