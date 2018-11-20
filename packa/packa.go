package packa

import (
	"fmt"
)

// coolPrivateFunction cannot be called from
// consumers of packa
func coolPrivateFunction(message string) {
	fmt.Println(message)
}

// CoolPublicFunction can be called from
// consumers of packa
func CoolPublicFunction(message string) {
	coolPrivateFunction(message) 
}
