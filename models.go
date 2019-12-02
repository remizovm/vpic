package vpic

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
