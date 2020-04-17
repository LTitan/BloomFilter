package fe

type (
	// CPUMemoryInfo .
	CPUMemoryInfo struct {
		Legend []string `json:"legend"`
		YAxis  []string `json:"y_axis"`
		XAxis  []string `json:"x_axis"`
		Series struct {
			CPU    []interface{} `json:"cpu"`
			Memory []interface{} `json:"memory"`
		} `json:"series"`
	}
	// ProductInfo .
	ProductInfo struct {
		Legend []string      `json:"legend"`
		YAxis  []string      `json:"y_axis"`
		XAxis  []string      `json:"x_axis"`
		Series []interface{} `json:"series"`
	}
	// NameAndValue .
	NameAndValue struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	}
)
