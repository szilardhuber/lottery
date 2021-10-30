# Lottery

Given an input file data of 5-10m lottery tickets find the number of winning tickets in a couple of seconds after getting the winning numbers on stdin.

# Building and running

`go run . <<input_file_name>>`

When run without any additional settings, the output is reduced to follow the required format in the problem statement. If the environment has a `DEBUG=1` variable then extra information is displayed as well (like invalid tickets or time for building the main data model or calculating the results)

# Getting to current results

#### Bash

My very first idea was that it must be possible with linux tools like sed, awk, grep. I came upon the following oneliner which could be used to find if there is line with 5 matching numbers in any order.

```ggrep -P '(?=.*?76)(?=.*73)(?=.*?19)(?=.*?88)(?=.*?78)^.*$' 10m-v2.txt | wc -l ```

Time for this: `3.01s user 0.04s system 99% cpu 3.082 total`

Besides it being slow, the other problem is that I would need to generate all possible combinations of 2,3,4 matching numbers.

#### Go first iteration

My first solution was to go through all the lines in the file and just stream process them without storing anything in memory. Also I didn't deal with converting the strings to numbers or checking the validity of the input. Besides returining wrong results for malformed input this solution also took couple of seconds on my laptop per "lottery show" which was in the given range but felt a bit slow.

Time: `3.88s user 0.19s system 106% cpu 3.811 total`

#### Go second iteration

After that I thought that all numbers can be represented as `uint8` and 10m times 5 times 1 is about 50MB, so I should be able to store everything in memory and avoid reading up a file for each "lottery show". The first very naive implementation had several optimization errors and kept growing the slice capacity one by one so took around 800MB of memory which didn't cause any troubles but I felt I should optimize. Building up the "main" data structure takes around 5s on my laptop and evaluating results for a "lottery show" is just below 300ms

Optimization steps I took:
1. Set the initial capacity of the in-memory representation of the numbers of the players to 10m. Still using append to add new data points so the solution would work with more numbers but might be not optimized in terms of memory allocation. - This took down the memory usage to around 450MB
2. Set the "inner" array of the "main" data structure to `[5]uint8` instead of `[]uint8`. This made the code a bit more complicated but had a significant impact on memory usage. - Memory usage further reduced to around 100MB
3. Force GC manually. I don't really believe in doing these kinds of things I just wanted to understand if it is GC that causes the "doubled" memory usage compared to my expectations. - Memory usage reduced to around 74MB (Which funny enough didn't really answer my question :) )

# Future improvements

Optimization:
I believe the solution could handle an order of magnitude more input data but if we had even more that that I would think about just splitting the input file to smaller chunks, distributing it on different machines and doing the procesing parallel with some parallelization solutions. (For example with MPI but I haven't yet used MPI with go, only with python)

Quality assurance:
I haven't written any unit tests. To be honest I have spent more time with playing around the reduction of the memory usage than I wanted to (2-3 hours total for the whole problem) so I stop it here. I think most methods could have some tests but the only one I believe is complex enough to provide actual value in covering with tests is the `split` method.
