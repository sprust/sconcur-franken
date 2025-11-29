//go:build franken
// +build franken

//export_php:namespace SConcur\Extension
package sconcur

// #include <Zend/zend_types.h>
import "C"

import (
	"unsafe"

	"github.com/sprust/sconcur-franken/internal/services/ping_service"

	"github.com/dunglas/frankenphp"
)

//export_php:function ping(string $string): string
func ping(s *C.zend_string) unsafe.Pointer {
	str := frankenphp.GoString(unsafe.Pointer(s))

	result := ping_service.Ping(str)

	return frankenphp.PHPString(result, false)
}
