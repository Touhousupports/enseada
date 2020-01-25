// Copyright 2019-2020 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package maven

import (
	"context"

	"github.com/chartmuseum/storage"
	"github.com/enseadaio/enseada/internal/couch"
	"github.com/enseadaio/enseada/internal/maven"
	"github.com/enseadaio/enseada/pkg/log"
	"github.com/go-kivik/kivik"
)

type Module struct {
	logger log.Logger
	Maven  *maven.Maven
	data   *kivik.Client
}

type Deps struct {
	Logger  log.Logger
	Data    *kivik.Client
	Storage storage.Backend
}

func NewModule(ctx context.Context, deps Deps) (*Module, error) {
	logger := deps.Logger
	data := deps.Data
	st := deps.Storage

	mvn := maven.New(logger, data, st)

	if err := couch.Transact(ctx, logger, data, migrateDB, couch.MavenDB); err != nil {
		return nil, err
	}

	return &Module{
		logger: logger,
		Maven:  mvn,
		data:   data,
	}, nil
}

func (m *Module) Start(ctx context.Context) error {
	m.logger.Info("started maven module")
	return nil
}

func (m *Module) Stop(ctx context.Context) error {
	m.logger.Info("stopped maven module")
	return nil
}

func migrateDB(ctx context.Context, logger log.Logger, data *kivik.Client) error {
	if err := couch.InitDb(ctx, logger, data, couch.MavenDB); err != nil {
		return err
	}

	if err := couch.InitIndex(ctx, logger, data, couch.MavenDB, "kind_index", map[string]interface{}{
		"fields": []string{"kind"},
	}); err != nil {
		return err
	}

	if err := couch.InitIndex(ctx, logger, data, couch.MavenDB, "file_index", map[string]interface{}{
		"fields": []string{"files"},
	}); err != nil {
		return err
	}

	return nil
}
