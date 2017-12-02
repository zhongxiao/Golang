/*
 *  @author zhongxiao.yzx
 *  @date   2017/12/2
 *  @description
 *   unit test
 */

// base_test
package base_test

import (
	"base"
	"testing"
)

func TestDivision(t *testing.T) {
	v, ok := base.Division(4, 5)
	t.Logf("%f", v)
	if ok != nil || v != 0.8 {
		t.Fail()
		t.Log("failure")
	}
}
