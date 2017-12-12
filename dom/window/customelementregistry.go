package window

import "github.com/matthewmueller/joy/macro"

// CustomElementRegistry struct
// js:"CustomElementRegistry,omit"
type CustomElementRegistry struct {
}

// Define fn
// js:"define"
func (*CustomElementRegistry) Define(name string, ctr interface{}) {
  macro.Rewrite("$_.define($1, $2)", name, ctr)
}
