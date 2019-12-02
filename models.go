package vpic

type Model struct {
	MakeID    int64  `json:"Make_ID"`
	MakeName  string `json:"Make_Name"`
	ModelID   int64  `json:"Model_ID"`
	ModelName string `json:"Model_Name"`
}

type EquipmentItem struct{}

type ManufacturerType struct{}

type VehicleType struct {
	MakeID   int64  `json:"MakeId"`
	MakeName string `json:"MakeName"`
	ID       int64  `json:"VehicleTypeId"`
	Name     string `json:"VehicleTypeName"`
}

type Manufacturer struct {
	Address            string             `json:"Address"`
	Address2           string             `json:"Address2"`
	City               string             `json:"City"`
	Email              string             `json:"ContactEmail"`
	Fax                string             `json:"ContactFax"`
	Phone              string             `json:"ContactPhone"`
	Country            string             `json:"Country"`
	EquipmentItems     []EquipmentItem    `json:"EquipmentItems"`
	LastUpdated        string             `json:"LastUpdated"`
	ManufacturerTypes  []ManufacturerType `json:"ManufacturerTypes"`
	CommonName         string             `json:"Mfr_CommonName"`
	ID                 int64              `json:"Mfr_ID"`
	Name               string             `json:"Mfr_Name"`
	Other              string             `json:"OtherManufacturerDetails"`
	PostalCode         string             `json:"PostalCode"`
	PrimaryProduct     string             `json:"PrimaryProduct"`
	PrincipalFirstName string             `json:"PrincipalFirstName"`
	PrincipalLastName  string             `json:"PrincipalLastName"`
	PrincipalPosition  string             `json:"PrincipalPosition"`
	StateProvince      string             `json:"StateProvince"`
	SubmittedName      string             `json:"SubmittedName"`
	SubmittedOn        string             `json:"SubmittedOn"`
	SubmittedPosition  string             `json:"SubmittedPosition"`
	VehicleTypes       []VehicleType      `json:"VehicleTypes"`
}

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
