//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entoas.NewExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	opts := []entc.Option{
		entc.FeatureNames(
			"schema/snapshot",
		),
		entc.Extensions(ex),
	}

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
		},
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
