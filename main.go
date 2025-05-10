package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NeverStop-cmy/go-practice/advanced"
	"github.com/NeverStop-cmy/go-practice/basics"
	"github.com/NeverStop-cmy/go-practice/concurrency"
	"github.com/NeverStop-cmy/go-practice/data_structures"
	"github.com/NeverStop-cmy/go-practice/error_handling"
	"github.com/NeverStop-cmy/go-practice/functions"
	"github.com/NeverStop-cmy/go-practice/oop"
	"github.com/NeverStop-cmy/go-practice/stdlib"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=== Go语言练习项目 ===")

	for {
		printMenu()
		fmt.Print("请选择要练习的模块（输入编号，0退出）：")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "0" {
			fmt.Println("退出练习，欢迎下次使用！")
			break
		}

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 10 {
			fmt.Println("⚠️ 无效输入，请输入数字1-10或0退出")
			continue
		}

		executeModule(choice, scanner)
	}
}

func printMenu() {
	fmt.Println(`
=== 模块菜单 ===
1. 基础语法
2. 函数
3. 数据结构
4. 面向对象
5. 并发
6. 错误处理
7. 标准库
8. 高级主题
9. 运行全部基础示例
10. 运行全部高级示例
0. 立即退出
-----------------`)
}

func executeModule(choice int, scanner *bufio.Scanner) {
	switch choice {
	case 0: // 新增直接退出选项
		fmt.Println("退出练习，欢迎下次使用！")
		os.Exit(0) // 立即终止程序
	case 1:
		runSection("基础语法", []func(){
			basics.RunVariablesDemo,
			basics.RunDataTypesDemo,
			basics.RunControlFlowDemo,
		}, scanner)

	case 2:
		runSection("函数", []func(){
			functions.RunBasicFunctionsDemo,
			functions.RunAdvancedFunctionsDemo,
		}, scanner)

	case 3:
		runSection("数据结构", []func(){
			data_structures.RunArraysSlicesDemo,
			data_structures.RunMapsDemo,
		}, scanner)

	case 4:
		runSection("面向对象", []func(){
			oop.RunStructsDemo,
			oop.RunMethodsInterfacesDemo,
		}, scanner)

	case 5:
		runSection("并发", []func(){
			concurrency.RunGoroutinesDemo,
			concurrency.RunChannelsDemo,
		}, scanner)

	case 6:
		runSection("错误处理", []func(){
			error_handling.RunErrorsDemo,
		}, scanner)

	case 7:
		runSection("标准库", []func(){
			stdlib.RunJSONProcessingDemo,
			stdlib.RunHttpServerDemo,
		}, scanner)

	case 8:
		runSection("高级主题", []func(){
			advanced.RunClosuresDemo,
			advanced.RunReflectionDemo,
			advanced.RunContextDemo,
		}, scanner)

	case 9:
		runAllBasicExamples(scanner)

	case 10:
		runAllAdvancedExamples(scanner)
	}
}

func runSection(title string, demos []func(), scanner *bufio.Scanner) {
	fmt.Printf("\n=== 执行 %s 模块 ===\n", title)
	for _, demo := range demos {
		if demo != nil {
			demo()
			fmt.Println("\n" + strings.Repeat("-", 30))
		}
	}
	waitForContinue(scanner)
}

func runAllBasicExamples(scanner *bufio.Scanner) {
	fmt.Println("\n=== 运行全部基础示例 ===")
	runSection("基础语法", []func(){basics.RunVariablesDemo}, scanner)
	runSection("函数", []func(){functions.RunBasicFunctionsDemo}, scanner)
	// 可继续添加其他基础模块...
	waitForContinue(scanner)
}

func runAllAdvancedExamples(scanner *bufio.Scanner) {
	fmt.Println("\n=== 运行全部高级示例 ===")
	runSection("并发", []func(){concurrency.RunChannelsDemo}, scanner)
	runSection("高级主题", []func(){advanced.RunContextDemo}, scanner)
	// 可继续添加其他高级模块...
	waitForContinue(scanner)
}

func waitForContinue(scanner *bufio.Scanner) {
	fmt.Print(`
按回车键返回菜单...
或输入以下命令：
  exit  - 立即退出程序
  clear - 清屏
  help  - 显示帮助
`)
	for {
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		switch input {
		case "":
			fmt.Print("\033[H\033[2J") // 清屏（Windows可用cls命令）
			return
		case "exit", "quit", "q":
			fmt.Println("退出练习，欢迎下次使用！")
			os.Exit(0)
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "help":
			fmt.Print(`
=== 操作指南 ===
回车键   - 返回主菜单
exit    - 立即退出
clear   - 清屏
help    - 显示帮助
`)
		default:
			fmt.Println("⚠️ 未知命令，输入回车返回菜单")
		}
	}
}
