package body

type BodyLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BodyChangePass struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
