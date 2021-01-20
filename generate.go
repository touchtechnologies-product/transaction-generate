package transaction_generate

func Generate(preString string, preId string, lastId string) string {

	return preString + lenZeroPreId(len(preId)) + preId + "-" +
		lenZeroLastId(len(lastId)) + lastId
}

func lenZeroLastId(len int) string {
	var zeroString string
	switch len {
	case 1:
		zeroString = "0000"
	case 2:
		zeroString = "000"
	case 3:
		zeroString = "00"
	case 4:
		zeroString = "0"
	default:
		zeroString = ""
	}
	return zeroString
}

func lenZeroPreId(len int) string {
	var zeroString string
	switch len {
	case 1:
		zeroString = "00"
	case 2:
		zeroString = "0"
	default:
		zeroString = ""
	}
	return zeroString
}
