package service

import (
	"context"
	"github.com/otanfener/congestion-controller/pkg/models"
	"github.com/otanfener/congestion-controller/repos"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalculateTax(t *testing.T) {
	mockRepo := &repos.RepoMock{
		GetCityFunc: func(ctx context.Context, city string) (models.City, error) {
			return models.City{
				Name: "Stockholm",
				Tariff: []models.Tariff{
					{From: "06:00", To: "06:29", Fee: 8},
					{From: "06:30", To: "06:59", Fee: 13},
					{From: "07:00", To: "07:59", Fee: 18},
					{From: "08:00", To: "08:29", Fee: 13},
					{From: "08:30", To: "14:59", Fee: 8},
					{From: "15:00", To: "15:29", Fee: 13},
					{From: "15:30", To: "16:59", Fee: 18},
					{From: "17:00", To: "17:59", Fee: 13},
					{From: "18:00", To: "18:29", Fee: 8},
					{From: "18:30", To: "05:59", Fee: 0}},
				Rules:          models.Rules{MaxChargePerDay: 60, SingleChargeInterval: 60, WorkingDays: []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY"}},
				ExemptVehicles: []string{"Bicycle", "Tractor"},
			}, nil
		},
	}
	service := New(mockRepo)

	testCases := []struct {
		name        string
		city        string
		vehicle     string
		dates       []models.CivilTime
		expected    models.Tax
		expectedErr error
	}{
		{
			name:    "Empty dates list",
			city:    "Stockholm",
			vehicle: "car",
			dates:   nil,
			expected: models.Tax{
				TotalFee: 0,
				History:  nil,
			},
			expectedErr: nil,
		},
		{
			name:    "Tax exempt vehicle",
			city:    "Stockholm",
			vehicle: "Bicycle",
			dates:   []models.CivilTime{{Time: time.Date(2023, 4, 13, 8, 0, 0, 0, time.UTC)}},
			expected: models.Tax{
				TotalFee: 0,
				History:  nil,
			},
			expectedErr: nil,
		},
		{
			name:    "Normal taxed vehicle",
			city:    "Stockholm",
			vehicle: "Car",
			dates:   []models.CivilTime{{Time: time.Date(2023, 4, 13, 8, 0, 0, 0, time.UTC)}, {Time: time.Date(2023, 4, 13, 8, 30, 0, 0, time.UTC)}, {Time: time.Date(2023, 4, 13, 18, 0, 0, 0, time.UTC)}},
			expected: models.Tax{
				TotalFee: 21,
				History: map[string]int{
					"2023-04-13": 21,
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tax, err := service.CalculateTax(context.Background(), tc.dates, tc.city, tc.vehicle)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected.TotalFee, tax.TotalFee)
			assert.Equal(t, tc.expected.History, tax.History)
		})
	}
}
