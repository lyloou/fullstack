

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
