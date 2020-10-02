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

// IntersectSortedUint16 returns index in X if X[i] exists in Y[]
func IntersectXUint16(x, y []uint16, ix []uint32) int {
	cnt := 0
	for i, j := 0, 0; i < len(x) && j < len(y); {
		if x[i] == y[j] {
			ix[cnt] = uint32(i)
			cnt++
			i++
			j++
		} else if x[i] > y[j] {
			j++
		} else {
			i++
		}
	}

	return cnt
}

// IntersectSortedUint16 returns index in X if X[i] exists in Y[]
func IntersectSortedUint16(x, y []uint16, ix, iy []uint32) int {
	cnt := 0
	for i, j := 0, 0; i < len(x) && j < len(y); {
		if x[i] == y[j] {
			ix[cnt] = uint32(i)
			iy[cnt] = uint32(j)
			cnt++
			i++
			j++
		} else if x[i] > y[j] {
			j++
		} else {
			i++
		}
	}

	return cnt
}

// IntersectSortedUint16 returns index in X if X[i] exists in Y[]
func IntersectSortedListUint16(x, y []uint16, ix, iy []uint32) int {
	cnt := 0
	for i, j := 0, 0; i < len(x) && j < len(y); {
		if x[i] == y[j] {
			ix[cnt] = uint32(i)
			iy[cnt] = uint32(j)
			cnt++
			for i < len(x)-1 && x[i] == x[i+1] {
				i++
				ix[cnt] = uint32(i)
				iy[cnt] = uint32(j)
				cnt++
			}
			i++
			j++
		} else if x[i] > y[j] {
			j++
		} else {
			i++
		}
	}

	return cnt
}
