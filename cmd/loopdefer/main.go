package main

import (
	"github.com/gostaticanalysis/loopdefer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(loopdefer.Analyzer) }
