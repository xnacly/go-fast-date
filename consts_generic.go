//go:build !arm

package gofastdate

const (
	eRAS    uint64 = 4726498270 // Use 14704 for 32–bit
	d_shift uint64 = 146097*eRAS - 719469
	y_shift uint64 = 400*eRAS - 1
	c1      uint64 = 505054698555331   // floor(2^64*4/146097)
	c2      uint64 = 50504432782230121 // ceil(2^64*4/1461)
	c3      uint64 = 8619973866219416  // floor(2^64/2140)
	scale   uint64 = 32                // SCALE = 32 for non-ARM builds
	shift_0 uint64 = 30556 * 32        // 30556 * SCALE = 977792
	shift_1 uint64 = 5980 * 32         // 5980  * SCALE = 191360
)
