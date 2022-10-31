package guipigSqlite

import (
	"fmt"

	guipigCore "github.com/vergieet/guipig-go/guipig-core"
)

func convertDataType(){
	
}

func Migrator() guipigCore.Runnable {
	return &guipigCore.Migrator{
		DbType: "sqlite",
		Migrate: func(schemas []guipigCore.Schema) {
			var columnsString []string;
			for _, schema := range schemas {
				for _, column := range schema.Columns {
					fmt.Printf("\t asdf %s\n", column)
					isDefaultValue := false
					if(column.DefaultValue != "") {
						isDefaultValue = true
					}	
					append(columnsString, " " + c.getName() + " " + convertDataType(c.getType(), c.getLength()) + " " +
					(c.getNullable() ? "":(null != c.getDefaultValue() && c.getDefaultValue().equalsIgnoreCase(":null")?"":"NOT NULL")) + " " +
					((null == c.getDefaultValue())?"":convertDefaultValue(c.getDefaultValue())) + " " +
					(c.getAutoIncrement()?"AUTO_INCREMENT":"");
				}
			}

		},
	}

}
