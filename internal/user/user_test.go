package user

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/stuckinforloop/shit/deps/timeutils"
	"github.com/stuckinforloop/shit/deps/ulid"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

const (
	testID    = "01JFYY7M4G06AFVGQT5ZYC0GEK"
	testName  = "test_user"
	testEmail = "test_email"
)

func setup(t *testing.T) (*DAO, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	close := func() {
		db.Close()
	}

	timeNow := func() time.Time {
		return timeutils.FoundingTimeUTC
	}

	rnd := rand.New(rand.NewSource(0))
	ulid := ulid.New(rnd, timeNow)
	dao := NewDAO(db, timeNow, ulid)

	return dao, mock, close
}

func TestUser(t *testing.T) {
	ctx := context.Background()
	dao, mock, close := setup(t)
	defer close()

	user := &User{
		ID:    testID,
		Name:  testName,
		Email: testEmail,
	}

	t.Run("create user", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO users").
			WithArgs(testID, testName, testEmail, timeutils.FoundingTimeUTC).
			WillReturnResult(sqlmock.NewResult(0, 1))
		_, err := dao.CreateUser(ctx, user)
		require.NoError(t, err)
	})

	t.Run("get user", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, email, created_at FROM users").
			WithArgs(testID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "email", "created_at"}).
					AddRow(testID, testName, testEmail, timeutils.FoundingTimeUTC),
			)

		u, err := dao.GetUser(ctx, testID)
		require.NoError(t, err)
		require.Equal(t, u.ID, testID)
		require.Equal(t, u.Name, testName)
		require.Equal(t, u.Email, testEmail)
		require.Equal(t, u.CreatedAt, timeutils.FoundingTimeUTC)
	})

	t.Run("get user -- sql no rows err", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, email, created_at FROM users").
			WithArgs(testID).
			WillReturnError(sql.ErrNoRows)
		u, err := dao.GetUser(ctx, testID)
		require.NoError(t, err)
		require.Nil(t, u)
	})

	t.Run("update user", func(t *testing.T) {
		mock.ExpectExec("UPDATE users").
			WithArgs(testName, testEmail, testID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		_, err := dao.UpdateUser(ctx, user)
		require.NoError(t, err)
	})

	t.Run("delete user", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM users").
			WithArgs(testID).WillReturnResult(sqlmock.NewResult(0, 1))
		err := dao.DeleteUser(ctx, testID)
		require.NoError(t, err)
	})
}
