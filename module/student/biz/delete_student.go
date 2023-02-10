package studentbiz 

import (
	"context"
	"errors"
	studentmodel "first-apis/module/student/model"
)

type DeleteStudentStore interface{
	FindDataWithCondition(ctx context.Context, 
		condition map[string]interface{},
		moreKeys ...string,
		) (*studentmodel.Student, error)
	Delete(ctx context.Context, id int) error
}

type deleteStudentBiz struct {
	store DeleteStudentStore
}

func NewDeleteStudentBiz(store DeleteStudentStore) *deleteStudentBiz {
	return &deleteStudentBiz{store: store}
}


func (b *deleteStudentBiz) DeleteStudent(ctx context.Context, id int) error {
	oldStudent, err := b.store.FindDataWithCondition(ctx,map[string]interface{}{"id": id})
	if err!= nil {
        return err
    }

	if oldStudent.Status == 0{
		return errors.New("student does not exist")
	}

	if err := b.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
