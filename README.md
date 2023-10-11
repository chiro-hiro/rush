# rush

Rust in Go = Rush

## Install

```bash
go get github.com/chiro-hiro/rush
```

### Usage

Ordinary code

```go
package main

func findValue[T int32 | float32](slices []T, search T) (T, error) {
  for _, slice := range slices {
    if slice == search {
      return slice, nil
    }
  }
  return 0, errors.New("not found")
}

func main(){
  slice := []int32{3, 4, 5, 6, 8, 9, 11}
  find, err := findValue(slice, 8)
  if err != nil {
    println(err.Error())
  } else {
    println(find)
  }
}
```

Rewrite with rush

```go
package main

import (
  "errors"

  "github.com/chiro-hiro/rush/option"
  "github.com/chiro-hiro/rush/result"
)

func findValueWithResult[T int32 | float32](slices []T, search T) result.Result[T] {
  for _, slice := range slices {
    if slice == search {
      return result.Ok(slice)
    }
  }
  return result.Err[T](errors.New("not found"))
}

func main() {
  slice := []int32{3, 4, 5, 6, 8, 9, 11}
  findWResult := findValueWithResult[int32](slice, 8).Unwrap()
  println(findWResult)
}

```

### Wrap go function

You can wrap Go function with `result.From[T]` and `option.From[T]`

```go
findWrapResult := result.From(findValue[int32](slice, 8))
println(findWrapResult.Unwrap())
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details

**built with ❤️**
