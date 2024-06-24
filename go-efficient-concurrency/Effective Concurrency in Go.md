Go, also known as Golang, has gained popularity for its execution speed albeit
being easy to learn and write, and for its robust approach to concurrent
programming. It was designed with concurrency as a core feature, making it an
ideal choice for modern, scalable applications. Created by Google engineers
Robert Griesemer, Rob Pike, and Ken Thompson, Go has revolutionized concurrent
programming since its inception in 2009. 

Unlike many other popular programming languages, go makes it extremely easy to
write concurrent programmes. We'll dive into key concepts like goroutines,
channels, and the select statement, demonstrating how they work together to
create powerful concurrent systems. Whether you're new to Go or looking to
optimize your existing code, you'll find practical tips for leveraging Go's
concurrency model effectively.

## Concurrency
If you have read this far, I think you already have some idea about concurrency.
It simply translates to running some tasks simultaneously. There's another similar concept
named parallelism. But that's another topic, if you wanna learn more keep your eye on
Vivasoft blog. Too keen to wait? It's just a Google search or ChatGPT prompt away. Come
back here after you do that.

## Waitgroup
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
→ go run .
Hello
```

## Channel