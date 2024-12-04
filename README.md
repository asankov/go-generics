# The Inner Workings of Go Generics

Repo for the materials for my presentation on The Inner Workings of Go Generics

## Build

To build the code and output a binary named `generics`:

```shell
go build -gcflags="-N -l" -o generics main.go
```

`-gcflags` is flags that will be passed to the compiler, e.g. `go tool compile`.

`-N` means disable optimization.
We use this because otherwise the compiler might see that we don't use the results of some of the functions and just delete them.

`-l` means disable inlining.
We use this so that we can see the compiled functions which otherwise might be inlined.

## Inspect

We can the use `go tool objdump` to disasemble the binary and look at the disassembled machine code:

```shell
go tool objdump generics > objdump
```

### Notable things

Interesting things that we observe in the dissambled machine code:

#### Number of `PrintAndReturn` function declarations

We can find three `PrintAndReturn` function declarations at the dissambled code.

- `main.PrintAndReturn[go.shape.*uint8]`
- `main.PrintAndReturn[go.shape.string]`
- `main.PrintAndReturn[go.shape.int]`

This is consistent to what we know about generics in Go.

The compiler has generated one function for every different GC Shape (memory type) that we invoke the `PrintAndReturn` function with.

This call:

```go
_ = PrintAndReturn(printableInt(1))
```

generates this function - `main.PrintAndReturn[go.shape.int]`.

(the compiler treats `int` and `printableInt` the same, because they have the same memory shape)

This call:

```go
_ = PrintAndReturn(printableString("string"))
```

generates this function - `main.PrintAndReturn[go.shape.string]`.

(the compiler treats `string` and `printableString` the same, because they have the same memory shape)

And these two calls:

```go
_ = PrintAndReturn(&A{})
_ = PrintAndReturn(&B{})
```

generate this function - `main.PrintAndReturn[go.shape.*uint8]`.

Even though `A` and `B` are different types with different memory shapes, we are passing them via pointers, which are actually the same type (`*uint8`).

## Slides

To download the slides from this presentation use this [link](https://drive.google.com/file/d/1x7vfwC7hwwIgakRlZ0eBUO3nGiDMsOZ2/view?usp=sharing).
