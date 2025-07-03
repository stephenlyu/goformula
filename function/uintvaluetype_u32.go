//go:build !u16
// +build !u16

package function

import "github.com/z-ray/log"

type UintVecValueType = uint32

func init() {
	log.Info("UintVecValueType = uint32")
}
