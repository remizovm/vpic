package vpic

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
