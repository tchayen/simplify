# simplify

<a href="https://travis-ci.org/tchayen/simplify"><img src="https://travis-ci.org/tchayen/simplify.svg?branch=master"></a>
<a href='https://coveralls.io/github/tchayen/simplify?branch=master'><img src='https://coveralls.io/repos/github/tchayen/simplify/badge.svg?branch=master'/></a>

Implementation of Visvalingam’s algorithm for line simplification.

Runs in `O(n*log(n))` time.

## Installation
```bash
go get github.com/tchayen/simplify
```

## Usage
> _🚧 The lib is currently in progress, usage is not recommended yet 🚧_

## TODO
- decide what metrics should be available to user (i.e. target percentage of
points, target fixed number of points etc.)
- design API with different use cases in mind (both processing and outputting
filtered arrays)
- provide examples (maybe some interactive website?)

## API draft

```go
// Process points without filtering:
Process(points, simplify.NORMALIZE) // Put Z-coordinate in range [0,1]
Process(points, simplify.AREAS) // Use Z-coordinate for accurate area storage.

// Simply return array of points:
p, err := Simplify(points, T{percentage: 30})
p, err := Simplify(points, T{amount: 1520})
p, err := Simplify(points, T{area: 1.5})
```

## Contributing

Feel free to join! Start with posting an issue with a functionality that is
missing in your opinion (or with suggestion for some improvement).

Then we will discuss how to change that into living code.