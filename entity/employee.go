package entity

type Employee struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Team  string `json:"team"`
	Email string `json:"email"`
}
