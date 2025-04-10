package models

type Report struct {
	ReportId    string `json:"report_id"`
	ReportType  string `json:"report_type"`
	Description string `json:"description"`
}
