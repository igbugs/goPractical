package data

import (
	"day8/homework/models"
	"testing"
)

func init() {
	dns := "root:123456@tcp_chat(192.168.247.133:3306)/library_mgr?parseTime=True"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertStudent(t *testing.T) {
	var tests = []*models.Student{
		&models.Student{
			StudentId: "000000001",
			Name:      "lio",
			Age:       "10",
			Grade:     "F",
		},
		&models.Student{
			StudentId: "000000002",
			Name:      "lio",
			Age:       "10",
			Grade:     "F",
		},
		&models.Student{
			StudentId: "000000003",
			Name:      "lio",
			Age:       "10",
			Grade:     "F",
		},
		&models.Student{
			StudentId: "000000004",
			Name:      "lio",
			Age:       "10",
			Grade:     "F",
		},
	}

	for _, tt := range tests {
		err := InsertStudent(tt)
		if err != nil {
			t.Errorf("insert student failed, err:%v", err)
			return
		}
		t.Logf("insert student succ")
	}
}

func TestUpdateStudent(t *testing.T) {
	var student = models.Student{
		StudentId: "000000001",
		Name:      "lio",
		Age:       "10",
		Grade:     "F",
	}

	err := UpdateStudent(&student)
	if err != nil {
		t.Errorf("update student failed, err:%v", err)
		return
	}

	t.Logf("update student succ")
}

func TestQueryStudent(t *testing.T) {
	var student *models.Student
	studentId := "000000002"
	student, err := QueryStudent(studentId)
	if err != nil {
		t.Errorf("query student failed, err:%v", err)
		return
	}

	t.Logf("query student succ, book:%#v", student)
}

func TestDeleteStudent(t *testing.T) {
	studentId := "000000002"
	err := DeleteStudent(studentId)
	if err != nil {
		t.Errorf("delete student failed, err:%v", err)
		return
	}

	t.Logf("delete student succ")
}
