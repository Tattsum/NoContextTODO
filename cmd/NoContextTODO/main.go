package main

import (
	"github.com/Tattsum/NoContextTODO"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(NoContextTODO.Analyzer) }
