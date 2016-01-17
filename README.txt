package count // import "github.com/carlosjhr64/count"

sync/atomic counting, and wating for count to return to one.

const VERSION string = "0.0.0"

var Interval time.Duration = 100 * time.Millisecond

func Minus()
func Plus()
func Wait()
