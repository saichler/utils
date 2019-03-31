package common

type Query struct {
	ormQuery string
	tableName string
}

func NewQuery(ormQuery string) *Query {
	q:=&Query{}
	q.ormQuery = ormQuery
	q.parse()
	return q
}

func (q *Query) parse(){
	q.tableName = "Node"
}

func (q *Query) TableName() string {
	return q.tableName
}