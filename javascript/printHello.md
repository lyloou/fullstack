
1.
```js
console.log("Hello World!");
```

1.
```js
jQuery.globalEval(console.log("Hello World!"));
```


小技巧：
可以通过传入`console.log`来作为`jQuery.map()`等类似的工具函数的第二个参数，
来打印`jQuery.map()`的参数。（如果不清楚参数有什么的话）
