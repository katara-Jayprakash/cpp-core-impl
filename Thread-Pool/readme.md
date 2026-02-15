# Thread Pool 

## The Problem

You have 10,000 tasks to do. You think: "I will create 10,000 threads. Each thread does one task. Easy!"

But this does not work. Here is why.

---

## Why Creating 10,000 Threads is Bad

### Problem 1: You Only Have 8 CPU Cores

Your computer has 8 cores. This means only 8 things can run at the same time.

```
You create: 10,000 threads
Your computer has: 8 cores

What happens:
- 8 threads run (one per core)
- 9,992 threads wait
- They just sit there doing nothing
- But they still use memory!
```

### Problem 2: Each Thread Uses Memory

Each thread needs space in memory. Usually 1 MB per thread.

```
10,000 threads × 1 MB = 10,000 MB = 10 GB

Your computer needs 10 GB just to store empty threads!
```

### Problem 3: Creating Threads is Slow

Every time you create a thread:
- It takes about 2 milliseconds
- The operating system does work
- Memory is allocated

```
10,000 threads × 2 milliseconds = 20 seconds

```

### Problem 4: Switching Between Threads is Expensive

Your computer switches between threads very fast. But this switching costs time.

```
With 10,000 threads:
- Computer spends more time switching
- Less time doing actual work
- Everything becomes slow
```


## The Solution: Thread Pool

Smart people thought: "Why not create a small number of threads and reuse them?"

This is called a **thread pool**.

### How Thread Pool Works

```
Instead of:
- Create 10,000 threads
- Each thread does one task
- Destroy all threads

Do this:
- Create 8 threads (one for each core)
- Give them 10,000 tasks
- Each thread does many tasks
- Keep threads alive
```



### Benefits

- Only 8 threads (8 MB memory, not 10 GB)
- Create threads once (takes 16 ms, not 20 seconds)
- All 8 threads always working
- Less switching between threads
- Computer runs fast

---

## How Go Does This (Goroutines)

In languages like C++, you must build a thread pool yourself. Go does it for you automatically.

### You Write This

```go
for i := 0; i < 10000; i++ {
    go doWork(i)
}
```

You just created 10,000 goroutines! But everything is fine how?

### What Go Does

Go is smart. Here is what happens:

```
Your code creates: 10,000 goroutines
Go runtime creates: 8-10 actual threads
Your computer has: 8 cores

Go runtime:
- Takes 10,000 goroutines
- make 8-10 thread that we can called gothread
- 8-10 gothreads now pick these goroutines from queues
- Each thread runs many goroutines
```
### The Go Scheduler
Go uses something called the **GMP model**:

- **G** = Goroutine (your task)
- **M** = Machine (actual thread, made by operating system)
- **P** = Processor (one per CPU core, like a work queue)

```
You write:
10,000 goroutines (G)
     ↓
Go creates:
8 processors (P) - one per core
     ↓
Go creates:
8-10 threads (M) - actual threads
     ↓
Operating system:
Uses your 8 CPU cores
```

---

## The Three Parts of a Thread Pool

Every thread pool has three parts:

### 1. Tasks (Jobs)
 job is something that can be executable, 

### 2. Queue (Where Tasks Wait)

Queue is a line where tasks wait their turn.
basically means where we can store incoming query while our worker thread are busy


### 3. Workers (Threads)

 we limit the creation of thread, we are going to use them again and again this can save us from destroying and creating new thread.

```
Worker 1: Takes Task1, does it, takes Task4, does it ...
Worker 2: Takes Task2, does it, takes Task5, does it ...
Worker 3: Takes Task3, does it, takes Task6, does it ...
```

Workers never stop. They keep taking tasks and doing them.

### How They Work Together

```
       TASKS          QUEUE              WORKERS
    
Task1 ──┐                               ┌──→ Worker 1
Task2 ──┼──→ [T1][T2][T3] ──→ Pick ──→  ├──→ Worker 2
Task3 ──┘                               └──→ Worker 3
                                             ↓
                                        Do the work
```

1. Tasks come in
2. Tasks go into queue
3. Workers take tasks from queue
4. Workers do the tasks
5. Workers go back for more tasks

---
### How much worker thread should we included in threadPool or how should we tune thread
From my prospective, 
I dont think there is certain number for using this specific number of worker-thread, 
   i mean its totally depends on 

  1. what kind of task are you doing like 
     if its i/o  bound or network related use then use more number of worker-Thread 
     because go has special way of handling them go does not count them in blocking task
     and if task is more cpu Intensive then use less number of worker thread 
     because thread are going to busy maximum time.
  2. another thing is it totally depends on your computer hardware,
---
