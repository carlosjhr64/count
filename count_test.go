package count

import "fmt"
import "time"
import "testing"
import "sync/atomic"

func do_nothing_for_a_second(threads *Threads) {
  time.Sleep(time.Second)
  threads.Minus()
}

func TestCount(test *testing.T) {
  bad := test.Error
  var n int
  fmt.Println("Should not take more than a couple seconds.")

  threads := New(3)

  n = threads.Plus()
  if n != 2 { bad("Should have 2 threads.") }
  go do_nothing_for_a_second(threads)

  n = threads.Plus()
  if n != 3 { bad("Should have 3 threads.") }
  go do_nothing_for_a_second(threads)

  n = threads.Plus()
  if n != 4 { bad("Should have 4 threads.") }
  go do_nothing_for_a_second(threads)

  n = threads.Plus()
  if n != 4 { bad("Should still have 4 threads.") }
  go do_nothing_for_a_second(threads)

  threads.Wait()
  n = int(atomic.LoadInt32(&threads.count))
  if n != 1 { bad("Should just have 1 thread.") }
}
