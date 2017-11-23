package window

import (
	"github.com/matthewmueller/golly/dom2/childnode"
	"github.com/matthewmueller/golly/js"
)

// js:"DocumentType,omit"
type DocumentType struct {
	Node
	childnode.ChildNode
}

// Entities
func (*DocumentType) Entities() (entities *NamedNodeMap) {
	js.Rewrite("$<.Entities")
	return entities
}

// InternalSubset
func (*DocumentType) InternalSubset() (internalSubset string) {
	js.Rewrite("$<.InternalSubset")
	return internalSubset
}

// Name
func (*DocumentType) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Notations
func (*DocumentType) Notations() (notations *NamedNodeMap) {
	js.Rewrite("$<.Notations")
	return notations
}

// PublicID
func (*DocumentType) PublicID() (publicId string) {
	js.Rewrite("$<.PublicID")
	return publicId
}

// SystemID
func (*DocumentType) SystemID() (systemId string) {
	js.Rewrite("$<.SystemID")
	return systemId
}
