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