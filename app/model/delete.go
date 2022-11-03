package model

func DeleteSQL(dataSource string, conditions map[string]string) error {
	query := "delete from "
	query += dataSource
	query += " where "
	for key, value := range conditions {
		query += key + "=" + value
		query += " and"
	}
	query = query[:len(query)-4]
	if _, err := Db.Exec(query); err != nil {
		return err
	}
	return nil
}
