go mod 使用在go版本1.11+
无需配置GOPATH， 安装好即用


### go mod 简单使用

1，初始化项目
```go
go mod init [moduleName]

#初始化项目模块为Test,其实就是在根目录生成go.mod文件，这个文件用来记载将要需要安装的依赖
eg: go mod init Test
```
2，下载依赖包

```go
go get [Pageckage]

# 下载thrift (会下载到默认的GOPATH路径里，不是某一个指定的项目里)，会生成go.sum文件记录包的版本信息
go get thrift
```

3，拉取依赖包

```go

# 根据项目目录里的go.mod 和 go.sum去拉
go mod download
```

### go mod 其他

1, go依赖包加载方式
从go安装路径的系统包里查询，从GOPATH里安装的第三发依赖包里查询,都没有就报错

2, 常用命令

```go
go mod命令

download    download modules to local cache (下载依赖的module到本地cache))
edit        edit go.mod from tools or scripts (编辑go.mod文件)
graph       print module requirement graph (打印模块依赖图))
init        initialize new module in current directory (再当前文件夹下初始化一个新的module, 创建go.mod文件))
tidy        add missing and remove unused modules (增加丢失的module，去掉未用的module)
vendor      make vendored copy of dependencies (将依赖复制到vendor下)
verify      verify dependencies have expected content (校验依赖)
why         explain why packages or modules are needed (解释为什么需要依赖)

初始化mod

go mod init [module]可以创建一个go.mod，只有一行信息module。

# 查看对go modules 的介绍（官方长文）
go help modules 
```

3，编辑器

推荐goland， 但要下载到2018版本以后的，才支持go module ，否则无法import安装的包，编辑器会报红线错误,这块具体可以搜索引擎解决




