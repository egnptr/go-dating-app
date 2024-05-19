package http

type (
	responseDefault struct {
		Header Header      `json:"header"`
		Data   interface{} `json:"data"`
	}

	Header struct {
		ProcessTime float64  `json:"process_time"`
		Messages    []string `json:"messages"`
		Reason      string   `json:"reason"`
		ErrorCode   []string `json:"error_code"`
	}
)
