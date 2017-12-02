// base_test
package base_test

import (
	"base/set"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var hash_set *set.HashSet = set.NewHashSet()

	setI := set.NewHashSet()
	fmt.Println("%s", setI.String())
	if setI == nil {
		t.Fail()
		return
	}
	t.Log("suc")
}
func TestFail(t *testing.T) {
	t.Log("fail")
}
func TestError(t *testing.T) {
	t.Log("error")

}
