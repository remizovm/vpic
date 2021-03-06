// +build integration

package vpic

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestModelsByMakeIDAndYearVehicleType(t *testing.T) {
	Convey("ModelsByMakeIDAndYearVehicleType method", t, func() {
		Convey("should return "+ErrArgsInvalid.Error(), func() {
			c := Client{}
			resp, err := c.ModelsByMakeIDAndYearVehicleType(context.Background(), 0, 474, "")
			So(err, ShouldEqual, ErrArgsInvalid)
			So(resp, ShouldBeNil)
		})
		Convey("should return "+ErrYearInvalid.Error(), func() {
			c := Client{}
			resp, err := c.ModelsByMakeIDAndYearVehicleType(context.Background(), 1994, 474, "")
			So(err, ShouldEqual, ErrYearInvalid)
			So(resp, ShouldBeNil)
		})
		Convey("should work properly in normal case", func() {
			c := Client{}
			resp, err := c.ModelsByMakeIDAndYearVehicleType(context.Background(), 2015, 474, "")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestModelsByMakeAndYearVehicleType(t *testing.T) {
	Convey("ModelsByMakeAndYearVehicleType method", t, func() {
		Convey("Should return "+ErrArgsInvalid.Error(), func() {
			c := Client{}
			resp, err := c.ModelsByMakeAndYearVehicleType(context.Background(), 0, "honda", "")
			So(err, ShouldEqual, ErrArgsInvalid)
			So(resp, ShouldBeNil)
		})
		Convey("Should return "+ErrYearInvalid.Error(), func() {
			c := Client{}
			resp, err := c.ModelsByMakeAndYearVehicleType(context.Background(), 1994, "honda", "")
			So(err, ShouldEqual, ErrYearInvalid)
			So(resp, ShouldBeNil)
		})
		Convey("should work properly in normal case", func() {
			c := Client{}
			resp, err := c.ModelsByMakeAndYearVehicleType(context.Background(), 2015, "honda", "")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestDecodeVINFlatBatch(t *testing.T) {
	Convey("DecodeVINFlatBatch method", t, func() {
		c := Client{}
		req := []*VINBatchRequest{
			&VINBatchRequest{
				VIN:  "5UXWX7C5*BA",
				Year: 2011,
			},
			&VINBatchRequest{
				VIN: "5YJSA3DS*EF",
			},
		}
		resp, err := c.DecodeVINFlatBatch(context.Background(), req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestCanadianSpecs(t *testing.T) {
	Convey("CanadianSpecs method", t, func() {
		Convey("should return "+ErrYearInvalid.Error(), func() {
			c := Client{}
			resp, err := c.CanadianVehicleSpecs(context.Background(), 1960, "Acura", "", nil)
			So(err, ShouldEqual, ErrYearInvalid)
			So(resp, ShouldBeNil)
		})
		Convey("should work properly in normal case", func() {
			c := Client{}
			resp, err := c.CanadianVehicleSpecs(context.Background(), 2011, "Acura", "", nil)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestEquipmentPlantCodes(t *testing.T) {
	Convey("EquipmentPlantCodes method", t, func() {
		Convey("should return "+ErrYearInvalid.Error(), func() {
			c := Client{}
			resp, err := c.EquipmentPlantCodes(context.Background(), 2017, EquipmentTypeTires, ReportTypeAll)
			So(err, ShouldEqual, ErrYearInvalid)
			So(resp, ShouldBeNil)
		})
		Convey("should work properly in normal case", func() {
			c := Client{}
			resp, err := c.EquipmentPlantCodes(context.Background(), 2015, EquipmentTypeTires, ReportTypeAll)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestMakesByManufacturerByYear(t *testing.T) {
	Convey("MakesByManufacturerByYear", t, func() {
		Convey("and ByName method", func() {
			c := Client{}
			resp, err := c.MakesByManufacturerNameAndYear(context.Background(), "mer", 2013)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
		Convey("and ByID method", func() {
			c := Client{}
			resp, err := c.MakesByManufacturerIDAndYear(context.Background(), 988, 2013)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestMakesByManufacturer(t *testing.T) {
	Convey("MakesByManufacturer", t, func() {
		Convey("ByName method", func() {
			c := Client{}
			resp, err := c.MakesByManufacturerName(context.Background(), "honda")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
		Convey("ByID method", func() {
			c := Client{}
			resp, err := c.MakesByManufacturerID(context.Background(), 988)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestManufacturers(t *testing.T) {
	Convey("Manufacturers method", t, func() {
		c := Client{}
		resp, err := c.Manufacturers(context.Background(), "", 0)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestGetParts(t *testing.T) {
	Convey("GetParts method", t, func() {
		c := Client{}
		dtFrom := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
		dtTo := time.Date(2015, time.May, 5, 0, 0, 0, 0, time.UTC)
		resp, err := c.GetParts(context.Background(), 565, dtFrom, dtTo, 0)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestMakesByVehicleTypeName(t *testing.T) {
	Convey("MakesByVehicleTypeName method", t, func() {
		c := Client{}
		resp, err := c.MakesByVehicleTypeName(context.Background(), "car")
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestVehicleTypesByMakeID(t *testing.T) {
	Convey("VehicleTypesByMakeID method", t, func() {
		c := Client{}
		resp, err := c.VehicleTypesByMakeID(context.Background(), 450)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestVehicleTypesByMake(t *testing.T) {
	Convey("VehicleTypesByMake method", t, func() {
		c := Client{}
		resp, err := c.VehicleTypesByMake(context.Background(), "merc")
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestModelsByMake(t *testing.T) {
	Convey("ModelsByMake method", t, func() {
		c := Client{}
		resp, err := c.ModelsByMake(context.Background(), "honda")
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestModelsByMakeID(t *testing.T) {
	Convey("ModelsByMakeID method", t, func() {
		c := Client{}
		resp, err := c.ModelsByMakeID(context.Background(), 440)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestManufacturerDetails(t *testing.T) {
	Convey("ManufacturerDetails", t, func() {
		Convey("ByName method", func() {
			c := Client{}
			resp, err := c.ManufacturerDetailsByName(context.Background(), "honda")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
		Convey("ByID method", func() {
			c := Client{}
			resp, err := c.ManufacturerDetailsByID(context.Background(), 989)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestMakes(t *testing.T) {
	Convey("Makes method", t, func() {
		c := Client{}
		resp, err := c.Makes(context.Background())
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestGetWMIList(t *testing.T) {
	Convey("GetWMIList method", t, func() {
		c := Client{}
		resp, err := c.GetWMIList(context.Background(), "hon")
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestDecodeWMI(t *testing.T) {
	Convey("DecodeWMI method", t, func() {
		c := Client{}
		resp, err := c.DecodeWMI(context.Background(), "1FD")
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestVehicleVariblesList(t *testing.T) {
	Convey("VehicleVariablesList method", t, func() {
		c := Client{}
		resp, err := c.VehicleVariablesList(context.Background())
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestVehicleVarialbeValuesList(t *testing.T) {
	Convey("VehicleVariableValues", t, func() {
		Convey("by name", func() {
			c := Client{}
			resp, err := c.VehicleVariableValuesListByName(context.Background(), "battery%20type")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
		Convey("by id", func() {
			c := Client{}
			resp, err := c.VehicleVariableValuesListByID(context.Background(), 2)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeEmpty)
		})
	})
}

func TestDecodeVINExtendedFlat(t *testing.T) {
	Convey("DecodeVINFlat method", t, func() {
		c := Client{}
		resp, err := c.DecodeVINExtendedFlat(context.Background(), "5UXWX7C5*BA", 2011)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestDecodeVINFlat(t *testing.T) {
	Convey("DecodeVINFlat method", t, func() {
		c := Client{}
		resp, err := c.DecodeVINFlat(context.Background(), "5UXWX7C5*BA", 2011)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestDecodeVINExtended(t *testing.T) {
	Convey("DecodeVINExtended method", t, func() {
		c := Client{}
		resp, err := c.DecodeVINExtended(context.Background(), "5UXWX7C5*BA", 2011)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}

func TestDecodeVIN(t *testing.T) {
	Convey("DecodeVIN method", t, func() {
		c := Client{}
		resp, err := c.DecodeVIN(context.Background(), "5UXWX7C5*BA", 2011)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeEmpty)
	})
}
