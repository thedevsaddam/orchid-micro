package main

import (
	"app/migrations"
	"app/helpers/console"
	"app/helpers/general"
	"fmt"
)

func main() {
	fmt.Println("Commands:")
	fmt.Println("\t0. Exit command mode")
	fmt.Println("\t1. Migrate database")
	fmt.Println("\t2. Make example migration")
	switch general.Scanner() {
	case "0":
		console.ConsoleWarning("Exited from command mode!")
		return
	case "1":
		migrations.Migrate()
	case "2":
		fmt.Println("Another command")
	default:
		fmt.Println("Undefined command!")
	}
}
