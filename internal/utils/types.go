package utils

type PageData struct {
	Authenticated bool
	Files         []FileInfo
	Folders       []DirectoryInfo
}

type DeleteRequest struct {
	Name string `json: name`
}

type FileInfo struct {
	Name string
	Path string
}

type DirectoryInfo struct {
	Name string
	Path string
}

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}
