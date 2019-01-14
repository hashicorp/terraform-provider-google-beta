package google

import (
	accesscontextmanager "google.golang.org/api/accesscontextmanager/v1beta"
)

type AccessContextManagerOperationWaiter struct {
	Service *accesscontextmanager.OperationsService
	CommonOperationWaiter
}

func (w *AccessContextManagerOperationWaiter) QueryOp() (interface{}, error) {
	return w.Service.Get(w.Op.Name).Do()
}

func accessContextManagerOperationWaitTime(service *accesscontextmanager.Service, op *accesscontextmanager.Operation, activity string, timeoutMinutes int) error {
	w := &AccessContextManagerOperationWaiter{
		Service: service.Operations,
	}
	if err := w.SetOp(op); err != nil {
		return err
	}
	return OperationWait(w, activity, timeoutMinutes)
}
