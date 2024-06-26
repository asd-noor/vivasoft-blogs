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

If you have read this far, I think you already have some idea about concurrency.
If not, it simply translates to running multiple tasks simultaneously. There's
another similar concept named parallelism. But that's another topic, if you
wanna learn more keep your eye on Vivasoft blog. Too keen to wait? It's just a
Google search or ChatGPT prompt away. Come back here after you do that.

## Goroutines

 By definition, a _goroutine_ is a lightweight thread managed by the Go runtime.
 It's a function executing concurrently with other goroutines in the same
 address space. Goroutines allow for efficient use of multi-core processors and
 can significantly improve programme performance.

As previously mentioned, it's extremely easy to write concurrent programmes in
Go. All you need to remember is the keyword `go` and call a function. It
magically spawns the goroutines as needed and do the work, no need to
create and manage the threads manually.

Let's see it in action.

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
```shell
⯈ go run .
Hi!
Bye!
```

What just happened? Even though `inform()` was called, we didn't see its
reflection. There are numerous explanations why this happened on the internet
filled with technical jargon. Since it's redundant to say those words again,
let's understand the scenario with an analogy.

Assume that you are a freelance delivery person (like Jason Statham in
Transporter movie). Now, you're hired by someone to deliver a package of meat to
a meat to a meat store, probably for a refund. You took the package and went to
the store. In the meantime the person hired you fled the country or whatever,
just became unavailable. You, on the other hand, who doesn't have any contract
to prove that you're doing your job nor someone is waiting to pay you, holding a
package of meat (probably rotten), inside a meat store where your delivery
package may not be taken back or thrown away.

Whatever happened to you, that's what has happened with the `greet()` function
and the meat package is the garbage value or memory leak in technical terms.

## Waitgroup

To prevent the depicted scenario, `WaitGroup` can be used which is a
synchronisation primitive provided by the `sync` package available in Go
standard library. `WaitGroup` is particularly useful in scenarios where you need
to perform multiple operations concurrently and wait for all of them to
complete, such as parallel data processing or graceful server shutdown.

Let's see it in action with the context of previous example:

```go
package main

import (
	"fmt"
	"sync"
)

func inform(w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Hello from the meat store!")
}

func main() {
	fmt.Println("Hi!")
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go inform(wg)

	wg.Wait()
	fmt.Println("Bye!")
}
```

Output:
```shell
⯈ go run .
Hi!
Hello from the meat store!
Bye!
```

Here, in the `main()` function, a WaitGroup (wg) is created to synchronise the main goroutine with the `inform` goroutine. Then, before spawing another goroutine for `inform` function, `wg.Add(1)` which increases the WaitGroup's counter by 1, indicating that we're waiting for one goroutine to complete. After launching a separate goroutine to execute `inform()`,  main goroutine waits (by invoking `wg.Wait()`) for the WaitGroup's counter to be zero, in other words, it awaits all goroutines for which the counter was incremented have been executed.

In the `inform()` function, pointer to the WaitGroup was passed. `w.Done()` tells the WaitGroup to decrement its counter by one. `defer` was used to invoke `w.Done()` before exiting the function.

### Communication between goroutines by using pointer


## Channel