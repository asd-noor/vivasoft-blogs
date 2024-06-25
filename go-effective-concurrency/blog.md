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
It simply translates to running some tasks simultaneously. There's another
similar concept named parallelism. But that's another topic, if you wanna learn
more keep your eye on Vivasoft blog. Too keen to wait? It's just a Google search
or ChatGPT prompt away. Come back here after you do that.

## Goroutines

 By definition, a _goroutine_ is a lightweight thread managed by the Go runtime.
 It's a function executing concurrently with other goroutines in the same
 address space. Goroutines allow for efficient use of multi-core processors and
 can significantly improve programme performance.

As previously mentioned, it's extremely easy to write concurrent programmes in
Go. All you need to remember is the keyword `go` and call a function. It
magically spawns the threads (goroutines) as needed and do the work, no need to
create and manage the threads manually.

Let's see some shiny code.
```go
func greet() {
	fmt.Println("Hello from the other side")
}

func main() {
	go greet()
	fmt.Println("Hello")
}
```

Output:
```shell
⯈ go run .
Hello
```

What just happened? Even though `greet()` was called, we didn't see any
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
package may not be accepted or thrown away.

Whatever happened to you, that's what has happened with that `greet()` function
and the meat package is the garbage value or memory leak in technical terms.

## Waitgroup

## Channel