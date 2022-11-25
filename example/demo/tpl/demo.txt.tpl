1. 原样生成基本内容
hello world

2. 定义/覆盖数据
#{Define .a = 1}
#{Define .b = #{.a}}
#{Define .c = x}

3. 生成数据
#{.a}
#{.b}
#{.a .c}
haha #{.b} = #{.a}

4. 也可以修改文件定义数据
#{Define file.name = test}
#{Define file.suffix = .t}

5. 插入一个文件，常作为文件头或文件尾循环利用（支持循环插入检测，不用担心会死循环）
#{Insert header.i}

6. 插入文件中定义的变量也会带到此文件中，你可以像引用头文件那样灵活使用 Insert 和 Define
.d 在 header.i 中定义 #{.d}

7. 可以使用 GetKey 遍历所有变量
#{Define .map1.a=a}
#{Define .map1.b=b}
#{Define .map2.aa=aaa}
#{Define .map2.bb=bbb}
#{Define .map2.cc=ccc}
#{GetKey .map1 2} // 输出 .map1.b 的值
#{GetKey .map2 3} // 输出 .map1.cc 的值

8. 可以使用 Loop + EndLoop 关键字定义循环
#{Loop 0}
will not print
#{EndLoop}

#{Loop 3}
Hello World
hhhhhhhh
#{EndLoop}

9. 循环进阶，可以将当前循环次数存入数据中，可以嵌套循环
#{Loop #{.map1.*Length} index=map1.index}
index1: #{map1.index} index2: #{map2.index}
map1 content: #{GetKey .map1 #{map1.index}}
#{  Loop #{.map2.*Length} index=map2.index  }
index1: #{map1.index} index2: #{map2.index}
map2 content: #{GetKey .map2 #{map2.index}}
#{  EndLoop  }
#{EndLoop}

10. 可以使用 If + EndIf 关键字定义条件表达式
#{If .a}
.a has a value, and .a != "" ("" will equal to undefined in If block)
#{EndIf}

#{If .noset}
will not print
#{EndIf}

#{If true}
also support "true" text
#{EndIf}

#{Define .flag=true}
#{If #{.flag}}
Logically, you can set a value "true" in any key
#{EndIf}

#{If false}
will not print
#{EndIf}

11. 可以使用 Else 在 If 和 EndIf 之间，以达到区块选择的效果
#{If .money}
We have money: $#{.money}.
#{Else}
We don't have any money.
#{EndIf}

12. 灵活组合他们，必要时可以通过提供的接口扩展他们。想要增加一个关键字？想要改变文件输出方式？都能实现
您可以继续查看其他 Demo，了解真实场景下如何使用 CodeGenerator，或者即刻开始使用。