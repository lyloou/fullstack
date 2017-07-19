## 获取服务器时间
原理：通过获取链接的Header信息来获取时间。
注意：网络请求需要在多线程中执行
```java
URL url = new URL("http://lyloou.com");
URLConnection uc = url.openConnection();// 获取连接对象
uc.connect();// 连接
long webTimeMillis = uc.getDate();
SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss", Locale.CHINA);
String formatWebTime = simpleDateFormat.format(new Date(webTimeMillis));
System.out.println(formatWebTime);
```

## 添加顺序的注释
> 参考资料:深入探索Android热修复技术原理7.3Q.pdf p107

// %% Part 1. 创建了新对象；  
// %% Part 2. 找到旧对象的引用；  
// %% Part 3. 用新对象赋值给旧对象的引用；  


## 对象转换成字符串
在不确定对象是否为空时，通过`String.valueOf(object)`的方法，  
而不是直接调用：`object.toString();`方法

## 封装
`当类似的代码多次出现的时候，就可以考虑将其封装起来。`

## 日期、时间格式的转换
>
- [GankBeautyResultToItemsMapper.java](https://github.com/lyloou/RxJavaSamples/blob/master/app/src/main/java/com/rengwuxian/rxjavasamples/util/GankBeautyResultToItemsMapper.java)

```java
String strDate = "2017-05-09T14:28:32.974Z";
SimpleDateFormat inputFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SS'Z'");
SimpleDateFormat outputFormat = new SimpleDateFormat("yy/MM/dd HH:mm:ss");
try {
    Date inputDate = inputFormat.parse(strDate);
    String outputDate = outputFormat.format(inputDate);
    System.out.println(outputDate);
} catch (ParseException e) {
    e.printStackTrace();
}
```

## 打印出好看的`list`
```java
System.out.println(Arrays.toString(list.toArray()));
```

## HashMap 用来缓存对象
```java
private static final Map<String, Object> objectsCache = new HashMap<>();

// put in
public void putIn(String key, Object obj) {
  objectsCache.put(key, obj);
}

// get out
public Object getOut(String key, Object default) {
  if(objectsCache.containsKey(key)) {
    Object obj = objectsCache.get(key);
    if(obj == null) {
      objectsCache.put(key, default);
    }
    return obj;
  } else {
    objectsCache.put(key, default);
    return default;
  }
}
```

# 资源

## 环境变量配置
```cmd
CLASSPATH =.;%JAVA_HOME%\lib\dt.jar;%JAVA_HOME%\lib\tools.jar;
JAVA_HOME =C:\jdk8.0
Path=%JAVA_HOME%\bin;%JAVA_HOME%\jre\bin;
```

## java tutorial 搜索
https://docs.oracle.com/javase/tutorial/search.html
