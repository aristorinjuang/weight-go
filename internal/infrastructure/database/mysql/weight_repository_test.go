package mysql

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aristorinjuang/weight-go/internal/domain/repository"
	"github.com/aristorinjuang/weight-go/test/data"
	"github.com/stretchr/testify/suite"
)

type weightRepositoryMySQLTest struct {
	suite.Suite
	db               *sql.DB
	mock             sqlmock.Sqlmock
	weightRepository repository.WeightRepository
}

func (t *weightRepositoryMySQLTest) SetupSuite() {
	t.db, t.mock, _ = sqlmock.New()
	t.weightRepository = NewWeightRepositoryMySQL(t.db)
}

func (t *weightRepositoryMySQLTest) TestList() {
	query := "SELECT max, min, date FROM weights"

	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{"max", "min", "date"}).AddRow(
			data.Weights[0].Max(),
			data.Weights[0].Min(),
			data.Weights[0].Date(),
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		weights, err := t.weightRepository.List()

		t.NotNil(weights)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("failed list weights"))

		weights, err := t.weightRepository.List()

		t.Nil(weights)
		t.Error(err)
	})
}

func (t *weightRepositoryMySQLTest) TestCreate() {
	query := "INSERT INTO weights"

	t.Run("success", func() {
		t.mock.ExpectExec(query).WithArgs(
			data.Weights[0].Max(),
			data.Weights[0].Min(),
			data.Weights[0].Date(),
		).WillReturnResult(sqlmock.NewResult(1, 1))

		t.NoError(t.weightRepository.Create(data.Weights[0]))
	})

	t.Run("failed", func() {
		t.mock.ExpectExec(query).WithArgs(
			data.Weights[0].Max(),
			data.Weights[0].Min(),
			data.Weights[0].Date(),
		).WillReturnError(errors.New("failed create a weight"))

		t.Error(t.weightRepository.Create(data.Weights[0]))
	})
}

func (t *weightRepositoryMySQLTest) TestRead() {
	query := "SELECT max, min FROM weights WHERE date = ?"

	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{"max", "min"}).AddRow(
			data.Weights[0].Max(),
			data.Weights[0].Min(),
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		weight, err := t.weightRepository.Read(data.Weights[0].Date())

		t.NotNil(weight)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("failed read a weight"))

		weight, err := t.weightRepository.Read(data.Weights[0].Date())

		t.Nil(weight)
		t.Error(err)
	})
}

func (t *weightRepositoryMySQLTest) TestUpdate() {
	query := "UPDATE weights"

	t.Run("success", func() {
		t.mock.ExpectExec(query).WithArgs(
			data.Weights[0].Max(),
			data.Weights[0].Min(),
			data.Weights[0].Date(),
		).WillReturnResult(sqlmock.NewResult(1, 1))

		t.NoError(t.weightRepository.Update(data.Weights[0]))
	})

	t.Run("success", func() {
		t.mock.ExpectExec(query).WithArgs(
			data.Weights[0].Max(),
			data.Weights[0].Min(),
			data.Weights[0].Date(),
		).WillReturnError(errors.New("failed update the weight"))

		t.Error(t.weightRepository.Update(data.Weights[0]))
	})
}

func (t *weightRepositoryMySQLTest) TestDelete() {
	query := "DELETE FROM weights"

	t.Run("success", func() {
		t.mock.ExpectExec(query).WithArgs(data.Weights[0].Date()).WillReturnResult(sqlmock.NewResult(1, 1))

		t.NoError(t.weightRepository.Delete(data.Weights[0].Date()))
	})

	t.Run("success", func() {
		t.mock.ExpectExec(query).WithArgs(data.Weights[0].Date()).WillReturnError(errors.New("failed delete the weight"))

		t.Error(t.weightRepository.Delete(data.Weights[0].Date()))
	})
}

func TestWeightRepositoryMySQL(t *testing.T) {
	suite.Run(t, new(weightRepositoryMySQLTest))
}
