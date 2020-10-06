package fuse

import "hash/fnv"

func FilepathHash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
