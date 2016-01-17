package count // import "github.com/carlosjhr64/count"

sync/atomic counting, blocking on max count, and waiting for count to return
to one.

Example:

    var threads = count.New(4)
    //...
    func main() {
      //...
      for {
        //...many times...
        threads.Plus()
        go run_stuff()
      }
      threads.Wait()
      //...
    }
    //...
    func run_stuff(){
      //...
      threads.Minus()
    }

const VERSION string = "0.1.0"

func New(n int) *Threads
type Threads struct { ... }
