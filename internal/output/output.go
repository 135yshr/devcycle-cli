package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v3"
)

type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
	FormatYAML  Format = "yaml"
)

func ParseFormat(s string) Format {
	switch strings.ToLower(s) {
	case "json":
		return FormatJSON
	case "yaml", "yml":
		return FormatYAML
	default:
		return FormatTable
	}
}

type Printer struct {
	format Format
	writer io.Writer
}

func NewPrinter(format Format) *Printer {
	return &Printer{
		format: format,
		writer: os.Stdout,
	}
}

func (p *Printer) SetWriter(w io.Writer) {
	p.writer = w
}

func (p *Printer) Print(data any) error {
	switch p.format {
	case FormatJSON:
		return p.printJSON(data)
	case FormatYAML:
		return p.printYAML(data)
	default:
		return p.printTable(data)
	}
}

func (p *Printer) printJSON(data any) error {
	encoder := json.NewEncoder(p.writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (p *Printer) printYAML(data any) error {
	encoder := yaml.NewEncoder(p.writer)
	encoder.SetIndent(2)
	return encoder.Encode(data)
}

func (p *Printer) printTable(data any) error {
	if tableData, ok := data.(TableData); ok {
		return p.renderTable(tableData)
	}
	return p.printJSON(data)
}

type TableData interface {
	Headers() []string
	Rows() [][]string
}

func (p *Printer) renderTable(data TableData) error {
	w := tabwriter.NewWriter(p.writer, 0, 0, 2, ' ', 0)

	headers := data.Headers()
	fmt.Fprintln(w, strings.Join(headers, "\t"))

	separator := make([]string, len(headers))
	for i, h := range headers {
		separator[i] = strings.Repeat("-", len(h))
	}
	fmt.Fprintln(w, strings.Join(separator, "\t"))

	for _, row := range data.Rows() {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}

	return w.Flush()
}

func (p *Printer) PrintError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}

func (p *Printer) PrintSuccess(message string) {
	fmt.Fprintln(p.writer, message)
}
