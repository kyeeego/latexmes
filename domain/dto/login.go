package dto

type LoginDto struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Fingerprint string `json:"fingerprint"`
}

type LogoutDto struct {
	Fingerprint string `json:"fingerprint"`
}

type RefreshDto struct {
	Token       string `json:"token"`
	Fingerptint string `json:"fingerprint"`
}

type TokenPair struct {
	Jwt          string `json:"jwt"`
	RefreshToken string `json:"reftesh_token"`
}
