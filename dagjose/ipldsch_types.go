package dagjose

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	"github.com/ipld/go-ipld-prime/datamodel"
)

var _ datamodel.Node = nil // suppress errors when this dependency is not referenced
// Type is a struct embeding a NodePrototype/Type for every Node implementation in this package.
// One of its major uses is to start the construction of a value.
// You can use it like this:
//
// 		dagjose.Type.YourTypeName.NewBuilder().BeginMap() //...
//
// and:
//
// 		dagjose.Type.OtherTypeName.NewBuilder().AssignString("x") // ...
//
var Type typeSlab

type typeSlab struct {
	Any              _Any__Prototype
	Any__Repr        _Any__ReprPrototype
	Bool             _Bool__Prototype
	Bool__Repr       _Bool__ReprPrototype
	Bytes            _Bytes__Prototype
	Bytes__Repr      _Bytes__ReprPrototype
	Float            _Float__Prototype
	Float__Repr      _Float__ReprPrototype
	Int              _Int__Prototype
	Int__Repr        _Int__ReprPrototype
	JOSE             _JOSE__Prototype
	JOSE__Repr       _JOSE__ReprPrototype
	List             _List__Prototype
	List__Repr       _List__ReprPrototype
	Map              _Map__Prototype
	Map__Repr        _Map__ReprPrototype
	Recipient        _Recipient__Prototype
	Recipient__Repr  _Recipient__ReprPrototype
	Recipients       _Recipients__Prototype
	Recipients__Repr _Recipients__ReprPrototype
	Signature        _Signature__Prototype
	Signature__Repr  _Signature__ReprPrototype
	Signatures       _Signatures__Prototype
	Signatures__Repr _Signatures__ReprPrototype
	String           _String__Prototype
	String__Repr     _String__ReprPrototype
}

// --- type definitions follow ---

// Any matches the IPLD Schema type "Any".
// Any has union typekind, which means its data model behaviors are that of a map kind.
type Any = *_Any
type _Any struct {
	x _Any__iface
}
type _Any__iface interface {
	_Any__member()
}

func (_Bool) _Any__member()   {}
func (_String) _Any__member() {}
func (_Bytes) _Any__member()  {}
func (_Int) _Any__member()    {}
func (_Float) _Any__member()  {}
func (_Map) _Any__member()    {}
func (_List) _Any__member()   {}

// Bool matches the IPLD Schema type "Bool".  It has bool kind.
type Bool = *_Bool
type _Bool struct{ x bool }

// Bytes matches the IPLD Schema type "Bytes".  It has bytes kind.
type Bytes = *_Bytes
type _Bytes struct{ x []byte }

// Float matches the IPLD Schema type "Float".  It has float kind.
type Float = *_Float
type _Float struct{ x float64 }

// Int matches the IPLD Schema type "Int".  It has int kind.
type Int = *_Int
type _Int struct{ x int64 }

// JOSE matches the IPLD Schema type "JOSE".  It has struct type-kind, and may be interrogated like map kind.
type JOSE = *_JOSE
type _JOSE struct {
	aad         _String__Maybe
	ciphertext  _String__Maybe
	iv          _String__Maybe
	payload     _String__Maybe
	protected   _String__Maybe
	recipients  _Recipients__Maybe
	signatures  _Signatures__Maybe
	tag         _String__Maybe
	unprotected _Map__Maybe
}

// List matches the IPLD Schema type "List".  It has list kind.
type List = *_List
type _List struct {
	x []_Any
}

// Map matches the IPLD Schema type "Map".  It has map kind.
type Map = *_Map
type _Map struct {
	m map[_String]*_Any
	t []_Map__entry
}
type _Map__entry struct {
	k _String
	v _Any
}

// Recipient matches the IPLD Schema type "Recipient".  It has struct type-kind, and may be interrogated like map kind.
type Recipient = *_Recipient
type _Recipient struct {
	header        _Map__Maybe
	encrypted_key _String__Maybe
}

// Recipients matches the IPLD Schema type "Recipients".  It has list kind.
type Recipients = *_Recipients
type _Recipients struct {
	x []_Recipient
}

// Signature matches the IPLD Schema type "Signature".  It has struct type-kind, and may be interrogated like map kind.
type Signature = *_Signature
type _Signature struct {
	header    _Map__Maybe
	protected _String__Maybe
	signature _String
}

// Signatures matches the IPLD Schema type "Signatures".  It has list kind.
type Signatures = *_Signatures
type _Signatures struct {
	x []_Signature
}

// String matches the IPLD Schema type "String".  It has string kind.
type String = *_String
type _String struct{ x string }
