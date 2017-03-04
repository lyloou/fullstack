

| 缓动函数，默认的缓动函数有“swing”，还有线性缓动函数"linear"。
还可以自定义缓动函数：`jQuery.easing['squareroot'] = Math.sqrt` --p547


| 动画选项：
- 默认情况下，所有动画都是队列化的。（将`queue`属性设置为false可以取消队列化） --p546

| 动画只支持数值属性：对于颜色，字体或display等枚举属性是无法实现动画效果的。
（jQuery其他实现方案：传入自定义的CSS属性变化函数） --p545

| jQuery禁用动画：
```js
jQuery.fx.off = true;
```

| 实时事件依赖于事件冒泡。--p541

| `unbind()`只注销用`bind()`和相关jQuery方法注册的事件处理程序。--p536

| jQuery另一种事件处理程序的方法：`one()`，工作原理同`bind()`一样，
除此之外，使用`one()`注册的事件处理器永远只会触发一次。--p536

| 调用`$('a').hover(f)`等同于调用`$('a').bind('mouseenter mouseleave', f)` --p536

| 可以通过直接使用`bind()`方法来使用事件注册的高级特性。 --p535
- 最简形式：`$('p').bind('click', f)`
- 三个参数：中间一个参数作为Event的data属性。（方便进行数据交互）
- 为多个时间类型同时注册程序处理函数。（用空格分开：`$('a').bind('mouseenter mouseleave', f)`）
- 允许为注册的事件处理程序指定命名空间。f
- 第一个参数可以是对象。`$('a').bind({mouseenter:f, mouseleave:g})`

| jQuery事件处理程序
需要知道的最重要的一件事情是，每个事件处理程序都传入一个jQuery事件对象作为第一个参数。--p533

| 注意，在`replaceAll()`运行后，该jQuery对象中的元素将不再存在于文档中。--p528

| 注意：传递对象给`data()`时，该对象的属性将替换掉与元素相关联的旧数据。
  与很多其他setter方法不同，data()不接受函数参数。当将函数作为第二参数传递给data()时，
  该函数会存储，就和其他值一样。--p527

| jQuery只会把定位元素作为偏移父元素，
jQuery对象的offsetParent()方法则会把每个元素映射到最近的定位祖先元素或`<body>`元素。--p526

| css()方法允许在CSS样式名中使用连字符（"background-color"）
或使用驼峰格式JavaScript样式名（"backgroundColor"） --p523

| jQuery的setter和getter
- `attr()`
- `css()`
- `val()`
- `text()/html()`
- `offset()`
- `data()`

| 查询
- `$("div").each()`
- `$("div").map()`
- `$("div").index()`
- `$("div").is()`

| jQuery对象的三个有趣的属性 --p520
- selector: 创建jQuery对象时的选择器字符串（如果有的话）；
- context: 上下文对象，默认是document对象；
- jquery: 通过判断该属性是否存在来与其他类数组对象区分开。（值是字符串形式的jQuery版本号）

| $()的返回值是一个jQuery对象。jQuery对象是类数组：
他们拥有length属性和介于0~length-1之间的数值属性。--p519

| jQuery术语
- jQuery对象：是由jQuery函数返回的对象。
- jQuery函数：定义在jQuery命名空间中的函数。 --p519
- jQuery方法：由jQuery函数返回的jQuery对象的方法。

| jQuery工具函数（又称静态方法，命名空间函数）：
- `jQuery.noConflict()`
- `jQuery.each()`
- `jQuery.parseJSON()`


| jQuery调用方式 --p516
- 传递css选择器给`$()`方法；
- 传递一个Element、Document或Window对象给`$()`方法；
- 传递HTML文本字符串给`$()`方法；
- 传入一个函数给`$()`方法；

| `jQuery()`是工厂函数，不是构造函数，它返回一个新创建的对象，但并没有和new关键字一起使用。

## 第十九章：jQuery类库

| XMLHttpRequest规范列出了这个令人困惑名字的不足之处： --p486
对象名`XMLHttpRequest`是为了兼容Web，虽然这个名字的每个部分都可能造成误导。
首先，这个对象支持包括XML在内的任何基于文本的格式。
其次，它能用于http和https请求。
最后，它所支持的请求是一个广义的概念，指的是对于定义的HTTP方法的设计HTTP请求或响应的所有活动。

| 使用JSONP，JSON响应数据（理论上）是合法的JavaScript代码，当它到达时浏览器将执行它。--p506


| 使用`<script>`元素进行Ajax传输的一个主要原因是，它不受同源策略的影响，
因此可以使用他们从其他服务器请求数据。
第二个原因是包含JSON编码数据的响应体会自动解码（即，执行）

| 测试`withCredentials`的存在性是测试浏览器是否支持CORS的一种方法。--p504

| 对于XMLHttpRequest对象x，设置x.progress以监控响应的下载进度，
并且设置x.upload.onprogress以监控请求的上传进度；--p501

| HTTP进度事件，Event对象属性
- type
- timestamp
- loaded
- total
- lengthComputable
```js
request.onprogresss = function (e) {
  if(e.lengthComputable) {
    progress.innerHTML = Math.round(100*e.loaded/e.total) + "% Complete";
  }
}
```

| 测试浏览器是否支持progress事件：
```js
if("onprogress" in (new XMLHttpRequest())) {
  // 支持progress事件
}
```

| 对于任何具体的请求，浏览器将只会触发load、abort、timeout和error事件中的一个。--p500

| HTTP请求无法完成有3种情况：
- 请求超时，会触发`timeout`事件；
- 请求中止，会触发`abort`事件；
- 网络错误（例如：重定向），会触发`error`事件；

| 如果上传的主体是XML文档或纯文本，那么可以不用显示设置HEAD；--p498

| 表单数据编码格式有一个正式的MIME类型：
```js
application/x-www-form-urlencoded
```

| 对表单数据使用的编码方案：--p495
对每个表单元素的名字和值执行普通的URL编码，使用等号把编码后的名字和值分开，
并使用“&”符号分开名/值对。

| 可以使用`setOverrideMimeType()`让`XMLHttpRequest`知道它不需要把文件解析成XML文档：
--p494
```js
request.overrideMimeType("text/plain; charset=utf-8");
```

| 记住`<script>`元素能发起跨域HTTP请求，而XMLHttpRequest API则禁止。 --p494

| XMLHttpRequest也支持同步响应。如果把false作为第三个参数传递给`open()`，
那么`send()`方法将阻塞直到请求完成。 --p492

| 用POST方法发送纯文本给服务器 --p490
```js
function postMessage(msg) {
  var request = new XMLHttpRequest();
  request.open("POST", "/log.php");
  request.setRequestHeader("Content-Type", "text/plain;charset=UTF-8");
  request.send(msg);
}
```

| HTTP架构：请求/响应 --p488

一个HTTP请求由4部分组成：
- HTTP请求方法或“动作”（verb）
- 正在请求的URL
- 一个可选的请求主体

服务器返回的HTTP响应包含3部分：
- 一个数字和文字组成的状态码，用来显示请求的成功和失败
- 一个响应头集合
- 响应主体

| 虽然它的名字叫`XMLHttpRequest API`，但并未限定只能使用XML文档，
它能获取任何类型的文本类型。--p486

| 一段时间以来，所有浏览器都支持XMLHttpRequest对象，他定义了用脚本操纵HTTP的API。
除了常用的GET请求，这个API还包含实现POST请求的能力，
同时它能用文本或Document对象的形式返回服务器的响应。 --p486


| 在Ajax中，客户端从服务端“拉”数据，而在Comet中，服务端向客户端“推”数据。--p484

| Web应用可以使用Ajax技术把用户的交互数据记录到服务器中，也可以开始只显示简单页面，
之后按需加载额外的数据和页面组件来提升应用的启动时间。--p484

| 术语Ajax描述了一种主要使用脚本操纵HTTP的Web应用架构。
Ajax应用的主要特点是使用脚本操纵HTTP和Web服务器进行数据交换，不会导致页面重载。--p484

## 第十八章：脚本化HTTP

| 事件传播
- 事件冒泡是事件传播的第三个“阶段”；
- 目标对象本身的事件处理程序调用是第二个阶段；
- 第一个阶段甚至发生在目标处理程序调用之前，称为“捕获”阶段；

| 通过时间处理程序的返回值来实现：确认离开当前页？
```js
window.onbeforeunload = function () {
  console.log("==> onbeforeload");
  return confirm("确认跳转？");
}
```

| 相对`addEventListener()`的是`removeEventListener()`方法。--p453

| 能通过多次调用addEventListener()为同一个对象注册同一事件类型的多个处理程序函数。
当对象上发生事件时，所有该事件类型的注册处理程序都会按照注册的顺序调用。
使用相同的参数在同一个对象上多次调用addEventListener()是没用的，处理程序仍然只注册一次，
同时重复调用也不会改变调用处理程序的顺序。 --p453

| 按照约定，事件处理程序属性的名字由“on”后面跟着事件名组成：nclick、onchange、onload、
onmouseover等。注意这些属性名是区分大小写的，所有都是小写，即使事件类型是由多个词组成。--p451

| 客户端编程的通用风格是保持HTML内容和JavaScript行为分离，
遵循这条规则的程序员应禁止（或至少避免）使用HTML事件处理程序属性，
因为这些属性直接混合了JavaScript和HTML；

| 注册事件处理程序： --p451
- 给事件目标对象或文档元素设置属性（Window对象、HTML标签属性）；
- 将事件处理程序传递给对象或元素的一个方法；（`addEventListenrer()`）

| 移动设备事件： --p450
- gesturestart
- gestureend
- gesturechange
- 属性：scale、rotation
- touchstart
- touchend
- touchmove
- 属性：changedTouches
- 旋转事件：orientationchanged事件


| DOM事件、HTML5事件、触摸屏和移动设备事件。

| 跨文档通信API允许一台服务器上的文档脚本能和另一台服务器上的文档脚本交换信息。
其工作受限于同源策略这一安全方式。 --p449

| textinput事件不是键盘特定事件，无论通过键盘、剪切和粘贴、拖放等方式，
每当发生文本输入时就会触发它。--p448

| 键盘事件 --p447
- keyCode
- keyDown
- keyUp
- keyPress
- altKey
- ctrlKey
- metaKey
- shiftKey

| 鼠标事件 --p446
- mousedown
- mouseup
- mouseover
- mousewheel(鼠标滚轮)

| Window事件 --p445
- load
- onload
- onerror
- abort
- focus、blur
- resize
- scroll



| 表单事件 --p444
- submit
- reset
- change
- focus
- blur
- input(键盘、剪切粘贴等文本操作会触发，另一个替代方案：textinput)

## 第十七章：事件处理

| CSS动画
原理：使用`setTimeout()`和`setInterval()`重复调用函数来修改元素的内联样式达到目的。--p429

| 开启和关闭样式表：--p436
```js
document.styleSheets[0].disabled = true; // 关闭样式表
document.styleSheets[0].disabled = false; // 开启样式表
```

| overflow属性允许指定当内容超出元素边框时该如何显示，而clip属性确切地指定了应该显示元素的
哪个部分，它不管元素是否溢出。 --p424

| 颜色值：
颜色值由16进制来表示红绿蓝（RGB）。
每个颜色最低位为0（16进制为00），最高位为255（16进制为FF）
16进制值的写法为#号后面跟三个或六个十六进制字符。
三位数表示法：`#RGB`，转换为6位数表示法：`#RRGGBB`
- [HTML 颜色值\_w3cschool](http://www.w3cschool.cn/html/html-colorvalues.html)

| visibility和display的区别：
对于一个常规布局流中的元素，设置visibility属性为hidden使得元素不可见，
但是在文档布局中仍保留了它的空间。
如果给display设置了none属性，在文档布局中不再给他分配空间，它各边的元素会合拢，
就当作不存在。--p422

| 边框盒模型：
通过给元素指定box-sizing属性`box-sizing:border-box`（默认是`box-sizing:content-box`）
浏览器会为元素应用IE的盒模型，即width和height属性将包含边框和内边距。--p422

| [CSS 盒子模型\_w3cschool](http://www.w3cschool.cn/css/css-boxmodel.html)

| 第三个维度：`z-index`
当两个或多个元素重叠在一起的时候，他们是按照从低到高的`z-index`顺序绘制的。
如果值一样，那么按照在文档中出现的顺序绘制，即最后一个重叠的元素显示在最上面。--p418
注意：
`z-index`只对兄弟元素应用堆叠效果。

| 可以通过指定`width`和`height`属性来设置元素的宽度和高度。另外可以通过设置
`left`、`right`、`top`、`bottom`来指定，
注意：
如果同时指定了left、right和width，那么width属性将覆盖right属性。
如果同时指定了top、bottom和height，那么height属性将覆盖bottom属性。 --p418


| 静态定位`static`的元素不能使用top、left和类似其他属性定位。
欲对文档元素使用CSS定位技术，必先将其position属性设置为除此之外的其他三个属性值：
`absolute、fixed、relative`  --p417

| 非标准属性 --p413
在属性名前加一个厂商前缀。
- Firefox使用`-moz-`
- Chrome使用`-webkit-`
- IE使用`-ms-`
- 例如：
```css
.radius10 {
  border-radius: 10px;
  -moz-border-radius: 10px;
  -webkit-border-radius: 10px;
}
```

## 第16章：脚本化CSS

| 开启文档编辑模式： --p408
- `document.body.contenteditable=true`
- `document.designMode="on"`

| `document.write()` 只有在解析文档时才能使用`write()`方法输出到HTML到文档中。--p405

| 密码输入元素只能防止眼睛窥视，但在提交表单时未经任何加密。 --p402

| [HTML(5) 不要求标签自闭合 – 大魔 I'm Png](http://www.impng.com/web-dev/html-tags-without-self-closing.html)
`input`标签不要求自闭合。

| 表单元素在收到键盘的焦点时会触发focus事件，失去焦点时会触发blur事件。 --p400

| onsubmit事件处理程序只能通过单击“提交”按钮来触发。直接调用表单的submit()方法不触发onsubmit事件处理程序。
  onreset事件处理程序只能通过单击“提交”按钮来触发。直接调用表单的reset()方法不触发onreset事件处理程序。 --p400

| HTML表单
在服务器程序中，表单必须有一个“提交”按钮，否则它就没有用处。
在客户端编程中，“提交”按钮不是必须的。
服务端程序是基于表单提交动作的，客户端程序是基于事件的。 --p396

| scrollBy滚动偏移量 --p393
```js
// 每隔200毫秒向下滚动10像素
var result = setInterval(function(){scrollBy(0, 10)}, 200);
// 停止滚动
clearInterval(result);
```

| 如果调用`appendChild()`或`insertBefore()`将已存在文档中的一个节点再次插入，
那个节点将自动从当前的位置删除并在新的位置重新插入：没有必要显示删除该节点。
（也就是说可以通过这两个方法重新排列节点） --p383

| 对于没有指定`src`属性的`script`标签，是一个嵌入任意文本内容的理想地方。 --p380
```js
<script id="sscript">
  我是脚本里的内容。
</script>

var ss = document.getElementById("sscript");
console.log(ss.text);
```

| 只有`Element`节点定义了`outerHTML`属性，`Document`节点则无。 --p379

| 对innnerHTML属性用`+=`操作符重复追加一小段文本通常效率低下，因为它既要序列化又要解析。 --p379

| 元素的内容
- 使用HTML表示: `innerHTML()`
- 纯文本表示: `innerText()/textContent()`，注意两者的支持情况 --p380
- 树状表示；`nodeValue()`、`nodeType`

| nodeType  --p372
- 1 代表Element节点
- 3 代表Text节点
- 8 代表Comment节点
- 9 代表Document节点
- 11 代表DocumentFragment节点

| DOM选取元素：--p364
- 用指定的id属性：`document.getElementById('id');`
- 用指定的name属性：`document.getElementsByName('name');`
- 用指定的标签名字：`document.getElementsByTagName('span');`
- 用指定的CSS类：`document.getElementsByCLassName('warning')`
- 匹配指定的CSS选择器：`document.querySelector()`、`document.querySelectorAll()`
  （中间用逗号`,`分割）

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

| 错误处理
```js
<script>
window.onerror = function (msg, url, line) {
  var errMsg;
  console.log(errMsg = msg + "  => " + url + " => " + line);
  alert(errMsg);
  return true; // 返回true表示已经处理，不需要再继续向上传递了。
}

console.log(kk.jj); // kk未定义会抛出异常，异常的处理会被交给onerror回调处理。
</script>
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

| p448
> 然后当使用XMLHttpRequest进行Ajax编程时 => 然而当使用XMLHttpRequest进行Ajax编程时

| p490
> 例如，setRequestHeader()方法的调用必须在调用open()之前但在调用send()之后，否则它将抛出异常。
=>
  例如，setRequestHeader()方法的调用必须在调用open()之后但在调用send()之前，否则它将抛出异常。
  （可以通过下面的代码例子确定这个错误）

| p492
> 如果这样，它会检查响应状态码来取保请求成功。
  => 如果这样，它会检查响应状态码来确保请求成功。

| p493
> 但是还是其他方式来处理服务器的响应。
 => 但是还有其他方式来处理服务器的响应。


| p498
> "ext/plain; charset=UTF-8"
 => "text/plain; charset=UTF-8"

| p548
> `top()`方法接受两个可选的布尔值。
  => `stop()`方法接受两个可选的布尔值。

| p549
> 可以使用queque()方法；
 => 可以使用queue()方法;
