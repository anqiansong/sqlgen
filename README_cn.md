# sqlgen

[English](README.md) | 中文

[![Go](https://github.com/anqiansong/sqlgen/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/anqiansong/sqlgen/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/anqiansong/sqlgen/branch/main/graph/badge.svg?token=8mLCFUqD2l)](https://codecov.io/gh/anqiansong/sqlgen)
[![Go Reference](https://pkg.go.dev/badge/github.com/anqiansong/sqlgen.svg)](https://pkg.go.dev/github.com/anqiansong/sqlgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/anqiansong/sqlgen)](https://goreportcard.com/report/github.com/anqiansong/sqlgen)
[![Release](https://img.shields.io/github/v/release/anqiansong/sqlgen.svg?style=flat-square)](https://github.com/anqiansong/sqlgen)
[![GitHub license](https://img.shields.io/github/license/anqiansong/sqlgen?style=flat-square)](https://github.com/anqiansong/sqlgen/blob/main/LICENSE)

sqlgen 是一个 SQL 代码生成工具，其支持 **bun**， **gorm**， **sql**， **sqlx**， **xorm** 的代码生成，灵感来自于：

- [go-zero](https://github.com/zeromicro/go-zero)
- [goctl](https://github.com/zeromicro/go-zero/tree/master/tools/goctl)
- [sqlc](https://github.com/kyleconroy/sqlc).


# 安装

```bash
go install github.com/anqiansong/sqlgen@latest
```

# 示例

 见 [example](https://github.com/anqiansong/sqlgen/tree/main/example)

#  SQL 查询编写规则
## 1. 函数名称
你可以通过在查询语句上方添加一个单行注释，用 `fn` 关键字来声明一个函数名称，例如：

```sql
-- fn: my_func
SELECT * FROM user;
```

其生成后代码格式为:

```go
func (m *UserModel) my_func (...) {
    ...
}
```

## 2. 查询一条记录
当你只想要查询一条记录的需求时，你必须明确地指定 `limit 1`，sqlgen 通过此表达式来判断当前查询是单记录查询还是多记录查询，例如：

```sql
-- fn: FindOne
select * from user where id = ? limit 1;
```

## 3. 使用 '?' 还是具体值？
在 SQL 查询语句的编写中，你可以用 `?` 来替代一个参数，也可以是具体值，他们最终都会被 sqlgen 转换成一个变量，下列示例中的两个查询是等价的。

> 注意: 此规则不适用于规则 2

```sql
-- fn: FindLimit
select * from user where id = ?;

-- fn: FindLimit
select * from user where id = 1;

```

## 4. SQL 内置函数支持
sqlgen 支持 SQL 内置的聚合函数查询，除此之外的其他函数暂不支持，聚合函数查询的列必须要用 `AS` 来起一个别名，例如：

```sql
-- fn: CountAll
select count(*) as count from user;
```

更多查询示例, 你可以点击 [example.sql](https://github.com/anqiansong/sqlgen/blob/main/example/example.sql) 查看详情.

#  sqlgen 使用步骤
1. 创建一个 SQL 文件
2. 编写 SQL 查询语句，如建表语句、查询语句等
3. 使用 `sqlgen` 工具，生成代码

# 注意
1. 目前只支持 MYSQL 代码生成
3. 不支持多表操作
4. 不支持联表查询