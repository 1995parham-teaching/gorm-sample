package main

import (
	"io"
	"log"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/1995parham-teaching/gorm-sample/internal/domain/model"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(new(model.User))
	if err != nil {
		log.Fatalf("failed to load gorm schema: %v", err)
	}

	_, _ = io.WriteString(os.Stdout, stmts)
}
