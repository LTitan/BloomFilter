package fe

type (
	// CPUMemoryInfo .
	CPUMemoryInfo struct{
		Legend []string `json:"legend"`
		YAxis []string `json:"y_axis"`
		Series struct{
			CPU []int `json:"cpu"`
			Memory []int `json:"memory"`
		}`json:"series"`
	}
)