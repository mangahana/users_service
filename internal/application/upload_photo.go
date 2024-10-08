package application

import (
	"context"
	"os"
)

func (u *useCase) UploadPhoto(ctx context.Context, userId int, filename, mime string, data []byte) error {
	if err := os.WriteFile(u.uploadFolder+filename, data, 0777); err != nil {
		return err
	}

	return u.repo.UpdatePhoto(ctx, userId, filename)
}
