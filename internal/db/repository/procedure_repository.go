package repository

import (
	"database/sql"

	"github.com/StevenD2002/Medical-Bill-Insights/internal/db/models"
)

type ProcedureRepository interface {
	GetProcedurePricesByCode(code string) (*models.Procedure, error)
}

type procedureRepo struct {
	db *sql.DB
}

func NewProcedureRepository(db *sql.DB) ProcedureRepository {
	return &procedureRepo{db: db}
}

func (r *procedureRepo) GetProcedurePricesByCode(code string) (*models.Procedure, error) {
	// implementation
	var procedure models.Procedure

	query := `SELECT c.billing_code AS billing_code, r.negotiated_rate AS negotiated_rate, r.rate_metadata_id AS rate_metadata_id FROM code as c  

  INNER JOIN rate AS r ON c.id = r.code_id WHERE c.billing_code = $1;`
	// cast types to string before assigning to the procedure struct
	err := r.db.QueryRow(query, code).Scan(&procedure.Code, &procedure.BillingCodeType)
	if err != nil {
		return nil, err
	}
	return &procedure, nil
}
