/*
  sync/atomic counting,
  blocking on max count, and
  waiting for count to return to one.

  Example:
    var threads = count.New(4)
    //...
    func main() {
      //...
      for {
        //...many time...
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

*/
package count

import "sync/atomic"

const VERSION string = "1.0.0"

type Threads struct {
  plus chan bool
  minus chan int
  count int32
}

func New(n int) *Threads {
  if n < 2 { panic("Does not make sense for threads to be less than 2.") }
  var count int32 = 1
  n = n-1 // The thread count starts at 1.
  plus := make(chan bool, n)
  minus := make(chan int, n)
  return &Threads{plus, minus, count}
}

func (threads *Threads) flush() {
  var b bool = true
  for b {
    select {
    case <-threads.minus:
      // repeat
    default:
      b = false
    }
  }
}

func (threads *Threads) Plus() int {
  threads.flush()
  threads.plus <- true
  n := int(atomic.AddInt32(&threads.count, 1))
  if n < 2 { panic("Count must have been at least one.") }
  return n
}

func (threads *Threads) Minus() int {
  n := int(atomic.AddInt32(&threads.count, -1))
  if n < 1 { panic("Count must be at least one.") }
  <-threads.plus
  threads.minus <- n
  return n
}

func (threads *Threads) Wait() {
  var n int
  for {
    n = <-threads.minus
    if n == 1 { break }
  }
}
