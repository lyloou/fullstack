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

