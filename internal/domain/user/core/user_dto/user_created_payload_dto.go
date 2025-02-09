package user_dto

type UserCreatedPayloadDto struct {
	Id              string                   `json:"user_id"`
	SuccessResponse chan *SessionResponseDto `json:"session_response"`
}

type SessionResponseDto struct {
	Token   string     `json:"token"`
	Session SessionDto `json:"session"`
}

type SessionDto struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   bool   `json:"is_admin"`
}
