package main

import (
	"fmt"
	"io"
	ss "module/pkg/repo/repo_students"
	s "module/pkg/student"
)

func main() {
	storage := ss.NewStudentsStorage()
	for {
		nS := s.NewStudent()
		fmt.Println("Введите имя, возраст и оценку студента:")
		_, err := fmt.Scan(&nS.Name, &nS.Age, &nS.Grade)
		if err == io.EOF {
			break
		}
		if err = storage.Put(nS); err != nil {
			fmt.Println("error:", err.Error())
			return
		}
	}
	fmt.Println("Хранилище студентов:")
	for _, v := range storage.Get() {
		fmt.Println(v.Name, v.Age, v.Grade)
	}
}
