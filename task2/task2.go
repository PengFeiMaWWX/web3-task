package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// 指针
// 题目1：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
// incrementByTen 接收一个整数指针作为参数，并将其指向的值增加10
// 考察点 ：指针的使用、值传递与引用传递的区别
func incrementByTen(ptr *int) {
	*ptr += 10 // 通过解引用指针修改原始值
}

// 题目2：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
// 考察点 ：指针运算、切片操作。
func doubleSlice(slicePtr *[]int) {

	// 解引用指针获取原始切片
	slice := *slicePtr

	// 写法1：遍历每个切片，并将每个元素乘以2
	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}

	// 写法2：
	for i := range slice {
		slice[i] *= 2

	}
}

// 题目3：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 打印从1到10的奇数
func printOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知waitGroup // 修改原始wg的状态 Done()调用会实际减少 WaitGroup 的计数器
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数：%d \n", i)
	}

}

// 打印从2到10的偶数
func printEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知waitGroup
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数%d \n", i)
	}
}

// 题目4
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度
// task 表示一个可执行的任务
type Task struct {
	Name     string       // 任务名称
	Function func() error // 任务函数
}

// TaskResult 存储任务执行的结果
type TaskResult struct {
	Name      string        // 任务名称
	Duration  time.Duration // 任务执行时间
	StartTime time.Time     // 开始时间
	EndTime   time.Time     // 结束时间
	Error     error         // 执行异常
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks []Task // 待执行的任务列表
}

// NewScheduler 创建新的调度器实例
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// AddTask 添加任务到调度器
func (s *Scheduler) AddTask(name string, function func() error) {
	s.tasks = append(s.tasks,
		Task{
			Name:     name,
			Function: function,
		})
}

// Run 并发执行所有任务并返回任务执行结果
func (s *Scheduler) Run() []TaskResult {

	var wg sync.WaitGroup
	resultChan := make(chan TaskResult, len(s.tasks))

	// 启动协程执行每个任务
	for i, task := range s.tasks {
		wg.Add(1)
		go func(idx int, task Task) {
			defer wg.Done()
			start := time.Now()
			err := task.Function()
			duration := time.Since(start)
			end := time.Now()

			resultChan <- TaskResult{
				Name:      task.Name,
				Duration:  duration,
				StartTime: start,
				EndTime:   end,
				Error:     err,
			}
		}(i, task)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	var allResults []TaskResult
	for result := range resultChan {
		allResults = append(allResults, result)
	}
	return allResults
}

// 题目五：
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
// Shape 接口定义几何形状的基本操作
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width, Height float64
}

// 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func main() {

	// 题目1
	num := 5
	fmt.Printf("原始值%d \n", num)
	// 传递变量的地址给函数
	incrementByTen(&num)
	fmt.Printf("修改后的值%d \n", num)

	// 题目2
	// 创建并初始化一个整数切片
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("原始切片%d \n", numbers)

	// 传递切片的地址给函数
	doubleSlice(&numbers)
	fmt.Printf("修改后的切片%d \n", numbers)

	// 题目3 打印奇偶数
	var wg sync.WaitGroup
	wg.Add(2) // 等待两个协程完成

	// 启动打印奇数的协程
	go printOddNumbers(&wg)
	// 启动打印偶数的协程
	go printEvenNumbers(&wg)
	wg.Wait() // 能正确感知协程完成状态
	fmt.Printf("所有数字打印完成")

	// 题目4
	// 创建任务调度器
	scheduler := NewScheduler()
	// 添加示例任务
	scheduler.AddTask("任务A", func() error {
		time.Sleep(1 * time.Second)
		fmt.Print("任务A执行完成")
		return nil
	})
	scheduler.AddTask("任务B", func() error {
		time.Sleep(1 * time.Second)
		fmt.Print("任务B执行完成")
		return nil
	})
	scheduler.AddTask("任务C", func() error {
		time.Sleep(1 * time.Second)
		fmt.Print("任务C执行完成")
		return nil
	})
	results := scheduler.Run()
	fmt.Print("\n执行任务统计")
	for _, result := range results {
		status := "成功"
		if result.Error != nil {
			status = "失败"
		}
		fmt.Printf("任务: %-8s 耗时: %-12v 状态: %-4s 开始时间: %s\n",
			result.Name, result.Duration, status, result.StartTime.Format("15:04:05.000"))
	}

	// 题目5
	// 创建矩形实例
	rect := Rectangle{Width: 10, Height: 5}

	// 创建原型实例
	circle := Circle{Radius: 5}

	// 使用接口处理不同形状 创建 Shape接口类型的切片 存储不同具体类型（Rectangle 和 Circle） 通过接口统一调用方法，实现多态
	// 数组需要指定长度
	shapes := []Shape{rect, circle}

	for i, shape := range shapes {

		switch s := shape.(type) {
		case Circle:
			fmt.Printf("形状 %d: 圆形 (半径=%.1f)\n", i+1, s.Radius)
		case Rectangle:
			fmt.Printf("形状 %d: 矩形 (宽=%.1f, 高=%.1f)\n", i+1, s.Width, s.Height)

		}

		// 通过接口调用方法
		//%.2f由三部分组成：
		//
		//格式化说明符的开始
		//精度控制，表示保留 2 位小数
		//表示浮点数格式
		fmt.Printf("面积：%.2f\n", shape.Area())
		fmt.Printf("周长: %.2f\n", shape.Perimeter())
		fmt.Println("------")

	}
}
