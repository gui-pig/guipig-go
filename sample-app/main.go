package main

import (
	guipigService "github.com/vergieet/guipig-go/guipig-service"
	guipigSqlite "github.com/vergieet/guipig-go/guipig-sqlite"
)

func main() {
	guipigService.
		CreateGuiPigServer().
		AddRunBeforeScript(guipigSqlite.Migrator()).
		RegisterSchemas().
		RunServer()

	// fmt.Printf("\nIgloo House Door Type: %s\n", gps.Schemas)

	// guipigSqlite.Connect()
	// guipigCore.ReadJSON("schemas/unit.json")
	// guipigService.RunServer()
}
