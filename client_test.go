package vpic

import (
	"context"
	"testing"

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
	})
}
