package runs

import (
	"fmt"
	"runtime"
	"strings"
)

// RuntimeFrame represents a program counter inside a runtime stack frame.
type RuntimeFrame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f RuntimeFrame) pc() uintptr {
	return uintptr(f) - 1
}

func (f RuntimeFrame) Func() *runtime.Func {
	return runtime.FuncForPC(f.pc())
}

func (f RuntimeFrame) FileLine() (string, int) {
	fn := f.Func()
	if fn == nil {
		return "<unknown>", 0
	}
	return fn.FileLine(f.pc())
}

func (f RuntimeFrame) Name() string {
	fn := f.Func()
	if fn == nil {
		return "<unknown>"
	}
	return fn.Name()
}

// Format formats the frame according to the fmt.Formatter interface.
//    %v    github.com/subchen/go-stack/runs/runtime_stack:106
//    %+v   gstack.RuntimeCallers() (github.com/subchen/go-stack/runs/runtime_stack:106)
func (f RuntimeFrame) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		fn := f.Func()
		if fn == nil {
			fmt.Fprint(s, "<unknown source>")
			return
		}

		file, line := fn.FileLine(f.pc())
		file, name := trimGOPATH(file, fn.Name())
		if s.Flag('+') {
			fmt.Fprintf(s, "%s() (%s:%d)", name, file, line)
		} else {
			fmt.Fprintf(s, "%s:%d", file, line)
		}
	}
}

// RuntimeStack represents a runtime stack of program counters.
type RuntimeStack []uintptr

// Format formats the runtime stack according to the fmt.Formatter interface.
//    %v    at github.com/subchen/go-stack/runs/runtime_stack:106
//          at example/main.go:9
//          at example/main.go:16
//    %+v   at gstack.RuntimeCallers() (github.com/subchen/go-stack/runs/runtime_stack:106)
//          at main.test() (example/main.go:9)
//          at main.main() (example/main.go:16)
func (stack RuntimeStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			for _, pc := range stack {
				fmt.Fprintf(s, "\tat %+v\n", RuntimeFrame(pc))
			}
		} else {
			for _, pc := range stack {
				fmt.Fprintf(s, "\tat %v\n", RuntimeFrame(pc))
			}
		}
	}
}

// Frames returns runtime stack frames
func (stack RuntimeStack) Frames() []RuntimeFrame {
	frames := make([]RuntimeFrame, len(stack))
	for i, pc := range stack {
		frames[i] = RuntimeFrame(pc)
	}
	return frames
}

// RuntimeCaller returns a runtime stack frame for caller
func RuntimeCaller(skip int) (RuntimeFrame, bool) {
	if pc, _, _, ok := runtime.Caller(skip); ok {
		return RuntimeFrame(pc), true
	} else {
		return RuntimeFrame(0), false
	}
}

// RuntimeCallers returns a runtime stack for callers
func RuntimeCallers(skip int) RuntimeStack {
	pcs := make([]uintptr, 32) // max depth is 32
	n := runtime.Callers(skip, pcs)
	return RuntimeStack(pcs[0:n])
}

func trimGOPATH(file, name string) (string, string) {
	// Here we want to get the source file path relative to the compile time
	// GOPATH. As of Go 1.6.x there is no direct way to know the compiled
	// GOPATH at runtime, but we can infer the number of path segments in the
	// GOPATH. We note that fn.Name() returns the function name qualified by
	// the import path, which does not include the GOPATH. Thus we can trim
	// segments from the beginning of the file path until the number of path
	// separators remaining is one more than the number of path separators in
	// the function name. For example, given:
	//
	//    GOPATH     /home/user
	//    file       /home/user/src/pkg/module/file.go
	//    fn.Name()  pkg/module.Type.Method
	//
	//    Example for fn.Name():
	//               github.com/subchen/app.requestNotFound
	//               github.com/subchen/app.(*Router).route
	//               github.com/subchen/app.(*Router).(github.com/subchen/app.route)-fm
	//
	// We want to produce:
	//
	//    file       pkg/sub/file.go
	//    fn.name    module.Type.Method
	//
	// From this we can easily see that fn.Name() has one less path separator
	// than our desired output. We count separators from the end of the file
	// path until it finds two more than in the function name and then move
	// one character forward to preserve the initial path segment without a
	// leading separator.

	path := name
	if pos := strings.Index(path, ".("); pos != -1 {
		path = path[0 : pos-1]
	}
	count := strings.Count(path, "/") + 2

	pairs := strings.Split(file, "/")
	size := len(pairs)
	if i := size - count; i > 0 {
		file = strings.Join(pairs[i:size], "/")
	}

	if pos := strings.LastIndex(path, "/"); pos != -1 {
		name = name[pos+1:]
	}
	return file, name
}
