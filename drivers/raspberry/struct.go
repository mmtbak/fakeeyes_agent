package raspberry

// ResponseMessage  Http API response data struct
type ResponseMessage struct {
	Success      bool        `json:"Success"`
	Data         interface{} `json:"Data"`
	ErrorCode    string      `json:"ErrorCode"`
	ErrorMessage string      `json:"ErrorMessage"`
}
