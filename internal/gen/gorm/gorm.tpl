// generated by sqlgen, do not edit.

package model

import (
    "context"
    )

// {{UpperCamel $.Table.Name}} represents a {{$.Table.Name}} struct data.
type {{UpperCamel $.Table.Name}} struct { {{range $.Table.Columns}}
    {{UpperCamel .Name}} {{.Go}} `gorm:"{{if IsPrimary .Name}}primaryKey;{{end}}column:{{.Name}}" json:"{{LowerCamel .Name}}"`{{end}}
}

func ({{UpperCamel $.Table.Name}}) TableName() string {
    return "{{$.Table.Name}}"
}