// generated by sqlgen, do not edit.

package model

import (
    "time"
)

// {{UpperCamel .Name}} represents a {{.Name}} struct data.
type {{UpperCamel .Name}} struct { {{range .Columns}}
    {{UpperCamel .Name}} {{.Go}} `gorm:"{{if IsPrimary .Name}}primaryKey;{{end}}column:{{.Name}}" json:"{{LowerCamel .Name}}"`{{end}}
}

func ({{UpperCamel .Name}}) TableName() string {
    return "{{.Name}}"
}

{{range .InsertStmt}}// sql: {{. SQL}}}
func (m *{{UpperCamel .Name}}Model)Insert
{{end}}
