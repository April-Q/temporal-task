// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package serviceerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// WorkflowExecutionAlreadyStarted represents workflow execution already started error.
	WorkflowExecutionAlreadyStarted struct {
		Message        string
		StartRequestId string
		RunId          string
		st             *status.Status
	}
)

// NewWorkflowExecutionAlreadyStarted returns new WorkflowExecutionAlreadyStarted error.
func NewWorkflowExecutionAlreadyStarted(message, startRequestId, runId string) error {
	return &WorkflowExecutionAlreadyStarted{
		Message:        message,
		StartRequestId: startRequestId,
		RunId:          runId,
	}
}

// Error returns string message.
func (e *WorkflowExecutionAlreadyStarted) Error() string {
	return e.Message
}

func (e *WorkflowExecutionAlreadyStarted) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.AlreadyExists, e.Message)
	st, _ = st.WithDetails(
		&errordetails.WorkflowExecutionAlreadyStartedFailure{
			StartRequestId: e.StartRequestId,
			RunId:          e.RunId,
		},
	)
	return st
}

func newWorkflowExecutionAlreadyStarted(st *status.Status, errDetails *errordetails.WorkflowExecutionAlreadyStartedFailure) error {
	return &WorkflowExecutionAlreadyStarted{
		Message:        st.Message(),
		StartRequestId: errDetails.GetStartRequestId(),
		RunId:          errDetails.GetRunId(),
		st:             st,
	}
}
