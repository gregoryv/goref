package main

import (
	"os/exec"
	"testing"

	"github.com/gregoryv/golden"
)

func Test_mkset(t *testing.T) {
	out, _ := exec.Command("go", "generate", ".").Output()
	golden.Assert(t, string(out))
}

//go:generate mkset -t Car,Boat
type Car struct {
	Name string

	model string
	make  int
}

type Boat struct {
	color int
}