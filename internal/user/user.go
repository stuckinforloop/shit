package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (d *DAO) CreateUser(ctx context.Context, u *User) (*User, error) {
	id, err := d.ulid.Generate()
	if err != nil {
		return nil, fmt.Errorf("generate id: %w", err)
	}
	createdAt := d.timeNow()

	u.ID = id
	u.CreatedAt = createdAt

	query := `
		INSERT INTO users (id, name, email, created_at)
		VALUES (?, ?, ?, ?)
	`

	if _, err := d.db.ExecContext(ctx, query, u.ID, u.Name, u.Email, u.CreatedAt); err != nil {
		return nil, fmt.Errorf("insert user: %w", err)
	}

	return u, nil
}

func (d *DAO) GetUser(ctx context.Context, id string) (*User, error) {
	user := &User{}

	query := `
		SELECT id, name, email, created_at
		FROM users
		WHERE id = ?
	`
	if err := d.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}

func (d *DAO) UpdateUser(ctx context.Context, u *User) (*User, error) {
	query := `
		UPDATE users
		SET name = ?,
			email = ?,
		WHERE id = ?
	`
	if _, err := d.db.ExecContext(ctx, query, u.Name, u.Email, u.ID); err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	return u, nil
}

func (d *DAO) DeleteUser(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE id = ?
	`
	if _, err := d.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}
