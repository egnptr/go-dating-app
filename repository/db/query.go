package db

const (
	insertUserTable = `
	CREATE TABLE "users" (
		"id" integer PRIMARY KEY,
		"username" varchar NOT NULL,
		"password" varchar NOT NULL,
		"full_name" varchar NOT NULL,
		"email" varchar UNIQUE NOT NULL,
		"is_premium" bool NOT NULL DEFAULT (false),
		"created_at" timestamptz NOT NULL DEFAULT (date()),
		"updated_at" timestamptz
	);
	`

	createUser = `
	INSERT INTO users (
		username,
		password,
		full_name,
		email
	) VALUES (
		$1, $2, $3, $4
	)
	`

	updatePremiumStatus = `
		UPDATE users SET 
			is_premium = $1,
			updated_at = date()
		WHERE id = $2
	`

	getUser = `
		SELECT password, full_name, email, is_premium FROM users
		WHERE username = $1 LIMIT 1
	`

	getUserByID = `
		SELECT password, full_name, email, is_premium FROM users
		WHERE id = $1 LIMIT 1
	`

	getRelatedUserBasedOnID = `
		SELECT id, full_name, email, is_premium FROM users
		WHERE id <> $1
	`
)
