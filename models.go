package cloudreve

import "time"

// APIResponse is the general outer structure for
// Cloudreve API responses. It uses Go 1.18 generics
// and can encapsulate a data field of any type.
type APIResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// An User represents an user object.
// It includes a user's basic information, such as
// ID, UserName, Nickname and so on.
type User struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    int    `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	Preferred string `json:"preferred_theme"`
	Anonymous bool   `json:"anonymous"`
	Group     Group  `json:"group"`
	Tags      []any  `json:"tags"`
}

// A Group represents a group info of a User object
type Group struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	AllowShare           bool   `json:"allowShare"`
	AllowRemoteDownload  bool   `json:"allowRemoteDownload"`
	AllowArchiveDownload bool   `json:"allowArchiveDownload"`
	ShareDownload        bool   `json:"shareDownload"`
	Compress             bool   `json:"compress"`
	WebDAV               bool   `json:"webdav"`
	Source               int    `json:"sourceBatch"` // Didn't know what this is for
	AdvanceDel           bool   `json:"advanceDelete"`
	AllowWebDAVProxy     bool   `json:"allowWebDAVProxy"`
	// 注意：根据文档，max_storage 的类型是字符串，如 "10 GB"
	// MaxStorage        string `json:"max_storage"`
}

// An ObjectType represents a type of a file object.
type ObjectType string

const (
	Filetype ObjectType = "file"
	Dirtype  ObjectType = "dir"
)

// An Object represents a file or directory object.
type Object struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Type string    `json:"type"`
	Date time.Time `json:"date"`
	Size uint64    `json:"size"`
	Path string    `json:"path"`
	Pic  string    `json:"pic"`
}

type LoginRequest struct {
	Username string `json:"userName"`
	Password string `json:"Password"`
	Captcha  string `json:"captchaCode,omitempty"`
}
