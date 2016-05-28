package gormgql

import (
	"github.com/graphql-go/graphql"
)

type GGQL struct {
	Title        string      // 标题
	Type         string      // 单条记录还是多条记录
	Table        string      // 关联的数据表名
	QueryFields  []GGQLField // 允许查询的字段
	ReturnFields []GGQLField // 查询时返回的字段
	Resolve      interface{} // 处理方法
}

type GGQLField struct {
	Field  string      // 字段名
	Format interface{} // 字段的格式化方法
}
