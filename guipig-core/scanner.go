package guipigCore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type RelationType uint

const (
	NoRelation RelationType = iota
	OneToOne
	OneToMany
	ManyToOne
	ManyToMany
)

type Relation struct {
	Column   string
	Table    string
	Type     string
	OnDelete string
	OnUpdate string
}

type Column struct {
	Name      string
	Relations []Relation
}

type Schema struct {
	Table   string
	Columns []Column
}

func ScanSchemas(schemasDir ...string) []Schema {
	var schemas []Schema
	dir := "schemas/"
	if len(schemasDir) > 0 {
		dir = schemasDir[0]
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		fileContents, err := ioutil.ReadFile(dir + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		var result Schema
		json.Unmarshal(fileContents, &result)
		// if result.Table != "" {
		// 	for _, column := range *result.Columns {
		// 		if column.Relations != nil {
		// 			fmt.Print(&column.Relations)
		// 			for _, relation := range *column.Relations {
		// 				rel := &relation
		// 				if len(relation.Type) > 0 {
		// 					relation.Type = "one-to-many"
		// 				}
		// 				// 	if len(column.Relations[i].OnDelete) > 0 {
		// 				// 		column.Relations[i].OnDelete = "cascade"
		// 				// 	}
		// 				// 	if len(column.Relations[i].OnUpdate) > 0 {
		// 				// 		column.Relations[i].OnUpdate = "cascade"
		// 				// 	}
		// 			}
		// 		}
		// 	}
		schemas = append(schemas, result)
		// }
	}
	return schemas
}
