| tab悬浮效果实现：
搜索关键字：`Android tab悬浮`
开源库（如果布局嵌套很复杂，则不一定适合）：https://github.com/carlonzo/StikkyHeader
另一种解决思路：
- 在布局中，使用一个同tab1完全一样的tab2（并设置显示属性为gone）；
- 获取 tab1 到外层父视图顶部的距离（通过`tab1.getTop()`获取）；
- 当滚动的距离超过这个距离时，让外层的tab2显示出来，否则隐藏tab2；（即：`(scrollY - tab1.getTop()) > 0`）
- tab1 和 tab2 在逻辑上做相同的变化；

```xml
<FrameLayout>
    <ScrollTabView>
      <LinearLayout>
        <...>
        <tab1>
        <...>
      </LinearLayout>
    </ScrollTabView>

    <tab2 visibility="gone">
</FrameLayout>
```

```java
// 重写ScrollView
// 参考资料：http://www.jianshu.com/p/8ee4b0896b22
// https://github.com/aohanyao/Advanced/blob/master/code/CustomView/ScollTabView/app/src/main/java/aohanyao/com/scolltabview/ScrollLevitateTabView.java
public class ScrollTabView extends ScrollView {
    OnScrollListener mOnScrollListener;

    public ScrollTabView(Context context) {
        this(context, null);
    }

    public ScrollTabView(Context context, AttributeSet attrs) {
        this(context, attrs, 0);
    }

    public ScrollTabView(Context context, AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
        setUp();
    }

    public void setOnScrollListener(OnScrollListener onScrollListener) {
        mOnScrollListener = onScrollListener;
    }

    private void setUp() {
        if (mOnScrollListener == null) {
            return;
        }

        getViewTreeObserver().addOnGlobalLayoutListener(new ViewTreeObserver.OnGlobalLayoutListener() {
            @Override
            public void onGlobalLayout() {
                mOnScrollListener.onScroll(getScrollY());
            }
        });
    }

    @Override
    protected void onScrollChanged(int l, int t, int oldl, int oldt) {
        if (mOnScrollListener != null) {
            mOnScrollListener.onScroll(t);
        }
        super.onScrollChanged(l, t, oldl, oldt);
    }

    public interface OnScrollListener {
        void onScroll(int scrollY);
    }
}

// 运用
ScrollTabView scrollTabView = (ScrollTabView) mView.findViewById(R.id.scroll_tab_view);
scrollTabView.setOnScrollListener(new ScrollTabView.OnScrollListener() {
    @Override
    public void onScroll(int scrollY) {
        // ScrollView的高度变化决定了tab2的显示与否
        int deltaY = scrollY - tab1.getTop();
        tab2.setVisibility(deltaY > 0 ? View.VISIBLE : View.GONE);
    }
});
```

| GridLayoutManager 均分两列的装饰器ItemDecoration
```java
@Override
public void getItemOffsets(Rect outRect, View view, RecyclerView parent, RecyclerView.State state) {
    int position = parent.getChildLayoutPosition(view);
    outRect.top = space;
    outRect.left = space;
    if (position % 2 == 1) {
        outRect.right = space;
    }
}
```

| 规范
*资源*
- `color.xml` 由基本色构成，而不是针对每个模块或组件分别设置。
- `dimens.xml` 同上。
- `string.xml` 由于语言的多样性，属性名具体一点会更好。另外，value 值不要全部使用大写（用`textAllCaps`）。
- `layout` 中的层级不要太多。
> https://github.com/futurice/android-best-practices#resources

| RecyclerView
- 针对多种类型的情况，可以创建多个 `ViewHolder`和设置多个 `type`
- 分页加载（加载更多）
> http://www.gadgetsaint.com/android/recyclerview-header-footer-pagination/#.WRwxJGh96Hs

*方案一：*
```java
mRecyclerView.addOnScrollListener(new RecyclerView.OnScrollListener() {
    @Override
    public void onScrollStateChanged(RecyclerView recyclerView, int newState) {
        super.onScrollStateChanged(recyclerView, newState);

        int lastvisibleitemposition = mLayoutManager.findLastVisibleItemPosition();

        if (lastvisibleitemposition == mAdapter.getItemCount() - 1) {

            if (!loading && !isLastPage) {

                loading = true;
                fetchData((++pageCount));
                // Increment the pagecount everytime we scroll to fetch data from the next page
                // make loading = false once the data is loaded
                // call mAdapter.notifyDataSetChanged() to refresh the Adapter and Layout

            }


        }
    }
});
```

*方案二：*
```java
// http://www.jianshu.com/p/4feb0c16d1b5
private RecyclerView.OnScrollListener mListener = new RecyclerView.OnScrollListener() {
    @Override
    public void onScrollStateChanged(RecyclerView recyclerView, int newState) {
        super.onScrollStateChanged(recyclerView, newState);
    }

    @Override
    public void onScrolled(RecyclerView recyclerView, int dx, int dy) {
        super.onScrolled(recyclerView, dx, dy);

        LinearLayoutManager layoutManager = (LinearLayoutManager) recyclerView.getLayoutManager();
        int lastVisibleItem = layoutManager.findLastVisibleItemPosition();
        int totalItemCount = layoutManager.getItemCount();

        if (totalItemCount < 250 && lastVisibleItem >= totalItemCount - 4) {
            // 注意：要限制请求，否则请求太多次数，导致服务器崩溃或者服务器拒绝请求（罪过，罪过）。
            if (mIsLoading) {
                Log.i(TAG, "onScrolled: " + "加载中---------");
            } else {
                Log.i(TAG, "onScrolled: " + "加载更多了=======》");
                loadSubject();
            }

        }

        Log.d(TAG, "onScrolled: lastVisibleItem=" + lastVisibleItem);
        Log.d(TAG, "onScrolled: totalItemCount=" + totalItemCount);

    }
};
```

| Avoid Maven dynamic dependency resolution. (such as `2.1.+`)
  this result in different and unstable builds or subtle, untracked difference
  in behavior between builds.

| Stetho - A debug bridge for Android applications
> Stetho is a debug bridge for Android applications from Facebook that integrates with the Chrome desktop browser's Developer Tools. With Stetho you can easily inspect your application, most notably the network traffic. It also allows you to easily inspect and edit SQLite databases and the shared preferences in your app. You should, however, make sure that Stetho is only enabled in the debug build and not in the release build variant.
> https://github.com/futurice/android-best-practices#use-stetho

| 针对测试版本和发布版本使用不同的 appId, 这样两个版本就可以同时存在在一个设备上了。
可以通过前缀或后缀的方式来区分。
> https://github.com/futurice/android-best-practices#gradle-configuration

| 在配置 `build.gradle` 的时候，避免直接写入敏感信息（例如：密码），而是写入到版本控制工具
忽略的文件`gradle.properties`中。
> https://github.com/futurice/android-best-practices#gradle-configuration


## [Android Best Practices](https://github.com/futurice/android-best-practices)
after...
- http://blog.futurice.com/tech-pick-of-the-week-rx-for-net-and-rxjava-for-android
- https://github.com/square/mortar
- Use Stetho

| TextView 设置空格
```xml
&#160;&#160;&#160;&#8201;&#160;&#160;&#8201;
```

| TextView html 渲染（注意：html渲染的方式无法改变字体的大小，可以调整颜色、粗细、斜体等属性）
```java
public static void renderWithHtml(final TextView tv, String data) {
    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.N) {
        tv.setText(Html.fromHtml(data, Html.FROM_HTML_MODE_COMPACT));
    } else {
        tv.setText(Html.fromHtml(data));
    }
}
```

| TextView 渲染字体
> [Android TextView个别字体格式设置小结 - 简书](http://www.jianshu.com/p/2671e78089f9)

```java
SpannableString msp = new SpannableString("字体测试字体大小一半两倍前景色背景色正常粗体斜体粗斜体下划线删除线x1x2电话邮件网站短信彩信地图X轴综合");

//设置字体(default,default-bold,monospace,serif,sans-serif)
msp.setSpan(new TypefaceSpan("monospace"), 0, 2, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);
msp.setSpan(new TypefaceSpan("serif"), 2, 4, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);

//设置字体大小（绝对值,单位：像素）
msp.setSpan(new AbsoluteSizeSpan(20), 4, 6, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);
msp.setSpan(new AbsoluteSizeSpan(20, true), 6, 8, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //第二个参数boolean dip，如果为true，表示前面的字体大小单位为dip，否则为像素，同上。

//设置字体大小（相对值,单位：像素） 参数表示为默认字体大小的多少倍
msp.setSpan(new RelativeSizeSpan(0.5f), 8, 10, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //0.5f表示默认字体大小的一半
msp.setSpan(new RelativeSizeSpan(2.0f), 10, 12, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //2.0f表示默认字体大小的两倍

//设置字体前景色
msp.setSpan(new ForegroundColorSpan(Color.MAGENTA), 12, 15, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //设置前景色为洋红色

//设置字体背景色
msp.setSpan(new BackgroundColorSpan(Color.CYAN), 15, 18, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //设置背景色为青色

//设置字体样式正常，粗体，斜体，粗斜体
msp.setSpan(new StyleSpan(android.graphics.Typeface.NORMAL), 18, 20, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //正常
msp.setSpan(new StyleSpan(android.graphics.Typeface.BOLD), 20, 22, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //粗体
msp.setSpan(new StyleSpan(android.graphics.Typeface.ITALIC), 22, 24, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //斜体
msp.setSpan(new StyleSpan(android.graphics.Typeface.BOLD_ITALIC), 24, 27, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);  //粗斜体

//设置下划线
msp.setSpan(new UnderlineSpan(), 27, 30, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);

//设置删除线
msp.setSpan(new StrikethroughSpan(), 30, 33, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);

//设置上下标
msp.setSpan(new SubscriptSpan(), 34, 35, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);     //下标
msp.setSpan(new SuperscriptSpan(), 36, 37, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);   //上标

//超级链接（需要添加setMovementMethod方法附加响应）
msp.setSpan(new URLSpan("tel:4155551212"), 37, 39, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);     //电话
msp.setSpan(new URLSpan("mailto:webmaster@google.com"), 39, 41, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);     //邮件
msp.setSpan(new URLSpan("http://www.baidu.com"), 41, 43, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);     //网络
msp.setSpan(new URLSpan("sms:4155551212"), 43, 45, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);     //短信   使用sms:或者smsto:
msp.setSpan(new URLSpan("mms:4155551212"), 45, 47, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);     //彩信   使用mms:或者mmsto:
msp.setSpan(new URLSpan("geo:38.899533,-77.036476"), 47, 49, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE);     //地图

//设置字体大小（相对值,单位：像素） 参数表示为默认字体宽度的多少倍
msp.setSpan(new ScaleXSpan(2.0f), 49, 51, Spanned.SPAN_EXCLUSIVE_EXCLUSIVE); //2.0f表示默认字体宽度的两倍，即X轴方向放大为默认字体的两倍，而高度不变
//SpannableString对象设置给TextView
tokenTv.setText(msp);
//设置TextView可点击
tokenTv.setMovementMethod(LinkMovementMethod.getInstance());
```

| 引入开源库的时候，注意要及时引入 proguard （以免忘记）

| 通过调用 fragment 的 `isAdded()` 方法，来判断当前的 fragment 是否阵亡。
- 具体用法参考：`com.example.android.architecture.blueprints.todoapp.tasks.TasksContract.isActive();`
- `isAdded()`文档：Return true if the fragment is currently added to its activity.


| 通过图像 URL 设置圆角图像：
```java
private void setCircleIcon(ImageView view) {
    String url = "https://img.gcall.com/dca5/M00/10/8E/wKhoNlggetaENWylAAAAAAAAAAA457.jpg";
    final int w = Uscreen.dp2Px(mContext, 48);
    Glide.with(mContext)
            .load(url)
            .centerCrop()
            .fitCenter()
            .thumbnail(0.1f)
            .placeholder(R.mipmap.ic_launcher)
            .crossFade()
            .override(w, w)
            .transform(new BitmapTransformation(mContext) {
                @Override
                protected Bitmap transform(BitmapPool pool, Bitmap toTransform, int outWidth, int outHeight) {
                    return Uview.getBitmapByXfermode(mContext, toTransform,
                            Color.parseColor("#dddddd"),
                            w,
                            w, PorterDuff.Mode.SRC_IN);
                }

                @Override
                public String getId() {
                    return getClass().getName();
                }
            })
            .into(view);
}
```

| apk 安装路径： --p220(q1)
- `/system/app` : 系统及别的 apk;
- `/data/app` : 用户安装的 apk;

| pm(package manager)： 主宰着应用的包管理；
  am(activity manager): 主宰着应用的活动管理； --p210(q1)

| 如果说系统信息好比是国家的 GDP， 那么 Apk 应用信息则像是对个人的经济普查。 --p210(q1)

| 一个 Task 中的 Activity 可以来自不同的 APP，
同一个 App 的  Activity 也可能不再一个 Task 中。 --p200(q1)

| 一个合理的任务调度栈不仅是性能的保证，更是提供性能的基础。 --p200(q1)

| 如果你得自定义 View 需要频繁刷新，或者刷新时数据处理量比较大，
那么就可以考虑使用 SurfaceView 来取代 View 了。 --p155(q1)

| canvas 的方法： --p118(q1)
- save(): 保存画布；作用是将之前的所有已绘制图像保存起来，让后续的操作就好像在新图层上操作一样;
- restore(): 合并图层；作用是将 save 之后的和 save 之前的图像进行合并；
- translate(): 坐标系的平移；
- rotate(): 坐标系的翻转；


| scrollTo、scrollBy 方法移动的是 View 的 content, 即让 View 的内容移动( content )，
如果在 ViewGroup 中使用 scrollTo、 scrollBy 方法，那么移动的将是所有子 View；
例如： TextView, content 就是它的文本; ImageView, content 就是它的 drawable. --p95(q1)

| 形象的事件处理机制： --p62(q1)
- 事件的传递顺序：
    总经理(MyViewGroupA) --> 部长(MyViewGroupB) --> 你(View)。
    事件传递的时候，先执行 dispatchTouchEvent() 方法，再执行 onInterceptTouchEvent()
    方法（注意：View 没有 onInterceptTouchEvent方法，只有 ViewGroup 有这个方法）。
- 事件的处理顺序：
    你(View) --> 部长(MyViewGroupB) --> 总经理(MyViewGroupB)。
    事件处理都是执行 onTouchEvent() 方法。
- 事件传递的返回值： true, 拦截；false, 不拦截, 继续流程；    
- 事件处理的返回值： true, 处理了，不用审核； false, 给上级处理。
- 初始情况下，返回值都是 false。

| 动态控制UI模板的方法：接口回调、提供 public 方法。 --p48(q1)
- 接口回调
```java
interface Ilistener {
  void click(View v);  
}

private void innerMethod(View v) {
  mIlistener.click(v);
}

// 声明和初始化
Ilistener mIlistener;
public void register(Ilistener listener){
  mIlistener = listener;
}
```
- 提供 public 方法
```java
private String mStr;
public void publicMethod(String str) {
  mStr = str;
}
```

| 解释 bitmap 和 Canvas  --p38(q1)
```java
// 装载画布：将 bitmap 装载到画布 canvas 上,
// 这个 bitmap 存储了所有绘制在 Canvas.drawXXX 方法的像素信息
Canvas canvas = new Canvas(bitmap);
```
代码解释：
```java
Canvas canvas;
Bitmap bitmap1, bitmap2;
private void initBitmap() {
  bitmap1 = new Bitmap(...);
  bitmap2 = new Bitmap(...);
}

private void initCanvas() {
  this.canvas = new Canvas(bitmap2);
}

@Override
protected void onDraw(Canvas canvas) {
  canvas.drawBitmap(bitmap1, 0, 0, null);
  canvas.drawBitmap(bitmap2, 0, 0, null);
}

private void changeBitmap2(){

}
```


| Canvas 就像一个画板，使用 Paint 就可以在上面作画了。 --p37(q1)

| 其实，Android 就好像那个蒙着眼睛画画的人，你必须精确地告诉它如何去画，
它才能绘出你想要的图形；--p34(q1)|

| Intent - Android上的信使，信息传递的载体；
--p5(q1)



## 说明
- (q1) --《Android群英传》






### tips
| 在 Android Studio 中的 app module 中运行 Java 测试代码，会发现很慢；
可以尝试创建 java library 模块来运行 java 测试项目；

| 使用 Refector （style）的姿势：
将鼠标定位到要 refector 的标签中（不要选中任何代码），
然后右键 `Refector --> Extract --> Style...`
