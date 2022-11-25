# CodeGenerator

English | [简体中文](README_CN.md)

Builders use code generator to auto coding, coding to build a virtual world.

## What is CodeGenerator?

CodeGenerator read template files (```*.tpl``` by default) and generate code. It's just like a lightweight template compiler, which following the generation rules to generate several files.

We provide "package main" as a basic entry. You can build this (```go build```) and execute the program.

But, it is more recommended to using this project as a library. You can realize any logic key or data loader fast and easy.

## Why using CodeGenerator？

Several cases are using code generation tools to generate codes.

* [django](https://github.com/django/django) (web backend project)
* [create-react-app](https://www.github.com/facebook/create-react-app) (create react project)
* [go-zero](https://github.com/zeromicro/go-zero) (generate microservice)

We explain this problem from the following two points - why we need to automatically generate code, and the difference between this library and other implementations.

### 1. The benefits of code generation tools

* Reduce repetitive code writing
* Provide a set of code templates for development (microservices/web/app)
* Easy to maintain (sometimes only need to modify the template file)

### 2. Difference between CodeGenerator and other code generation tools

Most code generation tools use the solution ***"template + hard code"*** to generate specific code. CodeGenerator generates code through a specific template language. It is not only satisfied with text generation and value replacement, but also supports complex logic such as ***loops / conditions***.

* Quickly generate any project code, only need to follow the same set of rules
* It is very easy to increase or decrease logic(reserved word), almost all logic can be expanded, and fully realize your own logic
* You don't need to look at the logic, you can know all the file generation rules just by reading the template file

For example, by using the ```Loop``` keyword to define loop printing in the template file, instead of writing infinite loop logic hard code in the generation tool
``` text
#{Loop 3}
hello world!
#{EndLoop}
```

## Example

* [demo](example/demo) (basic rules)
* [microservice](example/service) (generate microservice, provide grpc service)
* [k8s](example/k8s) (generate kubernetes deployment files)
* [ue4](example/ue4) (generate ue4 server plugin)

## Extension

* [proto](extension/dataloader/proto) (read .proto files as data)
* [datetime](extension/logic/datetime) (customize ```DateTime``` reserved word, which will print the current time)

## Step

1. Using several DataLoader to load a key-value map (dictionary in code)
2. Create a mission, it will read all template files, and every template will create a task to generate content.
3. Print content to several files

## Install

Run the following command：
``` shell
go get -u github.com/finishy1995/codegenerator
```