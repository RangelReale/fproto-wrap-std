package pbwrappers_ptypes

type DoubleValue struct {
	Valid bool
	Value float64
}

type FloatValue struct {
	Valid bool
	Value float32
}

type Int64Value struct {
	Valid bool
	Value int64
}

type UInt64Value struct {
	Valid bool
	Value uint64
}

type Int32Value struct {
	Valid bool
	Value int32
}

type UInt32Value struct {
	Valid bool
	Value uint32
}

type BoolValue struct {
	Valid bool
	Value bool
}

type StringValue struct {
	Valid bool
	Value string
}

type BytesValue struct {
	Valid bool
	Value []byte
}
