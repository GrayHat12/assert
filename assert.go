package assert

import "fmt"

type Callback func(name string)

func AssertWithCustomCallback(name string, condition bool, callback Callback) {
	if !condition {
		callback(name)
	}
}

func Assert(name string, condition bool) {
	AssertWithCustomCallback(name, condition, func(assertionName string) {
		panic(fmt.Sprintf("Assertion \"%s\" failed", assertionName))
	})
}
