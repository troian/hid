//go:build hidraw && cgo
// +build hidraw,cgo

package hid

import (
	_ "github.com/zondax/hid/hidapi"
	_ "github.com/zondax/hid/hidapi/linux"
)
