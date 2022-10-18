# tern

A ternary package for Go

## Why

Go is missing one crucial feature that C even has, ternary.\
Go author have no intention to add this feature for the potential "abuse" they
said.

## Example

```go
PORT = tern.E(production, 80, 8000)
```
