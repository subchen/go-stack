package gstack

func IIfString(expr bool, trueValue, falseValue string) string {
	if expr {
		return trueValue
	} else {
		return falseValue
	}
}

func IIfInt(expr bool, trueValue, falseValue int) int {
	if expr {
		return trueValue
	} else {
		return falseValue
	}
}

func IIfInt64(expr bool, trueValue, falseValue int64) int64 {
	if expr {
		return trueValue
	} else {
		return falseValue
	}
}

func IIfUint(expr bool, trueValue, falseValue uint) uint {
	if expr {
		return trueValue
	} else {
		return falseValue
	}
}

func IIfUint64(expr bool, trueValue, falseValue uint64) uint64 {
	if expr {
		return trueValue
	} else {
		return falseValue
	}
}

func IIfFloat32(expr bool, trueValue, falseValue float32) float32 {
	if expr {
		return trueValue
	} else {
		return falseValue
	}
}

func IIfFloat64(expr bool, trueValue, falseValue float64) float64 {
	if expr {
		return trueValue
	} else {
		return falseValue
	}
}
