package constants

import "database/sql"

func Int32ToSQLNullInt32(s int32) sql.NullInt32 {
	if s != 0 {
		return sql.NullInt32{
			Int32: s,
			Valid: true,
		}
	}
	return sql.NullInt32{}
}