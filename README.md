
# DATA STRUCTURE AND ALGORITHM IN GOLANG 
-------------------------------------------------------

> DATA STRUCTURE 
---------------------
  Data structures are a way of organizing and storing data in a computer 
  so that it can be accessed and used efficiently. Algorithms are 
  step-by-step procedures for solving a problem or accomplishing a task.

> Some common data structures include arrays, linked lists, trees, 
  graphs, hash tables, and queues. Each data structure has its own 
  advantages and disadvantages, depending on the specific problem being solved.

> There are many different algorithms for performing common tasks such as 
  sorting, searching, and traversing data structures. Some popular algorithms 
  include bubble sort, quicksort, binary search, depth-first search, 
  and breadth-first search


# Memory Allocation and De-Allocation
-----------------------------

 1. Memory allocation is the process of reserving a block of memory for use by a program. 
    The allocated memory is then available for the program to use to store data.
   
 > To allocate memory in Go, you can use the built-in make function or the new function. 
   The make function is used for creating slices, maps, and channels, while the new function 
   is used for creating pointers to a new zeroed value of a given type.

     mySlice := make([]int, 10) // allocate a slice of 10 integers
     myIntPointer := new(int) // allocate a pointer to an integer

 2. memory de-allocation it is the process of releasing memory that is no longer 
    needed by a program This frees up the memory for other programs to use.

 > When it comes to deallocation of memory, Go handles this automatically through its 
   garbage collector. The garbage collector periodically scans the program's memory to 
   identify and free up memory that is no longer being used. This means that as a Go 
   programmer, you don't need to worry about explicitly deallocating memory in most cases. 
   However, if you are working with resources that are not managed by the garbage collector, 
   such as file handles or network sockets, you will need to explicitly close or release 
   them when you are done using them to avoid leaking resources.


# Bit and byte convertion
-------------------
1. 8 bit =  1 byte
2. 1024 byte = 1 kb [Kilo-byte]
3. 1024 kb = 1 mb [mega-byte]
4. 1024 mb = 1 bg [giga-byte]
5. 1025 gb = 1 pb [peta-byte]


# Asymptotic Notation 
    --------------
It is used to describe the running time of an algorithm - how much time an algorithm takes with a given input, n. There are three different notations: big O, big Theta (Θ), and big Omega (Ω).

1. Big O Notation - calculate the worst case complexity
2. Theta Notation - calculate the avarage case complexity
3. Omega Notation - calculate the best case complexity

// Diagrom image
![alt text](https://github.com/abhinandpn/dsa-golang/blob/main/diagrom.png)

# Oprations and Complexity
1. set
2. get
3. init
4. traverse
5. insert
6. delete

> Array

    Set opration
        Time complexity = o(1)T 
        space complexity = o(1)S

    Get opration
        Time complexity = o(n)T
        space complexity = o(1)S

    Init opration
        Time complexity = o(n)T
        space complexity = o(n)S
    
    Traverse opration
        Time complexity = 
        space complexity =

    Insert opration
        Time complexity = 
        space complexity =

    Insert opration
        Time complexity = 
        space complexity =

    Delete opration
        Time complexity = 
        space complexity =


>Linked list

    Set opration
        Time complexity = 
        space complexity = 

    Get opration
        Time complexity = 
        space complexity = 

    Init opration
        Time complexity = o(n)T
        space complexity = o(n)S
    
    Traverse opration
        Time complexity = 
        space complexity =

    Insert opration
        Time complexity = 
        space complexity =

    Insert opration
        Time complexity = 
        space complexity =

    Delete opration
        Time complexity = 
        space complexity =
