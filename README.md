# Karate Chop

Karate chop is an implementation of a binary search that finds the position of value in a sorted array of values or -1 
if it's not found. The idea is to re-implement same problem every day, 5 times. See more details here: [Karate Chop](http://codekata.com/kata/kata02-karate-chop/)

# Run the program and tests
Please make sure you have go installed 
Clone the repo and go to wordchain-breadth-first-search/src directory

```bash
$ git clone git@github.com:gohargasparyan/go-kata.git
$ cd go-kata
```

# Run tests

to run the tests

```bash
$  go test ./... -test.v
```
# Benchmarking
.. todo

# Notes on experience
# Version 1 ( iterative approach)
Coded relatively fast, had to take a moment to write down on paper an example to calculate median and indexes. Didn't end 
up with index out of bounds :)  

# Version 2 (Recursion)
Similar experience, the code is easier for me to come back and understand, prefer over version 1.

# Version 3 (Recursion with passing slices of array)
This version is adding unnecessary step of calculating the position of the slice in the initial array. 
It is complicated version of previous one with not much of an added value. 

# to be continued