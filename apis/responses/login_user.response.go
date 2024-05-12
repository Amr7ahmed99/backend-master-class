package responses

import db "backend-master-class/db/sqlc"

type LoginUserResponse struct {
	AccessToken string  `json:"access_token"`
	User        db.User `json:"user"`
}
