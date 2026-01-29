package util

import (
	"github.com/applejag/firefly-go-math/ffrand"
)

func RandomSliceElem[E any](slice []E) E {
	return slice[ffrand.Intn(len(slice))]
}
