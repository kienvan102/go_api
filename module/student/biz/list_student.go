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
	if paging.Page < 0{
		paging.Page = 1
	}
	if paging.Limit < 0{
        paging.Limit = 5
    }
	data, err := b.store.List(ctx, paging)
	if err != nil {
        return nil, err
    }

	return data, nil
}