package maven

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/chartmuseum/storage"
	"github.com/enseadaio/enseada/internal/couch"
	"github.com/enseadaio/enseada/internal/maven"
	"github.com/go-kivik/kivik"
	"github.com/labstack/echo"
)

func Boot(ctx context.Context, logger echo.Logger, e *echo.Echo, data *kivik.Client, store storage.Backend, enforcer *casbin.Enforcer) error {
	mvn := maven.New(logger, data, store, enforcer)
	mvn.MountRoutes(e)

	if err := couch.Transact(ctx, data, migrateDB, couch.MavenDB); err != nil {
		return err
	}

	return nil
}

func migrateDB(ctx context.Context, data *kivik.Client) error {
	if err := couch.InitDb(ctx, data, couch.MavenDB); err != nil {
		return err
	}

	if err := couch.InitIndex(ctx, data, couch.MavenDB, "kind_index", map[string]interface{}{
		"fields": []string{"kind"},
	}); err != nil {
		return err
	}

	if err := couch.InitIndex(ctx, data, couch.MavenDB, "file_index", map[string]interface{}{
		"fields": []string{"files"},
	}); err != nil {
		return err
	}

	return nil
}
