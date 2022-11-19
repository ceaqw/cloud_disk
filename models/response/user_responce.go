package response

import "time"

type UserAllResponse struct {
	Id       uint64    `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
	Identity string    `json:"identity"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
	Deleted  time.Time `json:"deleted_at"`
}
