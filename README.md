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
