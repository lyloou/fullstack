Go语言里面提供了一个完善的net/http包，通过http包可以很方便的就搭建起来一个可以运行的Web服务。
同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作。

万变不离其宗，Go的Web服务工作也离不开我们第一小节介绍的Web工作方式。

Go为了实现高并发和高性能, 使用了goroutines来处理Conn的读写事件, 这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是Go高效的保证。
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/03.4.md


## 分页
```go
o := orm.NewOrm()
qs = o.QueryTable("user");
var pageSize = 50 // 一页50条
var pageNo = 3 // 准备取第三页
// qs.Limit("[取多少个]","[偏移行数]");
qs = qs.Limit(pageSize, pageNo*pageSize)
```
- [Mysql 分页语句Limit用法 - 柳北风儿~~~~~~~欲宇仙炅 - ITeye博客](http://qimo601.iteye.com/blog/1634748)

## 排序
```go
// 顺序
sort.Ints(a []int);
sort.Strings(a []string);
// 逆序 https://stackoverflow.com/questions/18343208/how-do-i-reverse-sort-a-slice-of-integer-go
sort.Sort(sort.Reverse(sort.IntSlice(keys)))
sort.Sort(sort.Reverse(sort.StringSlice(keys)))
```

## GET & POST
- [Golang Web编程的Get和Post请求发送与解析](http://blog.csdn.net/typ2004/article/details/38669949)
```go

type AddressController struct {
	baseController
}

func (ac AddressController) Address() {
	userId := ac.GetString("user_id")
	var url = cfg.String("third::user_url") + "/user/address?user_id=" + userId
	body, err := get(url)
	if err != nil {
		logs.Error("get Address error:", err)
	}
	ac.Data["json"] = string(body)
	ac.ServeJSON()
}

func (ac AddressController) AddAddress() {
	userId := ac.GetString("user_id")
	var url = cfg.String("third::user_url") + "/user/add_address?user_id=" + userId

	body, err := post(url, ac.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("post AddAddress error:", err)
	}
	ac.Data["json"] = string(body)
	ac.ServeJSON()
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func post(url string, data []byte) ([]byte, error) {
	body := bytes.NewBuffer(data)
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

```