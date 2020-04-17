package app

type (
	// AddRequest .
	AddRequest struct {
		Key     string   `json:"key"`
		Strings []string `json:"strings"`
	}
	// ApplyRequest .
	ApplyRequest struct {
		Size       uint64 `json:"size"`
		Expiration string `json:"expiration"`
	}
)
