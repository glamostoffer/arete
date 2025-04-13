package repository

import (
	"context"

	"github.com/glamostoffer/arete/auth/internal/domain"
	"github.com/glamostoffer/arete/auth/pkg/errlist"
)

func (r *repository) InsertUser(ctx context.Context, user domain.User) (userID int64, err error) {
	err = r.db.GetContext(
		ctx,
		&userID,
		queryInsertUser,
		user.Login,
		user.Email,
		user.HashPassword,
		user.RegistrationDate,
	)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (r *repository) GetUserByLoginOrEmail(ctx context.Context, login string) (user domain.User, err error) {
	err = r.db.GetContext(
		ctx,
		&user,
		queryGetUserByLoginOrEmail,
		login,
	)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *repository) CheckUserExists(ctx context.Context, login, email string) (bool, error) {
	var exists bool

	err := r.db.GetContext(
		ctx,
		&exists,
		queryCheckUserExists,
		login,
		email,
	)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (r *repository) UpdateUser(ctx context.Context, user domain.User) error {
	res, err := r.db.ExecContext(
		ctx,
		queryUpdateUser,
		user.ID,
		user.Login,
		user.Email,
		user.HashPassword,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errlist.ErrInvalidAffectedRows
	}

	return nil
}

func (r *repository) GetUser(ctx context.Context, userID int64) (user domain.User, err error) {
	err = r.db.GetContext(ctx, &user, queryGetUser, userID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
