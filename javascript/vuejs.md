官方文档：https://cn.vuejs.org

视频教程：[vue.js入门教程](http://www.imooc.com/learn/694)

## 注意事项
- webpack结合vue的时候，less中的元素重名会导致渲染错误。可以通过添加前缀来区别：
```less
.a{
    .container {
        background:#fff;
    }

    .nihao {
        color:#2f0;
        background:#c38;
    }
}

.b{
    .container {
        background:#000;
    }

    .nihao {
        color: #333;
    }
}
```