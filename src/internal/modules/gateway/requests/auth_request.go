package gateway_requests

type AuthRequest struct {
	Jwt  string `json:"jwt"`
	Code string `json:"code"`
}
