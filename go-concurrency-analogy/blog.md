# Effective Concurrency in Go

Go, also known as Golang, has gained popularity for its execution speed albeit
being easy to learn and write, and for its robust approach to concurrent
programming. It was designed with concurrency as a core feature, making it an
ideal choice for modern, scalable applications. Created by Google engineers
Robert Griesemer, Rob Pike, and Ken Thompson, Go has revolutionised concurrent
programming since its inception in 2009. 

Unlike many other popular programming languages, go makes it extremely easy to
write concurrent programmes. We'll dive into key concepts like goroutines,
channels, and the select statement, demonstrating how they work together to
create powerful concurrent systems. Whether you're new to Go or looking to
optimise your existing code, you'll find practical tips for leveraging Go's
concurrency model effectively.

## Concurrency

Since you're reading this blog, I think, you already have some idea about
concurrency. If not, it simply translates to running multiple tasks or processes
simultaneously. There's another similar concept named parallelism. But that's
another topic, if you wanna learn more in detail with technical jargon, keep
your eye on Vivasoft blog or subscribe to our newsletter. Too keen to wait? It's
just a Google search or ChatGPT prompt away. Come back here after you do that.

## Goroutines

 By definition, a _goroutine_ is a lightweight thread managed by the Go runtime.
 It's a function executing concurrently with other goroutines in the same
 address space. Goroutines allow for efficient use of multi-core processors and
 can significantly improve programme performance.

As previously mentioned, it's extremely easy to write concurrent programmes in
Go. All you need to remember is the `go` keyword and call a function. It
magically spawns the goroutines as needed and do the work, no need to
create and manage the threads manually.

Let's see some code.

```go
package main

import "fmt"

func inform() {
	fmt.Println("Hello from the meat store!")
}

func main() {
	fmt.Println("Hi!")
	go inform()
	fmt.Println("Bye!")
}
```

Output:
```
➜ go run .
Hi!
Bye!
```

What just happened? Even though `inform()` was called, we didn't see its
reflection. There are numerous explanations why this happened on the internet
filled with geekspeak. Since it's redundant to say those words again,
let's understand the scenario with an analogy.

Assume that you are a freelance delivery person (like Jason Statham in
Transporter movie). Now, you're hired by someone to deliver a package of meat to
a meat to a meat store, probably for a refund. You took the package and went to
the store. In the meantime the person hired you fled the country or whatever,
just became unavailable. You, on the other hand, who doesn't have any contract
to prove that you're doing your job nor someone is waiting to pay you, holding a
package of meat (probably rotten), inside a meat store where your delivery
package may not be taken back or thrown away.

Whatever happened to you, that's what has happened with the `inform()` function
and the meat package is the garbage value or memory leak in technical terms.

## Waitgroup

To prevent the depicted scenario, `WaitGroup` can be used which is a
synchronisation primitive provided by the `sync` package available in Go
standard library. `WaitGroup` is particularly useful in scenarios where you need
to perform multiple operations concurrently and wait for all of them to
complete, such as parallel data processing or graceful server shutdown.

Let's see it in action:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchTickets(w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Got your tickets. Coming...")
	time.Sleep(3 * time.Second)
	fmt.Println("Reached airport")
}

func main() {
	fmt.Println("I left my tickets at home, bring them to me!")
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fetchTickets(wg)

	wg.Wait()
	fmt.Println("Thanks! Adios, amigo.")
}
```

Output:
```
➜ go run .
Hi!
Hello from the meat store!
Bye!
```

Let's borrow the context from previous example and understand this code with another example. Imagine that person went to the airport but left his plane tickets at home. You're hired to bring him his tickets. He has to wait for you to board on the plane. After you finish your job, he carries on with his life. 

Here, in the `main` function, a WaitGroup (wg) is created to synchronise the
main goroutine with the `fetchTickets` goroutine. Then, before spawing another
goroutine for `fetchTickets` function, `wg.Add(1)` which increases the WaitGroup's
counter by 1, indicating that we're waiting for one goroutine to complete. After
launching a separate goroutine to execute `fetchTickets`,  main goroutine waits (by
invoking `wg.Wait()`) for the WaitGroup's counter to be zero, in other words, it
awaits all goroutines for which the counter was incremented have been executed.

In the `fetchTickets()` function, pointer to the WaitGroup was passed. `w.Done()`
tells the WaitGroup to decrement its counter by one. `defer` was used to invoke
`w.Done()` before exiting the function.

### Communication between goroutines by using pointer

It can be seen that how data can be shared by different goroutines by using pointers. `wg` was mutated in `fetchTickets` goroutine yet main goroutine knew about the change. 

Moving on, let's see you have 1000 employees. You assigned each of them to
collect a specific product and add it to the inventory. Every one of them goes
out, collects the item and increments the inventory count.

Let's translate this scenario into a very simple programme:

```go
package main

import (
	"fmt"
	"sync"
)

type Inventory struct {
	mu           sync.Mutex
	ProductCount int
}

func addToInventoryUnsafe(w *sync.WaitGroup, inventory *Inventory) {
	defer w.Done()
	inventory.ProductCount++
}

func main() {
	fmt.Println("Collecting the products...")
	wg := &sync.WaitGroup{}
	inventory := Inventory{
		ProductCount: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addToInventoryUnsafe(wg, &inventory)
	}

	wg.Wait()
	fmt.Printf("There are %d products in inventory\n\n", inventory.ProductCount)
}
```

The output:
```
➜ go run .
Collecting the products...
There are 983 products in inventory
```

Wait, what? There should be 1000 products since all goroutines did finish.  Let's run it 5 times and examine the output.

```
➜ for i in {1..5}; do go run . ; done
Collecting the products...
There are 986 products in inventory

Collecting the products...
There are 990 products in inventory

Collecting the products...
There are 984 products in inventory

Collecting the products...
There are 979 products in inventory

Collecting the products...
There are 980 products in inventory
```

Wait a minute?!! The results aren't consistent. Let's run this with `-race` and see what we get as output.

```
➜ go run -race .
Collecting the products...
==================
WARNING: DATA RACE
Read at 0x00c00010a038 by goroutine 10:
  main.addToInventoryUnsafe()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:15 +0x85
  main.main.gowrap1()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0x44

Previous write at 0x00c00010a038 by goroutine 6:
  main.addToInventoryUnsafe()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:15 +0x9d
  main.main.gowrap1()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0x44

Goroutine 10 (running) created at:
  main.main()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0xda

Goroutine 6 (finished) created at:
  main.main()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0xda
==================
==================
WARNING: DATA RACE
Write at 0x00c00010a038 by goroutine 13:
  main.addToInventoryUnsafe()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:15 +0x9d
  main.main.gowrap1()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0x44

Previous write at 0x00c00010a038 by goroutine 20:
  main.addToInventoryUnsafe()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:15 +0x9d
  main.main.gowrap1()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0x44

Goroutine 13 (running) created at:
  main.main()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0xda

Goroutine 20 (finished) created at:
  main.main()
      /Users/noor/work/Vivasoft/Blogs/go-concurrency-analogy/code/gopath/src/example3/main.go:35 +0xda
==================
There are 979 products in inventory

Found 2 data race(s)
exit status 66
```


## Mutex



## Channel