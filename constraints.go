package ad

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

type AnyComparable interface {
	Comparable | comparable
}

type Comparable interface {
	CompareTo(c any) int
}
