## 参考网址
http://www.w3school.com.cn
http://zh.learnlayout.com/

## CSS参考手册：
http://www.w3school.com.cn/cssref/index.asp

## 一些需要记住的
- 根据 W3C 的规范，元素内容占据的空间是由 width 属性设置的，而内容周围的 padding 和 border 值是另外计算的。
http://www.w3school.com.cn/css/css_boxmodel.asp

- 上下内边距与左右内边距一致；即上下内边距的百分数会相对于父元素宽度设置，而不是相对于高度。
http://www.w3school.com.cn/css/css_padding.asp

- 记住这一点非常重要。事实上，忘记声明边框样式是一个常犯的错误;
由于 border-style 的默认值是 none，如果没有声明样式，就相当于 border-style: none。因此，如果您希望边框出现，就必须声明一个边框样式。


- 绝对定位使元素的位置与文档流无关，因此不占据空间。
http://www.w3school.com.cn/css/css_positioning_absolute.asp

- 常常有元素可以应用 clear，但是有时候不得不为了进行布局而添加无意义的标记。
http://www.w3school.com.cn/css/css_positioning_floating.asp



## 难点：
- float：http://www.w3school.com.cn/cssref/pr_class_float.asp
- position（包含position,top,bottom,left,right,）: http://www.w3school.com.cn/css/css_positioning.asp
- flex: 
    http://www.ruanyifeng.com/blog/2015/07/flex-grammar.html
    http://www.ruanyifeng.com/blog/2015/07/flex-examples.html
- background-size: http://www.w3school.com.cn/css3/css3_background.asp
- box-sizing:http://zh.learnlayout.com/box-sizing.html
    当你设置一个元素为 box-sizing: border-box; 时，此元素的内边距和边框不再会增加它的宽度。
- border-image: http://www.w3school.com.cn/css3/css3_border.asp    

## 易出错点：
### padding位置缩写：
```
http://www.w3school.com.cn/cssref/pr_padding.asp

padding:10px 5px 15px 20px;
上内边距是 10px
右内边距是 5px
下内边距是 15px
左内边距是 20px

padding:10px 5px 15px;
上内边距是 10px
右内边距和左内边距是 5px
下内边距是 15px

padding:10px 5px;
上内边距和下内边距是 10px
右内边距和左内边距是 5px

padding:10px;
所有 4 个内边距都是 10px
```


## [块元素和内联元素](http://www.w3school.com.cn/html/html_blocks.asp)

块元素指的是占据全部可用宽度的元素，并且在其前后都会换行。
```html
<h1>, <p>, <ul>, <table>
```

内联元素，与块元素相反。
```html
<b>, <td>, <a>, <img>
```

## [设置某块区域上下滚动](http://www.w3school.com.cn/cssref/pr_overflow-y.asp)
```css
.wrap{
    height: 440rpx;
    overflow-y: auto;
}
```

## [CSS 实现隐藏滚动条同时又可以滚动](https://blog.niceue.com/front-end-development/hide-scrollbar-but-still-scrollable-using-css.html)
https://blogs.msdn.microsoft.com/kurlak/2013/11/03/hiding-vertical-scrollbars-with-pure-css-in-chrome-ie-6-firefox-opera-and-safari/
```css
div1::-webkit-scrollbar{
    display: none;
}
```