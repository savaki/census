package names

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsX(t *testing.T) {
	Convey("When I call IsMale with a Male name", t, func() {
		freq, ok := IsMale("andres")

		Convey("Then I expect it to be found", func() {
			So(ok, ShouldBeTrue)
		})

		Convey("And I expect a freq above 0", func() {
			So(freq, ShouldBeGreaterThan, 0.0)
		})
	})

	Convey("When I call IsMale with a Female name", t, func() {
		_, ok := IsMale("Sameera")

		Convey("Then I expect it to NOT be found", func() {
			So(ok, ShouldBeFalse)
		})
	})

	Convey("When I call IsFemale with a Female name", t, func() {
		_, ok := IsFemale("Sameera")

		Convey("Then I expect it to NOT be found", func() {
			So(ok, ShouldBeFalse)
		})
	})

	Convey("When I call IsFemale with a Female name", t, func() {
		freq, ok := IsFemale("Susan")

		Convey("Then I expect it to be found", func() {
			So(ok, ShouldBeTrue)
		})

		Convey("And I expect a freq above 0", func() {
			So(freq, ShouldBeGreaterThan, 0.0)
		})
	})

	Convey("When I call IsSurname with a real Surname", t, func() {
		freq, ok := IsSurname("Smith")

		Convey("Then I expect it to be found", func() {
			So(ok, ShouldBeTrue)
		})

		Convey("And I expect a freq above 0", func() {
			So(freq, ShouldBeGreaterThan, 0.0)
		})
	})
}
