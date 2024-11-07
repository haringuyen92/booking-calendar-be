package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	course_models "booking-calendar-server-backend/internal/modules/course/models"
	staff_models "booking-calendar-server-backend/internal/modules/staff/models"
	"booking-calendar-server-backend/internal/modules/store/models"
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	"fmt"
	"io"
	"os"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(
		&models.Store{},
		&staff_models.Staff{},
		&course_models.Course{},
		&user_models.User{},
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
