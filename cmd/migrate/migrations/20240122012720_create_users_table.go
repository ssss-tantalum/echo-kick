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
		_, err := db.NewCreateTable().Model((*model.User)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		_, err := db.NewDropTable().Model((*model.User)(nil)).IfExists().Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
