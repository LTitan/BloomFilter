package handler

type (
	ApplyRes struct {
		Yes bool   `json:"yes"`
		Key string `json:"key"`
	}
	UpdateFeild struct {
		UUID         string `json:"uuid"`
		Size         uint64 `json:"size"`
		ExpirationAt string `json:"expiration_at"`
	}
)
