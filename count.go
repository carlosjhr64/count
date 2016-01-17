// sync/atomic counting, and wating for count to return to one.
package count

import "sync/atomic"
import "time"

const VERSION string = "0.0.0"

// Time interval in Millsecond
var Interval time.Duration = 100*time.Millisecond

var threads int64 = 1
func Wait() {
  time.Sleep(Interval)
  for {
    t := atomic.LoadInt64(&threads)
    if t == 1 { break }
    time.Sleep(Interval)
  }
}

func Plus(){
  atomic.AddInt64(&threads, 1)
}

func Minus() {
  atomic.AddInt64(&threads, -1)
}
