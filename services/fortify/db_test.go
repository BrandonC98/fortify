package main

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAddCredsRecord1(t *testing.T) {
	db, mock, err := sqlmock.New()
	dsn := "TESTING_DB"
	assert.NoError(t, err)
	defer db.Close()
	r := CredentialRepository{
		user:     "user",
		host:     "host",
		name:     "name",
		password: "password",
	}

	var tests = []struct {
		name          string
		expectedCreds Credentials
	}{
		{"Successfully insert credentials from the db", Credentials{ID: 1,
			Name:   "abc",
			Passwd: "def",
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
			mock.ExpectExec("INSERT INTO `credentials`").
				WithArgs(test.expectedCreds.Name, test.expectedCreds.Passwd, 1).
				WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectCommit()

			r.AddCredsRecord(&test.expectedCreds)

			assert.Nil(t, mock.ExpectationsWereMet())
			assert.Equal(t, uint(1), test.expectedCreds.ID)
			assert.Equal(t, "abc", test.expectedCreds.Name)
			assert.Equal(t, "def", test.expectedCreds.Passwd)
		})
	}
}

func TestRetriveAllCreds(t *testing.T) {
	db, mock, err := sqlmock.New()
	dsn := "TESTING_DB"
	assert.NoError(t, err)
	defer db.Close()
	r := CredentialRepository{
		user:     "user",
		host:     "host",
		name:     "name",
		password: "password",
	}

	var tests = []struct {
		name          string
		expectedCreds Credentials
	}{
		{"Successfully retrived credentials from the db", Credentials{ID: 1,
			Name:   "abc",
			Passwd: "def",
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

			rows := sqlmock.NewRows([]string{"id", "name", "passwd"}).
				AddRow(test.expectedCreds.ID, test.expectedCreds.Name, test.expectedCreds.Passwd)

			// regexp.QuoteMeta is needed to escape some characters
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `credentials`")).WillReturnRows(rows)

			creds := r.retriveAllCreds()

			assert.Equal(t, test.expectedCreds.ID, creds[0].ID)
			assert.Equal(t, test.expectedCreds.Name, creds[0].Name)
			assert.Equal(t, test.expectedCreds.Passwd, creds[0].Passwd)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
