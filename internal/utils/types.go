package utils

type PageData struct {
	Files   []string
	Folders []string
}

type DeleteRequest struct {
	Name string `json: name`
}
