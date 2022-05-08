package mysql

func (m *MysqlDB) Insert(value any) *MysqlDB {
	m.Insert(value)
	return m
}

func (m *MysqlDB) Delete(value any) *MysqlDB {
	m.Delete(value)
	return m
}

func (m *MysqlDB) Update(value any) *MysqlDB {
	m.Update(value)
	return m
}

func (m *MysqlDB) Query(value any) *MysqlDB {
	m.Find(value)
	return m
}
