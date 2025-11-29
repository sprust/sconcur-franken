//export_php:namespace SConcur\Extension
//go:build franken
// +build franken

package adapter

// #include <Zend/zend_types.h>
import "C"

import (
	"sconcur/internal/services/ping_service"
	"unsafe"

	"github.com/dunglas/frankenphp"
)

//export_php:function ping(): string
func ping(s *C.zend_string) unsafe.Pointer {
	str := frankenphp.GoString(unsafe.Pointer(s))

	result := ping_service.Ping(str)

	return frankenphp.PHPString(result, false)
}
