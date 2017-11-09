[简洁的格式化](https://golang.org/doc/effective_go.html#formatting)
```
Some formatting details remain. Very briefly:

Indentation
    We use tabs for indentation and gofmt emits them by default. Use spaces only if you must.
Line length
    Go has no line length limit. Don't worry about overflowing a punched card. If a line feels too long, wrap it and indent with an extra tab.
Parentheses
    Go needs fewer parentheses than C and Java: control structures (if, for, switch) do not have parentheses in their syntax. Also, the operator precedence hierarchy is shorter and clearer, so
    x<<8 + y<<16
```

If you hava a question about how to approach a problem or 
how something might be implemented, the documentation, code 
and examples in the libray can provide answers, ideas and background.
https://golang.org/doc/effective_go.html

Executable commands must always use `package main`   
https://golang.org/doc/code.html#PackageNames

Go没有类。然而，仍然可以在结构体类型上定义方法。

一个结构体（struct）就是一个字段的集合

## 指针：
`&` 符号会生成一个指向其作用对象的指针；
`*` 符号表示指针指向的底层的值；
通过 `go build` 来编译，可以来检验 lib 函数是否有错误，不会生成文件到 pkg 中去。
在执行 `go install` 的时候，才会生成到 pkg 中。

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

## The Command
- `go env GOPATH` prints the effective current GOPATH; it prints the default location if the environment variable is unset.

