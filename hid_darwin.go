//go:build !ios && cgo
// +build !ios,cgo

package hid

import (
	_ "github.com/zondax/hid/hidapi"
	_ "github.com/zondax/hid/hidapi/mac"
)
