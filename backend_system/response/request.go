package response

/**
* description:
* author: wqh
* date: 2025/1/8
 */
type RegisterInfo struct {
	Phone            string `json:"phone,omitempty"`
	VerificationCode string `json:"verification_code,omitempty"`
	Password         string `json:"password,omitempty"`
	Username         string `json:"username,omitempty"`
	SchoolCode       string `json:"school_code,omitempty"`
	Gender           string `json:"gender,omitempty"`
	Birthday         string `json:"birthday,omitempty"`
	Avatar           string `json:"avatar,omitempty"`
}

type LoginInfo struct {
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}
