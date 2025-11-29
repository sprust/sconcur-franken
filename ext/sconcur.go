package sconcur

// #include <stdlib.h>
// #include "sconcur.h"
import "C"
import (
	"unsafe"

	"github.com/dunglas/frankenphp"
	"github.com/sprust/sconcur-franken/internal/services/ping_service"
)

func init() {
	frankenphp.RegisterExtension(unsafe.Pointer(&C.sconcur_module_entry))
}

//export ping
func ping(s *C.zend_string) unsafe.Pointer {
	str := frankenphp.GoString(unsafe.Pointer(s))

	result := ping_service.Ping(str)

	return frankenphp.PHPString(result, false)
}
