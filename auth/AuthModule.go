package auth

import (
	"database/sql"
	"main/shared"
)

func BuildAuthModule(db *sql.DB) *shared.Module {
	return &shared.Module{Routers: BuildRouter(db)}
}
