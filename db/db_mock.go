package db

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

const (
	// create the database

	CREATE = "CREATE"

	// drop the database
	DROP = "DROP"
)

func DbMock(operation string) {
	if UNATTENDED_DEPLOY := os.Getenv("UNATTENDED_DEPLOY"); UNATTENDED_DEPLOY == "" {

		_, fileName, _, _ := runtime.Caller(0)

		currentPath := filepath.Dir(fileName)

		// by default we point to the tools/env/.env file
		err := godotenv.Load(currentPath + "/../tools/env/.env")

		if err != nil {
			log.Fatalf("Error loading env.\n%v", err)
		}

		// get values from env vars on the host.

		pg_name := os.Getenv("POSTGRES_DB")
		pg_user := os.Getenv("POSTGRES_USER")
		pg_password := os.Getenv("POSTGRES_PASSWORD")
		pg_host := os.Getenv("POSTGRES_ADDRESS")

		var task string

		if operation == "CREATE" {
			task = "createdb"
		} else {
			task = "dropdb"
		}

		cmd := exec.Command(task, "-h", pg_host, "-U", pg_user, "-e", pg_name)
		cmd.Env = os.Environ()

		cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%v", pg_password))

		if err := cmd.Run(); err != nil {
			log.Fatalf("Error executing %v on %v. \n%v", task, pg_name, err)
		}

		// **
		// at the beginning of each test that uses this DbMock() function we will need:
		// DbMock(CREATE)
		// defer DbMock(DROP)
	}
}
