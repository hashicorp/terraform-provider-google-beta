package google

import (
	file "google.golang.org/api/file/v1beta1"
)

type FilestoreOperationWaiter struct {
	Service *file.ProjectsLocationsService
	CommonOperationWaiter
}

func (w *FilestoreOperationWaiter) QueryOp() (interface{}, error) {
	return w.Service.Operations.Get(w.Op.Name).Do()
}

func filestoreOperationWaitTime(service *file.Service, op *file.Operation, project, activity string, timeoutMinutes int) error {
	w := &FilestoreOperationWaiter{
		Service: service.Projects.Locations,
	}
	if err := w.SetOp(op); err != nil {
		return err
	}
	return OperationWait(w, activity, timeoutMinutes)
}
