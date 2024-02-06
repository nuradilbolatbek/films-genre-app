package FilmsProject

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"Name" binging:"required"`
	Username string `json:"Username" binging:"required"`
	Password string `json:"Password" binging:"required"`
}
