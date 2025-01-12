package typegen

import (
	"io"
	"text/template"

	declaration "github.com/ipld/go-ipld-prime/typed/declaration"
	wish "github.com/warpfork/go-wish"
)

type generateKindString struct {
	Name declaration.TypeName
	Type declaration.Type
}

func (gk generateKindString) EmitNodeMethodTraverseField(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) TraverseField(key string) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{ /* todo more content */ }
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodTraverseIndex(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) TraverseIndex(idx int) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{ /* todo more content */ }
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodMapIterator(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) MapIterator() ipld.MapIterator {
			return mapIteratorReject{ipld.ErrWrongKind{ /* todo more content */ }}
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodListIterator(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) ListIterator() ipld.ListIterator {
			return listIteratorReject{ipld.ErrWrongKind{ /* todo more content */ }}
		}
	`))).Execute(w, gk) // REVIEW: maybe that rejection thunk should be in main package?  don't really want to flash it at folks though.  very impl detail.
}

func (gk generateKindString) EmitNodeMethodLength(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) Length() int {
			return -1
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodIsNull(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) IsNull() bool {
			return false
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodAsBool(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) AsBool() (bool, error) {
			return false, ipld.ErrWrongKind{ /* todo more content */ }
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodAsInt(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) AsInt() (int, error) {
			return 0, ipld.ErrWrongKind{ /* todo more content */ }
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodAsFloat(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) AsFloat() (float64, error) {
			return 0, ipld.ErrWrongKind{ /* todo more content */ }
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodAsString(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func (x {{ .Name }}) AsString() (string, error) {
			return x.x, nil
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodAsBytes(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) AsBytes() ([]byte, error) {
			return nil, ipld.ErrWrongKind{ /* todo more content */ }
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodAsLink(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) AsLink() (ipld.Link, error) {
			return nil, ipld.ErrWrongKind{ /* todo more content */ }
		}
	`))).Execute(w, gk)
}

func (gk generateKindString) EmitNodeMethodNodeBuilder(w io.Writer) {
	template.Must(template.New("").Parse("\n"+wish.Dedent(`
		func ({{ .Name }}) NodeBuilder() ipld.NodeBuilder {
			return {{ .Name }}__NodeBuilder{}
		}
	`))).Execute(w, gk)
}
