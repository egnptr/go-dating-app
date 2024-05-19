package db

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/egnptr/dating-app/model"
)

func (*sqliteRepo) CreateUser(ctx context.Context, req model.User) (err error) {
	db, err := sql.Open("sqlite3", "./testing.db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err.Error())
		return
	}
	stmt, err := tx.Prepare(createUser)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(req.Username, req.Password, req.FullName, req.Email)
	if err != nil {
		log.Println(err.Error())
		return
	}

	tx.Commit()
	return
}

func (*sqliteRepo) UpdatePremiumStatus(ctx context.Context, req model.SubscribeRequest) (err error) {
	db, err := sql.Open("sqlite3", "./testing.db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err.Error())
		return
	}
	stmt, err := tx.Prepare(updatePremiumStatus)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(req.Subscribe, req.UserID)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		err = errors.New("error failed to update subscription status")
		return
	}

	tx.Commit()
	return
}
