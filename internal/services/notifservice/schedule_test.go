package notifservice

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"testing"
	"time"
)

func TestCalculateScheduledTime(t *testing.T) {
	logger.InitDefaultLogger(logger.WithDebug())
	now := time.Now()
	tolerance := time.Second * 2

	temp, _ := time.Parse("2006-01-02 15:04:05", "2025-03-08 21:45:33")

	tests := []struct {
		name          string
		lastScheduled time.Time
		interval      string
		expected      time.Time
	}{
		{
			name:          "NoRepeat",
			lastScheduled: now,
			interval:      IntervalNoRepeat,
			expected:      time.Time{}, // "нулевое" время
		},
		{
			name:          "Less than Minute",
			lastScheduled: now.Add(-10 * time.Second),
			interval:      IntervalMinute,
			expected:      now.Add(-10 * time.Second).Add(time.Minute),
		},
		{
			name:          "Over minute",
			lastScheduled: now.Add(-2*time.Minute - 10*time.Second),
			interval:      IntervalMinute,
			expected:      now.Add(-10 * time.Second),
		},
	}

	temp.Hour()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userPrefs := domains.UserPrefs{} // fixme
			result, err := calculateScheduledTime(tt.lastScheduled, tt.interval, userPrefs)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			logger.GetLogger().Infof("%v | %v | %v", tt.name, tt.lastScheduled, time.Now())

			// Проверяем, что результат находится в пределах допустимой погрешности
			if !result.IsZero() && (result.Sub(tt.expected) > tolerance || tt.expected.Sub(result) > tolerance) {
				t.Errorf("Expected time: %v, but got %v", tt.expected, result)
			}
		})
	}
}
