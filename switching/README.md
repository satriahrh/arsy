# switching

## Get started

Add to your dependency

```shell script
go get -u github.com/satriahrh/arsy/switching 
```

## Features

- Sequential multiple switch case: `MultipleCase` [go to example](#multiplecase)
- Sequential multiple switch case with error: `MultipleHeavyCase` [go to example](#multipleHeavycase)

### `MultipleCase`

In traditional `switch case`, only single case can be executed.
Make use of `fallthrough` cannot give you power to evaluate the case expression.
`MultipleCase` allow you to do so.
 
Examine this following example
```go
package main

import (
  "fmt"

  "github.com/satriahrh/arsy/switching"
)

func main()  {
  value := 2
  switching.MultipleCase(
    switching.NewCase(
      func() bool {
        return value < 4
      },
      func() {
        fmt.Println("case 1")
      },
    ),
    switching.NewCase(
      func() bool {
        return value == 2
      },
      func() {
        fmt.Println("case 2")
      },
    ),
  )
}
```

Above code example would return following output on your console.
```
case 1
case 2
```

`case 1` is printed before `case 2` because those cases is evaluated sequentially from the first case, second case, and so on.

### `MultipleHeavyCase`

Examine this following example
```go
package main

import (
  "errors"
  "fmt"

  "github.com/satriahrh/arsy/switching"
)

func main() {
  value := 2
  err := switching.MultipleHeavyCase(
  	switching.NewHeavyCase(
      func() (bool, error) {
        return value < 4, nil
      },
      func() error {
        fmt.Println("case 1")
        return errors.New("an error occurred")
      },
    ),
    switching.NewHeavyCase(
      func() (bool, error) {
        return value == 2, nil
      },
      func() error {
        fmt.Println("case 2")
        return nil
      },
    ),
  )
  if err != nil {
    fmt.Println(err)
  }
}

```

Above code example would return following output on your console.
```
case 1
an error occurred
```

`case 1` is still printed because `fmt.Println("case 1")` is still at the safe flow.
Right after that, it return an error which the `MultipleHeavyCase` is breaking the loop right away if it got an error from `command` or `evaluator`, even though the next condition evaluated true.
