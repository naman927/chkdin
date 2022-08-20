package modals

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type person struct {
	db        *sqlx.DB
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"firstname"`
	LastName  string `json:"last_name" db:"lastname"`
	Address   string `json:"address" db:"address"`
	ContactNo string `json:"contact_no" db:"contactno"`
	Email     string `json:"email" db:"email"`
}

func NewPerson() (*person, error) {
	db, err := SetUpDb()
	if err != nil {
		return nil, err
	}

	return &person{
		db: db,
	}, nil
}

func (p *person) GetPersonDetails(id int) error {
	defer p.db.Close()
	sqlStr := `SELECT * FROM persons p 
			   WHERE p.id = ?`

	if err := p.db.QueryRowx(sqlStr, id).StructScan(p); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}

func (p *person) CreatePerson() error {
	defer p.db.Close()
	sqlStr := `INSERT INTO persons (firstname,lastname,address,contactno,email)
				VALUES (:firstname,:lastname,:address,:contactno,:email)`

	_, err := p.db.NamedExec(sqlStr, p)
	if err != nil {
		return err
	}

	return nil
}

func (p *person) UpdatePerson() error {
	defer p.db.Close()
	sqlStr := `UPDATE persons
			   SET
			   firstname = :firstname, lastname = :lastname, address = :address, contactno = :contactno, email = :email
			   WHERE id = :id`

	_, err := p.db.NamedExec(sqlStr, p)
	if err != nil {
		return err
	}

	return nil
}

func (p *person) DeletePerson(id int) error {
	defer p.db.Close()
	sqlStr := `DELETE FROM persons p 
			   WHERE p.id = ?`

	if err := p.db.QueryRowx(sqlStr, id).Scan(struct{}{}); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}
