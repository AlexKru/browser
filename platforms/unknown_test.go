package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUnknown(t *testing.T) {
	Convey("Subject: #NewUnknown", t, func() {
		Convey("It should return a new Unknown instance", func() {
			So(NewUnknown(""), ShouldHaveSameTypeAs, &Unknown{})
		})
	})
}

func TestUnknownName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Unknown", func() {
			u := NewUnknown("")
			So(u.Name(), ShouldEqual, "Unknown")
		})
	})
}

func TestUnknownVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("It should return 0", func() {
			So(NewUnknown("").Version(), ShouldEqual, "0")
		})
	})
}

func TestUnknownMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("It should return true", func() {
			So(NewUnknown("").Match(), ShouldBeTrue)
		})
	})
}
