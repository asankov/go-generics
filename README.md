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

## Slides

To download the slides from this presentation use this [link](https://drive.google.com/file/d/1x7vfwC7hwwIgakRlZ0eBUO3nGiDMsOZ2/view?usp=sharing).
