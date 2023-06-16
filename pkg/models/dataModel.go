package models

type Data struct {
	ID           int    `json:"id"`
	Database_URI string `json:"databaseURL"`
	DatabaseName string `json:"databaseName"`
}
