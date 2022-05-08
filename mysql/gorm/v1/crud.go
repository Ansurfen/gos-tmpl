package mysql

func (m *MysqlDB) Insert(value any) *MysqlDB {
	m.Create(value)
	return m
}

func (m *MysqlDB) Delete(value any) *MysqlDB {
	m.Delete(value)
	return m
}

func (m *MysqlDB) Update(value any) *MysqlDB {
	m.Save(value)
	return m
}

func (m *MysqlDB) Query(value any) *MysqlDB {
	m.First(value)
	return m
}
