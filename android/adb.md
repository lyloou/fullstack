## 判断当前界面的所属 activity
```sh
adb shell dumpsys activity # 加上 -h 可以获取帮助信息
adb shell dumpsys activity top
adb shell dumpsys activity top | findstr ACTIVITY
```
- [移动测试基础 android 中 dumpsys 命令使用](https://testerhome.com/topics/1462)

## 停止微信
```sh
adb shell am force-stop com.tencent.mm
```
## 启动微信
```sh
adb shell am start -n com.tencent.mm/.ui.LauncherUI
```