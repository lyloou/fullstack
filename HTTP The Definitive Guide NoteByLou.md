
| Agents
User agents (or just agents) are client programs that make HTTP requests on the
user's behalf. (用来发出HTTP请求的客户端程序，例如：web浏览器) --p19

| Tunnels
HTTP tunnels are often used to transport non-HTTP data over one or more HTTP
connections, without looking at the data. --p19

| Gateways

| Caches

| Proxies
A proxy sits between a client and a server.
Proxies are often used for security, acting as trusted intermediaries
through which all web traffic flows. --p18

## Architectrual Components of the Web

| Protocol Versions --p16
HTTP/0.9 supports only the GET method.
1.0 was the first version of HTTP that was widely deployed.

| An HTTP transaction using telnet:
```sh
[root@localhost ~]# telnet www.lyloou.com 80
Trying 151.101.192.133...
Connected to www.lyloou.com.
Escape character is '^]'.
GET /index.html HTTP/1.1  # input this
Host: www.lyloou.com      # input this
                          # input this

HTTP/1.1 301 Moved Permanently
Server: GitHub.com
Content-Type: text/html
Location: http://lyloou.com/index.html
X-GitHub-Request-Id: 26A4:06EC:8E05E6C:B479C11:58CB8810
Content-Length: 178
Accept-Ranges: bytes
Date: Fri, 17 Mar 2017 06:54:08 GMT
Via: 1.1 varnish
Age: 0
Connection: keep-alive
X-Served-By: cache-lax8651-LAX
X-Cache: MISS
X-Cache-Hits: 0
X-Timer: S1489733648.903540,VS0,VE68
Vary: Accept-Encoding
X-Fastly-Request-ID: b8d3f5bbb2343450fd125b92f1dde483c5109377

<html>
<head><title>301 Moved Permanently</title></head>
<body bgcolor="white">
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx</center>
</body>
</html>
```

| Telnet is commonly used for remote terminal sessions, but it can
generally connect to any TCP server, including HTTP servers.

| Establish a TCP/IP connection between the client and server:
- Internet protocol (IP) addresses
- Port numbers

| TCP/IP
Once a TCP connections is established, messages exchanged between the client and
server computers will never be lost, damaged, or received out of order. --p12

| HTTP messages consist of three parts: --p11
1. Start line; (indicating what to do for a request and what happened for a
  response)
2. Header fields; (Note: The headers end with a blank line)
3. Body; (can contain arbitray binary data. eg., images, videos, audio tracks,
  software applications.)

| Messages
- request messages: sent from web clients to web servers.
- response messages: messages from servers to clients.
they are very similar.



| Status Codes
Every HTTP response message comes back with a status code. --p9

| Methods
The method tells the server what action to perform (e.g. fetch a web page,
  delete a file, etc.) --p8
- GET
- POST
- DELETE
- PUT
- HEAD

| Transactions
An HTTP transaction consists of: --p8
- a request command (sent from client to server).
- a response result (sent from the server back to the client)

| URNs
A URNs serves as a unique name for a particular piect of content,
independent of where the resource currently resides.
URNs are still experimental and not yet widely adopted.

| URLs
URLs describe the specific location of a resource on a particular server.
contains three main parts: --p7
- scheme: (e.g., http://)
- server Internet address (e.g., www.lyloou.com)
- names a resource on the web server (e.g., hello.gif)

| URIs
URIs come in two flavors: URLs and URNs.

| Media Types
MIME was originally designed to solve problems encountered in moving messages
between different electronic mail systems.
MIME worked so well for email that HTTP adopted it to describe and label its
own multimedia content. --p5
eg.
- text/HTML
- text/plain
- image/jpeg
- image/gif
- video/quicktime
- application/vnd.ms-powerpoint
https://zh.wikipedia.org/wiki/多用途互聯網郵件擴展

| Resources
- static file: text files, HTML files, Word files, JPEG files, AVI...
- dynamic content: generate content based on your identity, on what
  information you've requested, or on the time of day.(eg. show live image
    from a camera, trade stocks, search databases, buy gifts from online
    stroes)

| Because HTTP uses reliable data-transmission protocols, it guarantees that
your data will not be damaged or scrambled in transit, even when it comes from
the other side of the globe. --p3

| Throughout the book, we are careful to explain the "why" of HTTP, not just the
"how".


| There are many books that  explain how to use Web,
but this is the book that explains how the Web works.

| The Definitive Guide is in understanding how the Web works
and how to apply that knowledge to web programming and administration.
