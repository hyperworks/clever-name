Library generates clever user names for people
e.x Spicy-fox, slow-rhino, purple-horse


To use
```
import clever "github.com/hyperworks/clever-name"
var cleverDict *clever.Clever

func init() {
	cleverDict := clever.InitDict()
	cleverDict.EnableExtraNums = true //If you have a lot of users append random numbers
}

...

username := cleverDict.GetUsername()
```

=====
To run tests

```
export GOPATH=`pwd`/lib
github.com/stretchr/testify/assert
go test clever_test.go clever.go
```