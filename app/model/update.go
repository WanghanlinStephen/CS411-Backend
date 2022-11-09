package model

func UpdateSQL(dataSource string, updates map[string]string, conditions map[string]string) error {
	query := "update "
	query += dataSource
	query += " set "
	for key, value := range updates {
		query += key + "=" + "'" + value + "'"
		query += " ,"
	}
	query = query[:len(query)-2]

	if len(conditions) != 0 {
		query += " where "
		for key, value := range conditions {
			query += key + "=" + value
			query += " and"
		}
		query = query[:len(query)-4]
	}

	if _, err := Db.Exec(query); err != nil {
		return err
	}
	return nil
}
