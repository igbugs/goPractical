package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"day8/homework/models"
	"time"
)

func InsertStudent(student *models.Student) (err error) {
	if student == nil {
		err = fmt.Errorf("invalid Student parameter")
		return
	}

	sqlstr := "select student_id from student where student_id = ?"
	var studentId string
	err = Db.Get(&studentId, sqlstr, student.StudentId)
	if err == sql.ErrNoRows {
		// 插入操作
		sqlstr = `insert into student (
					  student_id, password, stu_name, age, sex, grade, create_time
				  ) values (?, ?, ?, ?, ?, ?, ?)`
		_, err = Db.Exec(sqlstr, student.StudentId, student.PassWord, student.Name, student.Age, student.Sex, student.Grade, time.Now())
		if err != nil {
			return
		}
		return
	}

	if err != nil {
		return
	}

	err = fmt.Errorf("student_id:%s is already exists", studentId)
	return
}

func UpdateStudent(student *models.Student) (err error) {
	if student == nil {
		err = fmt.Errorf("invalid Student parameter")
		return
	}

	// 更新操作
	sqlstr := `update student set 
					stu_name = ?, age = ?, grade = ?, update_time = ?
				where 
				    student_id = ?`
	result, err := Db.Exec(sqlstr, student.Name, student.Age, student.Grade, time.Now(), student.StudentId)
	if err != nil {
		return
	}

	affects, err := result.RowsAffected()
	if err != nil {
		return
	}

	if affects == 0 {
		err = fmt.Errorf("update Student failed, Student_id:%s, not found", student.StudentId)
		return
	}
	return
}

func QueryStudent(studentId string) (student *models.Student, err error) {
	// 查询操作
	sqlstr := `select
					stu_name, age, grade
				from
					student
				where
					student_id = ?`
	student = &models.Student{}
	err = Db.Get(student, sqlstr, studentId)
	if err != nil {
		return
	}
	return
}

func DeleteStudent(studentid string) (err error) {
	sqlstr := "select student_id from student where student_id=?"
	var studentId string
	err = Db.Get(&studentId, sqlstr, studentid)
	if err == sql.ErrNoRows {
		fmt.Printf("Studentid: %s not found", studentid)
		return err
	}

	if err != nil {
		return err
	}

	// 删除操作
	sqlstr = `delete from
					student
				where 
					student_id = ?`
	_, err = Db.Exec(sqlstr, studentid)
	if err != nil {
		return nil
	}
	return
}
