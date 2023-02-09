package studentbiz 

import (
	"context"
	"errors"
	studentmodel "first-apis/module/student/model"
)

type CreateStudentStore interface {
	Create(ctx context.Context, data *studentmodel.Student) error
}

type createStudentBiz struct{
	store CreateStudentStore
}

func NewCreateStudentBiz(store CreateStudentStore) *createStudentBiz {
	return &createStudentBiz{store: store}
}

func (biz *createStudentBiz) CreateStudent(ctx context.Context, data *studentmodel.Student) error {
    if data.FullName == ""{
		return errors.New("data is empty")
    }
	
	if err := biz.store.Create(ctx, data); err != nil {
	    return err
    }

	return nil
}