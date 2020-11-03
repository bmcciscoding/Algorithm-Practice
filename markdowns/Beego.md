# Beego 框架简单上手

## 安装 bee
bee 是 Beego 都脚手架工作，github 上给了自动安装都教程，但是很不幸，失败了。
于是我选择用手动都的方式来安装
1. 从 master 分支下载 zip
2. 解压之后 go build main.go，得到一个可执行文件
3. 拷贝到 /Users/xxxx/go/bin 下
4. 然后在 .zrshc 文件里，导出路径
   ```go
   export GOPATH=/Users/xxxx/go
   export GOBIN=$GOPATH/bin
   export PATH=$PATH:$GOBIN
   ```
5. 重启终端

成功！
![](https://tva1.sinaimg.cn/large/0081Kckwgy1gkc6tzj30qj3070072dg5.jpg)



