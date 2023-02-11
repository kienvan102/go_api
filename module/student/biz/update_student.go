package studentbiz

import (
	"context"
	studentmodel "first-apis/module/student/model"
	// studentmodel "first-apis/module/student/model"
)

type UpdateStudentStore interface {
	Update(ctx context.Context, id int, data *studentmodel.StudentUpdate)  error
}

type updateStudentBiz struct{
	store UpdateStudentStore
}

func NewUpdateStudentBiz(store UpdateStudentStore) *updateStudentBiz {
	return &updateStudentBiz{store: store}
}

func (biz *updateStudentBiz) UpdateStudent(ctx context.Context, id int, data *studentmodel.StudentUpdate) error {

	err := biz.store.Update(ctx, id, data)
	if err!= nil {
        return err
    }

	return nil
}