package handler

import (
	"testing"
)

func TestExec(t *testing.T) {
	file := "./test.sh"
	ExecShell(file)
}
