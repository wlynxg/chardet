package chardet

type StateMachineModel struct {
	Name         string
	Language     string
	ClassTable   []byte
	ClassFactor  byte
	StateTable   []byte
	CharLenTable []byte
}

func HzSmModel() *StateMachineModel {
	hzCls := []byte{
		1, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
		0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
		0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
		0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
		0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
		0, 0, 0, 0, 0, 0, 0, 0, // 28 - 2f
		0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
		0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
		0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
		0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
		0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
		0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
		0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
		0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
		0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
		0, 0, 0, 4, 0, 5, 2, 0, // 78 - 7f
		1, 1, 1, 1, 1, 1, 1, 1, // 80 - 87
		1, 1, 1, 1, 1, 1, 1, 1, // 88 - 8f
		1, 1, 1, 1, 1, 1, 1, 1, // 90 - 97
		1, 1, 1, 1, 1, 1, 1, 1, // 98 - 9f
		1, 1, 1, 1, 1, 1, 1, 1, // a0 - a7
		1, 1, 1, 1, 1, 1, 1, 1, // a8 - af
		1, 1, 1, 1, 1, 1, 1, 1, // b0 - b7
		1, 1, 1, 1, 1, 1, 1, 1, // b8 - bf
		1, 1, 1, 1, 1, 1, 1, 1, // c0 - c7
		1, 1, 1, 1, 1, 1, 1, 1, // c8 - cf
		1, 1, 1, 1, 1, 1, 1, 1, // d0 - d7
		1, 1, 1, 1, 1, 1, 1, 1, // d8 - df
		1, 1, 1, 1, 1, 1, 1, 1, // e0 - e7
		1, 1, 1, 1, 1, 1, 1, 1, // e8 - ef
		1, 1, 1, 1, 1, 1, 1, 1, // f0 - f7
		1, 1, 1, 1, 1, 1, 1, 1, // f8 - ff
	}

	hzSt := []byte{
		StartMachineState, ErrorMachineState, 3, StartMachineState,
		StartMachineState, StartMachineState, ErrorMachineState, ErrorMachineState, // 00-07
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState,
		ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, // 08-0f
		ItsMeMachineState, ItsMeMachineState, ErrorMachineState, ErrorMachineState,
		StartMachineState, StartMachineState, 4, ErrorMachineState, // 10-17
		5, ErrorMachineState, 6, ErrorMachineState,
		5, 5, 4, ErrorMachineState, // 18-1f
		4, ErrorMachineState, 4, 4,
		4, ErrorMachineState, 4, ErrorMachineState, // 20-27
		4, ItsMeMachineState, StartMachineState, StartMachineState,
		StartMachineState, StartMachineState, StartMachineState, StartMachineState, // 28-2f
	}

	hzCharLenTable := []byte{0, 0, 0, 0, 0, 0}

	return &StateMachineModel{
		Name:         HzModelName,
		Language:     ChineseLanguage,
		ClassTable:   hzCls,
		ClassFactor:  6,
		StateTable:   hzSt,
		CharLenTable: hzCharLenTable,
	}
}

func Iso2022cnSmModel() *StateMachineModel {
	Iso2022cnCls := []byte{
		2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
		0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
		0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
		0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
		0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
		0, 3, 0, 0, 0, 0, 0, 0, // 28 - 2f
		0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
		0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
		0, 0, 0, 4, 0, 0, 0, 0, // 40 - 47
		0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
		0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
		0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
		0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
		0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
		0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
		0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
		2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
		2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
		2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
		2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
		2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
		2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
		2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
		2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
		2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
		2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
		2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
		2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
		2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
		2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
		2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
		2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
	}

	Iso2022cnSt := []byte{
		StartMachineState, 3, ErrorMachineState, StartMachineState,
		StartMachineState, StartMachineState, StartMachineState, StartMachineState, // 00-07
		StartMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState,
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, // 08-0f
		ErrorMachineState, ErrorMachineState, ItsMeMachineState, ItsMeMachineState,
		ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, // 10-17
		ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ErrorMachineState,
		ErrorMachineState, ErrorMachineState, 4, ErrorMachineState, // 18-1f
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ItsMeMachineState,
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, // 20-27
		5, 6, ErrorMachineState, ErrorMachineState,
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, // 28-2f
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ItsMeMachineState,
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, // 30-37
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState,
		ErrorMachineState, ItsMeMachineState, ErrorMachineState, StartMachineState, // 38-3f
	}

	Iso2022cnCharLenTable := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}

	return &StateMachineModel{
		Name:         Iso2022cnModelName,
		Language:     ChineseLanguage,
		ClassTable:   Iso2022cnCls,
		ClassFactor:  9,
		StateTable:   Iso2022cnSt,
		CharLenTable: Iso2022cnCharLenTable,
	}
}

func Iso2022jpSmModel() *StateMachineModel {
	Iso2022jpCls := []byte{
		2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
		0, 0, 0, 0, 0, 0, 2, 2, // 08 - 0f
		0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
		0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
		0, 0, 0, 0, 7, 0, 0, 0, // 20 - 27
		3, 0, 0, 0, 0, 0, 0, 0, // 28 - 2f
		0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
		0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
		6, 0, 4, 0, 8, 0, 0, 0, // 40 - 47
		0, 9, 5, 0, 0, 0, 0, 0, // 48 - 4f
		0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
		0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
		0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
		0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
		0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
		0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
		2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
		2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
		2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
		2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
		2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
		2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
		2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
		2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
		2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
		2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
		2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
		2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
		2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
		2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
		2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
		2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
	}

	Iso2022jpSt := []byte{
		StartMachineState, 3, ErrorMachineState, StartMachineState, StartMachineState, StartMachineState, StartMachineState, StartMachineState, // 00-07
		StartMachineState, StartMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, // 08-0f
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, // 10-17
		ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ErrorMachineState, ErrorMachineState, // 18-1f
		ErrorMachineState, 5, ErrorMachineState, ErrorMachineState, ErrorMachineState, 4, ErrorMachineState, ErrorMachineState, // 20-27
		ErrorMachineState, ErrorMachineState, ErrorMachineState, 6, ItsMeMachineState, ErrorMachineState, ItsMeMachineState, ErrorMachineState, // 28-2f
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ItsMeMachineState, ItsMeMachineState, // 30-37
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ItsMeMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, // 38-3f
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState, ItsMeMachineState, ErrorMachineState, StartMachineState, StartMachineState, // 40-47
	}

	Iso2022jpCharLenTable := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	return &StateMachineModel{
		Name:         Iso2022jpModelName,
		Language:     JapaneseLanguage,
		ClassTable:   Iso2022jpCls,
		ClassFactor:  10,
		StateTable:   Iso2022jpSt,
		CharLenTable: Iso2022jpCharLenTable,
	}
}

func Iso2022krSmModel() *StateMachineModel {
	Iso2022krCls := []byte{
		2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
		0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
		0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
		0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
		0, 0, 0, 0, 3, 0, 0, 0, // 20 - 27
		0, 4, 0, 0, 0, 0, 0, 0, // 28 - 2f
		0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
		0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
		0, 0, 0, 5, 0, 0, 0, 0, // 40 - 47
		0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
		0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
		0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
		0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
		0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
		0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
		0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
		2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
		2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
		2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
		2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
		2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
		2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
		2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
		2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
		2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
		2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
		2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
		2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
		2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
		2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
		2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
		2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
	}

	Iso2022krSt := []byte{
		StartMachineState, 3, ErrorMachineState, StartMachineState,
		StartMachineState, StartMachineState, ErrorMachineState, ErrorMachineState, // 00-07
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState,
		ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, ItsMeMachineState, // 08-0f
		ItsMeMachineState, ItsMeMachineState, ErrorMachineState, ErrorMachineState,
		ErrorMachineState, 4, ErrorMachineState, ErrorMachineState, // 10-17
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ErrorMachineState,
		5, ErrorMachineState, ErrorMachineState, ErrorMachineState, // 18-1f
		ErrorMachineState, ErrorMachineState, ErrorMachineState, ItsMeMachineState,
		StartMachineState, StartMachineState, StartMachineState, StartMachineState, // 20-27
	}

	Iso2022krCharLenTable := []byte{0, 0, 0, 0, 0, 0}

	return &StateMachineModel{
		Name:         "",
		Language:     "",
		ClassTable:   Iso2022krCls,
		ClassFactor:  0,
		StateTable:   Iso2022krSt,
		CharLenTable: Iso2022krCharLenTable,
	}
}
