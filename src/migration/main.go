package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"booking-calendar-server-backend/internal/modules/store/models"
	"fmt"
	"io"
	"os"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(
		&models.Store{},
	)

	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error loading migration statements: %s\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error writing migration statements: %s\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

}