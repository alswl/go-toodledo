package toodledo

type ErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}