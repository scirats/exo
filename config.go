package exo

import (
	"fmt"
	"strings"
)

type Property struct {
	Name string
	Value any
}

func property(props []*Property, name string) *Property {
	for _, p := range props {
		if p.Name == name {
			return p	
		}
	}

	panic(fmt.Sprintf("`%s` property is not defined", name))
}

type Config struct {
	Properties 	[]*Property
	Blocks 		[]*Block
}

func (c *Config) Has(name string) bool {
	return property(c.Properties, name) != nil || c.Block(name) != nil
}

func (c *Config) String(name string) string {
	p := c.StringList(name)
	return strings.Join(p, "")
}

func (c *Config) StringList(name string) []string {
	p := property(c.Properties, name)
	return p.Value.([]string)
}

func (c *Config) Block(name string) *Block {
	for _, b := range c.Blocks {
		if b.Name == name {
			return b
		}
	}

	return nil
}

type Block struct {
	Name 		string
	Properties 	[]*Property
	Blocks 		[]*Block
}

func (b *Block) Has(name string) bool {
	return property(b.Properties, name) != nil || b.Block(name) != nil
}

func (b *Block) String(name string) string {
	p := b.StringList(name)
	return strings.Join(p, "")
}

func (b *Block) StringList(name string) []string {
	p := property(b.Properties, name)
	return p.Value.([]string)
}

func (b *Block) Block(name string) *Block {
	for _, sb := range b.Blocks {
		if sb.Name == name {
			return sb
		}
	}

	return nil
}
