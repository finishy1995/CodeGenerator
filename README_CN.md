# CodeGenerator 代码生成器

[English](README.md) | 简体中文

使用代码生成器自动生成代码，使用代码构建虚拟世界。

## 什么是代码生成器？

代码生成器借助模板文件生成代码。你可以把代码生成器看成一个轻量级的编译器，根据简单的模板语法，制定文件内容的生成规则。

我们提供了一个基础的入口，你可以通过编译这个项目得到可执行文件来直接使用。更推荐的做法是，将这个项目作为你编写自己项目生成工具的依赖库。

## 为什么要用代码生成器？

不少成熟的项目都附带代码生成工具。

* [django](https://github.com/django/django) (用来生成整个项目)
* [create-react-app](https://www.github.com/facebook/create-react-app) (用来生成 react 项目)
* [go-zero](https://github.com/zeromicro/go-zero) (用来生成每个 rpc 服务)

我们从两点说明这个问题——为什么需要自动生成代码，以及这个库与其他实现的区别。

### 1. 代码生成工具的好处

* 减少重复代码书写
* 为开发提供一套代码模板（微服务/网页）
* 维护简便（只需要修改模板文件）

### 2. 代码生成器与其他代码生成工具的区别

大部分的代码生成工具采用 “模板 + 代码” 的方式，生成特定的代码。代码生成器完全通过模板语言生成任意代码，不仅仅满足于原样生成和值的替换，而是支持循环/条件分支等复杂逻辑。

* 不针对某一特定场景或功能，能快速生成任意项目代码，只需要遵守同一套规则
* 极易增减功能，几乎所有逻辑都能扩展，完全实现自己的逻辑
* 无需看代码生成的逻辑，只通过阅读模板文件就能知晓全部文件生成规则

举例，通过在模板文件内使用 Loop 关键字定义循环打印，而不是在生成工具里写死循环逻辑
``` text
#{Loop 3}
hello world!
#{EndLoop}
```

## 示例

* [demo](example/demo) (演示基本语法)
* [microservice](example/service) (自动生成微服务，对外提供 grpc 接口)
* [k8s](example/k8s) (自动生成 kubernetes deployment 文件)
* [ue4](example/ue4) (自动生成 UE4 服务器插件)

## 扩展

* [proto](extension/dataloader/proto) (读取 .proto 文件并作为数据源)
* [datetime](extension/logic/datetime) (自定义关键字 DateTime，打印当前时间)

## 步骤

1. 通过数据加载器 dataloader 加载键值对数据
2. 创建生成任务，依次读取模板文件，按行逐步生成
3. 输出到目标文件

## 安装

执行下面的命令：
``` shell
go get -u github.com/finishy1995/codegenerator
```