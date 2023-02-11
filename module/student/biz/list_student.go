package studentbiz

import (
	"context"
	studentmodel "first-apis/module/student/model"
)

type ListStudentStore interface {
	List(ctx context.Context, paging studentmodel.Paging)  ([]studentmodel.Student,error)
}

type listStudentBiz struct{
	store ListStudentStore
}

func NewListStudentBiz(store ListStudentStore) *listStudentBiz {
	return &listStudentBiz{store: store}
}

func (b *listStudentBiz) ListStudent(ctx context.Context, paging studentmodel.Paging) ([]studentmodel.Student, error) {
	
	data, err := b.store.List(ctx, paging)
	if err!= nil {
        return nil, err
    }

	return data, nil
}