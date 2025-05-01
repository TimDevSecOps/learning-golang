// 场景：我们想要一个 Greet 函数，它具有以下需求：
// 	•	可指定 name，默认是 "World"
// 	•	可指定 times，默认是 1
// 	•	可指定是否 shout（大写输出），默认是 false

package main

import (
	"fmt"
	"strings"
)

// 1. 使用 结构体 + 默认值
type GreetOptions struct {
	Name  string
	Times int
	Shout bool
}

// 适合：参数较多、有明确字段名的调用场景
func Greet(opts GreetOptions) {
	name := opts.Name
	if name == "" {
		name = "World"
	}
	times := opts.Times
	if times <= 0 {
		times = 1
	}
	if opts.Shout {
		name = strings.ToUpper(name)
	}
	for i := 0; i < times; i++ {
		fmt.Printf("Hello, %s!\n", name)
	}
}

// 2. 使用 可变参数 + 手动判断
// 可变参数
// 适合：只有 1~2 个可选参数，简单快速
func GreetArgs(args ...string) {
	name := "World"
	if len(args) > 0 && args[0] != "" {
		name = args[0]
	}
	fmt.Printf("Hello, %s!\n", name)
}

// 3. 函数选项模式（Functional Options Pattern）
type greetConfig struct {
	name  string
	times int
	shout bool
}

type GreetOption func(*greetConfig)

func WithName(name string) GreetOption {
	return func(c *greetConfig) {
		c.name = name
	}
}
func WithTimes(times int) GreetOption {
	return func(c *greetConfig) {
		c.times = times
	}
}
func WithShout() GreetOption {
	return func(c *greetConfig) {
		c.shout = true
	}
}

func GreetV2(opts ...GreetOption) {
	conf := &greetConfig{name: "World", times: 1}
	for _, opt := range opts {
		opt(conf)
	}
	name := conf.name
	if conf.shout {
		name = strings.ToUpper(name)
	}
	for i := 0; i < conf.times; i++ {
		fmt.Printf("Hello, %s!\n", name)
	}
}

func main() {
	Greet(GreetOptions{}) // 默认值
	Greet(GreetOptions{Name: "Go"})
	Greet(GreetOptions{Times: 3})
	Greet(GreetOptions{Name: "gopher", Shout: true})

	GreetArgs()           // Hello, World!
	GreetArgs("ZhangSan") // Hello, ZhangSan!

	GreetV2()
	GreetV2(WithName("ZhangSan"))
	GreetV2(WithName("gopher"), WithTimes(3))
	GreetV2(WithName("golang"), WithShout(), WithTimes(2))
}
