| JavaScript检测文档以何种模式进行渲染，通过`document.compatMode`来判断：
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
