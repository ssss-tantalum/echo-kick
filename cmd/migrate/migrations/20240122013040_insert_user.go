package migrations

import (
	"context"
	"fmt"

	"github.com/ssss-tantalum/echo-kick/internal/model"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		user := model.User{
			Username: "ssss.tantalum",
			Email:    "ssss.tantalum@gmail.com",
		}
		_, err := db.NewInsert().Model(&user).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		var user model.User
		_, err := db.NewDelete().Model(&user).Where("username = ?", "ssss.tantalum").Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
