package main

/*
#cgo CFLAGS: -D_GNU_SOURCE
*/
import "C"
import (
	"github.com/sprust/sconcur-franken/internal/services/ping_service"
)

//export ping
func ping(str *C.char) *C.char {
	result := ping_service.Ping(C.GoString(str))
	return C.CString(result)
}

func main() {}
