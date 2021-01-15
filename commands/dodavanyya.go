package commands

import (
	"strconv"
	"fmt"
	"github.com/github.com/Labib17/lab4/tree/main/engine"
)

type Add struct {
	Arg1, Arg2 int
}

type Print struct {
	Arg string
}

func (a *Add) Execute(loop engine.Handler) {
	sum := a.Arg1 + a.Arg2
	loop.Post(&Print{Arg: strconv.Itoa(sum)})
}

func (p *Print) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
