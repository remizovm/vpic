package vpic

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
