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
	NamespaceState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Registered":  1,
		"Deprecated":  2,
		"Deleted":     3,
	}
)

// NamespaceStateFromString parses a NamespaceState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to NamespaceState
func NamespaceStateFromString(s string) (NamespaceState, error) {
	if v, ok := NamespaceState_value[s]; ok {
		return NamespaceState(v), nil
	} else if v, ok := NamespaceState_shorthandValue[s]; ok {
		return NamespaceState(v), nil
	}
	return NamespaceState(0), fmt.Errorf("%s is not a valid NamespaceState", s)
}

var (
	ArchivalState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Disabled":    1,
		"Enabled":     2,
	}
)

// ArchivalStateFromString parses a ArchivalState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ArchivalState
func ArchivalStateFromString(s string) (ArchivalState, error) {
	if v, ok := ArchivalState_value[s]; ok {
		return ArchivalState(v), nil
	} else if v, ok := ArchivalState_shorthandValue[s]; ok {
		return ArchivalState(v), nil
	}
	return ArchivalState(0), fmt.Errorf("%s is not a valid ArchivalState", s)
}

var (
	ReplicationState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Normal":      1,
		"Handover":    2,
	}
)

// ReplicationStateFromString parses a ReplicationState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ReplicationState
func ReplicationStateFromString(s string) (ReplicationState, error) {
	if v, ok := ReplicationState_value[s]; ok {
		return ReplicationState(v), nil
	} else if v, ok := ReplicationState_shorthandValue[s]; ok {
		return ReplicationState(v), nil
	}
	return ReplicationState(0), fmt.Errorf("%s is not a valid ReplicationState", s)
}
