package ir

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/value"
)

// === [ Terminators ] =========================================================

// --- [ ret ] -----------------------------------------------------------------

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a void return.
func (block *BasicBlock) NewRet(x value.Value) *TermRet {
	term := NewRet(x)
	block.Term = term
	return term
}

// --- [ br ] ------------------------------------------------------------------

// NewBr returns a new unconditional br terminator based on the given target
// basic block.
func (block *BasicBlock) NewBr(target *BasicBlock) *TermBr {
	term := NewBr(target)
	block.Term = term
	return term
}

// --- [ conditional br ] ------------------------------------------------------

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target basic blocks.
func (block *BasicBlock) NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *TermCondBr {
	term := NewCondBr(cond, targetTrue, targetFalse)
	block.Term = term
	return term
}

// --- [ switch ] --------------------------------------------------------------

// NewSwitch returns a new switch terminator based on the given control
// variable, default target basic block and switch cases.
func (block *BasicBlock) NewSwitch(x value.Value, targetDefault *BasicBlock, cases ...*Case) *TermSwitch {
	term := NewSwitch(x, targetDefault, cases...)
	block.Term = term
	return term
}

// --- [ indirectbr ] ----------------------------------------------------------

// NewIndirectBr returns a new indirectbr terminator based on the given target
// address (derived from a blockaddress constant) and set of valid target basic
// blocks.
func (block *BasicBlock) NewIndirectBr(addr *BlockAddress, validTargets ...*BasicBlock) *TermIndirectBr {
	term := NewIndirectBr(addr, validTargets...)
	block.Term = term
	return term
}

// --- [ invoke ] --------------------------------------------------------------

// NewInvoke returns a new invoke terminator based on the given callee, function
// arguments and control flow return points for normal and exceptional
// execution.
func (block *BasicBlock) NewInvoke(callee value.Value, args []ll.Arg, normal, exception *BasicBlock) *TermInvoke {
	term := NewInvoke(callee, args, normal, exception)
	block.Term = term
	return term
}

// --- [ resume ] --------------------------------------------------------------

// NewResume returns a new resume terminator based on the given exception
// argument to propagate.
func (block *BasicBlock) NewResume(x value.Value) *TermResume {
	term := NewResume(x)
	block.Term = term
	return term
}

// --- [ catchswitch ] ---------------------------------------------------------

// NewCatchSwitch returns a new catchswitch terminator based on the given
// exception scope, exception handlers and unwind target.
func (block *BasicBlock) NewCatchSwitch(scope ll.ExceptionScope, handlers []*BasicBlock, unwindTarget ll.UnwindTarget) *TermCatchSwitch {
	term := NewCatchSwitch(scope, handlers, unwindTarget)
	block.Term = term
	return term
}

// --- [ catchret ] ------------------------------------------------------------

// NewCatchRet returns a new catchret terminator based on the given exit
// catchpad and target basic block.
func (block *BasicBlock) NewCatchRet(from *InstCatchPad, to *BasicBlock) *TermCatchRet {
	term := NewCatchRet(from, to)
	block.Term = term
	return term
}

// --- [ cleanupret ] ----------------------------------------------------------

// NewCleanupRet returns a new cleanupret terminator based on the given exit
// cleanuppad and unwind target.
func (block *BasicBlock) NewCleanupRet(from *InstCleanupPad, to ll.UnwindTarget) *TermCleanupRet {
	term := NewCleanupRet(from, to)
	block.Term = term
	return term
}

// --- [ unreachable ] ---------------------------------------------------------

// NewUnreachable returns a new unreachable terminator.
func (block *BasicBlock) NewUnreachable() *TermUnreachable {
	term := NewUnreachable()
	block.Term = term
	return term
}
