package pbwrappers_ptypes

// field can't be named "Value" because of conflicts with the database/sql driver.Valuer interface

type DoubleValue struct {
	Valid  bool
	WValue float64
}

type FloatValue struct {
	Valid  bool
	WValue float32
}

type Int64Value struct {
	Valid  bool
	WValue int64
}

type UInt64Value struct {
	Valid  bool
	WValue uint64
}

type Int32Value struct {
	Valid  bool
	WValue int32
}

type UInt32Value struct {
	Valid  bool
	WValue uint32
}

type BoolValue struct {
	Valid  bool
	WValue bool
}

type StringValue struct {
	Valid  bool
	WValue string
}

type BytesValue struct {
	Valid  bool
	WValue []byte
}
