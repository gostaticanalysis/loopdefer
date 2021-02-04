package loopdefer

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "loopdefer find using defer in loop"

// Analyzer find using defer in loop
var Analyzer = &analysis.Analyzer{
	Name: "loopdefer",
	Doc:  doc,
	Run:  new(analyzer).run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

type analyzer struct {
	done map[*ssa.BasicBlock]bool
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, f := range funcs {
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				a.done = make(map[*ssa.BasicBlock]bool)
				if _, ok := instr.(*ssa.Defer); ok && a.isLoop(b, b.Succs) {
					pass.Reportf(instr.Pos(), "defer should not use in a loop")
				}
			}
		}
	}
	return nil, nil
}

func (a *analyzer) isLoop(root *ssa.BasicBlock, blocks []*ssa.BasicBlock) bool {
	for _, b := range blocks {
		if a.done[b] {
			continue
		}
		a.done[b] = true
		if b == root || a.isLoop(root, b.Succs) {
			return true
		}
	}
	return false
}
