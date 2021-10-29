### Very first idea: It must be possible with linux tools like sed, awk, grep

```ggrep -P '(?=.*?76)(?=.*73)(?=.*?19)(?=.*?88)(?=.*?78)^.*$' 10m-v2.txt | wc -l ```

Time for this: `3.01s user 0.04s system 99% cpu 3.082 total`

The other problem is that we would need to generate all possible combinations of 2,3,4 occurance numbers.

### Go implementation

Naive one: Don't convert strings to numbers, just do a plain split and count the occurances.

Time: `3.88s user 0.19s system 106% cpu 3.811 total` But this at least seems to return all needed results.


### Optimization ideas

1. Sort each line upfront. As there is around an hour preparation time before the show, we can do some preparation. Naive implementation: 
`while read -r line;do echo $line | tr ' ' '\n' | sort -n | paste -sd' ' -  ;done < 10k-v2.txt > 10k_sorted.txt`

Takes around 48s on my laptop, which means that for the 10m rows it would be 1000 times that much which is too long.
