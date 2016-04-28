package shutdown

import "sync"

// Shutdown provides way to create goroutines that can be
// all shutdown at once and waits for them to all return before
// continuing
type Shutdown struct {
	wg       *sync.WaitGroup
	shutdown chan int
}

// NewShutdown returns a SyncUtil struct
func NewShutdown() *Shutdown {
	s := new(Shutdown)
	s.wg = new(sync.WaitGroup)
	s.shutdown = make(chan int)

	return s
}

// Go replaces the use of a goroutine to add the ability to shutdown the
// routine with Shutdown(). If you need a for loop us GoFor()
func (s *Shutdown) Go(f func()) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		f() // Now call function arg
		<-s.shutdown
		return
	}()
}

// GoFor calls your function infinitely until Shutdown() is called
func (s *Shutdown) GoFor(f func()) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			default:
				f() // Now call function arg
			case <-s.shutdown:
				return
			}
		}
	}()
}

// Shutdown closes the shutdown chan and then waits for all goroutines to return
func (s *Shutdown) Shutdown() {
	close(s.shutdown)
	s.wg.Wait()
}
