// Copyright 2019 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package http

import (
	"net/http"

	"contrib.go.opencensus.io/exporter/prometheus"

	"github.com/go-kivik/kivik"

	"github.com/enseadaio/enseada/internal/cachecontrol"

	"github.com/labstack/echo"
)

type HealthCheckResponse struct {
	Status   string `json:"status"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Msg      string `json:"message"`
}

type ObservabilityHandler struct {
	data *kivik.Client
	exp  *prometheus.Exporter
}

func (h *ObservabilityHandler) health(c echo.Context) error {
	req := c.Request()
	res := HealthCheckResponse{
		Status:   "UP",
		Protocol: req.Proto,
		Host:     req.Host,
		Method:   req.Method,
		Path:     req.URL.Path,
		Msg:      "all systems operational",
	}
	sc := http.StatusOK

	up, err := h.data.Ping(req.Context())
	if err != nil || !up {
		c.Logger().Error("database unavailable", err)
		res.Status = "DOWN"
		res.Msg = "database unavailable"
		sc = http.StatusServiceUnavailable
	}

	cc := cachecontrol.NoStore(true)
	cc.Write(c.Response().Writer)
	return c.JSON(sc, res)
}

func (h *ObservabilityHandler) metrics(c echo.Context) error {
	r := c.Request()
	w := c.Response()

	cc := cachecontrol.NoStore(true)
	cc.Write(w)
	h.exp.ServeHTTP(w, r)
	return nil
}
