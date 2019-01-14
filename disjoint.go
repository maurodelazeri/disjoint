package disjoint

import (
	listPkg "container/list"
)

// UnionFindSets data structure.
type UnionFindSets map[interface{}]*listPkg.List

// New return a new UnionFindSets instance.
func New() UnionFindSets {
	return make(UnionFindSets)
}

// Union objects into the same set, including the existing set members of objects
func (ufs UnionFindSets) Union(objects ...interface{}) {
	for i := 1; i < len(objects); i++ {
		ufs.union(objects[i-1], objects[i])
	}
}

func (ufs UnionFindSets) union(a, b interface{}) {
	if a == b {
		return
	}
	la, lb := ufs[a], ufs[b]
	switch {
	case la == nil && lb == nil:
		ufs.init(a, b)
	case la != nil && lb == nil:
		la.PushBack(b)
		ufs[b] = la
	case lb != nil && la == nil:
		lb.PushBack(a)
		ufs[a] = lb
	case la != lb:
		ufs.concat(la, lb)
	default:
		// la == lb, so a and b is already in the same sets.
	}
}

func (ufs UnionFindSets) init(a, b interface{}) {
	list := listPkg.New()
	list.PushBack(a)
	list.PushBack(b)
	ufs[a] = list
	ufs[b] = list
}

func (ufs UnionFindSets) concat(la, lb *listPkg.List) {
	if la.Len() < lb.Len() {
		la, lb = lb, la
	}
	for b := lb.Front(); b != nil; b = b.Next() {
		la.PushBack(b.Value)
		ufs[b.Value] = la
	}
}

// Find all set members of objects, with no duplicates.
func (ufs UnionFindSets) Find(objects ...interface{}) (result []interface{}) {
	if len(objects) > 1 {
		objects = ufs.Distinct(objects...)
	}
	for i := 0; i < len(objects); i++ {
		result = append(result, ufs.find(objects[i])...)
	}
	return
}

func (ufs UnionFindSets) find(object interface{}) (result []interface{}) {
	list := ufs[object]
	if list == nil {
		return nil
	}
	for e := list.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value)
	}
	return
}

// Distinct removes the duplicate objects in any same set.
func (ufs UnionFindSets) Distinct(objects ...interface{}) []interface{} {
	result := []interface{}{objects[0]}
	for i := 1; i < len(objects); i++ {
		if !ufs.InSameSetSlice(objects[i], result) {
			result = append(result, objects[i])
		}
	}
	return result
}

// InSameSetSlice check if object is in the same set with any member of slice.
func (ufs UnionFindSets) InSameSetSlice(object interface{}, slice []interface{}) bool {
	for _, e := range slice {
		if ufs.InSameSet(object, e) {
			return true
		}
	}
	return false
}

// InSameSet check if a and b are in the same set
func (ufs UnionFindSets) InSameSet(a, b interface{}) bool {
	la := ufs[a]
	if la == nil {
		return false
	}
	lb := ufs[b]
	if lb == nil {
		return false
	}
	return la == lb
}

// RemoveSet remove the set object belongs to. returns the removed set members.
func (ufs UnionFindSets) RemoveSet(object interface{}) (result []interface{}) {
	list := ufs[object]
	if list == nil {
		return nil
	}
	for e := list.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value)
		delete(ufs, e.Value)
	}
	return
}
