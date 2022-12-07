package repository

import (
	"database/sql"

	"github.com/petruspierre/cartola-consolidation/internal/infra/db"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}
