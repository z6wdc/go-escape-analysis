# Escape Analysis

A small project for quickly experimenting with Go's **escape analysis**.

## Build

Run in the project root to inspect escape analysis results:

```bash
go build -gcflags -m=2 github.com/z6wdc/go-escape-analysis/examples
```

## How to Read Escape Analysis Output

When running with `-gcflags -m=2`, the Go compiler prints detailed escape analysis and inlining diagnostics.  
Here’s an example:

```
# github.com/z6wdc/go-escape-analysis/examples
examples/escape.go:3:6: can inline ReturnPointer with cost 8 as: func() *int { x := 42; return &x }
examples/escape.go:8:6: can inline ReturnValue with cost 7 as: func() int { x := 42; return x }
examples/escape.go:4:2: x escapes to heap:
examples/escape.go:4:2:   flow: ~r0 = &x:
examples/escape.go:4:2:     from &x (address-of) at examples/escape.go:5:9
examples/escape.go:4:2:     from return &x (return) at examples/escape.go:5:2
examples/escape.go:4:2: moved to heap: x
```

### Explanation

- **`can inline ReturnPointer with cost 8 as: ...`**  
  The function `ReturnPointer` can be inlined (its body can replace the call site) with a cost of 8.  
  The compiler provides the inlined version.

- **`can inline ReturnValue with cost 7 as: ...`**  
  Similarly, `ReturnValue` can also be inlined with a cost of 7.

- **`x escapes to heap:`**  
  The local variable `x` cannot safely remain on the stack. It must be allocated on the heap.

- **`flow: ~r0 = &x:`**  
  `~r0` represents the return value slot of the function.  
  The analysis shows that the return value is assigned `&x`.

- **`from &x (address-of)` / `from return &x (return)`**  
  Traces how the value flows: you took the address of `x` (`&x`), then returned it.  
  This makes `x` live beyond the function call.

- **`moved to heap: x`**  
  Final decision: the variable `x` will be heap-allocated.

### In short
The compiler is saying: *"You returned the address of a local variable.  
To keep that pointer valid after the function returns, I must move `x` to the heap."*
