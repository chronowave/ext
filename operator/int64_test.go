/*
 *  Copyright 2020 ChronoWave Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  Package parser declares an expression parser with support for macro
 *  expansion.
 */

package operator

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var (
	data []int64
	rs   []uint32
)

func TestMain(m *testing.M) {
	data = make([]int64, 1024*1024)
	for i := range data {
		data[i] = rand.Int63()
	}
	rs = make([]uint32, len(data))
	os.Exit(m.Run())
}
func TestLessInt64(t *testing.T) {
	a, b := byte(0), byte(252)
	fmt.Printf("%b %b", (a), (0 - b))
	type args struct {
		data []int64
		v    int64
		rs   []uint32
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LessThanInt64(tt.args.data, tt.args.v, tt.args.rs); got != tt.want {
				t.Errorf("LessInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLessInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessThanInt64(data, 10, rs)
	}
}

func BenchmarkBetweenInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BetweenInt64(data, 10, 30, rs)
	}
}
