package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"wireflow/internal/infra"
	"wireflow/monitor"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	peerManager := infra.NewPeerManager()
	runner := monitor.NewMonitorRunner(peerManager)
	runner.Run(ctx)
	
}
