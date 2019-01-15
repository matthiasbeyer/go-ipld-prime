package ipld

import (
	cid "github.com/ipfs/go-cid"
)

// NodeBuilder is an interface that describes creating new Node instances.
//
// The Node interface is entirely read-only methods; a Node is immutable.
// Thus, we need a NodeBuilder system for creating new ones; the builder
// is mutable, and when we're done accumulating mutations, we take the
// accumulated data and produce an immutable Node out of it.
//
// Separating mutation into NodeBuilder and keeping Node immutable makes
// it possible to perform caching (or rather, memoization, since there's no
// such thing as cache invalidation for immutable systems) of computed
// properties of Node; use copy-on-write algorithms for memory efficiency;
// and to generally build pleasant APIs.
//
// Each package in `go-ipld-prime//impl/*` that implements ipld.Node also
// has a NodeBuilder implementation that produces new nodes of that same
// package's type.
//
// Most Node implementations also have a method which returns a NodeBuilder
// that produces more nodes of their same concrete implementation type.
// This is useful for algorithms that work on trees of nodes: this NodeBuilder
// getter will be used when an update deep in the tree causes a need to
// create several new nodes to propagate the change up through parent nodes.
//
// The NodeBuilder retrieved from a Node can also be used to do *updates*:
// consider the AmendMap and AmendList methods.  These methods are useful
// not just for programmer convenience, but also because they can reuse memory,
// sharing any common segments of memory with the earlier Node.
// (In the NodeBuilder exposed by the `go-ipld-prime//impl/*` packages, these
// methods are equivalent to their Create* counterparts.  As there's no
// "existing" node for them to refer to, it's treated the same as amending
// an empty node.)
type NodeBuilder interface {
	CreateMap() MapBuilder
	AmendMap() MapBuilder
	CreateList() ListBuilder
	AmendList() ListBuilder
	CreateNull() Node
	CreateBool(bool) Node
	CreateInt() Node
	CreateFloat(float64) Node
	CreateString(string) Node
	CreateBytes([]byte) Node
	CreateLink(cid.Cid) Node
}

type MapBuilder interface {
	InsertAll(map[Node]Node) MapBuilder
	Insert(k, v Node) MapBuilder
	Delete(k Node) MapBuilder
	Build() Node
}

type ListBuilder interface {
	AppendAll([]Node) ListBuilder
	Append(v Node) ListBuilder
	Set(idx int, v Node) ListBuilder
	Build() Node
}

// future: add AppendIterator() methods (when we've implemented iterators!)

// future: add InsertConverting(map[string]interface{}) and similar methods.
//  (some open questions about how useful that is, given ipldbind should likely be more efficient, depending on use case.)

// future: define key ordering semantics during map insertion.
//  methods for re-ordering will probably be wanted someday.