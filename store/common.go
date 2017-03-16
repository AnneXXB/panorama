package store

import (
	. "deephealth/health"
)

type ViewStorage interface {
	// Add a subject to the observing subject list
	ObserveSubject(subject EntityId, reply *bool) error

	// Stop observing a particular subject, all the reports
	// concerning this subject will be ignored
	StopObservingSubject(subject EntityId, reply *bool) error

	// Add a report to the view storage
	AddReport(report *HReport, reply *int) error
}