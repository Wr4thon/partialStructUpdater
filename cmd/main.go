package main

import (
	"os"
	"strings"

	"github.com/Wr4thon/partialStructUpdater/pkg/generation"
	"github.com/Wr4thon/partialStructUpdater/pkg/jobs"
	"github.com/Wr4thon/partialStructUpdater/pkg/loader"
	// _ "github.com/Wr4thon/partialStructUpdater/testData"
)

func main() {
	typeName := os.Args[1]
	typeLoader := loader.NewBinaryTypeloader()

	dir := "/home/johannes/go/src/github.com/Wr4thon/partialStructUpdater/testData"

	cleanup := jobs.NewCleanupJob(dir)

	generator, err := generation.NewTemplateGenerator()
	if err != nil {
		panic(err)
	}

	writer := generation.NewFileWriter(dir)

	t, err := typeLoader.LoadByTypename(typeName)
	if err != nil {
		panic(err)
	}

	builder := &strings.Builder{}
	if err = generator.GenerateFile(t, builder); err != nil {
		panic(err)
	}

	err = writer.WriteFile(t.TypeName(), builder.String())
	if err != nil {
		panic(err)
	}

	if err := cleanup.Run(); err != nil {
		panic(err)
	}
}
