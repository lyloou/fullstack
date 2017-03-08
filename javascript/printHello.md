
1.
```js
console.log("Hello World!");
```

1.
```js
jQuery.globalEval(console.log("Hello World!"));
```


小技巧：
在使用回调函数的地方传入`console.log`，可以知道这个回调函数需要哪些参数。
也就是说，
可以通过传入`console.log`来作为一个回调函数（例如：`jQuery.map()`）的参数，
来打印参数信息。

小技巧：
可以粘贴多行语句到调试台，然后回车查看结果。
