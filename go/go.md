
# Go语言

## 特点：
简洁 快速 安全
并行 有趣 开源
内存管理，数组安全，编译迅速

在Go中，首字母大写的名称是被导出的。
http://go-tour-zh.appspot.com/basics/3

`package main` 包表示它是一个可独立运行的包，它在编译后会产生可执行文件。
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.1.md#详解
除了main包之外，其它的包最后都会生成*.a文件（也就是包文件）并放置在$GOPATH/pkg/$GOOS_$GOARCH中
> 每一个可独立运行的Go程序，必定包含一个package main，在这个main包中必定包含一个入口函数main，而这个函数既没有参数，也没有返回值。

Go使用package（和Python的模块类似）来组织代码。main.main()函数(这个函数位于主包）是每一个独立的可运行程序的入口点。Go使用UTF-8字符串和标识符(因为UTF-8的发明者也就是Go的发明者之一)，所以它天生支持多语言。

:=这个符号直接取代了var和type,这种形式叫做简短声明。不过它有一个限制，那就是它只能用在函数内部；在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量。

_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。    


## 问题
- [ ] make和new的区别 https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md#makenew操作
- [ ] make 用法
- [ ] strconv.Itoa
- [ ] 神奇之处：在java中，必须现有interface，然后才有其继承类；而在go中，可以现有类，然后才有接口（当然，现有接口再有类更是可以的）；https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
- [ ] 反射，laws of reflection http://golang.org/doc/articles/laws_of_reflection.html

## 再多读几遍
- 接口https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md#26-interface
- 并行 https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md

[如何理解 Golang 中“不要通过共享内存来通信，而应该通过通信来共享内存”？ - 知乎](https://www.zhihu.com/question/58004055)
- 

## Go程序设计的一些规则
Go之所以会那么简洁，是因为它有一些默认的行为：
大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

## 那么到底传指针有什么好处呢？
传指针使得多个函数能操作同一个对象。
传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）

## [import](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.3.md#import)
- 点操作
- 别名操作
- `_`操作


## 匿名字段
当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。  
struct不仅仅能够将struct作为匿名字段、自定义类型、内置类型都可以作为匿名字段，而且可以在相应的字段上面进行函数操作

如果内层与继承层都有相同的字段，则：
- 最外层优先访问； 例如：Bob.phone访问的是 Bob 层的字段
- 可以通过匿名字段访问继承层；Bob.Human.phone 访问的是Human层的字段

## method
method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面增加了一个receiver(也就是method所依从的主体)。

```go
type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
```
在使用method的时候重要注意几点：  
- 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
- method里面可以访问接收者的字段
- 调用method通过`.`访问，就像struct里面访问字段一样

指针作为Receiver会对实例对象的内容发生操作,而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作。

那是不是method只能作用在struct上面呢？当然不是咯，他可以定义在任何你自定义的类型、内置类型、struct等各种类型上面。这里你是不是有点迷糊了，什么叫自定义类型，自定义类型不就是struct嘛，不是这样的哦，struct只是自定义类型里面一种比较特殊的类型而已，还有其他自定义类型申明，可以通过如下这样的申明来实现。
`type typeName typeLiteral`
```go
type ages int
type money float32
type months map[string]int
m := months {
	"January":31,
	"February":28,
	...
	"December":31,
}
```

## 面向对象
- 如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
- 如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method

Go里面的面向对象是如此的简单，没有任何的私有、公有关键字，通过大小写来实现(大写开头的为公有，小写开头的为私有)，方法也同样适用这个原则。

## 接口
空`interface`在我们需要存储任意类型的数值的时候相当有用，因为他可以存储任意类型的数值。


## 并发
goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

我们可以看到go关键字很方便的就实现了并发编程。 上面的多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。

runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。

必须使用`make`来创建channel

