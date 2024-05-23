package models

// Procedure represents a procedure that appears on a bill
type Procedure struct {
	Code            string `db:"billing_code"`
	BillingCodeType string `db:"billing_code_type"`
}
