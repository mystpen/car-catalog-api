package model

type CarInfo struct {
	ID     int64  `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year,omitempty"`
	Owner  Person `json:"owner"`
}

type Person struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type Filters struct {
	Mark     string
	Model    string
	Year     int
	Page     int
	PageSize int
}
