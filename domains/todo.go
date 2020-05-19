package domains

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding: "required"`
}