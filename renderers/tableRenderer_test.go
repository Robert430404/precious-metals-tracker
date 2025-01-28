package renderers

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRenderValueTable(t *testing.T) {
	renderer := &TableRenderer{}

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

	renderer.RenderValueTable(mockData)

	writer.Close()
	out, _ := io.ReadAll(reader)
	os.Stdout = rescueStdout

	stringOut := string(out)

	expectedToContain := "Metal Type"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	expectedToContain = "Current Value"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	expectedToContain = "Current Spot Price"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	expectedToContain = "Total Holding Weight"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	expectedToContain = "one"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	expectedToContain = "two"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	expectedToContain = "three"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	expectedToContain = "four"
	if !strings.Contains(stringOut, expectedToContain) {
		t.Fatalf("Output does not have: %s", expectedToContain)
	}

	splitOut := strings.Split(stringOut, "\n")
	splitExpected := []string{
		"┌────────────────────────────────────────────────────────────────────────┐",
		"│ Metal Type │ Current Value │ Current Spot Price │ Total Holding Weight │",
		"├────────────────────────────────────────────────────────────────────────┤",
		"│ one        │ two           │ three              │ four                 │",
		"└────────────────────────────────────────────────────────────────────────┘",
	}

	firstLine := renderer.stripAnsiCodes(strings.Trim(splitOut[0], " "))
	if firstLine != splitExpected[0] {
		t.Fatalf("First line did not match: \n received: %s \n expected: %s", firstLine, splitExpected[0])
	}

	secondLine := renderer.stripAnsiCodes(strings.Trim(splitOut[1], " "))
	if secondLine != splitExpected[1] {
		t.Fatalf("Second line did not match: \n received: %s \n expected: %s", secondLine, splitExpected[1])
	}

	thirdLine := renderer.stripAnsiCodes(strings.Trim(splitOut[2], " "))
	if thirdLine != splitExpected[2] {
		t.Fatalf("Third line did not match: \n received: %s \n expected: %s", thirdLine, splitExpected[2])
	}

	fourthLine := renderer.stripAnsiCodes(strings.Trim(splitOut[3], " "))
	if fourthLine != splitExpected[3] {
		t.Fatalf("Fourth line did not match: \n received: %s \n expected: %s", fourthLine, splitExpected[3])
	}

	fifthLine := renderer.stripAnsiCodes(strings.Trim(splitOut[4], " "))
	if fifthLine != splitExpected[4] {
		t.Fatalf("Fifth line did not match: \n received: %s \n expected: %s", fifthLine, splitExpected[4])
	}
}
