package sql

import "fmt"

func (m *SQL) DiskSpace(info string) (interface{}, error) {
	if m.Driver == "postgres" {
		table := ""
		msg, query := cfg.GetDiskSpace(info)
		res, err := m.Session.Query(query)
		if err != nil {
			return table, err
		}
		for res.Next() {
			res.Scan(&table)
		}
		return fmt.Sprintf("%s, Disk Size: %s", msg, table), nil
	}
	return fmt.Sprintf("%s : Disk status not available", m.Driver), nil
}
