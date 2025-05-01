package models

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEvent_Sync(t *testing.T) {
	startTime := time.Now()
	endTime := time.Now().Add(time.Hour)

	tests := []struct {
		name          string
		dest          Event
		source        Event
		expectedEvent Event
	}{
		{
			name: "overwrite dest event with source",
			dest: Event{
				ICalUID:     "foo",
				ID:          "bar",
				Title:       "Dest",
				Description: "Should stay",
				StartTime:   time.Now(),
				EndTime:     time.Now().Add(2 * time.Hour),
				AllDay:      false,
				Metadata: &Metadata{
					SyncID: "foo",
				},
			},
			source: Event{
				ICalUID:     "New ID",
				ID:          "New UUID",
				Title:       "Source",
				Description: "Should become",
				StartTime:   startTime,
				EndTime:     endTime,
				AllDay:      true,
				Metadata: &Metadata{
					SyncID: "foo",
				},
			},
			expectedEvent: Event{
				ICalUID:     "foo",
				ID:          "bar",
				Title:       "Source",
				Description: "Should become",
				StartTime:   startTime,
				EndTime:     endTime,
				AllDay:      true,
				Metadata: &Metadata{
					SyncID: "foo",
				},
			},
		},
		{
			name: "overwrite with colorID",
			dest: Event{
				ICalUID:     "foo",
				ID:          "bar",
				Title:       "Dest",
				StartTime:   time.Now(),
				EndTime:     time.Now().Add(2 * time.Hour),
				Metadata: &Metadata{
					SyncID: "foo",
				},
			},
			source: Event{
				ICalUID:     "New ID",
				ID:          "New UUID",
				Title:       "Source",
				StartTime:   startTime,
				EndTime:     endTime,
				ColorID:     "5",
				Metadata: &Metadata{
					SyncID: "foo",
				},
			},
			expectedEvent: Event{
				ICalUID:     "foo",
				ID:          "bar",
				Title:       "Source",
				StartTime:   startTime,
				EndTime:     endTime,
				ColorID:     "5",
				Metadata: &Metadata{
					SyncID: "foo",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.dest.Overwrite(tt.source)

			assert.Equal(t, tt.expectedEvent, actual)
		})
	}
}

func TestIsSameEvent(t *testing.T) {
	testTime := time.Now()
	endTime := testTime.Add(30 * time.Minute)

	check := []struct {
		name     string
		a        Event
		b        Event
		expected bool
	}{
		{
			name: "identical events",
			a: Event{
				Title:     "Title",
				StartTime: testTime,
				EndTime:   endTime,
			},
			b: Event{
				Title:     "Title",
				StartTime: testTime,
				EndTime:   endTime,
			},
			expected: true,
		},
		{
			name: "different title",
			a: Event{
				Title:     "A",
				StartTime: testTime,
				EndTime:   endTime,
			},
			b: Event{
				Title:     "B",
				StartTime: testTime,
				EndTime:   endTime,
			},
			expected: false,
		},
		{
			name: "different ColorID",
			a: Event{
				Title:     "Title",
				StartTime: testTime,
				EndTime:   endTime,
				ColorID:   "3",
			},
			b: Event{
				Title:     "Title",
				StartTime: testTime,
				EndTime:   endTime,
				ColorID:   "7",
			},
			expected: false,
		},
		{
			name: "one event has ColorID, other doesn't",
			a: Event{
				Title:     "Title",
				StartTime: testTime,
				EndTime:   endTime,
				ColorID:   "5",
			},
			b: Event{
				Title:     "Title",
				StartTime: testTime,
				EndTime:   endTime,
			},
			expected: false,
		},
	}

	for _, c := range check {
		t.Run(c.name, func(t *testing.T) {
			result := IsSameEvent(c.a, c.b)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestReminders_Sort(t *testing.T) {
	now := time.Now()
	tt := []struct {
		name      string
		Reminders Reminders
	}{
		{
			name: "Sort all reminders",
			Reminders: Reminders{
				{
					Actions: 3,
					Trigger: ReminderTrigger{
						PointInTime: now.Add(time.Hour),
					},
				},
				{
					Actions: 2,
					Trigger: ReminderTrigger{
						PointInTime: now.Add(time.Minute),
					},
				},
				{
					Actions: 1,
					Trigger: ReminderTrigger{
						PointInTime: now.Add(time.Second),
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sort.Sort(&tc.Reminders)

			for i := 1; i <= 3; i++ {
				assert.Equal(t, ReminderActions(i), tc.Reminders[i-1].Actions)
			}
		})
	}
}