# Hello World Genetic Algorithm

## Introduction

https://www.mathworks.com/help/gads/what-is-the-genetic-algorithm.html

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [References](#references)
- [TODO](#todo)

## Installation

Make sure you have a working Golang env, then:

```bash
go get github.com/rms1000watt/hello-world-genetic-algorithm
cd $(go env GOPATH)/src/github.com/rms1000watt/hello-world-genetic-algorithm
```

## Usage

```bash
go run main.go
```

You should see an output like:

```
Population Size: 10000
Iterations: 1000
Start Grade: 49
Done Grade: 85
Run Time: 3.43785641s
```

## References

- https://lethain.com/genetic-algorithms-cool-name-damn-simple/

## TODO

- [ ] Create concurrent example
