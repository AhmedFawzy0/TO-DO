package repos

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/AhmedFawzy0/TO-DO/app/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepoTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

type Repo struct {
	db *gorm.DB
}

type anyTime struct{}

func (a anyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (s *RepoTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	s.NoError(err)

	s.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	s.NoError(err)

}

func (s *RepoTestSuite) TestCreateUser() {

	repo := NewRepo(s.db)

	username_test := "user"
	password_test := "password"

	s.mock.ExpectBegin()
	s.mock.ExpectQuery("INSERT INTO \"users\" \\(\"created_at\",\"updated_at\",\"deleted_at\",\"username\",\"password\"\\) VALUES \\(\\$1,\\$2,\\$3,\\$4,\\$5\\) RETURNING \"id\",\"username\",\"password\"").
		WithArgs(anyTime{}, anyTime{}, nil, "user", "password").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()

	user, err := CreateUser(username_test, password_test, repo.db)

	s.NoError(err)
	s.Equal(username_test, user.Username)
	s.Equal(password_test, user.Password)
	s.NoError(s.mock.ExpectationsWereMet())

}

func (s *RepoTestSuite) TestFindUserEmpty() {
	repo := NewRepo(s.db)
	user_test := "user"
	user_found := &models.User{}
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE Username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(user_test).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err := FindUser(user_found, user_test, repo.db)

	s.NoError(err)
	s.Equal(user_found.Username, "")
	s.NoError(s.mock.ExpectationsWereMet())

}

func (s *RepoTestSuite) TestFindExistingUser() {
	repo := NewRepo(s.db)

	username_create := "user"
	password_create := "password"

	s.mock.ExpectBegin()
	s.mock.ExpectQuery("INSERT INTO \"users\" \\(\"created_at\",\"updated_at\",\"deleted_at\",\"username\",\"password\"\\) VALUES \\(\\$1,\\$2,\\$3,\\$4,\\$5\\) RETURNING \"id\",\"username\",\"password\"").
		WithArgs(anyTime{}, anyTime{}, nil, username_create, password_create).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()

	user_find := "user"
	user_found := &models.User{}
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE Username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(user_find).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	user_created, err_created := CreateUser(username_create, password_create, repo.db)
	err_find := FindUser(user_found, user_find, repo.db)

	s.NoError(err_find)
	s.NoError(err_created)
	s.Equal(username_create, user_created.Username)
	s.Equal(password_create, user_created.Password)
	s.NoError(s.mock.ExpectationsWereMet())

}

func (s *RepoTestSuite) TestFindEmptyTasks() {
	repo := NewRepo(s.db)
	user_test := &models.User{Username: "user", Password: "password"}
	user_test.ID = 1

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."user_id" = $1 AND "tasks"."deleted_at" IS NULL`)).
		WithArgs(user_test.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	task_find := new([]models.Task)
	err := FindUserTasks(user_test, task_find, repo.db)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *RepoTestSuite) TestDeleteTask() {
	repo := NewRepo(s.db)
	task_id := uint(1)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(`UPDATE "tasks" SET "deleted_at"=\$1 WHERE "tasks"."id" = \$2 AND "tasks"."deleted_at" IS NULL`).
		WithArgs(anyTime{}, task_id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	err := TaskDelete(task_id, repo.db)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())

}

func (s *RepoTestSuite) TestUpdateTask() {
	//Flips the task finished status!

	repo := NewRepo(s.db)
	task_model := &models.Task{Finished: false, Detail: "run"}
	task_model.ID = uint(0)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(`UPDATE "tasks" SET "updated_at"=\$1,"finished"=\$2,"detail"=\$3 WHERE id = \$4 AND "tasks"."deleted_at" IS NULL`).
		WithArgs(anyTime{}, !task_model.Finished, task_model.Detail, task_model.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := UpdateTask(task_model, repo.db)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())

}

func TestRepoTestSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}
