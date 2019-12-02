package vpic

type Make struct {
	ID   int64  `json:"Make_ID"`
	Name string `json:"Make_Name"`
}

type WMI struct {
	CanceledDate string `json:"CanceledDate"`
	CreatedOn    string `json:"CreatedOn"`
	UpdatedOn    string `json:"UpdatedOn"`
	VehicleType  string `json:"VehicleType"`
	Country      string `json:"Country"`
	ID           int    `json:"Id"`
	Name         string `json:"Name"`
	WMI          string `json:"WMI"`
}

type DecodeWMIResult struct {
	CanceledDate          string `json:"CanceledDate"`
	CommonName            string `json:"CommonName"`
	CreatedOn             string `json:"CreatedOn"`
	DateAvailableToPublic string `json:"DateAvailableToPublic"`
	Make                  string `json:"Make"`
	ManufacturerName      string `json:"ManufacturerName"`
	ParentCompanyName     string `json:"ParentCompanyName"`
	URL                   string `json:"URL"`
	UpdatedOn             string `json:"UpdatedOn"`
	VehicleType           string `json:"VehicleType"`
}

type VehicleVariable struct {
	Name        string `json:"Name"`
	DataType    string `json:"DataType"`
	Description string `json:"Description"`
	ID          int    `json:"Id"`
}

type VehicleVariableValues struct {
	ElementName string `json:"ElementName"`
	ID          int    `json:"Id"`
	Name        string `json:"Name"`
}

type DecodeVINResult struct {
	Value      string `json:"Value"`
	ValueID    string `json:"ValueId"`
	Variable   string `json:"Variable"`
	VariableId int    `json:"VariableId"`
}
