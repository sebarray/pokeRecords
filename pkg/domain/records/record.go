package records

type CreateRecordRequest struct {
	Name  string `json:"name"  db:"NAME"`
	Color string `json:"color"   db:"COLOR"`
	Value int64  `json:"value"   db:"VALUE" `
}

type CreateRecordResponse struct {
	RecordID   string `json:"recordID"`
	StatusCode int    `json:"statusCode"`
}
