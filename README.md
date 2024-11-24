# Assert for golang

## Install
```sh
go get github.com/GrayHat12/assert
```

## Usage
```go
import "github.com/GrayHat12/assert"
...
assert.Assert("This should work", true)
assert.Assert("Assert Test Name", false) // The code will panic and terminate
...
assert.Assert("Name of assertion", someval == expectedVal)

// Takes in a callback that you can customise what to do
// if assertion fails instead of the default panic behaviour
assert.AssertWithCustomCallback("Assertion Test A", condition, func(name string) {
	fmt.Printf("Assertion %s failed. Using custom callback\n", name)
	...
})
```
