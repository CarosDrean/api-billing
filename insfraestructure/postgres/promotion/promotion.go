package promotion

import (
	"api-billing/insfraestructure/postgres"
	"api-billing/model"
	"database/sql"
)

const table = "promotions"

var fields = []string{
	"description",
	"percentage",
	"start_date",
	"finish_date",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type Promotion struct {
	db *sql.DB
}

func New(db *sql.DB) Promotion {
	return Promotion{db}
}

func (p Promotion) Create(m *model.Promotion) error {
	stmt, err := p.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Description,
		m.Percentage,
		m.StartDate,
		m.FinishDate,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		if errPsql := postgres.CheckError(err); errPsql != nil {
			return errPsql
		}
		return err
	}

	return nil
}

func (p Promotion) GetWhere(filter model.Fields, sort model.SortFields) (model.Promotion, error) {
	conditions, args := postgres.BuildSQLWhere(filter)
	query := psqlGetAll + " " + conditions

	sorts := postgres.BuildSQLOrderBy(sort)
	query += " " + sorts

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return model.Promotion{}, err
	}
	defer stmt.Close()

	return p.scanRow(stmt.QueryRow(args...))
}

func (p Promotion) GetAllWhere(filter model.Fields, sort model.SortFields, pag model.Pagination) (model.Promotions, error) {
	conditions, args := postgres.BuildSQLWhere(filter)
	query := psqlGetAll + " " + conditions

	sorts := postgres.BuildSQLOrderBy(sort)
	query += " " + sorts

	pagination := postgres.BuildSQLPagination(pag)
	query += " " + pagination

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Promotions, 0)
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Promotion) scanRow(sq postgres.RowScanner) (model.Promotion, error) {
	m := model.Promotion{}
	updatedAtNull := sql.NullTime{}

	err := sq.Scan(
		&m.ID,
		&m.Description,
		&m.Percentage,
		&m.StartDate,
		&m.FinishDate,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return model.Promotion{}, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
