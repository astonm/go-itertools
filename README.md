# go-itertools - Python's itertools for Go iterators

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-007d9c)
[![License](https://img.shields.io/github/license/astonm/go-itertools)](./LICENSE)

**`astonm/go-itertools` is a library styled on the Python [itertools](https://docs.python.org/3/library/itertools.html) library which works using [Go 1.23+ iterators](https://pkg.go.dev/iter).**

By convention, all of the functions in this library return either [`iter.Seq`](https://pkg.go.dev/iter#Seq) or [`iter.Seq2`](https://pkg.go.dev/iter#Seq2). Almost all of the functions operate on [`iter.Seq`](https://pkg.go.dev/iter#Seq) as input. The most notable exception is for the combinatoric functions which take a slice, constraining their use to finite sequences.

## Install

```sh
go get github.com/astonm/go-itertools
```

## Usage


This package can be imported under its default name of `itertools` like so:

```go
import "github.com/astonm/go-itertools"
```

Then use one of the helpers below, e.g.

```go
counter := itertools.Count()
// 0, 1, 2, ...
```

If you are using a large number of library functions or are nesting calls, the lengthy `itertools` can be a bit much. When brevity is better than clarity, we recommend the short name `it`, i.e.


```go
import it "github.com/astonm/go-itertools"
```

Then use like so:


```go
cycle := it.Cycle(it.FromSlice([]int{1,2,3}))
// 1, 2, 3, 1, 2, 3, ...
```

## Documentation

GoDoc: [https://pkg.go.dev/github.com/astonm/go-itertools](https://pkg.go.dev/github.com/astonm/go-itertools)

## License

Copyright 2024 [Aston Motes](https://github.com/astonm).

This project is under [MIT](./LICENSE) license.

