Ema's supercomputer

https://www.hackerrank.com/challenges/two-pluses

Essentially, this solution revolves around parsing the entire grid into an internal structure and then searching for all "plus" structures inside the grid
After all pluses are found, they are sorted according to their size
Then, all the pluses are compared with each other for overlaps, a 2D array stores overlap information
Again, all the plus structures are copmpared against each other, if two plus do not overlap, their element size is multiplied with each other
and if their product is greater than the maximum product (initilized as 0), the maximum product is updated to their product.
Finally the maximum product i.e. the biggest possible product of two non-overlapping pluses in the grid is calculated.

