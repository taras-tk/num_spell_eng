package num_spell_eng

import "testing"
import . "github.com/smartystreets/goconvey/convey"


func TestConvert(t *testing.T) {
	Convey("Spell number", t, func() {
		Convey("Number less than twenty", func() {
			So(NumberSpell(0), ShouldEqual, "zero")
			So(NumberSpell(1), ShouldEqual, "one")
			So(NumberSpell(8), ShouldEqual, "eight")
			So(NumberSpell(10), ShouldEqual, "ten")
			So(NumberSpell(13), ShouldEqual, "thirteen")
			So(NumberSpell(19), ShouldEqual, "nineteen")
		})
		Convey("Number less than one hundred", func() {
			So(NumberSpell(20), ShouldEqual, "twenty")
			So(NumberSpell(30), ShouldEqual, "thirty")
			So(NumberSpell(40), ShouldEqual, "forty")
			So(NumberSpell(50), ShouldEqual, "fifty")
			So(NumberSpell(70), ShouldEqual, "seventy")
			So(NumberSpell(90), ShouldEqual, "ninety")
		})
		Convey("Two word number", func() {
			So(NumberSpell(23), ShouldEqual, "twenty-three")
			So(NumberSpell(31), ShouldEqual, "thirty-one")
			So(NumberSpell(42), ShouldEqual, "forty-two")
			So(NumberSpell(54), ShouldEqual, "fifty-four")
			So(NumberSpell(99), ShouldEqual, "ninety-nine")
		})
		Convey("Bigger numbers", func() {
			So(NumberSpell(100), ShouldEqual, "one hundred")
			So(NumberSpell(200), ShouldEqual, "two hundred")
			So(NumberSpell(300), ShouldEqual, "three hundred")
			So(NumberSpell(123), ShouldEqual, "one hundred and twenty-three")
			So(NumberSpell(999), ShouldEqual, "nine hundred and ninety-nine")
			So(NumberSpell(1028), ShouldEqual, "one thousand, twenty-eight")
			So(NumberSpell(1100123), ShouldEqual, "one million, one hundred thousand, one hundred and twenty-three")
		})
		Convey("Negative numbers", func() {
			So(NumberSpell(-123), ShouldEqual, "minus one hundred and twenty-three")
		})
	})
}

