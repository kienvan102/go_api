package studentbiz 

import (
	"context"
	"errors"
	studentmodel "first-apis/module/student/model"
)

type CreateStudentStore interface {
	Create(ctx context.Context, data *studentmodel.StudentCreate) error
}

type createStudentBiz struct{
	store CreateStudentStore
}

func NewCreateStudentBiz(store CreateStudentStore) *createStudentBiz {
	return &createStudentBiz{store: store}
}

func (biz *createStudentBiz) CreateStudent(ctx context.Context, data *studentmodel.StudentCreate) error {
    if data.FullName == ""{
		return errors.New("fullname could not be empty")
    }

	if data.Sex == ""{
		return errors.New("gentle could not be empty")
    }
	
	if err := biz.store.Create(ctx, data); err != nil {
	    return err
    }

	return nil
}