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

func EqualInt64(data []int64, v int64, rs []uint32) uint {
	j := uint(0)
	for i, d := range data {
		if d == v {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func NotEqualInt64(data []int64, v int64, rs []uint32) uint {
	j := uint(0)
	for i, d := range data {
		if d != v {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func LessThanInt64(data []int64, v int64, rs []uint32) uint {
	j := uint(0)
	for i, d := range data {
		if d < v {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func LessEqualInt64(data []int64, v int64, rs []uint32) uint {
	j := uint(0)
	for i, d := range data {
		if d <= v {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func GreaterThanInt64(data []int64, v int64, rs []uint32) uint {
	j := uint(0)
	for i, d := range data {
		if d > v {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func GreaterEqualInt64(data []int64, v int64, rs []uint32) uint {
	j := uint(0)
	for i, d := range data {
		if d >= v {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}

func BetweenInt64(data []int64, x, y int64, rs []uint32) uint {
	j := uint(0)
	for i, d := range data {
		if x <= d && d <= y {
			rs[j] = uint32(i)
			j++
		}
	}
	return j
}
