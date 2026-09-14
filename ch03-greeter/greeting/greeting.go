package greeting

import "fmt"

var Polite = false

func Greet(name string) string {
	if name == "" {
		name = "friend"
	}
	return fmt.Sprintf("%s, %s!", salutation(), name)
}
