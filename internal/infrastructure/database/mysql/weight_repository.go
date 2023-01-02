package mysql

import (
	"database/sql"
	"time"

	"github.com/aristorinjuang/weight-go/internal/aggregate"
	"github.com/aristorinjuang/weight-go/internal/entity"
	"github.com/aristorinjuang/weight-go/internal/repository"
	_ "github.com/go-sql-driver/mysql"
)

type weightRepositoryMySQL struct {
	db *sql.DB
}

func (d *weightRepositoryMySQL) List() (*aggregate.Weights, error) {
	rows, err := d.db.Query("SELECT max, min, date FROM weights")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	weights := aggregate.NewWeights(nil)

	for rows.Next() {
		var (
			max, min float64
			date     time.Time
		)
		rows.Scan(&max, &min, &date)

		weights.AddWeight(entity.RebuildWeight(date, max, min))
	}

	return weights, nil
}

func (d *weightRepositoryMySQL) Create(weight *entity.Weight) error {
	_, err := d.db.Exec(
		"INSERT INTO weights(max, min, date) VALUES(?, ?, ?)",
		weight.Max(),
		weight.Min(),
		weight.Date(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *weightRepositoryMySQL) Read(date time.Time) (*entity.Weight, error) {
	rows, err := d.db.Query(
		"SELECT max, min FROM weights WHERE date = ?",
		date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var max, min float64
	for rows.Next() {
		rows.Scan(&max, &min)
	}

	return entity.RebuildWeight(date, max, min), nil
}

func (d *weightRepositoryMySQL) Update(weight *entity.Weight) error {
	_, err := d.db.Exec(
		"UPDATE weights SET max = ?, min = ? WHERE date = ?",
		weight.Max(),
		weight.Min(),
		weight.Date(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *weightRepositoryMySQL) Delete(date time.Time) error {
	_, err := d.db.Exec(
		"DELETE FROM weights WHERE date = ?",
		date,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewWeightRepositoryMySQL(db *sql.DB) repository.WeightRepository {
	return &weightRepositoryMySQL{db}
}
