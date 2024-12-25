package user

import (
	"database/sql"

	"github.com/stuckinforloop/shit/deps/timeutils"
	"github.com/stuckinforloop/shit/deps/ulid"
)

type DAO struct {
	db      *sql.DB
	timeNow timeutils.TimeNow
	ulid    *ulid.Source
}

func NewDAO(db *sql.DB, timeNow timeutils.TimeNow, ulid *ulid.Source) *DAO {
	dao := &DAO{
		db,
		timeNow,
		ulid,
	}

	return dao
}
