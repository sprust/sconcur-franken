//go:build franken
// +build franken

// export_php:namespace SConcur\Extension
package sconcur

// #include <Zend/zend_types.h>
import "C"

import (
	"unsafe"

	"github.com/dunglas/frankenphp"
	"github.com/sprust/sconcur-franken/internal/services/features/ping_feature"
	"github.com/sprust/sconcur-franken/internal/services/flow"
)

var flowService *flow.Service

func init() {
	flowService = flow.NewFlow()
}

// export_php:function ping(string $string): string
func ping(tId *C.zend_string, s *C.zend_string) unsafe.Pointer {
	str := frankenphp.GoString(unsafe.Pointer(s))

	flowService.
		result := ping_feature.Ping(str)

	return frankenphp.PHPString(result, false)
}

// export_php:method Flow::usleep(milliseconds int): string
func usleep(tId *C.zend_string, milliseconds int64) {
	taskId := frankenphp.GoString(unsafe.Pointer(tId))

	flowService.Usleep(taskId, milliseconds)
}
