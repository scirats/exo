package main

import (
	"os"
)

type Type int

const (
	TypeString = iota
	TypeStringList
)

const (
	rootBlock = ""
)

func ParseFile(file string) (*Config, error) {
	d, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return Parse(string(d))
}

func Parse(s string) (*Config, error) {
	t := NewTokenizer(s)

	b, err := parseBlock(t, rootBlock)
	if err != nil {
		return nil, err
	}

	return &Config{b.Properties, b.Blocks}, nil
}

func parseBlock(t *Tokenizer, name string) (*Block, error) {
	var props 	[]*Property
	var blocks 	[]*Block
	var closed 	bool = name == rootBlock

	for t.HasNext() {
		n, err := t.Next()
		if err != nil {
			return nil, newError(t.Line(), err.Error())
		}

		if name != rootBlock && n.Name == NameBlockEnd {
			closed = true
			break
		}

		if n.Name != NameIdent {
			return nil, newError(t.Line(), "identifier token expected")
		}

		op, err := t.Next()
		if err != nil {
			return nil, newError(t.Line(), err.Error())
		}

		switch op.Name {
			case NameEq:
				if !t.HasNext() {
					return nil, newError(t.Line(), "value expected")
				}

				v, err := t.Next()
				if err != nil {
					return nil, newError(t.Line(), err.Error())
				}

				var lst []string
				lst = append(lst, v.Value)
				
				for {
					if !t.HasNext() {
						break
					}

					tk, err := t.Next()
					if err != nil {
						return nil, newError(t.Line(), err.Error())
					}

					if tk.Name != NameComma {
						t.Unread()
						break
					}

					if !t.HasNext() {
						return nil, newError(t.Line(), "unexpected EOF")
					}

					tk, err = t.Next()
					if err != nil {
						return nil, newError(t.Line(), err.Error())
					}

					if tk.Name != NameString {
						return nil, newError(t.Line(), "string value expected")
					}

					lst = append(lst, tk.Value)
				}

				props = append(props, &Property{
					Name:	n.Value,
					Value:	lst,
				})
			case NameBlockStart:
				b, err := parseBlock(t, n.Value)
				if err != nil {
					return nil, err
				}

				blocks = append(blocks, b)
			default:
				return nil, newError(t.Line(), "`=` or `{` expected")
		}
	}

	if !closed {
		return nil, newError(t.Line(), "`}` expected")
	}

	return &Block{Name: name, Properties: props, Blocks: blocks}, nil
}

