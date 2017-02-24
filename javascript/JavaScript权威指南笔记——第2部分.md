
| 通过`open()` 的返回值，来获取通过脚本打开的新窗口的window对象。 --p355
  在新窗口window对象里，通过`window.opener`可以拿到来源窗口的window对象。
```js
var w = window.open("http://url1.com");
w.location = "http://url2.com"; // 通过w来修改新窗口的url
```
在新窗口中，可以通过`opener`来操作打开它的脚本的Window对象。
```js
var w2 = window.opener;
w2.location = "http://url3.com"; // 修改引用窗口的url
```

| `window.open()` --p355

| `showModalDialog()`模态对话框：
- [BX9036: Opera 和 Chrome 对模态对话框（showModalDialog）的支持有缺陷，且非 IE 浏览器均不支持非模态对话框（showModelessDialog） - W3Help](http://w3help.org/zh-cn/causes/BX9036)
> showModalDialog 方法与 showModelessDialog 方法均不被 W3C 支持，虽然 showModalDialog 方法目前已比较广泛的被支持，但还是应避免使用它。因为模态对话框会阻塞父窗口的输入，使其失去焦点，这将对用户体验不利。
 - 对于 showModalDialog 方法可以使用模拟模态对话框的方式，比如通过半透明 DIV 元素遮住页面主体，在其之上显示 "对话框 " 内容。
 - 对于 showModelessDialog 方法可以使用 window.open 代替。

- [Chrome不支持showModalDialog模态对话框和无法返回returnValue的问题 - 快乐乔巴 - 博客园](http://www.cnblogs.com/chopper/archive/2012/06/25/2556266.html)

| History对象：--p345
- `back()` 后退；
- `forward()` 前进；
- `go()` 接受整数参数，正参数向前，负参数向后跳转任意多个页。


| 载入新的文档：--p334
- `location.assign("./newFile.html")` 可后退；
- `location.replace("./newFile.html")` 不可后退；
- `location.reload()` 重新加载；

| `setTimeout()`和`setInterval()`的第一个参数可以作为字符串传入（相当于执行`eval()`）。--p342
```js
setTimeout("alert('Hello World!')", 3000);
```

| `setInterval()`和`setTimeout()`类似，只不过会在指定毫秒数的间隔里重复调用。
也可以通过`clearInterval()`来取消后续函数的调用。--p342

| `setTimeout()`设置一个函数在指定的毫秒后执行，返回一个值，这个值可以传递给`clearTimeout()`用于取消这个函数的执行；--p342

## 第十四章：Window对象


| 使用框架的好处： --p340
- 用更简洁的代码完成更复杂的功能；
- 完善的框架已经对兼容性、安全性和可访问性进行了处理；

| 客户端框架：对Web浏览器提供的标准和专用API进行了封装向上提供更高级的API，用以高效地进行客户端编程。--p339

| 脚本本身的来源和同源策略并不相关，相关的是脚本所嵌入的文档的来源（脚本所依赖的宿主环境）。

| 针对安全性问题，JavaScript对一些能力和和已有功能做了限制： --p335
- 不支持一些功能（没有权限操作计算机上的文件和目录）。
- 没有通用的网络能力。
- 限制打开窗口的能力：只是为了响应鼠标单击这样的用户触发事件才使用。
- 同源策略。

| JavaScript的角色应当是增加信息的表现力，而不是信息的呈现。--p334

| JavaScript检测文档以何种模式进行渲染，通过`document.compatMode`来判断： --p325
- 如果值为`CSS1Compat`，说明浏览器工作在标准模式。
- 如果值为`BackCompat`或其他情况，说明浏览器工作在怪异模式。

| [document.readyState - Web API 接口 | MDN](https://developer.mozilla.org/zh-CN/docs/Web/API/Document/readyState)
> 当document文档正在加载时,返回"loading"。当文档结束渲染但在加载内嵌资源时，返回"interactive"，并引发DOMContentLoaded事件。当文档加载完成时,返回"complete"，并引发load事件。

| JavaScript线程模型：单线程执行的机制（让编程变得简单）。--p324
在H5中，通过WebWorker实现了并发的控制方式（可以计算密集任务而不冻结用户界面的后台线程--p325）。

| 事件驱动阶段里发生的第一个事件是load事件，指示文档已经完全载入，并可以操作。--p320

| JavaScript的执行包括两个阶段： --p320
1. 载入文档——按照在文档中出现的顺序执行。
2. 事件驱动——
  - 用户输入（鼠标点击或经过、键盘按下）
  - 网络活动
  - 运行时间
  - JavaScript错误

| 在标签script中使用`src`属性时，将会忽略`<script>`和`</script>`标签之间的任何内容。 --p315

| JavaScript文件的扩展名通常是以`.js`结尾的。包含纯粹的JavaScript代码，其中既没有`<script>`标签，也没有其他HTML标签。 --p315

|　在HTML里嵌入JavaScript代码： --p313
1. 内联，放置在`<script>`和`</script>`标签对之间。
2. 放置在`<script>`标签的`src`属性指定的外部文件中。
3. 放置在HTML事件处理程序中，例如：onclick、onmouseover。
4. 放在URL中，这个URL使用`javascript:`协议（bookmarklet）。

| Window有一个指向自身的对象叫做`window`。--p310

| Window对象处于作用域链的顶部，它的属性和方法实际上是全局变量和全局函数。--p310

## 第十三章Web浏览器中的JavaScript



## 书中的错误

| p336
> 对于防止脚本窃取似有的信息来说 => 对于防止脚本窃取私有的信息来说

| p348
> cookieEnable()  => cookieEnabled
