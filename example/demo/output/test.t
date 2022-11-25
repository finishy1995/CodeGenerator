1. 原样生成基本内容
hello world

2. 定义/覆盖数据

3. 生成数据
1
1
1x
haha 1 = 1

4. 也可以修改文件定义数据

5. 插入一个文件，常作为文件头或文件尾循环利用（支持循环插入检测，不用担心会死循环）
// Copyright XXX. All Rights Reserved
x // 可以引用上级文件的变量

6. 插入文件中定义的变量也会带到此文件中，你可以像引用头文件那样灵活使用 Insert 和 Define
.d 在 header.i 中定义 hhh

7. 可以使用 GetKey 遍历所有变量
b // 输出 .map1.b 的值
ccc // 输出 .map1.cc 的值

8. 可以使用 Loop + EndLoop 关键字定义循环

Hello World
hhhhhhhh
Hello World
hhhhhhhh
Hello World
hhhhhhhh

9. 循环进阶，可以将当前循环次数存入数据中，可以嵌套循环
index1: 1 index2: 
map1 content: a
index1: 1 index2: 1
map2 content: aaa
index1: 1 index2: 2
map2 content: bbb
index1: 1 index2: 3
map2 content: ccc
index1: 2 index2: 3
map1 content: b
index1: 2 index2: 1
map2 content: aaa
index1: 2 index2: 2
map2 content: bbb
index1: 2 index2: 3
map2 content: ccc

10. 可以使用 If + EndIf 关键字定义条件表达式
.a has a value, and .a != "" ("" will equal to undefined in If block)


also support "true" text

Logically, you can set a value "true" in any key


11. 可以使用 Else 在 If 和 EndIf 之间，以达到区块选择的效果
We don't have any money.

12. 灵活组合他们，必要时可以通过提供的接口扩展他们。想要增加一个关键字？想要改变文件输出方式？都能实现
您可以继续查看其他 Demo，了解真实场景下如何使用 CodeGenerator，或者即刻开始使用。