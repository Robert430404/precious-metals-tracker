package renderers

import (
	"io"
	"os"
	"testing"

	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/models"
)

func TestJsonRendererRenderHoldingList(t *testing.T) {
	renderer := &JsonRenderer{}

	mockData := []entities.Holding{
		{
			Name:              "name",
			Source:            "source",
			PurchaseSpotPrice: "spot-price",
			TotalUnits:        "total-units",
			UnitWeight:        "unit-weight",
			Type:              models.Gold,
		},
	}

	rescueStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	renderer.RenderHoldingList(mockData)

	writer.Close()
	out, _ := io.ReadAll(reader)
	os.Stdout = rescueStdout

	expectedOutput := "[{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Name\":\"name\",\"Source\":\"source\",\"PurchaseSpotPrice\":\"spot-price\",\"TotalUnits\":\"total-units\",\"UnitWeight\":\"unit-weight\",\"Type\":\"Gold\"}]"
	if string(out) != expectedOutput {
		t.Fatalf("Output does not match: %s", out)
	}
}

func TestJsonRendererRenderValueList(t *testing.T) {
	renderer := &JsonRenderer{}

	mockData := [][]string{
		{
			"one",
			"two",
			"three",
			"four",
		},
	}

	rescueStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	renderer.RenderValueList(mockData)

	writer.Close()
	out, _ := io.ReadAll(reader)
	os.Stdout = rescueStdout

	expectedOutput := "[{\"type\": \"one\", \"currentValue\": \"two\", \"currentSpotPrice\": \"three\", \"totalHoldingWeight\": \"four\"}]"
	if string(out) != expectedOutput {
		t.Fatalf("Output does not match: %s", out)
	}
}
