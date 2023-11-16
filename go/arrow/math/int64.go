// Code generated by type.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package math

import (
	"github.com/apache/arrow/go/v15/arrow/array"
)

type Int64Funcs struct {
	sum func(a *array.Int64) int64
}

var (
	Int64 Int64Funcs
)

// Sum returns the summation of all elements in a.
func (f Int64Funcs) Sum(a *array.Int64) int64 {
	if a.Len() == 0 {
		return int64(0)
	}
	return f.sum(a)
}

func sum_int64_go(a *array.Int64) int64 {
	acc := int64(0)
	for _, v := range a.Int64Values() {
		acc += v
	}
	return acc
}
