package random

import (
	"fmt"
	"math/rand"
	"time"
)

func Code(length int) string {
	rand.NewSource(time.Now().UnixNano())
	return fmt.Sprintf("%4v", rand.Intn(10000))
}
