// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"

	"google.golang.org/api/googleapi"
)

type ServiceUsageOperationWaiter struct {
	Config     *Config
	UserAgent  string
	Project    string
	retryCount int
	CommonOperationWaiter
}

func (w *ServiceUsageOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("%s%s", w.Config.ServiceUsageBasePath, w.CommonOperationWaiter.Op.Name)

	return sendRequest(w.Config, "GET", w.Project, url, w.UserAgent, nil)
}

func (w *ServiceUsageOperationWaiter) IsRetryable(err error) bool {
	// Retries errors on 403 3 times if the error message
	// returned contains `has not been used in project`
	maxRetries := 3
	if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 403 {
		if w.retryCount < maxRetries && strings.Contains(gerr.Body, "has not been used in project") {
			w.retryCount += 1
			log.Printf("[DEBUG] retrying on 403 %v more times", w.retryCount-maxRetries-1)
			return true
		}
	}
	return false
}

func createServiceUsageWaiter(config *Config, op map[string]interface{}, project, activity, userAgent string) (*ServiceUsageOperationWaiter, error) {
	w := &ServiceUsageOperationWaiter{
		Config:    config,
		UserAgent: userAgent,
		Project:   project,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

// nolint: deadcode,unused
func serviceUsageOperationWaitTimeWithResponse(config *Config, op map[string]interface{}, response *map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	w, err := createServiceUsageWaiter(config, op, project, activity, userAgent)
	if err != nil {
		return err
	}
	if err := OperationWait(w, activity, timeout, config.PollInterval); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func serviceUsageOperationWaitTime(config *Config, op map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w, err := createServiceUsageWaiter(config, op, project, activity, userAgent)
	if err != nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
