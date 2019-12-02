package vpic

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
