package app

type (
	// AddRequest .
	AddRequest struct {
		Key     uint32   `json:"key"`
		Strings []string `json:"strings"`
	}
	// ApplyRequest .
	ApplyRequest struct {
		Size uint64 `json:"size"`
	}
)
