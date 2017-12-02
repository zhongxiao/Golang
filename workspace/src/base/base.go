/*
 *  @author zhongxiao.yzx
 *  @date   2017/12/2
 *  @description
 *    basic test method for go language which is imported
 *    by main program; UT program and benchmark programe
 */
// base.go
package base

import (
	"errors"
	"fmt"
)

// the go module initialization method(go builtin)
func init() {
	fmt.Println("base module init")
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("x/0")
	}
	return a / b, nil
}

func PanicRecover() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("recover from a panic %#v\n", v)
		}
	}()
	fmt.Printf("before from a panic\n")
	panic(errors.New("trigger a panic"))
	fmt.Printf("after from a panic\n")
}

func Swap() {
	a, b := 10, 20
	fmt.Printf("a = %v; b = %v!\n", a, b)
	a, b = b, a
	fmt.Printf("a = %v; b = %v!\n", a, b)
}

func RangeEmulate() {
	ints := []int{1, 2, 3, 4}
	for id, v := range ints {
		fmt.Printf("%v %v!\n", id, v)
	}
}
