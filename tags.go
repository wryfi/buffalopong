package buffalopong

import (
	"github.com/flosch/pongo2"
	"bytes"
)

type rootPath struct {
	path	string
}

func (rootPath *rootPath) Execute(context *pongo2.ExecutionContext, buffer *bytes.Buffer) *pongo2.Error {
	return nil
}

func rootPathParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	rootPath := &rootPath{path: "/"}
	return rootPath, nil
}
