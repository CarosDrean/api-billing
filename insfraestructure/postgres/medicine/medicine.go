package medicine

import (
	"database/sql"

	"api-billing/insfraestructure/postgres"
	"api-billing/model"
)

const table = "medicines"

var fields = []string{
	"name",
	"price",
	"location",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type Medicine struct {
	db *sql.DB
}

func New(db *sql.DB) Medicine {
	return Medicine{db}
}

func (p Medicine) Create(m *model.Medicine) error {
	stmt, err := p.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		m.Price,
		m.Location,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		if errPsql := postgres.CheckError(err); errPsql != nil {
			return errPsql
		}
		return err
	}

	return nil
}

func (p Medicine) GetWhere(filter model.Fields, sort model.SortFields) (model.Medicine, error) {
	conditions, args := postgres.BuildSQLWhere(filter)
	query := psqlGetAll + " " + conditions

	sorts := postgres.BuildSQLOrderBy(sort)
	query += " " + sorts

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return model.Medicine{}, err
	}
	defer stmt.Close()

	return p.scanRow(stmt.QueryRow(args...))
}

func (p Medicine) GetAllWhere(filter model.Fields, sort model.SortFields, pag model.Pagination) (model.Medicines, error) {
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

	ms := make(model.Medicines, 0)
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Medicine) scanRow(sq postgres.RowScanner) (model.Medicine, error) {
	m := model.Medicine{}
	updatedAtNull := sql.NullTime{}

	err := sq.Scan(
		&m.ID,
		&m.Name,
		&m.Price,
		&m.Location,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return model.Medicine{}, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
