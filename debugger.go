package gojdb

type JavaDebugger interface {
	LineBreakpoint()
	MethodBreakpoint()
	ExceptionBreakpoint()
	Status()
}

type JaveBreakpoint interface {
	Backtrace()
	Continue() error
	Step()
	Next()
}
