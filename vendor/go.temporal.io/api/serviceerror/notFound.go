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
	// NotFound represents not found error.
	NotFound struct {
		Message        string
		CurrentCluster string
		ActiveCluster  string
		st             *status.Status
	}
)

// NewNotFound returns new NotFound error.
func NewNotFound(message string) error {
	return &NotFound{
		Message: message,
	}
}

// Error returns string message.
func (e *NotFound) Error() string {
	return e.Message
}

func (e *NotFound) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.NotFound, e.Message)
	st, _ = st.WithDetails(
		&errordetails.NotFoundFailure{
			CurrentCluster: e.CurrentCluster,
			ActiveCluster:  e.ActiveCluster,
		},
	)
	return st
}

func newNotFound(st *status.Status, errDetails *errordetails.NotFoundFailure) error {
	return &NotFound{
		Message:        st.Message(),
		CurrentCluster: errDetails.GetCurrentCluster(),
		ActiveCluster:  errDetails.GetActiveCluster(),
		st:             st,
	}
}
