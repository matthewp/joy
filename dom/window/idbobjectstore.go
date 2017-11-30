package window

import (
	"github.com/matthewmueller/golly/dom/domstringlist"
	"github.com/matthewmueller/golly/dom/idbcursordirection"
	"github.com/matthewmueller/golly/dom/idbindexparameters"
	"github.com/matthewmueller/golly/js"
)

// IDBObjectStore struct
// js:"IDBObjectStore,omit"
type IDBObjectStore struct {
}

// Add fn
// js:"add"
func (*IDBObjectStore) Add(value interface{}, key *interface{}) (i IDBRequest) {
	js.Rewrite("$_.add($1, $2)", value, key)
	return i
}

// Clear fn
// js:"clear"
func (*IDBObjectStore) Clear() (i IDBRequest) {
	js.Rewrite("$_.clear()")
	return i
}

// Count fn
// js:"count"
func (*IDBObjectStore) Count(key *interface{}) (i IDBRequest) {
	js.Rewrite("$_.count($1)", key)
	return i
}

// CreateIndex fn
// js:"createIndex"
func (*IDBObjectStore) CreateIndex(name string, keyPath string, optionalParameters *idbindexparameters.IDBIndexParameters) (i *IDBIndex) {
	js.Rewrite("$_.createIndex($1, $2, $3)", name, keyPath, optionalParameters)
	return i
}

// Delete fn
// js:"delete"
func (*IDBObjectStore) Delete(key interface{}) (i IDBRequest) {
	js.Rewrite("$_.delete($1)", key)
	return i
}

// DeleteIndex fn
// js:"deleteIndex"
func (*IDBObjectStore) DeleteIndex(indexName string) {
	js.Rewrite("$_.deleteIndex($1)", indexName)
}

// Get fn
// js:"get"
func (*IDBObjectStore) Get(key interface{}) (i IDBRequest) {
	js.Rewrite("$_.get($1)", key)
	return i
}

// Index fn
// js:"index"
func (*IDBObjectStore) Index(name string) (i *IDBIndex) {
	js.Rewrite("$_.index($1)", name)
	return i
}

// OpenCursor fn
// js:"openCursor"
func (*IDBObjectStore) OpenCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	js.Rewrite("$_.openCursor($1, $2)", rng, direction)
	return i
}

// Put fn
// js:"put"
func (*IDBObjectStore) Put(value interface{}, key *interface{}) (i IDBRequest) {
	js.Rewrite("$_.put($1, $2)", value, key)
	return i
}

// IndexNames prop
// js:"indexNames"
func (*IDBObjectStore) IndexNames() (indexNames *domstringlist.DOMStringList) {
	js.Rewrite("$_.indexNames")
	return indexNames
}

// KeyPath prop
// js:"keyPath"
func (*IDBObjectStore) KeyPath() (keyPath string) {
	js.Rewrite("$_.keyPath")
	return keyPath
}

// Name prop
// js:"name"
func (*IDBObjectStore) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// Transaction prop
// js:"transaction"
func (*IDBObjectStore) Transaction() (transaction *IDBTransaction) {
	js.Rewrite("$_.transaction")
	return transaction
}