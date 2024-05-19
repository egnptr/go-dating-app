package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/egnptr/dating-app/model"
)

func (*sqliteRepo) GetUser(ctx context.Context, username string) (*model.User, error) {
	db, err := sql.Open("sqlite3", "./testing.db")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var hashedPassword string
	var fullName string
	var email string
	var isPremium bool

	if err := db.QueryRow(getUser, username).Scan(
		&hashedPassword,
		&fullName,
		&email,
		&isPremium,
	); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user := model.User{
		Username:  username,
		Password:  hashedPassword,
		FullName:  fullName,
		Email:     email,
		IsPremium: isPremium,
	}

	return &user, nil
}

func (*sqliteRepo) GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	db, err := sql.Open("sqlite3", "./testing.db")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var username string
	var fullName string
	var email string
	var isPremium bool

	if err := db.QueryRow(getUserByID, userID).Scan(
		&username,
		&fullName,
		&email,
		&isPremium,
	); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user := model.User{
		Username:  username,
		FullName:  fullName,
		Email:     email,
		IsPremium: isPremium,
	}

	return &user, nil
}

func (*sqliteRepo) GetRelatedUser(ctx context.Context, id int64) ([]model.User, error) {
	db, err := sql.Open("sqlite3", "./testing.db")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	rows, err := db.Query(getRelatedUserBasedOnID, id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var id int64
		var fullName string
		err = rows.Scan(&id, &fullName)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		user := model.User{
			UserID:   id,
			FullName: fullName,
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return users, nil
}
