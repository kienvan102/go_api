package studentbiz 

import (
	"context"
)

type DeleteStudentStore interface{
	Delete(ctx context.Context, id int) error
}

type deleteStudentBiz struct {
	store DeleteStudentStore
}

func NewDeleteStudentBiz(store DeleteStudentStore) *deleteStudentBiz {
	return &deleteStudentBiz{store: store}
}


func (b *deleteStudentBiz) DeleteStudent(ctx context.Context, id int) error {
	if err := b.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
