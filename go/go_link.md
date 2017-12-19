
- [x] http://fuxiaohei.me/2016/9/20/go-and-http-server.html

- [ ] http://wiki.jikexueyuan.com/project/go-command-tutorial/0.0.html

- [x] https://www.calhoun.io/5-tips-for-using-strings-in-go-2/

## [零值](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md#%E9%9B%B6%E5%80%BC)
## [fmt格式](https://golang.org/pkg/fmt/)

## [make & new](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md#makenew%E6%93%8D%E4%BD%9C)
> 1. new返回指针。  
> new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。
> 2. make返回初始化后的（非零）值。  
> make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。 在这些项目被初始化之前，slice为nil。

