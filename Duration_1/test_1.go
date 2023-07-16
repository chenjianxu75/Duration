package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type SystemStats struct {
	CPUUsage    float64
	MemoryUsage float64
}

func main() {
	// 读取 Ubuntu 文本文件
	contents, err := ioutil.ReadFile("C:/Users/17811/Desktop/testdata.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 将文件内容转换为字符串
	fileStr := string(contents)

	// 提取 CPU 使用率和内存使用量数据
	stats, err := extractStats(fileStr)
	if err != nil {
		fmt.Println("Error extracting stats:", err)
		return
	}

	// 输出数据点的表头
	fmt.Println("Data Point,CPU Usage (%),Memory Usage (%)")

	// 输出数据点的内容
	for i, stat := range stats {
		fmt.Printf("%d,%.2f,%.2f\n", i+1, stat.CPUUsage, stat.MemoryUsage)
	}
}

func extractStats(fileStr string) ([]SystemStats, error) {
	// 定义提取 CPU 使用率和内存使用量的正则表达式
	statsRegex := regexp.MustCompile(`\s+(\d+\.\d+)\s+(\d+\.\d+)\s+\d+\:\d+\.\d+\s+\w+`)

	// 使用正则表达式提取 CPU 使用率和内存使用量
	matches := statsRegex.FindAllStringSubmatch(fileStr, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("No stats found")
	}

	// 存储数据点的切片
	stats := make([]SystemStats, len(matches))

	for i, match := range matches {
		cpuUsage, err := strconv.ParseFloat(strings.TrimSpace(match[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("Error parsing CPU usage: %v", err)
		}

		memUsage, err := strconv.ParseFloat(strings.TrimSpace(match[2]), 64)
		if err != nil {
			return nil, fmt.Errorf("Error parsing memory usage: %v", err)
		}

		stats[i] = SystemStats{
			CPUUsage:    cpuUsage,
			MemoryUsage: memUsage,
		}
	}

	return stats, nil
}
