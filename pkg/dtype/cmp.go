package dtype

// TypesEqType returns true if the types are equal to each other.
func TypesEqType(x, y Type) bool {
	return x.ID() == y.ID()
}

// TypesEqSize returns true if both types have the same storage size.
func TypesEqSize(x, y Type) bool {
	return x.Size1() == y.Size1()
}

// TypesCanConv returns true if converting `x` to `y` is
// a valid operation.
func TypesCanConv(x, y Type) bool {
	// all types can be converted to string, and string can
	// possibly be converted to all other types
	if y.id == STRT || x.id == STRT {
		return true
	}
	// all numeric types can be converted to each other
	if x.IsNumericType() && y.IsNumericType() {
		return true
	}
	// numeric types can be converted to bool (e.g., x != 0 ? true : false)
	if x.IsNumericType() && y.id == BOOLT {
		return true
	}
	// bool types can be converted into numeric types (e.g., x ? 1|1.0 : 0|0.0)
	if x.id == BOOLT && y.IsNumericType() {
		return true
	}

	return false
}
