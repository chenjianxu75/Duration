# Duartion_1
#### **输出产物:**

- 使用 top 命令获得 0.1 秒刷新率的系统数据并重定向到一个文本文件
- 使用 golang 读取改文本文件并使用正则表达式获得 cpu 使用率和内存使用量的数据储存到数组里面
- 输出数组中的数据点,可以使用**列复制** 复制所有数据, 粘贴到 excel 中然后通过 excel 软件绘制图形 (当然如果你也可以使用第三方 excel 库直接生成 excel 文件)

直接上最终图表图

![Example Image](https://github.com/chenjianxu75/Duration/blob/main/Duration_1/%E5%9B%BE%E8%A1%A8%E6%88%AA%E5%9B%BE.png)

ps:这是第一版代码，相关细节方面还不完善，最大的问题是，代码输出只能将输出数据按照一定格式排列，暂时无法调用excel库来直接制作并输出表格（因本人在后续几天需要出差，在出差结束后会将后续代码问题修改并在20号之前上传）
代码输出图：
![Example Image](https://github.com/chenjianxu75/Duration/blob/main/Duration_1/%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE%EF%BC%881%EF%BC%89.png)
![Example Image](https://github.com/chenjianxu75/Duration/blob/main/Duration_1/%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE%EF%BC%882%EF%BC%89.png)
此次任务内所展示表格是使用excel手动生成

> 下面是代码的详细报告



#### SystemStats

这是一个自定义的结构类型 SystemStats，它包含了 CPU 使用率和内存使用量的两个字段：

```sh
type SystemStats struct {
	CPUUsage    float64
	MemoryUsage float64
}
```
## main

在 main 函数中，它使用 ioutil.ReadFile 函数读取指定路径下的 Ubuntu 文本文件的内容，并将内容存储在 contents 变量中。如果读取文件时发生错误，它将打印错误信息并返回。

```sh
func main() {
	// 读取 Ubuntu 文本文件
	contents, err := ioutil.ReadFile("C:/Users/17811/Desktop/testdata.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
```

##### 将文件内容转换为字符串

它使用 string 函数将文件内容转换为字符串形式，存储在 fileStr 变量中。

```sh
	// 将文件内容转换为字符串
	fileStr := string(contents)
```

##### 提取 CPU 使用率和内存使用量数据

它调用 extractStats 函数，将文件内容字符串传递给该函数进行处理。extractStats 函数用于从文本中提取 CPU 使用率和内存使用量的数据，并返回一个包含这些数据的 SystemStats 类型的切片。如果提取过程中出现错误，它会打印错误信息并返回.

```sh
// 提取 CPU 使用率和内存使用量数据
	stats, err := extractStats(fileStr)
	if err != nil {
		fmt.Println("Error extracting stats:", err)
		return
	}
```

##### 输出部分

它依次遍历从 extractStats 函数返回的数据切片 stats，并使用 fmt.Printf 函数按照指定格式输出每个数据点的序号、CPU 使用率和内存使用量。

```sh
	// 输出数据点的表头
	fmt.Println("Data Point,CPU Usage (%),Memory Usage (%)")

	// 输出数据点的内容
	for i, stat := range stats {
		fmt.Printf("%d,%.2f,%.2f\n", i+1, stat.CPUUsage, stat.MemoryUsage)
	}
}

```
