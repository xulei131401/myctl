#官方文档地址：https://cobra.dev/


#1.shorthand 只能是单个字符，虽然是字符串

#2.选项分类

1.persistent  //永久选项，它可以给子命令添加选项，同时也会被子命令的子命令所继承
//适用于全局性质的选项
//语法1  添加一个选项为foo，默认值为空，说明信息
//将这个信息赋予给Foo的变量
Foo = testCmd.PersistentFlags().String("foo", "", "A help for foo")



//语法2    先定义变量,再以指针的形式进行传递
var Print string
testCmd.PersistentFlags().StringVar(&Print, "print", "", "print")


//语法3    定义多个选项为相同功能，--show/-s
//默认值为false，其他同上
testCmd.PersistentFlags().BoolVarP(&show, "show", "s", false, "show")



2.local       //特定选项   只允许给特定的子命令去添加选项
//这个和上面的基本类似
//--showL / -S 如果后面有定义该选项则为true，如果没有定义则为false
showL = *testCmd.Flags().BoolP("showL", "S", false, "show")

//同上 语法不同
testCmd.Flags().StringVar(&PrintL, "printL", "", "print")

//同上 语法不同
FooL = testCmd.Flags().String("fooL", "", "A help for foo")


#1.go list命令的使用

go list -json -m
{
"Path": "holy-cmd",
"Main": true,
"Dir": "/Users/xulei/jungle/golangworkspace/holy-cmd",
"GoMod": "/Users/xulei/jungle/golangworkspace/holy-cmd/go.mod",
"GoVersion": "1.18"
}

go list -m -f '{{.GoMod}}'
/Users/xulei/jungle/golangworkspace/holy-cmd/go.mod

#2.file,tpl,cmd等

#3.embed包的使用
https://zhuanlan.zhihu.com/p/351931501

#4.antlr使用
antlr -Dlanguage=Go -o parser Calc.g4
给定规则，文本按照规则进行解析，解析成结构话数据进行分析

####################
目的：写一个通用的程序，提供好模板，传入变量，就能生成对应的文件到指定目录