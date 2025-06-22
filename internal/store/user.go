package store

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int64    `json:"id" `
	Username   string   `json:"username"`
	Email      string   `json:"email"`
	Password   password `json:"-"`
	Created_At string   `json:"created_at"`
	Updated_At string   `json:"updated_at"`
	Is_Active  bool     `json:"is_active"`
}

type UserStore struct {
	db *sql.DB
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare() error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(*p.text))
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
		INSERT INTO users(username, email, password, is_active)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	err := tx.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password.hash,
		user.Is_Active,
	).Scan(
		&user.ID,
		&user.Created_At,
		&user.Updated_At,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	query := `
  	SELECT id, username, email, created_at, updated_at FROM users	
		WHERE id = $1 AND is_active = true
	`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Created_At,
		&user.Updated_At,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserStore) Delete(ctx context.Context, userID int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) CreateAndInvite(ctx context.Context, user *User, token string, expiry time.Duration) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		// create user
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}
		// create invitation
		if err := s.createUserInvitation(ctx, tx, token, user.ID, expiry); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) createUserInvitation(ctx context.Context, tx *sql.Tx, token string, userId int64, expiry time.Duration) error {
	query := `
		INSERT INTO user_invitations(token, user_id, expiry)
		VALUES ($1, $2, $3)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, userId, time.Now().Add(expiry))
	if err != nil {
		return err
	}

	return nil
}
