// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
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

// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package enums

import (
	"fmt"
)

var (
	QueryResultType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Answered":    1,
		"Failed":      2,
	}
)

// QueryResultTypeFromString parses a QueryResultType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to QueryResultType
func QueryResultTypeFromString(s string) (QueryResultType, error) {
	if v, ok := QueryResultType_value[s]; ok {
		return QueryResultType(v), nil
	} else if v, ok := QueryResultType_shorthandValue[s]; ok {
		return QueryResultType(v), nil
	}
	return QueryResultType(0), fmt.Errorf("%s is not a valid QueryResultType", s)
}

var (
	QueryRejectCondition_shorthandValue = map[string]int32{
		"Unspecified":         0,
		"None":                1,
		"NotOpen":             2,
		"NotCompletedCleanly": 3,
	}
)

// QueryRejectConditionFromString parses a QueryRejectCondition value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to QueryRejectCondition
func QueryRejectConditionFromString(s string) (QueryRejectCondition, error) {
	if v, ok := QueryRejectCondition_value[s]; ok {
		return QueryRejectCondition(v), nil
	} else if v, ok := QueryRejectCondition_shorthandValue[s]; ok {
		return QueryRejectCondition(v), nil
	}
	return QueryRejectCondition(0), fmt.Errorf("%s is not a valid QueryRejectCondition", s)
}
