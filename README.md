## 什么是MarsBase
MarsBase是一个基于iris框架开发的一个基础web开发框架并包含常见的一些接口，比如用户管理、站点设置等。
## 使用
一般使用该库作为基础框架，所以使用的时候一般是需要修改库名(对应go中package名)的，修改库名就意味着源码里的*import*也需要相应修改。  
这里我们假设我们新项目package前缀为*github.com/yourname/yourprojectname*，则步骤如下：  
### 首先clone代码库
```
git clone git@github.com:yushuailiu/MarsBase.git yourprojectname
```
### 修改config下的配置

### 修改代码库包前缀
```
go get git@github.com:yushuailiu/gorename.git // 这个脚本用来修改项目包前缀名
cd yourprojectname
gorename github.com/yushuailiu/MarsBase github.com/yourname/yourprojectname // 根据提示填 Y 即可
make generate // 如果没有安装make，可执行 go generate 这里是把config目录下的配置文件生成go代码

go run main.go --env development // 运行server 访问 http://localhost:8082/api/hello 检测是否启动
```