package collections

import "GoStructAlgorithm/queue/base"

type Iterator interface {
    HashNext() bool
    Next() base.Object
}
