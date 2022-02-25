// Copyright 2011 The Graphics-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphics

import (
	"errors"
	"image"
	"image/draw"
	"math"

	"github.com/grokify/graphics-go/graphics/convolve"
)

// DefaultStdDev is the default blurring parameter.
const DefaultStdDev = 0.5
const DefaultSize = 0

// blurOptions are the blurring parameters.
// StdDev is the standard deviation of the normal, higher is blurrier.
// Size is the size of the kernel. If zero, it is set to Ceil(6 * StdDev).
type blurOptions struct {
	StdDev float64
	Size   int
}

// Blur produces a blurred version of the image, using a Gaussian blur.
// `stdDev`` is the standard deviation of the normal, higher is blurrier.
// For default value, use `DefaultStdDev` or `0.5`.
// `size`` is the size of the kernel. If zero, it is set to Ceil(6 * StdDev).
func Blur(dst draw.Image, src image.Image, stdDev float64, size int) error {
	if dst == nil {
		return errors.New("graphics: dst is nil")
	}
	if src == nil {
		return errors.New("graphics: src is nil")
	}

	if size < 1 {
		size = int(math.Ceil(stdDev * 6))
	}

	kernel := make([]float64, 2*size+1)
	for i := 0; i <= size; i++ {
		x := float64(i) / stdDev
		x = math.Pow(1/math.SqrtE, x*x)
		kernel[size-i] = x
		kernel[size+i] = x
	}

	// Normalize the weights to sum to 1.0.
	kSum := 0.0
	for _, k := range kernel {
		kSum += k
	}
	for i, k := range kernel {
		kernel[i] = k / kSum
	}

	return convolve.Convolve(dst, src, &convolve.SeparableKernel{
		X: kernel,
		Y: kernel,
	})
}
