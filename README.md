# Concurrency Using Golang
My personal notes, projects and best practices.

## Author
Aditya Hajare ([Linkedin](https://in.linkedin.com/in/aditya-hajare)).

## Current Status
WIP (Work In Progress)!

## License
Open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).

-----------

## Concurrency:
- `Concurrency` is about **multiple things happening at the same time** in random order.
- `Concurrency` is composition of **independent execution compiutations**, which may or may not run in parallel.
- `Parallelism` is ability to **execute multiple computations simultaneously**.
- **`Concurrency` enables the `Parallelism`.**
- `Concurrency` is about **dealing with lots of things at once**.
- `Parallelism` is about **doing lots of things at once**.
- `Concurrency` is about **structure** and `Parallelism` is about **execution**.
- `Memory Access Synchronization` tools reduce `Parallelism` and comes with their own limitations.

-----------

## Processes And Threads 101:
- **Process:**
    * `Process` is just an instance of a running program.
    * `Process` provides environment for program to execute.
    * When the program is executed, the Operating System creates a process and :
        - Allocates memory in a virtual address space.
        - The virtual address space will contain Code Segment which is a compiled machine code.
        - There is a `Data Region` which contains Global Variables.
        - `Heap Segment` is used for `Dynamic Memory Allocation`.
        - `Stack` is used for storing `Local Variables in Function`.
- **Operating System:**
    * The job of operating system is to give fair chance for all processes to access CPU, memory and other resources. There are times when higher priority tasks get precedence.
- **Threads:**
    * `Threads` are **smallest unit of execution** that CPU accepts.
    * Each `Process` has atleast one thread. That is `main thread`.
    * `Process` can have multiple `Threads`.
    * `Threads` share the same address space.
    * Each `Thread` has it's own `Stack`.
    * `Threads` run **independent of each other**.
    * **Operating System Scheduler makes scheduling decisions at thread level and not at the process level!**
    * `Threads` can run concurrently (with each thread taking turn on individual core) or in parallel (with each thread running on different cores at the same time).
    * `Threads` communicate between each other by sharing memory.
    * Sharing of memory between threads creates lot of complexity.
    * Concurrent access to to shared memory by two or more threads can lead to **Data Race** and outcome can be **Un-deterministic**.
    * The actual number of threads we can create are limited.

-----------

## Goroutines:
- We can think of Goroutines as **user space threads managed by Go runtime**.
- Goroutines are extremely lightweight. Goroutines starts with **2kb of stack**, which grows and shrinks as required.
- **Low CPU Overhead**: 3 instructions per function call.
- Can **create hundreds of thousands of goroutines** in the same address space.
- **Channels are used for communication of data** between Goroutines. Sharing of memory can be avoided.
- `Context Switching` between Goroutines is much cheaper than thread `Context Switching`.
- Go runtime can be more selective in what is persisted for retrieval, how it is persisted and when the persisting needs to occur.
- Go runtime creates OS threads.
- Goroutines runs in the context of OS threads.
- Many Goroutines can execute in a context of single OS threads.

-----------

## WaitGroups:
- Go follows a logical concurrency model called **Fork and Join**.
- Go statement `Forks` a Goroutine. When a Goroutine is `done` with it's job, it `Joins` back to the `main routine`. If `main routine` doesn't wait for the Goroutine, then it is highly likely that a program will finish before the Goroutine gets a chance to run.
- To create a `Join` point, we can use `sync.WaitGroup`.
- **Waitgroups deterministically blocks the main goroutine.**
- Psudo Code:
    ```go
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        // Do stuff..
    }

    wg.Wait()
    ```

-----------

## Goroutines And Closures:
- Goroutines execute within the **same address space** they are created in.
- Goroutines can directly modify variables in the enclosing lexical block.

-----------

## Go Scheduler:
- Go runtime has mechanism known as **MN Scheduler**.
- Fo Scheduler runs in user space.
- Go Scheduler uses OS threads to schedule Goroutines for execution.
- **Goroutine runs in the context of OS threads.**
- Go runtime creates number of worker OS threads, equals to **`GOMAXPROCS` environment variable value**.
- **`GOMAXPROCS` default value is number of processors/cores on machine.**
- It is a responsibility of Go Scheduler to distribute runnable Goroutines on over multiple OS threads that are created.

-----------

## Channels:
- Channels are used to communicate data between Goroutines.
- Channels can also help synchronize Goroutines execution.
- **Channels are typed**. They are used to send and receive values of a particular type.
- Channels are **Thread Safe!** Channel variables can be used to send and receive values concurrently by multiple Goroutines.
- Default value of Channel is `nil`.
- Example:
```go
var ch chan int
ch = make(chan int) // make() will allocate memory for channel and returns referance for the allocated Memory
// OR
ch := make(chan int) // Declares and allocates memory for channel
```
- Pointer operator is used for sending and receiving the value from channel.
- The "arrow" indicates the direction of data flow.
- **Channels are blocking!** i.e. The sending Goroutine is going to block until there is a corrosponding Goroutine ready to receive the value.
- It is responsibility of channel to make the Goroutine runnable again once it has the data.
- Sender Goroutine must indicate receiving Goroutine that it has no more values to send. We use `close(ch)` to close the channel.
- Receiver receives 2 values:
```go
value, ok = <- ch
// value: Received value from the sender
// ok: is a boolean value
//    True, value generated by write
//    False, value generated by close
```

-----------

## Range Over Channel:
- The receiver Goroutine can receive sequence of values from Channel and then it can range over those values.
- Loop automatically breaks when Channel is closed.
- Range does not return the second boolean value.

-----------

## Unbuffered Channels:
- They are Synchronous.
- Receiving Goroutine will block until there is sender and sender will block until there is receiver.
- To create Unbuffered Channel:
```go
ch := make(chan int)
```

-----------

## Buffered Channels:
- They are Asynchronous.
- They are in-memory FIFO queues.
- There is a buffer between sender and receiver Goroutine.
- We can specify capacity i.e. Buffer size, which indicates the number of elemenets that can be sent without the receiver being ready.
- Sender can keep sending the values until buffer gets full. When the buffer gets full, the sender will get blocked.
- Receiver will keep receiving values until buffer gets empty. When the buffer gets empty, the receiver will get blocked.

-----------

## Channel Direction:
- When using channels as a function parameters, we can specify if a channel is meant to only send or receive values.
- This will increase the Type Safety of our program.
- For e.g.
```go
func(in <-chan string, out chan<- string) {
    // in: This channel can only receive values of type string
    // out: This channel can only send values of type string
}
```

-----------

## Channel Default Values And Best Practices:
- Default value for channels is `nil`.
```go
var ch chan interface{}
```
- Reading and Writing to a `nil` channel will block forever.
```go
var ch chan interface{}
<- ch
ch <- struct{}{}
```
- Closing `nil` channel will `panic`.
```go
var ch chan interface{}
close(ch)
```
- Ensure the channels are initialized first.
- Owner of channel is a **Goroutine that instantiates, writes and closes a channel**.
- Channel **utilizers only have a read-only** view into the channel.

-----------

## Select Statement:
- It allows us to do operations on Channel which ever is ready and don't worry about the order.
- Select is like a `switch` statement.
- Each statement specifies `send` or `receive` on a specific channel and it has associated block of statements.
- Each cases specifies communication.
- All channel operations are considered simultaneously.
- **Select waits until some case is ready to proceed.** If none of the channels are ready, then entire Select statement is going to be blocked until some case is ready for the communication.
- When one channel is ready, that operation will proceed.
- If multiple channels are ready, it will pick one of the channels randomly and proceed.
- Select is helpful for implementing:
    * Timeouts
    * Non-blocking communications
- Select Statement syntax:
```go
select {
    case <- ch1:
        // Block of statements
    case <- ch2:
        // Block of statements
    case ch3 <- struct{}{}:
        // Block of statements
}
```
- **We can specify timeouts on channel operations as below**.
- In below example:
    * Select will wait until there is event on channel ch or until timeout is reached (after 3 seconds).
    * The `time.After()` function takes in a `time.Duration` argument and returns a channel that will send the current time after the duration we have specified.
```go
select {
    case v := <- ch:
        fmt.Println(v)
    case <- time.After(3 * time.Second):
        fmt.Println("Timeout!")
}
```
- **As we know channels are blocking, we can implement Non-Blocking operation using `select` by specifying the `default` case**.
- In below example:
    * Send or receive on a channel, but avoid blocking if channel is not ready.
    * `default` allows us to exit a `select` block without blocking.
```go
select {
    case m := <- ch:
        fmt.Println("Received message: ", m)
    default:
        fmt.Println("No message received")
}
```
- Empty `select` statement will block forever.
```go
select {}
```
- `select` on `nil` channel will block forever.
```go
// Will block forever
var ch chan string
select {
    case v := <- ch:
    case ch <- v:
}
```

-----------

## Mutex:
- Mutex is used to guard access to shared resource.
- sync.Mutex provides exclusive access to shared resource.
- If the Goroutine is just reading from the memory and not writing to the memory then we can use `READ WRITE MUTEX`.
- `sync.RWMutex` allows multiple readers. Writers get exclusive lock.

-----------

## When to use Channels vs. When to use Mutex:
- Channels:
    * They are made to implement communication between Goroutines.
    * Passing copy of data.
    * Distributing units of work.
    * Communicating asynchronous results.
- Mutex:
    * When we have data such as Caches, States, Registeries which are big to be sent over the channels and we want access to this data to be thread safe. This is where classic synchronization tool such as Mutex comes into the picture.

-----------

## sync.Atomic:
- `Automic` is used to performed low level automic operations on memory. It is used by other synchonization utilities.
- It is a `lockless` operation.

-----------

## sync.Cond:
- Condition variable is one of the synchronization mechanism.
- It is a `lockless` operation.
- A condition variable is basically a container of Goroutines that are waiting for a certain condition.
- Condition variables are type:
```go
var c *sync.Cond
```
- We use constructor method `sync.NewCond()` to create a conditional variable, it takes `sync.Locker` interface as input, which is usually `sync.Mutex`.
```go
mu := sync.Mutex{}
cond := sync.NewCond(&mu)
```
- Wait suspends the execution of Goroutine.
- Signal wakes one Goroutine waiting on `c`.
- Broadcast wakes all Goroutines waiting on `c`.

-----------