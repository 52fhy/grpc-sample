# grpc-sample

## 入坑预警

- 如果使用PHP、Python开发gRPC的客户端，需要编译gRPC命令行工具，生成proto的代码生成插件，否则proto里定义的service无法编译出来。编译需要使用GCC4.8级以上版本，否则报不支持C++11。然后需要龟速下周grpc源码，并下载一大堆第三方依赖。这个过程非常痛苦。使用golang、java的可以忽略。
- PHP还需要按照grpc的c扩展。编译需要使用GCC4.8级以上版本。
- 如果使用golang开发服务，依赖的第三方服务基本是下载不下来的，需要使用`go mod`增加映射规则到github仓库，github下载也是龟速。

## 运行

运行服务端

``` go 
go mod tidy 
go run main.go
```

运行单元测试：
``` go
go test -v client_test.go
```

运行php客户端：  
需要先安装扩展：
``` bash
pecl install protobuf
pecl install grpc

# 记得修改php.ini文件


cd tests
composer require grpc/grpc
cd ../
```
然后执行代码
``` php 
php tests/client_test.php
```

## 常见问题

**1、CentOS6使用 go mod获取第三方依赖包unknown revision xxx错误**  
解决：其实go mod调用链中会用到一些git指令，当git版本比较旧时，调用失败产生错误，并给出歧义的提示信息。方法就是升级git版本，CentOS6自带的git是1.7版本。升级完毕后，再尝试go mod。

快速升级方法：  
centos6:  
```
# 安装yum源
wget http://opensource.wandisco.com/centos/6/git/x86_64/wandisco-git-release-6-1.noarch.rpm && rpm -ivh wandisco-git-release-6-1.noarch.rpm

## 安装git 2.x
yum install git -y

## 验证
git --version
git version 2.14.1
```

**2、PHP报错：Fatal error: Class 'Google\Protobuf\Internal\Message' not found**   
解决：请安装PHP的protobuf c扩展。  

**3、PHP报错：Fatal error: Class '\Grpc\BaseStub' not found**   
解决：使用`composer require grpc/grpc`安装grpc。另外对应的grpc C扩展也要安装。 


