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