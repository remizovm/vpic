package vpic

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	endpoint = `https://vpic.nhtsa.dot.gov/api`
)

type EquipmentType int

const (
	EquipmentTypeTires      EquipmentType = 1
	EquipmentTypeBrakeHoses EquipmentType = 3
	EquipmentTypeGlazing    EquipmentType = 13
	EquipmentTypeRetread    EquipmentType = 16
)

type ReportType string

const (
	ReportTypeNew     ReportType = "New"
	ReportTypeUpdated ReportType = "Updated"
	ReportTypeClosed  ReportType = "Closed"
	ReportTypeAll     ReportType = "All"
)

type Units string

const (
	UnitsMetric Units = "Metric"
	UnitsUS     Units = "US"
)

var (
	ErrYearInvalid = errors.New("year is invalid")
	ErrArgsInvalid = errors.New("arguments are invalid")
)

type Client struct {
	HTTPClient http.Client
}

func (c Client) ModelsByMakeIDAndYearVehicleType(ctx context.Context, year int, id int64, vehicleType string) ([]Model, error) {
	if year < 1995 {
		return nil, ErrYearInvalid
	}
	var uri string
	switch {
	case year != 0 && vehicleType != "":
		uri = endpoint + "/vehicles/GetModelsForMakeIdYear/makeId/" + strconv.FormatInt(id, 10) + "/modelyear/" + strconv.Itoa(year) + "/vehicletype/" + vehicleType + "?format=json"
	case year != 0 && vehicleType == "":
		uri = endpoint + "/vehicles/GetModelsForMakeIdYear/makeId/" + strconv.FormatInt(id, 10) + "/modelyear/" + strconv.Itoa(year) + "?format=json"
	case year == 0 && vehicleType != "":
		uri = endpoint + "/vehicles/GetModelsForMakeIdYear/makeId/" + strconv.FormatInt(id, 10) + "/vehicletype/" + vehicleType + "?format=json"
	default:
		return nil, ErrArgsInvalid
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int     `json:"count"`
		Message        string  `json:"message"`
		SearchCriteria string  `json:"SearchCriteria"`
		Results        []Model `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) ModelsByMakeAndYearVehicleType(ctx context.Context, year int, makeName, vehicleType string) ([]Model, error) {
	if year < 1995 {
		return nil, ErrYearInvalid
	}
	var uri string
	switch {
	case year != 0 && vehicleType != "":
		uri = endpoint + "/vehicles/GetModelsForMakeYear/make/" + makeName + "/modelyear/" + strconv.Itoa(year) + "/vehicletype/" + vehicleType + "?format=json"
	case year != 0 && vehicleType == "":
		uri = endpoint + "/vehicles/GetModelsForMakeYear/make/" + makeName + "/modelyear/" + strconv.Itoa(year) + "?format=json"
	case year == 0 && vehicleType != "":
		uri = endpoint + "/vehicles/GetModelsForMakeYear/make/" + makeName + "/vehicletype/" + vehicleType + "?format=json"
	default:
		return nil, ErrArgsInvalid
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int     `json:"count"`
		Message        string  `json:"message"`
		SearchCriteria string  `json:"SearchCriteria"`
		Results        []Model `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) DecodeVINFlatBatch(ctx context.Context, request []*VINBatchRequest) ([]map[string]string, error) {
	var rawReq []string
	for _, r := range request {
		data := r.VIN
		if r.Year != 0 {
			data = data + "," + strconv.Itoa(r.Year)
		}
		rawReq = append(rawReq, data)
	}
	payload := strings.Join(rawReq, ";")

	data := url.Values{}
	data.Set("data", payload)
	data.Set("format", "json")

	uri := endpoint + "/vehicles/DecodeVINValuesBatch/"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int                 `json:"count"`
		Message        string              `json:"message"`
		SearchCriteria string              `json:"SearchCriteria"`
		Results        []map[string]string `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil

	return nil, nil
}

func (c Client) CanadianVehicleSpecs(ctx context.Context, year int, makeName, model string, units *Units) ([]Spec, error) {
	if year <= 1971 {
		return nil, ErrYearInvalid
	}
	values := url.Values{}
	values.Set("Year", strconv.Itoa(year))
	values.Set("Make", makeName)
	values.Set("Model", model)
	values.Set("Units", "")
	if units != nil {
		values.Set("Units", string(*units))
	}
	values.Set("format", "json")

	url := endpoint + "/vehicles/GetCanadianVehicleSpecifications/?" + values.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Spec `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) EquipmentPlantCodes(ctx context.Context, year int, equipmentType EquipmentType, reportType ReportType) ([]EquipmentPlantCode, error) {
	if year > 2016 {
		return nil, ErrYearInvalid
	}
	values := url.Values{}
	values.Set("year", strconv.Itoa(year))
	values.Set("equipmentType", strconv.Itoa(int(equipmentType)))
	values.Set("reportType", string(reportType))
	values.Set("format", "json")

	url := endpoint + "/vehicles/GetEquipmentPlantCodes?" + values.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int                  `json:"count"`
		Message        string               `json:"message"`
		SearchCriteria string               `json:"SearchCriteria"`
		Results        []EquipmentPlantCode `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) MakesByManufacturerNameAndYear(ctx context.Context, name string, year int) ([]Make, error) {
	values := url.Values{}
	values.Set("year", strconv.Itoa(year))
	values.Set("format", "json")

	url := endpoint + "/vehicles/GetMakesForManufacturerAndYear/" + name + "?" + values.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Make `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) MakesByManufacturerIDAndYear(ctx context.Context, id int64, year int) ([]Make, error) {
	values := url.Values{}
	values.Set("year", strconv.Itoa(year))
	values.Set("format", "json")

	url := endpoint + "/vehicles/GetMakesForManufacturerAndYear/" + strconv.FormatInt(id, 10) + "?" + values.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Make `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) MakesByManufacturerID(ctx context.Context, id int64) ([]Make, error) {
	url := endpoint + "/vehicles/GetMakeForManufacturer/" + strconv.FormatInt(id, 10) + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Make `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) MakesByManufacturerName(ctx context.Context, name string) ([]Make, error) {
	url := endpoint + "/vehicles/GetMakeForManufacturer/" + name + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Make `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) Manufacturers(ctx context.Context, mType string, page int) ([]Manufacturer, error) {
	values := url.Values{}
	values.Set("format", "json")
	if mType != "" {
		values.Set("ManufacturerType", mType)
	}
	if page != 0 {
		values.Set("page", strconv.Itoa(page))
	}
	uri := endpoint + "/vehicles/GetAllManufacturers?" + values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int            `json:"count"`
		Message        string         `json:"message"`
		SearchCriteria string         `json:"SearchCriteria"`
		Results        []Manufacturer `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) GetParts(ctx context.Context, partType int64, dtFrom, dtTo time.Time, page int) ([]Part, error) {
	values := url.Values{}
	values.Set("format", "json")
	values.Set("type", strconv.FormatInt(partType, 10))
	values.Set("fromDate", dtFrom.Format("1/2/2006"))
	values.Set("toDate", dtTo.Format("1/2/2006"))
	if page != 0 {
		values.Set("page", strconv.Itoa(page))
	}
	uri := endpoint + "/vehicles/GetParts?" + values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Part `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) MakesByVehicleTypeName(ctx context.Context, name string) ([]Make, error) {
	url := endpoint + "/vehicles/GetMakesForVehicleType/" + name + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Make `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) VehicleTypesByMakeID(ctx context.Context, id int64) ([]VehicleType, error) {
	url := endpoint + "/vehicles/GetVehicleTypesForMakeId/" + strconv.FormatInt(id, 10) + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int           `json:"count"`
		Message        string        `json:"message"`
		SearchCriteria string        `json:"SearchCriteria"`
		Results        []VehicleType `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) VehicleTypesByMake(ctx context.Context, name string) ([]VehicleType, error) {
	url := endpoint + "/vehicles/GetVehicleTypesForMake/" + name + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int           `json:"count"`
		Message        string        `json:"message"`
		SearchCriteria string        `json:"SearchCriteria"`
		Results        []VehicleType `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) ModelsByMake(ctx context.Context, name string) ([]Model, error) {
	url := endpoint + "/vehicles/getmodelsformake/" + name + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int     `json:"count"`
		Message        string  `json:"message"`
		SearchCriteria string  `json:"SearchCriteria"`
		Results        []Model `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) ModelsByMakeID(ctx context.Context, id int64) ([]Model, error) {
	url := endpoint + "/vehicles/GetModelsForMakeId/" + strconv.FormatInt(id, 10) + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int     `json:"count"`
		Message        string  `json:"message"`
		SearchCriteria string  `json:"SearchCriteria"`
		Results        []Model `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) ManufacturerDetailsByID(ctx context.Context, id int64) ([]Manufacturer, error) {
	url := endpoint + "/vehicles/getmanufacturerdetails/" + strconv.FormatInt(id, 10) + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int            `json:"count"`
		Message        string         `json:"message"`
		SearchCriteria string         `json:"SearchCriteria"`
		Results        []Manufacturer `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) ManufacturerDetailsByName(ctx context.Context, name string) ([]Manufacturer, error) {
	url := endpoint + "/vehicles/getmanufacturerdetails/" + name + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int            `json:"count"`
		Message        string         `json:"message"`
		SearchCriteria string         `json:"SearchCriteria"`
		Results        []Manufacturer `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) Makes(ctx context.Context) ([]Make, error) {
	url := endpoint + "/vehicles/getallmakes?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []Make `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) GetWMIList(ctx context.Context, manufacturer string) ([]WMI, error) {
	url := endpoint + "/vehicles/GetWMIsForManufacturer/" + manufacturer + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int    `json:"count"`
		Message        string `json:"message"`
		SearchCriteria string `json:"SearchCriteria"`
		Results        []WMI  `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) DecodeWMI(ctx context.Context, wmi string) ([]DecodeWMIResult, error) {
	url := endpoint + "/vehicles/decodewmi/" + wmi + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int               `json:"count"`
		Message        string            `json:"message"`
		SearchCriteria string            `json:"SearchCriteria"`
		Results        []DecodeWMIResult `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) VehicleVariablesList(ctx context.Context) ([]VehicleVariable, error) {
	url := endpoint + "/vehicles/getvehiclevariablelist?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int               `json:"count"`
		Message        string            `json:"message"`
		SearchCriteria string            `json:"SearchCriteria"`
		Results        []VehicleVariable `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) VehicleVariableValuesListByID(ctx context.Context, id int) ([]VehicleVariableValues, error) {
	url := endpoint + "/vehicles/getvehiclevariablevalueslist/" + strconv.Itoa(id) + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int                     `json:"count"`
		Message        string                  `json:"message"`
		SearchCriteria string                  `json:"SearchCriteria"`
		Results        []VehicleVariableValues `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) VehicleVariableValuesListByName(ctx context.Context, name string) ([]VehicleVariableValues, error) {
	url := endpoint + "/vehicles/getvehiclevariablevalueslist/" + name + "?format=json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int                     `json:"count"`
		Message        string                  `json:"message"`
		SearchCriteria string                  `json:"SearchCriteria"`
		Results        []VehicleVariableValues `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) DecodeVINExtendedFlat(ctx context.Context, vin string, modelyear int) (map[string]string, error) {
	url := endpoint + "/vehicles/DecodeVinValuesExtended/" + vin + "?format=json"
	if modelyear != 0 {
		url = url + "&modelyear=" + strconv.Itoa(modelyear)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int                 `json:"count"`
		Message        string              `json:"message"`
		SearchCriteria string              `json:"SearchCriteria"`
		Results        []map[string]string `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results[0], nil
}

func (c Client) DecodeVINExtended(ctx context.Context, vin string, modelyear int) ([]DecodeVINResult, error) {
	url := endpoint + "/vehicles/DecodeVinExtended/" + vin + "?format=json"
	if modelyear != 0 {
		url = url + "&modelyear=" + strconv.Itoa(modelyear)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int               `json:"count"`
		Message        string            `json:"message"`
		SearchCriteria string            `json:"SearchCriteria"`
		Results        []DecodeVINResult `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (c Client) DecodeVINFlat(ctx context.Context, vin string, modelyear int) (map[string]string, error) {
	url := endpoint + "/vehicles/DecodeVinValues/" + vin + "?format=json"
	if modelyear != 0 {
		url = url + "&modelyear=" + strconv.Itoa(modelyear)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int                 `json:"count"`
		Message        string              `json:"message"`
		SearchCriteria string              `json:"SearchCriteria"`
		Results        []map[string]string `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results[0], nil
}

func (c Client) DecodeVIN(ctx context.Context, vin string, modelyear int) ([]DecodeVINResult, error) {
	url := endpoint + "/vehicles/decodevin/" + vin + "?format=json"
	if modelyear != 0 {
		url = url + "&modelyear=" + strconv.Itoa(modelyear)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Count          int               `json:"count"`
		Message        string            `json:"message"`
		SearchCriteria string            `json:"SearchCriteria"`
		Results        []DecodeVINResult `json:"Results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}
