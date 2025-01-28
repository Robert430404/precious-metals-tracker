package services

import (
	"testing"

	"github.com/robert430404/precious-metals-tracker/db/entities"
)

func TestCalculateMetalWeight(t *testing.T) {
	calcService := GetCalculationService()

	tests := []struct {
		name           string
		holdings       []entities.Holding
		expectedResult float64
	}{
		{
			name:           "empty holdings",
			holdings:       []entities.Holding{},
			expectedResult: 0,
		},
		{
			name: "single valid holding",
			holdings: []entities.Holding{
				{UnitWeight: "5.0", TotalUnits: "10"},
			},
			expectedResult: 50.0,
		},
		{
			name: "multiple holdings with positive weights and units",
			holdings: []entities.Holding{
				{UnitWeight: "2.5", TotalUnits: "4"},
				{UnitWeight: "1.5", TotalUnits: "3"},
			},
			expectedResult: 10 + 4.5, // 15.5
		},
		{
			name: "holdings with invalid unit weight",
			holdings: []entities.Holding{
				{UnitWeight: "invalid", TotalUnits: "10"}, // Ignored due to error
				{UnitWeight: "4.0", TotalUnits: "5"},
			},
			expectedResult: 20,
		},
		{
			name: "holdings with invalid total units",
			holdings: []entities.Holding{
				{UnitWeight: "3.0", TotalUnits: "invalid"}, // Ignored due to error
				{UnitWeight: "2.0", TotalUnits: "8"},
			},
			expectedResult: 16,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calcService.CalculateMetalWeight(test.holdings)
			if result != test.expectedResult {
				t.Errorf("expected %f, got %f", test.expectedResult, result)
			}
		})
	}
}
