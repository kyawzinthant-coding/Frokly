package nullable

import "database/sql"

func StringPtr(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func String(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}
