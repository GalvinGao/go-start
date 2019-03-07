package main

import (
	"fmt"
	"os"
)

// SQLCredential holds a SQL Credential
type SQLCredential struct {
	username string
	password string
	protocol string
	location string
	database string
}

func (sql *SQLCredential) render() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", sql.username, sql.password, sql.protocol, sql.location, sql.database)
}

// SQLHandler handles sub-command `sql`
func SQLHandler(args interface{}) {
	if args == nil {
		fmt.Println("Pretending I am using the variable so go compiler will not go crazy.")
	}
	panic("Not Implemented yet.")
}

// GoHandler handles sub-command `go`
func GoHandler(args interface{}) {
	if args == nil {
		fmt.Println("Pretending I am using the variable so go compiler will not go crazy.")
	}
	panic("Not Implemented yet.")
}

// DebugHandler handles sub-command `debug`
func DebugHandler(args interface{}) {
	if args == nil {
		fmt.Println("Pretending I am using the variable so go compiler will not go crazy.")
	}
	panic("Not Implemented yet.")
}

// PresetHandler handles sub-command `preset`
func PresetHandler(args interface{}) {
	if args == nil {
		fmt.Println("Pretending I am using the variable so go compiler will not go crazy.")
	}
	panic("Not Implemented yet.")
}

func main() {
	args := os.Args[1:]
	subCommand := os.Args[0]
	switch subCommand {
	case "sql":
		SQLHandler(args)
	case "go":
		GoHandler(args)
	case "debug":
		DebugHandler(args)
	case "preset":
		PresetHandler(args)
	default:
		fmt.Println("Hopefully there's some help message here in the future. But not now my friend.")
	}
}
