/*
 *  @author zhongxiao.yzx
 *  @date   2017/12/2
 *  @description
 *   benchmark test
 */
// base_bench_test.go
package base_test

import (
	"base"
	"testing"
)

func Benchmark_Devision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base.Division(4, 5)
	}
}

func Benchmark_TimeConsumingDevision(b *testing.B) {
	//begin 调用该函数停止压力测试时间计数
	b.StopTimer()

	// TODO() ...
	// 执行测试准备，例如读取数据文件，数据库连接，网络连接等
	// 这里的操作不影响测试函数对目标操作的时间统计
	// end

	//[begin : end] 之间的操作不会计入Benchmark函数的时间统计
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.Division(4, 5)
	}
}
