package repository

import (
	"database/sql"
	"main/user/dto"
	"main/user/entities"
)

type UserRepository struct {
	database *sql.DB
}

func (this *UserRepository) Create(data dto.CreateUserDto) error {
	prepare, err := this.database.Prepare("INSERT INTO users (name, email, password ) values ($1, $2, $3)")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(data.Name, data.Email, data.Password)

	if err != nil {
		return err
	}

	return nil
}

func (this *UserRepository) FindOneByEmail(email string) (*entities.User, error) {
	prepare, err := this.database.Prepare("SELECT * FROM users WHERE email = $1 LIMIT 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(email))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this *UserRepository) scanEntity(r *sql.Row) (*entities.User, error) {
	var user = entities.User{}
	err := r.Scan(&user.Id, &user.Name, &user.Email, &user.EmailConfirmed, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}
