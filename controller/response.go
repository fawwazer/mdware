package user

type LoginResponse struct {
	Hp    int    `json:"hp"`
	Nama  string `json:"nama"`
	Token string `json:"token"`
}
