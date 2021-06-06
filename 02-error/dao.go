package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)


func LoadStudents(db *sql.DB) ([]Student, error){
	ctx, stop := context.WithCancel(context.Background())
	defer  stop()
	query := "SELECT student.name,student.age,student.comments FROM student"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}else{
			return nil, xerrors.Wrapf(err, "LoadStudents query fail")
		}
	}
	defer rows.Close()

	res := []Student{}

	for rows.Next(){
		var s Student
		err = rows.Scan(&s.Name, &s.Age, &s.Comments)
		if err != nil {
			return nil, xerrors.Wrapf(err, "LoadStudents rows scan fail")
		}
		res = append(res, s)
	}

	if err := rows.Err(); err != nil {
		return nil, xerrors.Wrapf(err, "LoadStudents rows next fail")
	}

	return res, nil
}


func LoadStudentByName(db *sql.DB, name string) (*Student, error){
	ctx, stop := context.WithCancel(context.Background())
	defer  stop()
	query := "SELECT student.name,student.age,student.comments FROM student where student.name=?"
	res := &Student{}
	err := db.QueryRowContext(ctx, query, name).Scan(&res.Name, &res.Age, &res.Comments)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFoundError{Err: err, Msg: "NOT Found"}
		}else{
			return nil, xerrors.Wrapf(err, "LoadStudentByName query fail")
		}
	}

	return res, nil
}