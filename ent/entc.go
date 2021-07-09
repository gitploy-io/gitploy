// +build ignore

package main

import (
	"fmt"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// Code generation hook
// https://entgo.io/docs/code-gen/#code-generation-hooks
// https://github.com/ent/ent/tree/master/examples/entcpkg
func main() {
	err := entc.Generate("./schema", &gen.Config{
		Hooks: []gen.Hook{
			TagFields("json"),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// TagFields tags all fields defined in the schema with the given struct-tag.
// To remove omitempty for json tag.
func TagFields(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					field.StructTag = fmt.Sprintf("%s:%q", name, field.Name)
				}
			}
			return next.Generate(g)
		})
	}
}
