package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseInsert(stmt *ast.InsertStmt) (*spec.InsertStmt, error) {
	var text = stmt.Text()
	comment, err := parseLineComment(text)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	sql, err := NewSqlScanner(text).ScanAndTrim()
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	var ret spec.InsertStmt
	ret.Comment = comment
	tableName, err := parseTableRefsClause(stmt.Table)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	columns, err := parseColumns(stmt.Columns, tableName)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	ret.Table = tableName
	ret.Action = spec.ActionCreate
	ret.SQL = sql
	ret.Columns = columns

	return &ret, nil
}
