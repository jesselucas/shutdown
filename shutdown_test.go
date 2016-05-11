package shutdown

import (
	"fmt"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	s := NewShutdown()

	for i := 0; i < 10; i++ {
		s.Go(func() { fmt.Println("Some code.") })
	}

	s.Shutdown()
	fmt.Println("All routines shutdown")
}

func TestGoFor(t *testing.T) {
	testLoop := func() {
		fmt.Println("loop forever")
	}

	s := NewShutdown()

	for i := 0; i < 10; i++ {
		s.GoFor(testLoop)
	}

	// Let's wait to see loop working
	time.Sleep(2 * time.Second)

	s.Shutdown()
	fmt.Println("All routines shutdown")
}
