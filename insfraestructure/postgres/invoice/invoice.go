package invoice

import (
	"database/sql"
	"encoding/json"

	"api-billing/insfraestructure/postgres"
	"api-billing/model"
)

const table = "invoices"

var fields = []string{
	"total_price",
	"promotion_id",
	"medicines_ids",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type Invoice struct {
	db *sql.DB
}

func New(db *sql.DB) Invoice {
	return Invoice{db}
}

func (p Invoice) Create(m *model.Invoice) error {
	stmt, err := p.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if m.MedicinesIDs != nil {
		medicinesIDs, err := json.Marshal(m.MedicinesIDs)
		if err != nil {
			return err
		}
		medicinesIDsRawJson := json.RawMessage(medicinesIDs)
		m.MedicinesIDsRawJSON = &medicinesIDsRawJson
	}

	err = stmt.QueryRow(
		m.TotalPrice,
		m.PromotionID,
		m.MedicinesIDsRawJSON,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		if errPsql := postgres.CheckError(err); errPsql != nil {
			return errPsql
		}
		return err
	}

	return nil
}

func (p Invoice) GetWhere(filter model.Fields, sort model.SortFields) (model.Invoice, error) {
	conditions, args := postgres.BuildSQLWhere(filter)
	query := psqlGetAll + " " + conditions

	sorts := postgres.BuildSQLOrderBy(sort)
	query += " " + sorts

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return model.Invoice{}, err
	}
	defer stmt.Close()

	return p.scanRow(stmt.QueryRow(args...))
}

func (p Invoice) GetAllWhere(filter model.Fields, sort model.SortFields, pag model.Pagination) (model.Invoices, error) {
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

	ms := make(model.Invoices, 0)
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Invoice) scanRow(sq postgres.RowScanner) (model.Invoice, error) {
	m := model.Invoice{}

	err := sq.Scan(
		&m.ID,
		&m.TotalPrice,
		&m.PromotionID,
		&m.MedicinesIDsRawJSON,
		&m.CreatedAt,
		&m.UpdatedAt,
	)
	if err != nil {
		return model.Invoice{}, err
	}

	if m.MedicinesIDsRawJSON != nil {
		if err := json.Unmarshal(*m.MedicinesIDsRawJSON, &m.MedicinesIDs); err != nil {
			return m, err
		}
	}

	return m, nil
}
