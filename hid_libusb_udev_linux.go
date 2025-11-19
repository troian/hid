//go:build !hidraw && !android && cgo
// +build !hidraw,!android,cgo

package hid

/*
#include "hidapi/libusb/hidapi_libusb.h"
*/
import "C"

import (
	_ "github.com/zondax/hid/libusb/libusb/os/udev"
)
