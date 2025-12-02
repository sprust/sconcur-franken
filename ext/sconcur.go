package sconcur

// #include <stdlib.h>
// #include "sconcur.h"
import "C"
import (
	"unsafe"

	"github.com/dunglas/frankenphp"
	"github.com/sprust/sconcur-franken/internal/services/features/ping_feature"
)

func init() {
	frankenphp.RegisterExtension(unsafe.Pointer(&C.sconcur_module_entry))
}

//export ping
func ping(s *C.zend_string) unsafe.Pointer {
	str := frankenphp.GoString(unsafe.Pointer(s))

	result := ping_feature.Ping(str)

	return frankenphp.PHPString(result, false)
}
