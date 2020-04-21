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

# Notes on experience
# Version 1 ( iterative approach)
Coded relatively fast, had to take a moment to write down on paper an example to calculate median and indexes. Didn't end 
up with index out of bounds :)  

# Version 2 (Recursion)
Similar experience, the code is easier for me to come back and understand, prefer over version 1.

# Version 3 (Recursion with passing slices of array)
This version is adding unnecessary step of calculating the position of the slice in the initial array. 
It is complicated version of previous one with not much of an added value. 

# Version 4 (iterative approach different variation)
Really hard to come up with new solutions, I am going to re-implement iterative approach slightly different.
Version 1 and Version 4 are quite similar, but if you look closely, after calculating median, version one checks if 
number is greater or not and then chops again, while in version 4 it also checks if number we are looking for is in median, 
which saves additional iteration. So in fact Version 4 is saving additional iteration and therefore is more effective.

# Version 5 (tailing Recursion)
Version 5 is implementation with tailing recursion, meaning the recursive call is the last thing executed by the function.
Interesting to see if there is difference in execution time with Version. It is better readable for me though than Version 2, 
since recursive call is done in one line.

# Benchmarking
In order to run benchmarks 
```bash
$  go test ./... -bench=.
```
* After benchmarking we see that iterative approaches (V1 and V4) are performing better and V4 is slightly better. 
* Tailing Recursion (V5) shows to be slightly more performanct than Recursion (V2)

# Conclusion
In terms of code readability and maintainability my favourites are V1, V4, V5. Since V4 performs best among them, 
then that would be the one going to production.