Go没有类。然而，仍然可以在结构体类型上定义方法。

一个结构体（struct）就是一个字段的集合

## 指针：
`&` 符号会生成一个指向其作用对象的指针；
`*` 符号表示指针指向的底层的值；


## 零值：
数值：0
布尔：false
指针：nil

常量使用关键字`const`表示，不可用`:=`语法定义。

var 语句定义了一个变量列表；
`:=`结构不能使用在函数外，函数外的每个语句都必须以关键字`var`, `func`开始。

## 进一步阅读：
https://blog.golang.org/gos-declaration-syntax


## 还存在的问题：
sampleZ()方法
http://go-tour-zh.appspot.com/methods/9

range和close
http://go-tour-zh.appspot.com/concurrency/4

where to go from here
http://go-tour-zh.appspot.com/concurrency/10
