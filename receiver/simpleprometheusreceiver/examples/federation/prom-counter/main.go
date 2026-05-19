// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/otel/attribute"
	api "go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

func initMeter() api.Meter { _ = "STUB: not implemented"; return *new(api.Meter) }

func main() {
	// set up prometheus
	meter := initMeter()
	// logging
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()
	logger.Info("Start Prometheus metrics app")
	valueRecorder, err := meter.Int64Histogram("prom_counter")
	if err != nil {
		log.Panicf("failed to initialize histogram %v", err)
	}
	ctx := context.Background()
	valueRecorder.Record(ctx, 0)
	commonLabels := []attribute.KeyValue{attribute.String("A", "1"), attribute.String("B", "2"), attribute.String("C", "3")}
	counter := int64(0)
	valueRecorder.Record(ctx, counter, api.WithAttributes(commonLabels...))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			time.Sleep(1 * time.Second)
			counter++
			valueRecorder.Record(ctx, counter, api.WithAttributes(commonLabels...))
			break
		case <-c:
			ticker.Stop()
			logger.Info("Stop Prometheus metrics app")
			return
		}
	}
}
