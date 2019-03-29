package transaction

type Transaction struct {
	tableData map[string][]*Record
}

func (tx *Transaction) AddRecord(tableName string) *Record {
	if tx.tableData == nil {
		tx.tableData = make(map[string][]*Record)
	}
	if tx.tableData[tableName]==nil {
		tx.tableData[tableName]=make([]*Record,0)
	}
	record:=&Record{}
	tx.tableData[tableName] = append(tx.tableData[tableName],record)
	return record
}

func (tx *Transaction) Records() map[string][]*Record {
	return tx.tableData
}