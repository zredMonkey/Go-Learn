# 一、名称

规则：开头是一个字母或者下划线，后面可以跟任意数量的字符、数字、下划线，并且区分大小写。


# 二、声明
有四个主要声明：变量(var)、常量(const)、类型(type)、函数(func)。


# 三、变量

var name type = expression

```
var b,f,s = true, 2.3, "four" // bool float64,string

var f, err = os.Open(name) // os.Open() 返回的一个文件和一个错误
```

## 1.短变量声明
name := expresion, name的类型由expression决定

```
i := 100 // := 表示声明，而不是赋值，=表示赋值
```

注：
（1）如果一些变量在同一个词法块中声明，那么对于那些变量，短声明的行为等同于赋值。
（2）短变量声明最少声明一个新变量，否则，代码编译无法通过。

## 2.指针
指针的值是一个变量的地址。一个指针指示值所保存的位置，但不是所有的值都有地址。

```
x := 1
p := &x // p是整型指针，指向x
fmt.PrintLn(*p) // "1"
*p = 2
fmt.PrintLn(x) // "2"

```
注：
（1）两个指针当且仅当指向同一个变量或者两者都是nil的情况下才相等。

## 3. new函数

new(T)创建一个未命名的T类型变量，初始化为T类型的零值，并返回其地址(地址类型为*T)

```
p := new(int) // *int 类型的p,指向未命名的int变量
fmt.PrintLn(*p) // 输出“0”
*p = 2 // 把未命名的int设置成2
fmt.Println(*p) // 输出“2”
```

# 四、赋值

# 五、类型声明

# 六、包和文件

# 七、作用域

# 八、语法
## 1. for循环

用法一：for 赋值表达式; 判断条件; 赋值同时控制变量增减 { }

```
for i:=0; i<10; i++ { 
    // 循环10次
}
```

用法二：for 条件 { }

```
i := 0
for i<10 {
    // 只要条件满足就循环，尝尝在这里修改循环条件来控制循环
}
```
用法三：for {}

```
for {
    // 死循环。需要配合if 判断，当达到条件 break 退出循环
}
```
用法四：For-each range 循环
这种循环可以对字符串、数组、切片、字典等进行迭代，获取元素。有不同应用形式：
- 只获取索引（字典就是键）
```
var a [3]int  
for i := range a {
    fmt.Printf("%d %d\n", i, a[i])    
}

ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
for name := range ages {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
```
- 获取索引及对应元素（字典就是键和值）

```
var a [3]int  

for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
}


ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}
// 如果只想取值，不需要索引，可以用"_"下划线占位索引
for _, age := range ages {
    fmt.Printf("%d\n", age)
}
```
