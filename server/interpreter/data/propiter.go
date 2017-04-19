/* Copyright 2017 Google Inc.
 * https://github.com/NeilFraser/CodeCity
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package data

// PropIter is an iterator which will iterate over the (non-deleted)
// properties of an object and its prototypes.
//
// FIXME: skip non-enumerable properties
// FIXME: perhaps we should guarantee iteration order, as most
// browsers (and ES6) do?
type PropIter struct {
	value Value
	keys  []string
	seen  map[string]bool
}

// NewPropIter takes any Value and returns an PropIter for it.
func NewPropIter(v Value) *PropIter {
	return &PropIter{v, v.OwnPropertyKeys(), make(map[string]bool)}
}

// Next returns the next non-deleted, non-shadowed property key, with
// ok == true, or returns ok == false if there are no more property
// keys to iterate over.
func (iter *PropIter) Next() (string, bool) {
	var key string
	for {
		for len(iter.keys) > 0 {
			key = iter.keys[0]
			iter.keys = iter.keys[1:]
			if iter.value.HasOwnProperty(key) && !iter.seen[key] {
				iter.seen[key] = true
				return key, true
			}
		}
		iter.value = iter.value.Proto()
		if iter.value == nil {
			return "", false
		}
		iter.keys = iter.value.OwnPropertyKeys()
	}
}