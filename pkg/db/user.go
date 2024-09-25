package db

import (
	"context"
	"errors"
	"online-pathsaala/model"
	"strings"

	"github.com/lib/pq"
)

func (dbs *Database) Register(ctx context.Context, payload model.RegisterPayload) (id string, err error) {
	fields := []string{"email", "password", "username", "roles"}
	insertquery := InsertQuery(USERS, fields, "id", 1)
	hash, err := HashPasword(payload.Password)
	if err != nil {
		return "", nil
	}
	err = dbs.Db.QueryRowContext(ctx, insertquery, payload.Email, hash, payload.UserName, pq.Array([]string{payload.UserRole})).Scan(&id)
	if err != nil {
		errorString := err.Error()
		if strings.Contains(errorString, "user_authenticate_email_key") {
			err = errors.New("email is already used")
		}
		return
	}
	return
}

func (dbs *Database) GetUser(ctx context.Context, payload model.LoginPayload) (user model.User, err error) {
	query := GetRecordQuery(USERS, []string{"id", "email", "password", "username", "roles"}, []string{"email"}, "=", 1)
	err = dbs.Db.QueryRowContext(ctx, query, payload.Email).Scan(&user.ID, &user.Email, &user.Password, &user.UserName, pq.Array(&user.Roles))
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
