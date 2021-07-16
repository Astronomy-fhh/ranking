package test

import (
	"github.com/fatih/color"
	"os"
	"testing"
)

func TestColor(t *testing.T) {

	str := "127.0.0.1:11917"
	c := color.New(color.FgCyan,color.Bold).FprintfFunc()
	c(os.Stdout, str)
}