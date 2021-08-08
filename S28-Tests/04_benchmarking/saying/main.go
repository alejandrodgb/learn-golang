package saying

import "fmt"

// Greet welcomes someone when called
func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}
