# avg

Takes arguments and calculates average or calculates the amount (i.e. count) to get that specific average.
* The first argument without ":" is taken as the expected price. If provided, it calculates the amount to be bought to
get a specific average.
* If all arguments provided is `price:count` pair, then it calculates the average.

## Installation

```shell
go build -o avg
cp avg /usr/local/bin/
```

## Run

```shell
avg 350 500:x 200:10 # prints 10

avg 500:10 200:10 # prints 350
```
