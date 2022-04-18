// Copyright 2011 The Graphics-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphics

import (
	"errors"
	"image"
	"image/draw"

	"github.com/grokify/graphics-go/graphics/interp"
)

// rotateOptions are the rotation parameters.
// Angle is the angle, in radians, to rotate the image clockwise.
type rotateOptions struct {
	Angle float64
}

// Rotate produces a rotated version of src, drawn onto dst.
// `angle` is in radians, to rotate the image clockwise.
func Rotate(dst draw.Image, src image.Image, angle float64) error {
	if dst == nil {
		return errors.New("graphics: dst is nil")
	}
	if src == nil {
		return errors.New("graphics: src is nil")
	}

	return I.Rotate(angle).TransformCenter(dst, src, interp.Bilinear)
}
