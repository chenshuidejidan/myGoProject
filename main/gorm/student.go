package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type student struct {
	gorm.Model        //id, CreatedAt, UpdatedAt,DeletedAt
	Sid        string `gorm:"index"`
	Name       string
	Sex        bool
	Age        int `gorm:"default:18"`
}

func (s student) String() string {
	sex := "男"
	if !s.Sex {
		sex = "女"
	}
	return fmt.Sprintf("student{ sid=%s, name=%s, sex=%s, age=%d, CreatedAt=%s, UpdatedAt=%s }", s.Sid, s.Name, sex, s.Age, s.CreatedAt.Format("2006-01-02 15:04:05"), s.UpdatedAt.Format("2006-01-02 15:04:05"))
}

func AddStudent(s ...student) bool {
	for _, student := range s {
		res := db.Create(&student)
		if res.Error != nil || res.ID == 0 {
			fmt.Errorf("addStudent error!")
			return false
		}
	}
	return true
}

func GetStudentById(s *student) (ok bool, err error) {
	ok = false
	err = db.First(&s, "sid = ?", s.Sid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if err != nil {
		fmt.Errorf("getStudent error")
		return
	}
	ok = true
	return
}

func DeleteStudentByStudentId(s ...student) {
	for _, student := range s {
		db.Where("sid=?", student.Sid).Delete(&student)
	}
}

//func (s student) BeforeCreate(tx *gorm.DB) (err error) {
//	if s.Age < 10 || s.Age > 30 {
//		return fmt.Errorf("invalid age")
//	}
//	return
//}

func main() {
	InitDB()
	db.AutoMigrate(&student{}) //自动创建表格
	//AddStudent(student{
	//	Sid:  "1120110128",
	//	Name: "滴滴28",
	//	Sex:  true,
	//	Age:  19,
	//})
	//queryStudent := student{
	//	Sid: "1120110128",
	//}
	//err := GetStudentById(queryStudent)
	//if err != nil {
	//	fmt.Errorf(err.Error())
	//}
	//fmt.Printf("%v\n", queryStudent)
	//deleteStudent := student{
	//	Sid: "1120110128",
	//}
	//DeleteStudentByStudentId(deleteStudent)

	//students := []student{
	//	{
	//		Sid:  "1120110110",
	//		Name: "滴滴10",
	//		Sex:  true,
	//	},
	//	{
	//		Sid:  "1120110111",
	//		Name: "滴滴11",
	//		Sex:  false,
	//		Age:  40,
	//	}}
	//db.Session(&gorm.Session{SkipHooks: true}).Create(&students)
	multi := 2
	db.Model(&student{}).Where("sex=1").Update("age", gorm.Expr("age * ?", multi))
}
