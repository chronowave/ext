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

import "math"

func InUint32(data, m []uint32, rs []uint32) uint {
	mm := make(map[uint32]bool, len(m))
	for _, v := range m {
		mm[v] = true
	}

	j := uint(0)
	for i, d := range data {
		if _, ok := mm[d]; ok {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func InInt64(data, m []int64, rs []uint32) uint {
	mm := make(map[int64]bool, len(m))
	for _, v := range m {
		mm[v] = true
	}

	j := uint(0)
	for i, d := range data {
		if _, ok := mm[d]; ok {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func InFloat64(data, m []float64, rs []uint32) uint {
	mm := make(map[uint64]bool, len(m))
	for _, v := range m {
		mm[math.Float64bits(v)] = true
	}

	j := uint(0)
	for i, d := range data {
		if _, ok := mm[math.Float64bits(d)]; ok {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}
