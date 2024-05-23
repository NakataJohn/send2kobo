package domain

type RequestID struct {
	ID string `json:"id"`
}

type RequestIDs struct {
	ids []RequestID
}
