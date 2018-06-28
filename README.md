# SimpleFileServer

文件管理服务器

## 使用

### 参数

- `-port`
    - 端口
    - 默认9410
- `-rootDir`  
    - 根目录路径
    - 为了安全，不填会直接报错

### 简单版本

使用最新的release包执行运行就好

### 复杂版本

- 在拉下来的代码目录下 `go build` 可以得到SimpleFileServer的可执行文件
- 以windows为例，在当前目录下`.\SimpleFileServer.exe -port=9410 -rootDir="F:\FromGithub"`
- 浏览 `http://127.0.0.1:9410`即可
