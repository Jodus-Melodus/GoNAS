package utils

type PageData struct {
	Authenticated bool
	Files         []string
	Folders       []string
}

type DeleteRequest struct {
	Name string `json: name`
}
