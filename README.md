#### 这个工具是做什么的
在查找了一圈github没有找到哪怕一个可以指明包含url的文件进行CDN预热的工具以后，我决定自己写一个。

#### 我该怎么安装
```
#~ sh ./install.sh
```

#### 我该怎么用
本工具提供了 `aliyun` 和 `tencent cloud` 两个平台的预热可能。详细使用方式如下：
```
Usage of cdnPushCache:
  -o    Whether to output the result to screen.
  -p string
        The CSP platform to select. (default "aliyun") | support "aliyun" & "tencent"
  -u string
        The url file to purge. Please use filename or location of file
  -h    CdnPushCache Help manual.
```
