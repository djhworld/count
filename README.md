# count

`uniq -c` requires incoming data to be sorted first, this can be a slow process.

`count` will take any line delimited input from a file, or stdin, and count unique lines. 

## Installation

```
go get github.com/djhworld/count
```

## Warning

`count` will store an entry for each unique item it finds in a map. If every line of your input is unique, memory usage will be the size of the input _plus_ the overhead for storing a key in the map. 

## Example usage

    $ echo "hello\ngoodbye\nhello" | count
    hello	2
    goodbye	1

    $ cat netflix-views.tsv | awk '{ print $1 }' | count
    arrested development	1838
    stranger things	2039
    house of cards	3985

    $ count items.txt
    item1	395
    item2	829
