package health

import "github.com/alexfalkowski/go-signal/v2"

// Module registers service health checks and HTTP and gRPC observers.
var Module = signal.Module(
	signal.Register(register),
	signal.Register(httpHealthObserver),
	signal.Register(httpLivenessObserver),
	signal.Register(httpReadinessObserver),
	signal.Register(grpcObserver),
)
