# webview External Dependency PoC for Go/Windows

Verify that build does not work due to missing `WebView2.h` header:

```bat
go build
```

Expected output:

```
# github.com/webview/webview
In file included from webview.cc:1:
webview.h:839:10: fatal error: WebView2.h: No such file or directory
  839 | #include "WebView2.h"
      |          ^~~~~~~~~~~~
compilation terminated.
```

Use generator to fetch dependencies and set up Go environment:

```bat
go generate
```

Optionally restart the terminal to make sure that there are no lingering variables that help make this work.

Try to build again:

```bat
go build
basic.exe
```
