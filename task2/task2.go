package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
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

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person     // 匿名嵌入 Person，实现组合（字段和方法自动提升）
	EmployeeID string
}

func (e Employee) PrintInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", e.Name, e.Age)
}

// 题目七：题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来
// produce 生产者函数，生成1到10的整数并发送到通道
func produce(ch chan<- int, wg *sync.WaitGroup) {

	defer wg.Done() // 确保函数退出时 ， 通知WaitGroup
	defer close(ch) //关闭通道 ，通知接收方， 数据已经发送

	for i := 1; i <= 10; i++ {
		ch <- i // 将整数发送到通道
		fmt.Printf("发送 %d \n", i)
	}
	fmt.Printf("生产者发送完数据")
}

// 消费者函数 ，从通道中接受函数 并打印
func consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("接收到的整数 %d \n", num)
	}
	fmt.Printf("消费者对于函数接收完成")
}

// 题目八：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制
// BufferedChannelDemo 演示带有缓冲通道的生产者-消费者模型
func BufferedChannelDemo(bufferSize, itemCount int) {
	fmt.Printf("=== 开始演示缓冲通道（缓冲区大小：%d，数据量：%d）===\n", bufferSize, itemCount)
	// 创建带有指定缓冲大小的通道
	ch := make(chan int, bufferSize)
	var wg sync.WaitGroup

	// 启动生产者
	wg.Add(1)
	go producer(ch, itemCount, &wg)

	// 启动消费者
	wg.Add(1)

}

// producer 生产者函数：向通道发送指定数量的整数
func producer(ch chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	fmt.Printf("生产者开始工作，将生产 %d 个整数\n", count)
	for i := 1; i <= count; i++ {
		ch <- i //发送整数到通道
		fmt.Printf("生产整数 %d \n", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // 模拟生产耗时
	}
	fmt.Printf("生产者工作完成")

}

// consumer 消费者函数：从通道接收并打印整数
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	receivedCount := 0
	// 从通道中接收数据 ，直到通道关闭
	for num := range ch {
		fmt.Printf("消费 %d \n", num)
		receivedCount++
		time.Sleep(100 * time.Millisecond) // 模拟消费耗时（比生产慢）
	}
	fmt.Printf("消费者完成工作，共接收 %d 个整数\n", receivedCount)

}

// 题目 9 ：题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// SharedCounter 使用Mutex保护的共享计数器
type SharedCount struct {
	value int
	mu    sync.Mutex
}

// Increment 安全的递增计数器
func (s *SharedCount) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value++
}

// GetValue 获取当前计数器值
func (c *SharedCount) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value

}

// 题目10 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
// AtomicCounter 使用原子操作的无锁计数器
type AtomicCounter struct {
	value int64
}

// 原子递增计数器
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// GetValue 获取当前的计数器
func (c *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

// incrementWorker 协程工作函数
func incrementWorker(counter *SharedCount, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		counter.Increment()
	}
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

	// 题目 6
	// 创建 Employee 实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "111",
	}
	emp.PrintInfo()

	// 题目七：用于生产者-消费者协程通信
	// 创建整形无缓冲通道
	ch := make(chan int)

	// 使用waitGroup等待两个协程
	var wg2 sync.WaitGroup
	wg2.Add(2) // 设置需要等待的协程的数量

	// 自动生产者协程
	go produce(ch, &wg2) //wg的内存地址（指针）
	// 自动消费者协程
	go consume(ch, &wg2)
	wg2.Wait()
	fmt.Printf("测试完成 ， 所有协程执行完毕")

	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建带有缓冲的通道，缓冲区大小设为10
	ch8 := make(chan int, 10)

	// 使用WaitGroup等待生产者和消费者完成
	var wg8 sync.WaitGroup

	// 启动生产者协程
	wg8.Add(1)
	go producer(ch8, 100, &wg8)

	// 启动消费者协程
	wg8.Add(1)
	go consumer(ch8, &wg8)

	// 等待所有协程完成
	wg8.Wait()
	fmt.Println("程序执行完毕")

	// 题目9
	counter := &SharedCount{}
	var wg9 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg9.Add(1)
		go func(id int) {
			defer wg9.Done()
			// 每个协程递增1000次
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			fmt.Printf("协程%d 递增完成", id)
		}(i)
	}
	wg9.Wait()
	// 输出最终结果
	fmt.Printf("最终计数器值: %d (期望: 10000)\n", counter.GetValue())

	// 题目10
	counter10 := &SharedCount{}
	var wg10 sync.WaitGroup

	//启动10个协程
	for i := 0; i < 10; i++ {
		wg10.Add(1)
		go incrementWorker(counter10, 1000, &wg10)
	}

	// 等待所有协程完成
	wg10.Wait()

	// 输出最总结果
	fmt.Printf("最终计数器 : %d (期望: 10000)\n", counter10.GetValue())
}
