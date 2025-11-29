package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseFormat(t *testing.T) {
	tests := []struct {
		input    string
		expected Format
	}{
		{"json", FormatJSON},
		{"JSON", FormatJSON},
		{"Json", FormatJSON},
		{"yaml", FormatYAML},
		{"YAML", FormatYAML},
		{"yml", FormatYAML},
		{"YML", FormatYAML},
		{"table", FormatTable},
		{"TABLE", FormatTable},
		{"", FormatTable},
		{"unknown", FormatTable},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ParseFormat(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPrinter_PrintJSON(t *testing.T) {
	data := map[string]string{"key": "value", "name": "test"}

	var buf bytes.Buffer
	printer := NewPrinter(FormatJSON)
	printer.SetWriter(&buf)

	err := printer.Print(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, `"key": "value"`) {
		t.Errorf("expected JSON output to contain key: value, got %s", output)
	}
	if !strings.Contains(output, `"name": "test"`) {
		t.Errorf("expected JSON output to contain name: test, got %s", output)
	}
}

func TestPrinter_PrintYAML(t *testing.T) {
	data := map[string]string{"key": "value", "name": "test"}

	var buf bytes.Buffer
	printer := NewPrinter(FormatYAML)
	printer.SetWriter(&buf)

	err := printer.Print(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "key: value") {
		t.Errorf("expected YAML output to contain key: value, got %s", output)
	}
	if !strings.Contains(output, "name: test") {
		t.Errorf("expected YAML output to contain name: test, got %s", output)
	}
}

type testTableData struct {
	items [][]string
}

func (d testTableData) Headers() []string {
	return []string{"NAME", "VALUE"}
}

func (d testTableData) Rows() [][]string {
	return d.items
}

func TestPrinter_PrintTable(t *testing.T) {
	data := testTableData{
		items: [][]string{
			{"item1", "value1"},
			{"item2", "value2"},
		},
	}

	var buf bytes.Buffer
	printer := NewPrinter(FormatTable)
	printer.SetWriter(&buf)

	err := printer.Print(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()

	// Check headers
	if !strings.Contains(output, "NAME") {
		t.Errorf("expected table output to contain NAME header, got %s", output)
	}
	if !strings.Contains(output, "VALUE") {
		t.Errorf("expected table output to contain VALUE header, got %s", output)
	}

	// Check rows
	if !strings.Contains(output, "item1") {
		t.Errorf("expected table output to contain item1, got %s", output)
	}
	if !strings.Contains(output, "value2") {
		t.Errorf("expected table output to contain value2, got %s", output)
	}

	// Check separator line exists
	if !strings.Contains(output, "----") {
		t.Errorf("expected table output to contain separator, got %s", output)
	}
}

func TestPrinter_PrintTable_NonTableData(t *testing.T) {
	// When data doesn't implement TableData, it should fall back to JSON
	data := map[string]string{"key": "value"}

	var buf bytes.Buffer
	printer := NewPrinter(FormatTable)
	printer.SetWriter(&buf)

	err := printer.Print(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, `"key": "value"`) {
		t.Errorf("expected JSON fallback output, got %s", output)
	}
}

func TestPrinter_PrintSuccess(t *testing.T) {
	var buf bytes.Buffer
	printer := NewPrinter(FormatTable)
	printer.SetWriter(&buf)

	printer.PrintSuccess("Operation completed")

	output := buf.String()
	if !strings.Contains(output, "Operation completed") {
		t.Errorf("expected success message, got %s", output)
	}
}

func TestPrinter_PrintSlice(t *testing.T) {
	data := []map[string]string{
		{"name": "item1"},
		{"name": "item2"},
	}

	var buf bytes.Buffer
	printer := NewPrinter(FormatJSON)
	printer.SetWriter(&buf)

	err := printer.Print(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "item1") {
		t.Errorf("expected output to contain item1, got %s", output)
	}
	if !strings.Contains(output, "item2") {
		t.Errorf("expected output to contain item2, got %s", output)
	}
}
