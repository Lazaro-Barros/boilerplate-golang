package application

type TodoIn struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	BoardUUID   string `json:"board_uuid" binding:"required"`
}
