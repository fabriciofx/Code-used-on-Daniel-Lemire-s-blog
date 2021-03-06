// +build amd64,!appengine

package ctz

import "github.com/klauspost/cpuid"

var useSSE42 bool

// Detect SSE 4.2 feature.
func init() {
	useSSE42 = cpuid.CPU.SSE42()
}

// countTrailingZeros counts the number of zeros
// from the least-significant bit up to the
// first set (1) bit. if x is 0, 64 is returned.
//
// references:
// a. https://en.wikipedia.org/wiki/Find_first_set
// b. TZCNTQ on amd64, page 364 of http://support.amd.com/TechDocs/24594.pdf
//
// *** the following function is defined in ctz_amd64.s
//
// TODO: possibly use "github.com/klauspost/cpuid"
// to check if cpuid.CPU.BMI1() is true before using the assembly version.
//
// The Go version is in ctz_generic.go.
//
func countTrailingZeros(x uint64) int
