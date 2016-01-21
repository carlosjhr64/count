/*
  sync/mutex counting,
  blocking on max count, and
  waiting for count to return to one.

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

*/
package count

import "sync"

const VERSION string = "2.0.0"

type Threads struct {
  max int
  count int
  send bool
  channel chan bool
  mutex *sync.Mutex
}


func New(n int) *Threads {
  if n < 2 { panic("Does not make sense for threads to be less than 2.") }
  max, count, send, channel, mutex := n, 1, false, make(chan bool, 1), &sync.Mutex{}
  return &Threads{max, count, send, channel, mutex}
}

func (threads *Threads) Plus() int {
  threads.mutex.Lock()
  count := threads.count
  // Although count can't be higher than max via the API,
  // include the possibility:
  if count >= threads.max {
    threads.send = true
    threads.mutex.Unlock()
    <-threads.channel
    // count remains the same.
    return count
  }
  count += 1
  threads.count = count
  threads.mutex.Unlock()
  return count
}

func (threads *Threads) Minus() int {
  threads.mutex.Lock()
  count := threads.count
  if threads.send {
    // A Plus (or Wait) call is waiting to proceed...
    threads.send = false // ...and this Minus handles it!
    // Allow the receiver to decide if count needs to decrement.
    threads.channel <- true
  } else {
    count -= 1
    threads.count = count
  }
  threads.mutex.Unlock()
  return count
}

func (threads *Threads) wait() {
  for {
    <-threads.channel
    threads.mutex.Lock()
    threads.count -= 1
    if threads.count < 2 { break }
    threads.send = true
    threads.mutex.Unlock()
  }
  threads.mutex.Unlock()
}

func (threads *Threads) Wait() {
  threads.mutex.Lock()
  count := threads.count
  if count > 1 {
    threads.send = true
    threads.mutex.Unlock()
    threads.wait()
    return
  }
  threads.mutex.Unlock()
}

func (threads *Threads) Count() int {
  threads.mutex.Lock()
  count := threads.count
  threads.mutex.Unlock()
  return count
}
