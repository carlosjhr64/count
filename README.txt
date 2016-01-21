package count // import "github.com/carlosjhr64/count"

sync/mutex counting, blocking on max count, and waiting for count to return
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
      theads.Minus()
    }

const VERSION string = "2.0.0"

func New(n int) *Threads

type Threads struct {
	// Has unexported fields.
}

func New(n int) *Threads
func (threads *Threads) Count() int
func (threads *Threads) Minus() int
func (threads *Threads) Plus() int
func (threads *Threads) Wait()
