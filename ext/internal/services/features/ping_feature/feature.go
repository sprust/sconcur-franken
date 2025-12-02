package ping_feature

import (
	"fmt"
)

func Ping(string string) string {
	return fmt.Sprintf("pong %s", string)
}
