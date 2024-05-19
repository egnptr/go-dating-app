# Backend Service for Dating App using Go

## How to run

### On Host machine

1. Install golang from https://golang.org/doc/install
2. Run `go mod vendor` to install required dependencies
3. Run by either directly running from source:

```
go run ./app/app.go
```

or by building and running the binary file from Makefile:

```
make run
```

### On docker

Or by simply using docker compose:

```
docker-compose up
```

# API endpoints

## GET

`/related-profiles` <br/>

## POST

`/user/sign-up` <br/>
`/user/login` <br/>
`/swipe` <br/>
`/subscribe-premium` <br/>
`/unsubscribe-premium` <br/>

---

### GET /related-profiles

Search for other dating profiles.

**Request Body**

```
{
    "user_id": 1
}
```

### POST /user/sign-up

Signs up for an account.

**Request Body**

```
{
    "username": "jdoe",
    "password": "test123",
    "full_name": "John Doe",
    "email": "john@doe.com"
}
```

---

### POST /user/login

Log in into an account.

**Request Body**

```
{
    "username": "jdoe",
    "password": "test123",
}
```

---

### POST /swipe

Swipes profile to pass (-1) or like (1).

**Request Body**

```
{
    "user_id": 1,
    "swiped_user_id": 2,
    "swipe_status": -1
}
```

---

### POST /subscribe-premium

Starts a new subscription or updates an existing one.

**Request Body**

```
{
    "user_id": 1
}
```

---

### POST /unsubscribe-premium

Cancels existing subscription.

**Request Body**

```
{
    "user_id": 1
}
```
