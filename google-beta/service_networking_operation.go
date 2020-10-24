package google

import (
	"encoding/json"
	"time"

	"google.golang.org/api/servicenetworking/v1"
)

type ServiceNetworkingOperationWaiter struct {
	Service *servicenetworking.APIService
	CommonOperationWaiter
}

func (w *ServiceNetworkingOperationWaiter) QueryOp() (interface{}, error) {
	return w.Service.Operations.Get(w.Op.Name).Do()
}

func serviceNetworkOperationWaitTimeWithResponse(config *Config, op *servicenetworking.Operation, response *servicenetworking.Subnetwork, activity, userAgent string, timeout time.Duration) error {
	w := &ServiceNetworkingOperationWaiter{
		Service: config.NewServiceNetworkingClient(userAgent),
	}

	if err := w.SetOp(op); err != nil {
		return err
	}
	if err := OperationWait(w, activity, timeout, config.PollInterval); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}


func serviceNetworkingOperationWaitTime(config *Config, op *servicenetworking.Operation, activity, userAgent string, timeout time.Duration) error {
	w := &ServiceNetworkingOperationWaiter{
		Service: config.NewServiceNetworkingClient(userAgent),
	}

	if err := w.SetOp(op); err != nil {
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
