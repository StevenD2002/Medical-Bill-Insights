package models

// Procedure represents a procedure that appears on a bill
type Procedure struct {
	Code            string  `db:"billing_code"`
	BillingCodeType string  `db:"billing_code_type"`
	NegotiatedRate  float64 `db:"negotiated_rate"`
	RateMetadataID  []uint8 `db:"rate_metadata_id"`
}
