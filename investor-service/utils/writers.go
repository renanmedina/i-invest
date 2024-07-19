package utils

import (
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
)

func NewTableWriter() table.Writer {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)
	return writer
}
