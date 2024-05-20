package operations_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/bradlsum/gscale/internal/operations"
)

func TestIntegerScale(t *testing.T) {

}

func TestImage2X(t *testing.T) {

}

func TestFindNearestColor(t *testing.T) {
	r := image.Rectangle{}
	img := image.NewNRGBA(r)
	c := operations.Coords{X: 0, Y: 0, C: color.Black}
	m := make(map[string]bool)
	operations.FindNearestColor(img, c, 0, m)
}
