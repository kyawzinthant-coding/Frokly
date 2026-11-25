package vendor

import (
	dbgen "github.com/kyawzinthant-coding/froky/db/generated"
)

type Service struct {
	q *dbgen.Queries
}

func NewService(q *dbgen.Queries) *Service {
	return &Service{q: q}
}
