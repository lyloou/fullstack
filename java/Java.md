

| `当类似的代码多次出现的时候，就可以考虑将其封装起来。`

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
private static final HashMap<String, Object> objectsCache = new Hasn<>();

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


## 环境变量配置
```cmd
CLASSPATH =.;%JAVA_HOME%\lib\dt.jar;%JAVA_HOME%\lib\tools.jar;
JAVA_HOME =C:\jdk8.0
Path=%JAVA_HOME%\bin;%JAVA_HOME%\jre\bin;
```
