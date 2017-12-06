## [设定参数必须传递](http://es6.ruanyifeng.com/#docs/function#应用)
```js
function throwIfMissing() {
  throw new Error('Missing parameter');
}

function foo(mustBeProvided = throwIfMissing()) {
  return mustBeProvided;
}

foo()
// Error: Missing parameter
```


## [为浮点数运算部署一个误差检查函数](http://es6.ruanyifeng.com/#docs/number#Number-EPSILON)
```js
function withinErrorMargin (left, right) {
  return Math.abs(left - right) < Number.EPSILON * Math.pow(2, 2);
}

0.1 + 0.2 === 0.3 // false
withinErrorMargin(0.1 + 0.2, 0.3) // true

1.1 + 1.3 === 2.4 // false
withinErrorMargin(1.1 + 1.3, 2.4) // true
```

## [正确返回字符串长度的函数（Unicode字符）](http://es6.ruanyifeng.com/#docs/regex#u-修饰符)
```js
function codePointLength(text) {
  var result = text.match(/[\s\S]/gu);
  return result ? result.length : 0;
}

var s = '𠮷𠮷';

s.length // 4
codePointLength(s) // 2
```

## apply & call
```js
var a = {x:1}
var b = {x:2}
function f(){ console.log(this.x)}
a.f = f
b.f = f
f.apply(a)  //1
f.apply(b)  //2
```

## 返回对象的类
```js
// p214
function classof(o) {
  return Object.prototype.toString.call(o).slice(8, -1);
}
```

## 输出对象的类：Object.prototype.toString.call
```js
Object.prototype.toString.call(1);
"[object Number]"

Object.prototype.toString.call(function(){});
"[object Function]"
```
typeof 得到的是类型；


## [How to pass url query params](https://github.com/github/fetch/issues/256)
```js

function getUrl(url, options) {
  if(!isEmptyObject(options)) {
      url += (url.indexOf('?') === -1 ? '?' : '&') + queryParams(options);
  }
  return url;
}

function queryParams(params) {
  return Object.keys(params)
      .map(k => encodeURIComponent(k) + '=' + encodeURIComponent(params[k]))
      .join('&');
}

// 判断对象是否为空 https://stackoverflow.com/questions/4994201/is-object-empty
function isEmptyObject (obj) {
  if (obj == null) return true;
  if (obj.length > 0) return false;
  if (obj.length === 0) return true;
  if (typeof obj != "object") return true;

  for (var name in obj) {
    if (obj.hasOwnProperty(name)) {
      return false;
    }
  }
  return true;
}

var url = getUrl("http://lyloou.com", {})
console.log(url)
```

