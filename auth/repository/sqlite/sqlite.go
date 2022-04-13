package sqlite

import (
	"database/sql"
	"rest/models"
)

const userTable = `CREATE TABLE user(
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	first_name TEXT,
	last_name TEXT
  );`

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB, table string) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(user *models.User) (int, error) {
	result, err := ur.DB.Exec("insert into user (first_name, last_name) values ($1, $2)",
		user.FirstName, user.LastName)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}
func (ur *UserRepository) GetUser(id int) (*models.User, error) {
	user := models.User{}
	row := ur.DB.QueryRow("select * from user where id = $1", id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)

	return &user, err
}
func (ur *UserRepository) PutUser(user *models.User) error {
	_, err := ur.DB.Exec("update user set first_name = $1, last_name = $2 where id = $3", user.FirstName, user.LastName, user.ID)
	return err
}
func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.DB.Exec("delete from user where id = $1", id)
	return err
}
