# `bspass`

## Usage

```bash
go build
./bspass              
mantle limiting so mazopathic mo
```

## Options

```bash
./bspass -h
Usage of ./bspass:
  -c	Copy to clipboard
  -copy
    	Copy to clipboard
  -l int
    	Number of characters (default 28)
  -length int
    	Number of characters (default 28)
  -maxc int
    	Maximum number of characters in a word (default 11)
  -minc int
    	Minimum number of characters in a word (default 2)
  -s string
    	Separator string (default " ")
  -separator string
    	Separator string (default " ")
  -v	Enable verbose output
  -verbose
    	Enable verbose output
  -w int
    	Number of words (default 5)
  -words int
    	Number of words (default 5)
```

### Entropy

With the default configuration, this generates passwords with about 130-160 bits of entropy. Verbose mode displays generated password's entropy.

```bash
❯ ./bspass -v
refixture busket athyrid ut ween
Entropy comparison:
  - bspass 📟  : 150.4 bits
  - 123456    : 19.9 bits
  - password  : 37.6 bits
  - hunter2   : 36.2 bits
Generation Time: 5.06 ms
```

If you want to have a more indepth look at the passwords this generates you can use

```bash
for i in $(seq 1 100); do ./bspass; done | column
```


## What is this

This program takes random words from [`/usr/share/dict/words`](https://en.wikipedia.org/wiki/Words_(Unix)) on Macos and puts them in sequence to form a password, conforming with the ideas of the `correct horse battery staple` XKCD. It has 235 976 words to chose from, although some are of lengths not included in generation by default.
