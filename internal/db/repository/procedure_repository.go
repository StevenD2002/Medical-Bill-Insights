package repository

import (
	"database/sql"

	"github.com/StevenD2002/Medical-Bill-Insights/internal/db/models"
)

type ProcedureRepository interface {
	GetProcedureByCode(code string) (*models.Procedure, error)
}

type procedureRepo struct {
	db *sql.DB
}

func NewProcedureRepository(db *sql.DB) ProcedureRepository {
	return &procedureRepo{db: db}
}

func (r *procedureRepo) GetProcedureByCode(code string) (*models.Procedure, error) {
	// implementation
	// query all items in the code table that have the same code inputted
	var procedure models.Procedure

	query := `SELECT billing_code, billing_code_type FROM code WHERE billing_code = $1`
	// cast types to string before assigning to the procedure struct
	err := r.db.QueryRow(query, code).Scan(&procedure.Code, &procedure.BillingCodeType)
	if err != nil {
		return nil, err
	}
	return &procedure, nil
}
