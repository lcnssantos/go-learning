package user

import (
	"database/sql"
	"main/shared"
)

func BuildUserModule(db *sql.DB) *shared.Module {
	return &shared.Module{Routers: BuildUserRouter(db)}
}
