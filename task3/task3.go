package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

// 用户模型
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`                        // 主键编号
	Name      string    `gorm:"size:64; not null" json:"name"`               // 用户姓名
	Email     string    `gorm:"size:255; uniqueIdex; not null" json:"email"` // 邮件地址
	PostCount int       `gorm:"default:0" json:"post_count"`                 // 文章统计数
	CreatedAt time.Time `json:"created_at"`                                  // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                                  // 更新时间

	// 一对多关系： 一个用户关联多个文章
	Posts []Post `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"posts,omitempty"`
}

// 文章模型
type Post struct {
	ID            uint      `gorm:"primarykey" json:"id"`                 // 主键编号
	Title         string    `gorm:"size:255; not null" json:"title"`      // 文章标题
	Content       string    `gorm:"size:255; not null" json:"content"`    // 文章内容
	UserID        uint      `gorm:"index" json:"user_id"`                 // 用户编号 外键， 关联用户模型
	CommentsCount uint      `gorm:"default:0" json:"comments_count"`      // 文章评论数量
	CommentStatus string    `gorm:"default:'没有评论'" json:"comment_status"` // 文章评论状态
	CreatedAt     time.Time `json:"created_at"`                           // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`                           // 更新时间

	// 多对一关系：多个文档属于用户
	User User `gorm:"foreignkey:UserID" json:"user,omitempty"`
	// 一对多关系：一个文文章有多个评论
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments,omitempty"`
}

// 文章评论
type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`              // 评论编号
	Content   string    `gorm:"size:255; not null" json:"content"` // 评论内容
	PostId    uint      `gorm:"index; not null" json:"post_id"`    // 文章编号
	UserID    uint      `gorm:"index; not null" json:"user_id"`    // 用户编号
	CreatedAt time.Time `json:"created_at"`                        //创建时间
	UpdatedAt time.Time `json:"updated_at"`                        //更新时间

	// 多对一关系 ： 多个评论属于一个文章
	Post Post `gorm:"foreignkey:PostId" json:"post,omitempty"`
	// 多对一关系： 多个评论属于一个用户
	User User `gorm:"foreignkey:UserID" json:"user,omitempty"`
}

var db *gorm.DB

// 初始化数据库
func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("数据库表迁移失败！")
	}

	fmt.Println("数据库连接成功， 表结构已经创建！")
}

// BeforeCreate Post 模型的钩子函数：在创建文档前更新用户文章数量
func (p *Post) BeforeSave(tx *gorm.DB) error {

	// 更新用户文章的数量
	result := tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在！")
	}
	return nil
}

// AfterdDelete Comment 函数：删除评论后检查文章的评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 查询文章的评论的数量
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count).Error; err != nil {
		return err
	}
	// 如果评论数量为0 ， 更新文章评论状态
	if count == 0 {
		if err := tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "无评论").Error; err != nil {
			return err
		}
	}
	// 更新文章的评论数量
	if err := tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comments_count", count).Error; err != nil {
		return err
	}
	return nil
}

// 题目2 ：关联查询函数
// GetUserPostsWithComments 查询某个用户发布的所有文章及其对应的评论信息
func GetUserPostsWithComments(userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("User").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User") // 同时预加载品论的
	}).Where("user_id = ?", userID).
		Find(&posts).Error
	return posts, err
}

// GetMostCommentedPost 查询评论数量最多的文章信息
func GetMostCommentedPost() (*Post, error) {
	var post Post
	err := db.Preload("User").
		Preload("Comments").
		Order("comments_count DESC").
		First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// createTestData 创建测试数据
func createTestData() {
	// 清空现有数据
	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM users")

	// 创建用户
	users := []User{
		{Name: "张三", Email: "zhangsan@example.com"},
		{Name: "李四", Email: "lisi@example.com"},
	}

	for i := range users {
		if err := db.Create(&users[i]).Error; err != nil {
			log.Fatal("创建用户失败:", err)
		}
	}

	// 创建文章
	posts := []Post{
		{Title: "Go语言入门", Content: "这是Go语言的入门教程...", UserID: users[0].ID},
		{Title: "Gorm使用指南", Content: "Gorm是Go语言的ORM框架...", UserID: users[0].ID},
		{Title: "Web开发实践", Content: "使用Go进行Web开发...", UserID: users[1].ID},
	}

	for i := range posts {
		if err := db.Create(&posts[i]).Error; err != nil {
			log.Fatal("创建文章失败:", err)
		}
	}

	// 创建评论
	comments := []Comment{
		{Content: "很好的文章！", PostId: posts[0].ID, UserID: users[1].ID},
		{Content: "学到了很多，谢谢！", PostId: posts[0].ID, UserID: users[0].ID},
		{Content: "期待更多内容", PostId: posts[1].ID, UserID: users[1].ID},
		{Content: "非常实用", PostId: posts[2].ID, UserID: users[0].ID},
	}

	for i := range comments {
		if err := db.Create(&comments[i]).Error; err != nil {
			log.Fatal("创建评论失败:", err)
		}
	}

	// 更新文章的评论数量
	for _, post := range posts {
		var count int64
		db.Model(&Comment{}).Where("post_id = ?", post.ID).Count(&count)
		db.Model(&post).Update("comments_count", count)
		if count > 0 {
			db.Model(&post).Update("comment_status", "有评论")
		}
	}

	fmt.Println("测试数据创建完成")
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

	// 初始化数据库
	initDB()

	// 创建测试数据
	createTestData()

	// 测试关联查询1 ：查询用户1的所有文章及评论
	fmt.Println("\n=== 测试1：查询用户1的所有文章及评论 ===")
	posts, err := GetUserPostsWithComments(1)
	if err != nil {
		log.Fatal("查询失败", err)
	}
	for _, post := range posts {
		fmt.Println("文章 %s （作者：%s , 评论数量：%d）\n", post.Title, post.Content, post.CommentsCount)
		for _, comment := range post.Comments {
			fmt.Printf("  - 评论: %s (评论者: %s)\n", comment.Content, comment.User.Name)

		}

	}
	// 测试关联查询2：查询评论最多的文章
	fmt.Println("\n=== 测试2：查询评论最多的文章 ===")
	mostCommentedPost, err := GetMostCommentedPost()
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	fmt.Printf("评论最多的文章: %s (作者: %s, 评论数: %d)\n",
		mostCommentedPost.Title, mostCommentedPost.User.Name, mostCommentedPost.CommentsCount)

	// 测试钩子函数：删除评论
	fmt.Println("\n=== 测试3：测试评论删除钩子函数 ===")
	var comment Comment
	db.First(&comment)
	fmt.Printf("删除评论前，文章ID=%d的评论状态: %s\n", comment.PostId, comment.Post.CommentStatus)

	// 删除评论
	db.Delete(&comment)

	// 重新查询文章状态
	var updatedPost Post
	db.First(&updatedPost, comment.PostId)
	fmt.Printf("删除评论后，文章ID=%d的评论状态: %s, 评论数: %d\n",
		updatedPost.ID, updatedPost.CommentStatus, updatedPost.CommentsCount)

}
