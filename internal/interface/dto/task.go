package dto

type TaskResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
