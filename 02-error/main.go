package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
	"net/http"
	"os"
)


type DB struct {
	driverName string
	dataSourceName string
}

type Service struct {
	db *sql.DB
	gin *gin.Engine
}

type Student struct {
	Name string  `"json:"name"`
 	Age int `"json:"age"`
	Comments string `"json:"comments"`
}


func NewDb() DB{
	return DB{
		driverName: "mysql",
		dataSourceName: "root:123456@tcp(127.0.0.1:3306)/test1",
	}
}
// 初始化数据库
func DbInit() (*sql.DB ,error){
	connInfo := NewDb()
	db, err := sql.Open(connInfo.driverName, connInfo.dataSourceName)
	if err != nil {
		return nil, xerrors.Wrapf(err, "database open fail")
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return db, nil
}

func main() {

	db, err := DbInit()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(-1)
	}

	r := gin.Default()
	s := &Service{db: db, gin: r}
	s.RegGetStudentList()
	s.RegGetStudentByName()
	s.Run()
}

func (s Service)Run(){
	err := s.gin.Run()
	if err != nil {
		fmt.Printf("serve run: %+v\n", err)
		os.Exit(-1)
	}
}

// 获取学生列表
func (s Service) RegGetStudentList(){
	s.gin.GET("/students", func(c *gin.Context){
		sts, err := LoadStudents(s.db)
		if err != nil{
			// 如果查询结果出错，通过http状态码500响应程序处理内部错误
			fmt.Printf("%+v\n", err)
			c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}else{
			sts_, _ := json.Marshal(sts)
			c.JSON(http.StatusOK, string(sts_))
		}
	})
}
// 通过学生名字获取学生的信息
// 此处学生名称唯一
func (s Service)RegGetStudentByName(){
	s.gin.GET("/student/:name", func(c *gin.Context){
		name := c.Params.ByName("name")
		sts, err := LoadStudentByName(s.db, name)
		if err != nil {
			if _,ok := err.(*NotFoundError); ok {
				// 如果通过名称无法找到，通过http状态码404响应此学生不存在
				fmt.Printf("student %s not found\n", name)
				c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			}else{
				// 如果通过名称无法找到，通过http状态码500响应程序处理内部错误
				c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				fmt.Printf("%+v\n", err)
			}
		}else{
			sts_, _ := json.Marshal(sts)
			c.JSON(http.StatusOK, string(sts_))
		}
	})
}
