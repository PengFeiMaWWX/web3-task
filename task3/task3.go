package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、
age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type student struct {
	Id    int64  `gorm:"primarykey;column:id;autoIncrement;comment:主键ID"`
	Name  string `gorm:"comment:姓名;size:64;not null"`
	Age   int    `gorm:"comment:年龄"`
	Grade string `gorm:"学生年级;size:64;not null"`
}
type studentService struct {
	db *gorm.DB
}

func NewStudentService() (*studentService, error) {
	// 数据库连接
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//log.Fatal("数据库连接失败" + err.Error()) 打印日志并调用 os.Exit(1) 严重错误，必须终止程序
		log.Println("数据库连接失败" + err.Error())
		//严重错误，必须终止程序
		return nil, fmt.Errorf("数据库连接失败 %w", err)
	}
	// 自动迁移（创建表并添加注释）
	err = db.AutoMigrate(&student{})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败 %w", err)
	}
	log.Printf("学生表添加成功！")
	return &studentService{db: db}, nil
}

// 根据学生id查询学生信息的方法
func (s *studentService) GetStudent(id int64) (*student, error) {
	var stu student
	err := s.db.Where("id=?", id).Error
	return &stu, err
}

// 添加学生的方法
func (s *studentService) CreateStudent(stu *student) error {
	err := s.db.Create(stu).Error
	return err
}

// 更新学生的方法
func (s *studentService) UpdateStudentByStruct(id int64, stu *student) error {
	//err := s.db.Model(stu).Where("id=?", id).Updates(stu).Error
	result := s.db.Model(stu).Where("id=?", id).Updates(stu)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("学生不存在或数据未变化")

	}
	return nil
}

// 删除学生的方法
func (s *studentService) remove(id int64) error {
	result := s.db.Where("id=?", id).Delete(&student{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func main() {

	studentService, err := NewStudentService()
	if err != nil {
		log.Fatal(err)
	}

	err = studentService.CreateStudent(&student{Name: "张三", Age: 20, Grade: "12"})
	if err != nil {
		log.Println("添加学生信息失败！")
	}
	// 使用service进行crud的操作
	stu, err := studentService.GetStudent(1)
	if err != nil {
		return
	}
	log.Printf("student:%+v\n", stu)

	updateData := &student{
		Age: 21,
	}
	err = studentService.UpdateStudentByStruct(3, updateData)
	if err != nil {
		return
	}
	err = studentService.remove(2)
	if err != nil {
		return
	}
}
