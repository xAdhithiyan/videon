package types

import "net/http"

type UserStore interface {
	CreateUser(User) error
	GetUserByEmail(email string) (*User, error)
}
type VideoFuns interface {
	ParseData(msg []byte, userID int) (int, bool)
}
type VideoStore interface {
	UploadS3(metaData MetaData, data []byte) error
	AddVideoDB(userId int, metaData MetaData) error
}

type Producer interface {
	SetupProducer(w http.ResponseWriter, r *http.Request)
}

type User struct {
	Id       int
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Cookie struct {
	Name  string
	Value string
}

type MetaData struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	VideoType   string `json:"videoType"`
	TotalChunks int    `json:"totalChunks"`
}

type VideoDB struct {
	UserId      int
	ID          int
	Name        string
	VideoType   string
	TotalChunks int

	CompressVideo     bool
	GeenrateThumbnail bool
	TranscodeVideo    bool
	AddWaterMark      bool
	VideoSummary      bool
}

type ChunkResponse struct {
	ChunkId  int
	Uploaded bool
}

type ProducerPayload struct {
	VideoId int            `json:"id" validate:"required"`
	Options ProducerNeeded `json:"options" validate:"required"`
}

type ProducerNeeded struct {
	Compress   bool `json:"compress" `
	Thumbnails bool `json:"thumbnails" `
	Transcode  bool `json:"transcode" `
	Watermark  bool `json:"watermark" `
	Summary    bool `json:"summary" `
}
