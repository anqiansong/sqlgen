// Code generated by sqlgen. DO NOT EDIT!

package model

import (
    "context"
    "fmt"
    "time"

    "xorm.io/xorm"

    "github.com/shopspring/decimal"
)

// {{UpperCamel $.Table.Name}}Model represents a {{$.Table.Name}} model.
type {{UpperCamel $.Table.Name}}Model struct {
    engine xorm.EngineInterface
}

// {{UpperCamel $.Table.Name}} represents a {{$.Table.Name}} struct data.
type {{UpperCamel $.Table.Name}} struct { {{range $.Table.Columns}}
{{UpperCamel .Name}} {{.GoType}} `xorm:"{{if IsPrimary .Name}}pk {{end}}{{if .AutoIncrement}}autoincr {{end}}'{{.Name}}'" json:"{{LowerCamel .Name}}"`{{if .HasComment}}// {{TrimNewLine .Comment}}{{end}}{{end}}
}

{{range $stmt := .SelectStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Having.IsValid}}{{$stmt.Having.ParameterStructure "Having"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}

{{if IsExtraResult $stmt.ReceiverName}}
{{$stmt.ReceiverStructure "xorm"}}

// TableName returns the table name. it implemented by gorm.Tabler.
func ({{$stmt.ReceiverName}}) TableName() string {
    return "{{$.Table.Name}}"
}{{end}}
{{end}}

{{range $stmt := .UpdateStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}

{{range $stmt := .DeleteStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}

func ({{UpperCamel $.Table.Name}}) TableName() string{
    return "{{$.Table.Name}}"
}

// New{{UpperCamel $.Table.Name}}Model returns a new {{$.Table.Name}} model.
func New{{UpperCamel $.Table.Name}}Model (engine xorm.EngineInterface) *{{UpperCamel $.Table.Name}}Model {
    return &{{UpperCamel $.Table.Name}}Model{engine: engine}
}

// Create creates  {{$.Table.Name}} data.
func (m *{{UpperCamel $.Table.Name}}Model) Create(ctx context.Context, data ...*{{UpperCamel $.Table.Name}}) error {
    if len(data)==0{
        return fmt.Errorf("data is empty")
    }

    var session = m.engine.Context(ctx)
    var list []interface{}
    for _, v := range data {
        list = append(list, v)
    }

    _,err := session.Insert(list...)
    return err
}
{{range $stmt := .SelectStmt}}
// {{.FuncName}} is generated from sql:
// {{LineComment $stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Having.IsValid}}, having {{$stmt.Having.ParameterStructureName "Having"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}})({{if $stmt.Limit.One}}*{{$stmt.ReceiverName}}, {{else}}[]*{{$stmt.ReceiverName}}, {{end}} error){
    var result {{if $stmt.Limit.One}} = new({{$stmt.ReceiverName}}){{else}}[]*{{$stmt.ReceiverName}}{{end}}
    var session = m.engine.Context(ctx)
    session.Select(`{{$stmt.SelectSQL}}`)
    {{if $stmt.Where.IsValid}}session.Where({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}})
    {{end }}{{if $stmt.GroupBy.IsValid}}session.GroupBy({{$stmt.GroupBy.SQL}})
    {{end}}{{if $stmt.Having.IsValid}}session.Having(fmt.Sprintf({{HavingSprintf $stmt.Having.SQL}}, {{$stmt.Having.Parameters "having"}}))
    {{end}}{{if $stmt.OrderBy.IsValid}}session.OrderBy({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}session.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}}{{if gt $stmt.Limit.Offset 0}}, {{$stmt.Limit.OffsetParameter "limit"}}{{end}})
    {{end}}{{if $stmt.Limit.One}}has, err := session.Get(result)
    if !has{
        return nil, sql.ErrNoRows
    }
    {{else}}err :=session.Find(&result){{end}}
    return result, err
}
{{end}}

{{range $stmt := .UpdateStmt}}
// {{.FuncName}} is generated from sql:
// {{LineComment $stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context, data *{{UpperCamel $.Table.Name}}{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}}) error {
    var session = m.engine.Context(ctx)
    session.Table(&{{UpperCamel $.Table.Name}}{})
    {{if $stmt.Where.IsValid}}session.Where({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}})
    {{end}}{{if $stmt.OrderBy.IsValid}}session.OrderBy({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}session.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}}{{if gt $stmt.Limit.Offset 0}}, {{$stmt.Limit.OffsetParameter "limit"}}{{end}})
    {{end}}_, err := session.Update(map[string]interface{}{
        {{range $name := $stmt.Columns}}"{{$name}}": data.{{UpperCamel $name}},
        {{end}}
    })
    return err
}
{{end}}

{{range $stmt := .DeleteStmt}}
// {{.FuncName}} is generated from sql:
// {{LineComment $stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}}) error {
    var session = m.engine.Context(ctx)
    {{if $stmt.Where.IsValid}}session.Where({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}})
    {{end}}{{if $stmt.OrderBy.IsValid}}session.OrderBy({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}session.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}}{{if gt $stmt.Limit.Offset 0}}, {{$stmt.Limit.OffsetParameter "limit"}}{{end}})
    {{end}}_, err := session.Delete(&{{UpperCamel $.Table.Name}}{})
    return err
}
{{end}}
