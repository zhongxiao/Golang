/*
 *  @author zhongxiao.yzx
 *  @date   2017/12/2
 *  @description
 *    basic test method for go language which is imported
 *    by main program; UT program and benchmark programe
 */
// main
package main

import (
	"base"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	v, ok := base.Division(4, 5)
	if ok == nil {
		fmt.Printf("%f\n", v)
	}
}
