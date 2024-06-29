package database

import (
	"regexp"
	"testing"

	"github.com/BrandonC98/fortify/services/fortify/internal/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAddRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	dsn := "TESTING_DB"
	assert.NoError(t, err)
	defer db.Close()
	r := SecretsRepository{
		user:     "user",
		host:     "host",
		name:     "name",
		password: "password",
	}

	var tests = []struct {
		name          string
		expectedCreds model.Secret
	}{
		{"Successfully insert credentials from the db", model.Secret{ID: 1,
			Name:  "abc",
			Value: "def",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gormDB, err := gorm.Open(mysql.New(mysql.Config{
				DSN:                       dsn,
				Conn:                      db,
				SkipInitializeWithVersion: true,
			}), &gorm.Config{})
			assert.NoError(t, err)

			r.DB = *gormDB

			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO `secrets`").
				WithArgs(test.expectedCreds.Name, test.expectedCreds.Value, 1).
				WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectCommit()

			r.AddRecord(&test.expectedCreds)

			assert.Nil(t, mock.ExpectationsWereMet())
			assert.Equal(t, uint(1), test.expectedCreds.ID)
			assert.Equal(t, "abc", test.expectedCreds.Name)
			assert.Equal(t, "def", test.expectedCreds.Value)
		})
	}
}

func TestRetriveAllRecords(t *testing.T) {
	db, mock, err := sqlmock.New()
	dsn := "TESTING_DB"
	assert.NoError(t, err)
	defer db.Close()
	r := SecretsRepository{
		user:     "user",
		host:     "host",
		name:     "name",
		password: "password",
	}

	var tests = []struct {
		name          string
		expectedCreds model.Secret
	}{
		{"Successfully retrived credentials from the db", model.Secret{ID: 1,
			Name:  "abc",
			Value: "def",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gormDB, err := gorm.Open(mysql.New(mysql.Config{
				DSN:                       dsn,
				Conn:                      db,
				SkipInitializeWithVersion: true,
			}), &gorm.Config{})
			assert.NoError(t, err)
			r.DB = *gormDB

			rows := sqlmock.NewRows([]string{"id", "name", "value"}).
				AddRow(test.expectedCreds.ID, test.expectedCreds.Name, test.expectedCreds.Value)

			// regexp.QuoteMeta is needed to escape some characters
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `secrets`")).WillReturnRows(rows)

			creds := r.RetriveAllRecords()

			assert.Equal(t, test.expectedCreds.ID, creds[0].ID)
			assert.Equal(t, test.expectedCreds.Name, creds[0].Name)
			assert.Equal(t, test.expectedCreds.Value, creds[0].Value)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
