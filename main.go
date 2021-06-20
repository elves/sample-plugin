package main

import (
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/vars"
)

var Ns = eval.NsBuilder{"foo": vars.NewReadOnly("bar")}.Ns()
