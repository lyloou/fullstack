

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
