package guipigService

import (
	"fmt"
	"log"
	"net/http"

	guipigCore "github.com/vergieet/guipig-go/guipig-core"
)

type GuiPigServer struct {
	Schemas     []guipigCore.Schema
	PreScripts  []guipigCore.Runnable
	PostScripts []guipigCore.Runnable
}

func CreateGuiPigServer() *GuiPigServer {
	return &GuiPigServer{}
}

func (b *GuiPigServer) RegisterSchemas() *GuiPigServer {
	b.Schemas = guipigCore.ScanSchemas()
	return b
}

func (b *GuiPigServer) AddRunBeforeScript(runnable guipigCore.Runnable) *GuiPigServer {
	b.PreScripts = append(b.PreScripts, runnable)
	return b
}

func (b *GuiPigServer) AddRunAfterScript(runnable guipigCore.Runnable) *GuiPigServer {
	b.PostScripts = append(b.PostScripts, runnable)
	return b
}

func (b *GuiPigServer) GetServer() GuiPigServer {
	return GuiPigServer{
		Schemas: b.Schemas,
	}
}

func (b *GuiPigServer) RunServer() {

	for _, runnable := range b.PreScripts {
		runnable.Run(b.Schemas)
	}

	for _, schema := range b.Schemas {
		http.HandleFunc("/"+schema.Table, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Heallo, %q", schema)
		})
	}

	log.Fatal(http.ListenAndServe(":8081", nil))

}
