// document.title="Lou223";

onmessage = function (e) {
  console.log(e.data);
  postMessage("已经成功打印！");
}
