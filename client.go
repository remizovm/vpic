package vpic

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	endpoint = `https://vpic.nhtsa.dot.gov/api`
)

type Client struct {
	HTTPClient http.Client
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
