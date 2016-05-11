# Shutdown
Shutdown provides way to create goroutines that can be all
shutdown at once and waits for them to all return before continuing

## Install
`go get -u github.com/jesselucas/shutdown`

## Usage
https://godoc.org/github.com/jesselucas/shutdown

```
s := NewShutdown()

for i := 0; i < 10; i++ {
  s.Go(func() { fmt.Println("Some code.") })
}

s.Shutdown()
fmt.Println("All routines shutdown")


// Or use GoFor to pass a function that will loop forever
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

```
