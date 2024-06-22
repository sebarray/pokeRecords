package records

type CreateRecordRequest struct {
	RecordName string `json:"recordName"`
	Email      string `json:"email" validate:"required,email"`
	Color      string `json:"color"`
}

type CreateRecordResponse struct {
	RecordID   string `json:"recordID"`
	StatusCode int    `json:"statusCode"`
}
