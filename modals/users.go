package modals

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	// in minute
	EXPIRE_DURATION = time.Minute * 1
)

type user struct {
	db       *sqlx.DB
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password string    `json:"password" db:"password"`
	Username string    `json:"username" db:"username"`
	Token    string    `json:"-" db:"token"`
	ExpireAt time.Time `json:"-" db:"expire_at"`
}

func NewUser() (*user, error) {
	db, err := SetUpDb()
	if err != nil {
		return nil, err
	}

	return &user{
		db: db,
	}, nil
}

func (u *user) ValidateUserForAuth() error {
	if u.Username == "" || u.Password == "" {
		return errors.New("username and password must not be null")
	}
	return nil
}

func (u *user) CreateUser() error {
	defer u.db.Close()
	// get user
	err := u.getUserByUsername()
	if err == nil {
		return errors.New("user already exist")
	}

	// encode the password
	password := base64.RawStdEncoding.EncodeToString([]byte(u.Password))

	// get token for auth
	token, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// set expire time for that token
	expireat := time.Now().Add(EXPIRE_DURATION)

	u.Password = password
	u.Token = string(token.String())
	u.ExpireAt = expireat

	sqlStr := `INSERT INTO users (name,password,username,token,expire_at)
				VALUES (:name,:password,:username,:token,:expire_at)`

	_, err = u.db.NamedExec(sqlStr, u)
	if err != nil {
		return err
	}

	return nil
}

func (u *user) AuthUser(token string) error {
	defer u.db.Close()
	sqlStr := `SELECT * FROM users u
			   WHERE u.token = ? `

	if err := u.db.QueryRowx(sqlStr, token).StructScan(u); err != nil {
		return errors.New("no user found")
	}

	if time.Now().UTC().After(u.ExpireAt) {
		return errors.New("invalid token, please login again")
	}

	return nil
}

func (u *user) Login() error {
	defer u.db.Close()
	// get user
	err := u.getUserByUsernamePassword()
	if err != nil {
		return errors.New("no user found please register")
	}

	// get token for auth
	token, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// set expire time for that token
	expireat := time.Now().UTC().Add(EXPIRE_DURATION)

	u.Token = string(token.String())
	u.ExpireAt = expireat

	sqlStr := `UPDATE users
			   SET 
			   token=:token,expire_at=:expire_at
			   WHERE id =:id `

	_, err = u.db.NamedExec(sqlStr, u)
	if err != nil {
		return err
	}

	return nil
}

// use this method while login to check if the username and password is correct
func (u *user) getUserByUsernamePassword() error {

	// encode the password
	password := base64.RawStdEncoding.EncodeToString([]byte(u.Password))
	u.Password = password

	sqlStr := `SELECT * FROM users WHERE username = ? AND password = ?`

	if err := u.db.QueryRowx(sqlStr, u.Username, u.Password).StructScan(u); err != nil {
		return err
	}

	return nil
}

// use this method to check where username is already in use
func (u *user) getUserByUsername() error {

	sqlStr := `SELECT * FROM users WHERE username = ?`

	if err := u.db.QueryRowx(sqlStr, u.Username).StructScan(u); err != nil {
		return err
	}

	return nil
}
