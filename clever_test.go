package clever

import (
	"fmt"
	a "github.com/stretchr/testify/assert"
	"testing"
	"strings"
)


func TestDictLoaded(t *testing.T) {
	c := InitDict()
	u1 := c.GetUsername()
	dashes := strings.Count(u1, "-")
	a.Equal(t, dashes, 1)
	fmt.Printf("Got username: %s\n", u1)

	c.EnableExtraNums = true
	u2 := c.GetUsername()
	dashes2 := strings.Count(u2, "-")
	fmt.Printf("Got username: %s\n", u2)
	a.Equal(t, dashes2, 2)
}