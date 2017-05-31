package style

import (
	"strconv"

	"github.com/youryharchenko/vecty"
)

type Size string

func Px(pixels int) Size {
	return Size(strconv.Itoa(pixels) + "px")
}

func Color(value string) vecty.Markup {
	return vecty.Style("color", value)
}

func Width(size Size) vecty.Markup {
	return vecty.Style("width", string(size))
}

func MinWidth(size Size) vecty.Markup {
	return vecty.Style("min-width", string(size))
}

func MaxWidth(size Size) vecty.Markup {
	return vecty.Style("max-width", string(size))
}

func Height(size Size) vecty.Markup {
	return vecty.Style("height", string(size))
}

func MinHeight(size Size) vecty.Markup {
	return vecty.Style("min-height", string(size))
}

func MaxHeight(size Size) vecty.Markup {
	return vecty.Style("max-height", string(size))
}

func Margin(size Size) vecty.Markup {
	return vecty.Style("margin", string(size))
}

type OverflowOption string

const (
	OverflowVisible OverflowOption = "visible"
	OverflowHidden  OverflowOption = "hidden"
	OverflowScroll  OverflowOption = "scroll"
	OverflowAuto    OverflowOption = "auto"
)

func Overflow(option OverflowOption) vecty.Markup {
	return vecty.Style("overflow", string(option))
}

func OverflowX(option OverflowOption) vecty.Markup {
	return vecty.Style("overflow-x", string(option))
}

func OverflowY(option OverflowOption) vecty.Markup {
	return vecty.Style("overflow-y", string(option))
}
