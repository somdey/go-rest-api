package gostrings
import "regexp"
import "fmt"
func Squish(text string) string {
	fmt.Println(text)
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(text, " ")
}