package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type StateMachineModel struct {
	Name         string
	Language     string
	ClassTable   []byte
	ClassFactor  byte
	StateTable   []consts.MachineState
	CharLenTable []byte
}

func HzSmModel() StateMachineModel {
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

	hzSt := []consts.MachineState{
		consts.StartMachineState, consts.ErrorMachineState, 3, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, 4, consts.ErrorMachineState, // 10-17
		5, consts.ErrorMachineState, 6, consts.ErrorMachineState, 5, 5, 4, consts.ErrorMachineState, // 18-1f
		4, consts.ErrorMachineState, 4, 4, 4, consts.ErrorMachineState, 4, consts.ErrorMachineState, // 20-27
		4, consts.ItsMeMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 28-2f
	}

	hzCharLenTable := []byte{0, 0, 0, 0, 0, 0}

	return StateMachineModel{
		Name:         consts.HzGB2312,
		Language:     consts.Chinese,
		ClassTable:   hzCls,
		ClassFactor:  6,
		StateTable:   hzSt,
		CharLenTable: hzCharLenTable,
	}
}

func Iso2022cnSmModel() StateMachineModel {
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

	Iso2022cnSt := []consts.MachineState{
		consts.StartMachineState, 3, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 00-07
		consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 08-0f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 10-17
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 4, consts.ErrorMachineState, // 18-1f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 20-27
		5, 6, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 28-2f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 30-37
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.StartMachineState, // 38-3f
	}

	Iso2022cnCharLenTable := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}

	return StateMachineModel{
		Name:         consts.ISO2022CN,
		Language:     consts.Chinese,
		ClassTable:   Iso2022cnCls,
		ClassFactor:  9,
		StateTable:   Iso2022cnSt,
		CharLenTable: Iso2022cnCharLenTable,
	}
}

func Iso2022jpSmModel() StateMachineModel {
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

	Iso2022jpSt := []consts.MachineState{
		consts.StartMachineState, 3, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 00-07
		consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 08-0f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 10-17
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 18-1f
		consts.ErrorMachineState, 5, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 4, consts.ErrorMachineState, consts.ErrorMachineState, // 20-27
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 6, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, // 28-2f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 30-37
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 38-3f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, // 40-47
	}

	Iso2022jpCharLenTable := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	return StateMachineModel{
		Name:         consts.ISO2022JP,
		Language:     consts.Japanese,
		ClassTable:   Iso2022jpCls,
		ClassFactor:  10,
		StateTable:   Iso2022jpSt,
		CharLenTable: Iso2022jpCharLenTable,
	}
}

func Iso2022krSmModel() StateMachineModel {
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

	Iso2022krSt := []consts.MachineState{
		consts.StartMachineState, 3, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 4, consts.ErrorMachineState, consts.ErrorMachineState, // 10-17
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 5, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 18-1f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 20-27
	}

	Iso2022krCharLenTable := []byte{0, 0, 0, 0, 0, 0}

	return StateMachineModel{
		Name:         consts.ISO2022KR,
		Language:     consts.Korean,
		ClassTable:   Iso2022krCls,
		ClassFactor:  6,
		StateTable:   Iso2022krSt,
		CharLenTable: Iso2022krCharLenTable,
	}
}

func UTF8SmModel() StateMachineModel {
	Utf8Cls := []byte{
		1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07  //allow 0x00 as a legal value
		1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
		1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
		1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
		1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
		1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
		1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
		1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
		1, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
		1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
		1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
		1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
		1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
		1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
		1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
		1, 1, 1, 1, 1, 1, 1, 1, // 78 - 7f
		2, 2, 2, 2, 3, 3, 3, 3, // 80 - 87
		4, 4, 4, 4, 4, 4, 4, 4, // 88 - 8f
		4, 4, 4, 4, 4, 4, 4, 4, // 90 - 97
		4, 4, 4, 4, 4, 4, 4, 4, // 98 - 9f
		5, 5, 5, 5, 5, 5, 5, 5, // a0 - a7
		5, 5, 5, 5, 5, 5, 5, 5, // a8 - af
		5, 5, 5, 5, 5, 5, 5, 5, // b0 - b7
		5, 5, 5, 5, 5, 5, 5, 5, // b8 - bf
		0, 0, 6, 6, 6, 6, 6, 6, // c0 - c7
		6, 6, 6, 6, 6, 6, 6, 6, // c8 - cf
		6, 6, 6, 6, 6, 6, 6, 6, // d0 - d7
		6, 6, 6, 6, 6, 6, 6, 6, // d8 - df
		7, 8, 8, 8, 8, 8, 8, 8, // e0 - e7
		8, 8, 8, 8, 8, 9, 8, 8, // e8 - ef
		10, 11, 11, 11, 11, 11, 11, 11, // f0 - f7
		12, 13, 13, 13, 14, 15, 0, 0, // f8 - ff
	}

	Utf8St := []consts.MachineState{
		consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 12, 10, // 00-07
		9, 11, 8, 7, 6, 5, 4, 3, // 08-0f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 10-17
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 18-1f
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 20-27
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 28-2f
		consts.ErrorMachineState, consts.ErrorMachineState, 5, 5, 5, 5, consts.ErrorMachineState, consts.ErrorMachineState, // 30-37
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 38-3f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 5, 5, 5, consts.ErrorMachineState, consts.ErrorMachineState, // 40-47
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 48-4f
		consts.ErrorMachineState, consts.ErrorMachineState, 7, 7, 7, 7, consts.ErrorMachineState, consts.ErrorMachineState, // 50-57
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 58-5f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 7, 7, consts.ErrorMachineState, consts.ErrorMachineState, // 60-67
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 68-6f
		consts.ErrorMachineState, consts.ErrorMachineState, 9, 9, 9, 9, consts.ErrorMachineState, consts.ErrorMachineState, // 70-77
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 78-7f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 9, consts.ErrorMachineState, consts.ErrorMachineState, // 80-87
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 88-8f
		consts.ErrorMachineState, consts.ErrorMachineState, 12, 12, 12, 12, consts.ErrorMachineState, consts.ErrorMachineState, // 90-97
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 98-9f
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 12, consts.ErrorMachineState, consts.ErrorMachineState, // a0-a7
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // a8-af
		consts.ErrorMachineState, consts.ErrorMachineState, 12, 12, 12, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // b0-b7
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // b8-bf
		consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // c0-c7
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // c8-cf
	}

	Utf8CharLenTable := []byte{0, 1, 0, 0, 0, 0, 2, 3, 3, 3, 4, 4, 5, 5, 6, 6}
	return StateMachineModel{
		Name:         consts.UTF8,
		Language:     "",
		ClassTable:   Utf8Cls,
		ClassFactor:  16,
		StateTable:   Utf8St,
		CharLenTable: Utf8CharLenTable,
	}
}

func Ucs2LeSmModel() *StateMachineModel {
	Ucs2leCls := []byte{
		0, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
		0, 0, 1, 0, 0, 2, 0, 0, // 08 - 0f
		0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
		0, 0, 0, 3, 0, 0, 0, 0, // 18 - 1f
		0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
		0, 3, 3, 3, 3, 3, 0, 0, // 28 - 2f
		0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
		0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
		0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
		0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
		0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
		0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
		0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
		0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
		0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
		0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
		0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
		0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
		0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
		0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
		0, 0, 0, 0, 0, 0, 0, 0, // a0 - a7
		0, 0, 0, 0, 0, 0, 0, 0, // a8 - af
		0, 0, 0, 0, 0, 0, 0, 0, // b0 - b7
		0, 0, 0, 0, 0, 0, 0, 0, // b8 - bf
		0, 0, 0, 0, 0, 0, 0, 0, // c0 - c7
		0, 0, 0, 0, 0, 0, 0, 0, // c8 - cf
		0, 0, 0, 0, 0, 0, 0, 0, // d0 - d7
		0, 0, 0, 0, 0, 0, 0, 0, // d8 - df
		0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
		0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
		0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
		0, 0, 0, 0, 0, 0, 4, 5, // f8 - ff
	}

	Ucs2leSt := []consts.MachineState{
		6, 6, 7, 6, 4, 3, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, 5, 5, 5, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, // 10-17
		5, 5, 5, consts.ErrorMachineState, 5, consts.ErrorMachineState, 6, 6, // 18-1f
		7, 6, 8, 8, 5, 5, 5, consts.ErrorMachineState, // 20-27
		5, 5, 5, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 5, 5, // 28-2f
		5, 5, 5, consts.ErrorMachineState, 5, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, // 30-37
	}

	Ucs2leCharLenTable := []byte{2, 2, 2, 2, 2, 2}
	return &StateMachineModel{
		Name:         consts.UTF16Le,
		Language:     "",
		ClassTable:   Ucs2leCls,
		ClassFactor:  6,
		StateTable:   Ucs2leSt,
		CharLenTable: Ucs2leCharLenTable,
	}
}

func Ucs2BeSmModel() *StateMachineModel {
	Ucs2beCls := []byte{
		0, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
		0, 0, 1, 0, 0, 2, 0, 0, // 08 - 0f
		0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
		0, 0, 0, 3, 0, 0, 0, 0, // 18 - 1f
		0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
		0, 3, 3, 3, 3, 3, 0, 0, // 28 - 2f
		0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
		0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
		0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
		0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
		0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
		0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
		0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
		0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
		0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
		0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
		0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
		0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
		0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
		0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
		0, 0, 0, 0, 0, 0, 0, 0, // a0 - a7
		0, 0, 0, 0, 0, 0, 0, 0, // a8 - af
		0, 0, 0, 0, 0, 0, 0, 0, // b0 - b7
		0, 0, 0, 0, 0, 0, 0, 0, // b8 - bf
		0, 0, 0, 0, 0, 0, 0, 0, // c0 - c7
		0, 0, 0, 0, 0, 0, 0, 0, // c8 - cf
		0, 0, 0, 0, 0, 0, 0, 0, // d0 - d7
		0, 0, 0, 0, 0, 0, 0, 0, // d8 - df
		0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
		0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
		0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
		0, 0, 0, 0, 0, 0, 4, 5, // f8 - ff
	}

	Ucs2beSt := []consts.MachineState{
		5, 7, 7, consts.ErrorMachineState, 4, 3, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, 6, 6, 6, 6, consts.ErrorMachineState, consts.ErrorMachineState, // 10-17
		6, 6, 6, 6, 6, consts.ItsMeMachineState, 6, 6, // 18-1f
		6, 6, 6, 6, 5, 7, 7, consts.ErrorMachineState, // 20-27
		5, 8, 6, 6, consts.ErrorMachineState, 6, 6, 6, // 28-2f
		6, 6, 6, 6, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, // 30-37
	}

	Ucs2beCharLenTable := []byte{2, 2, 2, 0, 2, 2}
	return &StateMachineModel{
		Name:         consts.UTF16Be,
		Language:     "",
		ClassTable:   Ucs2beCls,
		ClassFactor:  6,
		StateTable:   Ucs2beSt,
		CharLenTable: Ucs2beCharLenTable,
	}
}

func SjisSmModel() StateMachineModel {
	SjisCls := []byte{
		1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
		1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
		1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
		1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
		1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
		1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
		1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
		1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
		2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
		2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
		2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
		2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
		2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
		2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
		2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
		2, 2, 2, 2, 2, 2, 2, 1, // 78 - 7f
		3, 3, 3, 3, 3, 2, 2, 3, // 80 - 87
		3, 3, 3, 3, 3, 3, 3, 3, // 88 - 8f
		3, 3, 3, 3, 3, 3, 3, 3, // 90 - 97
		3, 3, 3, 3, 3, 3, 3, 3, // 98 - 9f
		// 0xa0 is illegal in sjis encoding, but some pages does
		// contain such byte. We need to be more error forgiven.
		2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
		2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
		2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
		2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
		2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
		2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
		2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
		2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
		3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
		3, 3, 3, 3, 3, 4, 4, 4, // e8 - ef
		3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
		3, 3, 3, 3, 3, 0, 0, 0, // f8 - ff
	}

	SjisSt := []consts.MachineState{
		consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, 3, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 10-17
	}

	SjisCharLenTable := []byte{0, 1, 1, 2, 0, 0}
	return StateMachineModel{
		Name:         consts.ShiftJis,
		Language:     "",
		ClassTable:   SjisCls,
		ClassFactor:  6,
		StateTable:   SjisSt,
		CharLenTable: SjisCharLenTable,
	}
}

func GB2312SmModel() StateMachineModel {
	Gb2312Cls := []byte{
		1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
		1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
		1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
		1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
		1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
		1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
		3, 3, 3, 3, 3, 3, 3, 3, // 30 - 37
		3, 3, 1, 1, 1, 1, 1, 1, // 38 - 3f
		2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
		2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
		2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
		2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
		2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
		2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
		2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
		2, 2, 2, 2, 2, 2, 2, 4, // 78 - 7f
		5, 6, 6, 6, 6, 6, 6, 6, // 80 - 87
		6, 6, 6, 6, 6, 6, 6, 6, // 88 - 8f
		6, 6, 6, 6, 6, 6, 6, 6, // 90 - 97
		6, 6, 6, 6, 6, 6, 6, 6, // 98 - 9f
		6, 6, 6, 6, 6, 6, 6, 6, // a0 - a7
		6, 6, 6, 6, 6, 6, 6, 6, // a8 - af
		6, 6, 6, 6, 6, 6, 6, 6, // b0 - b7
		6, 6, 6, 6, 6, 6, 6, 6, // b8 - bf
		6, 6, 6, 6, 6, 6, 6, 6, // c0 - c7
		6, 6, 6, 6, 6, 6, 6, 6, // c8 - cf
		6, 6, 6, 6, 6, 6, 6, 6, // d0 - d7
		6, 6, 6, 6, 6, 6, 6, 6, // d8 - df
		6, 6, 6, 6, 6, 6, 6, 6, // e0 - e7
		6, 6, 6, 6, 6, 6, 6, 6, // e8 - ef
		6, 6, 6, 6, 6, 6, 6, 6, // f0 - f7
		6, 6, 6, 6, 6, 6, 6, 0, // f8 - ff
	}

	Gb2312St := []consts.MachineState{consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, 3, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, // 10-17
		4, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 18-1f
		consts.ErrorMachineState, consts.ErrorMachineState, 5, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, // 20-27
		consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 28-2f
	}

	// To be accurate, the length of class 6 can be either 2 or 4.
	// But it is not necessary to discriminate between the two since
	// it is used for frequency analysis only, and we are validating
	// each code range there as well. So it is safe to set it to be
	// 2 here.
	Gb2312CharLenTable := []byte{0, 1, 1, 1, 1, 1, 2}
	return StateMachineModel{
		Name:         consts.GB2312,
		Language:     "",
		ClassTable:   Gb2312Cls,
		ClassFactor:  7,
		StateTable:   Gb2312St,
		CharLenTable: Gb2312CharLenTable,
	}
}

func EucTwSmModel() StateMachineModel {
	EucTwCls := []byte{
		2, 2, 2, 2, 2, 2, 2, 2, // 00 - 07
		2, 2, 2, 2, 2, 2, 0, 0, // 08 - 0f
		2, 2, 2, 2, 2, 2, 2, 2, // 10 - 17
		2, 2, 2, 0, 2, 2, 2, 2, // 18 - 1f
		2, 2, 2, 2, 2, 2, 2, 2, // 20 - 27
		2, 2, 2, 2, 2, 2, 2, 2, // 28 - 2f
		2, 2, 2, 2, 2, 2, 2, 2, // 30 - 37
		2, 2, 2, 2, 2, 2, 2, 2, // 38 - 3f
		2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
		2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
		2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
		2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
		2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
		2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
		2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
		2, 2, 2, 2, 2, 2, 2, 2, // 78 - 7f
		0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
		0, 0, 0, 0, 0, 0, 6, 0, // 88 - 8f
		0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
		0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
		0, 3, 4, 4, 4, 4, 4, 4, // a0 - a7
		5, 5, 1, 1, 1, 1, 1, 1, // a8 - af
		1, 1, 1, 1, 1, 1, 1, 1, // b0 - b7
		1, 1, 1, 1, 1, 1, 1, 1, // b8 - bf
		1, 1, 3, 1, 3, 3, 3, 3, // c0 - c7
		3, 3, 3, 3, 3, 3, 3, 3, // c8 - cf
		3, 3, 3, 3, 3, 3, 3, 3, // d0 - d7
		3, 3, 3, 3, 3, 3, 3, 3, // d8 - df
		3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
		3, 3, 3, 3, 3, 3, 3, 3, // e8 - ef
		3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
		3, 3, 3, 3, 3, 3, 3, 0, // f8 - ff
	}

	EucTwSt := []consts.MachineState{consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, 3, 3, 3, 4, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, // 10-17
		consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 18-1f
		5, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, // 20-27
		consts.StartMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 28-2f
	}

	EucTwCharLenTable := []byte{0, 0, 1, 2, 2, 2, 3}
	return StateMachineModel{
		Name:         consts.EucTw,
		Language:     "",
		ClassTable:   EucTwCls,
		ClassFactor:  7,
		StateTable:   EucTwSt,
		CharLenTable: EucTwCharLenTable,
	}
}

func EucKrSmModel() StateMachineModel {
	EuckrCls := []byte{
		1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
		1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
		1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
		1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
		1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
		1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
		1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
		1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
		1, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
		1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
		1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
		1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
		1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
		1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
		1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
		1, 1, 1, 1, 1, 1, 1, 1, // 78 - 7f
		0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
		0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
		0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
		0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
		0, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
		2, 2, 2, 2, 2, 3, 3, 3, // a8 - af
		2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
		2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
		2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
		2, 3, 2, 2, 2, 2, 2, 2, // c8 - cf
		2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
		2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
		2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
		2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
		2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
		2, 2, 2, 2, 2, 2, 2, 0, // f8 - ff
	}

	EuckrSt := []consts.MachineState{
		consts.ErrorMachineState, consts.StartMachineState, 3, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, // 08-0f
	}

	EuckrCharLenTable := []byte{0, 1, 2, 0}
	return StateMachineModel{
		Name:         consts.EucKr,
		Language:     "",
		ClassTable:   EuckrCls,
		ClassFactor:  4,
		StateTable:   EuckrSt,
		CharLenTable: EuckrCharLenTable,
	}
}

func EucJpSmModel() StateMachineModel {
	EucJpCls := []byte{
		4, 4, 4, 4, 4, 4, 4, 4, // 00 - 07
		4, 4, 4, 4, 4, 4, 5, 5, // 08 - 0f
		4, 4, 4, 4, 4, 4, 4, 4, // 10 - 17
		4, 4, 4, 5, 4, 4, 4, 4, // 18 - 1f
		4, 4, 4, 4, 4, 4, 4, 4, // 20 - 27
		4, 4, 4, 4, 4, 4, 4, 4, // 28 - 2f
		4, 4, 4, 4, 4, 4, 4, 4, // 30 - 37
		4, 4, 4, 4, 4, 4, 4, 4, // 38 - 3f
		4, 4, 4, 4, 4, 4, 4, 4, // 40 - 47
		4, 4, 4, 4, 4, 4, 4, 4, // 48 - 4f
		4, 4, 4, 4, 4, 4, 4, 4, // 50 - 57
		4, 4, 4, 4, 4, 4, 4, 4, // 58 - 5f
		4, 4, 4, 4, 4, 4, 4, 4, // 60 - 67
		4, 4, 4, 4, 4, 4, 4, 4, // 68 - 6f
		4, 4, 4, 4, 4, 4, 4, 4, // 70 - 77
		4, 4, 4, 4, 4, 4, 4, 4, // 78 - 7f
		5, 5, 5, 5, 5, 5, 5, 5, // 80 - 87
		5, 5, 5, 5, 5, 5, 1, 3, // 88 - 8f
		5, 5, 5, 5, 5, 5, 5, 5, // 90 - 97
		5, 5, 5, 5, 5, 5, 5, 5, // 98 - 9f
		5, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
		2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
		2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
		2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
		2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
		2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
		2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
		2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
		0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
		0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
		0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
		0, 0, 0, 0, 0, 0, 0, 5, // f8 - ff
	}

	EucJpSt := []consts.MachineState{3, 4, 3, 5, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // 08-0f
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 10-17
		consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 3, consts.ErrorMachineState, // 18-1f
		3, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 20-27
	}

	EucJpCharLenTable := []byte{2, 2, 2, 3, 1, 0}
	return StateMachineModel{
		Name:         consts.EucJp,
		Language:     "",
		ClassTable:   EucJpCls,
		ClassFactor:  6,
		StateTable:   EucJpSt,
		CharLenTable: EucJpCharLenTable,
	}
}

func CP949SmModel() StateMachineModel {
	Cp949Cls := []byte{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, // 00 - 0f
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, // 10 - 1f
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 20 - 2f
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 30 - 3f
		1, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, // 40 - 4f
		4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, // 50 - 5f
		1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, // 60 - 6f
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, // 70 - 7f
		0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, // 80 - 8f
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, // 90 - 9f
		6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, // a0 - af
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, // b0 - bf
		7, 7, 7, 7, 7, 7, 9, 2, 2, 3, 2, 2, 2, 2, 2, 2, // c0 - cf
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // d0 - df
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // e0 - ef
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, // f0 - ff
	}

	Cp949St := []consts.MachineState{
		consts.ErrorMachineState, consts.StartMachineState, 3, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, 4, 5, consts.ErrorMachineState, 6, // StartMachineState
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // ErrorMachineState
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // ItsMeMachineState
		consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 3
		consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 4
		consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 5
		consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 6
	}

	Cp949CharLenTable := []byte{0, 1, 2, 0, 1, 1, 2, 2, 0, 2}
	return StateMachineModel{
		Name:         consts.CP949,
		Language:     "",
		ClassTable:   Cp949Cls,
		ClassFactor:  10,
		StateTable:   Cp949St,
		CharLenTable: Cp949CharLenTable,
	}
}

func Big5SmModel() StateMachineModel {
	Big5Cls := []byte{
		1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07    //allow 0x00 as legal value
		1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
		1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
		1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
		1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
		1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
		1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
		1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
		2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
		2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
		2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
		2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
		2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
		2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
		2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
		2, 2, 2, 2, 2, 2, 2, 1, // 78 - 7f
		4, 4, 4, 4, 4, 4, 4, 4, // 80 - 87
		4, 4, 4, 4, 4, 4, 4, 4, // 88 - 8f
		4, 4, 4, 4, 4, 4, 4, 4, // 90 - 97
		4, 4, 4, 4, 4, 4, 4, 4, // 98 - 9f
		4, 3, 3, 3, 3, 3, 3, 3, // a0 - a7
		3, 3, 3, 3, 3, 3, 3, 3, // a8 - af
		3, 3, 3, 3, 3, 3, 3, 3, // b0 - b7
		3, 3, 3, 3, 3, 3, 3, 3, // b8 - bf
		3, 3, 3, 3, 3, 3, 3, 3, // c0 - c7
		3, 3, 3, 3, 3, 3, 3, 3, // c8 - cf
		3, 3, 3, 3, 3, 3, 3, 3, // d0 - d7
		3, 3, 3, 3, 3, 3, 3, 3, // d8 - df
		3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
		3, 3, 3, 3, 3, 3, 3, 3, // e8 - ef
		3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
		3, 3, 3, 3, 3, 3, 3, 0, // f8 - ff
	}

	Big5St := []consts.MachineState{consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, 3, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // 00-07
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ErrorMachineState, // 08-0f
		consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 10-17
	}

	Big5CharLenTable := []byte{0, 1, 1, 2, 0}
	return StateMachineModel{
		Name:         consts.Big5,
		Language:     "",
		ClassTable:   Big5Cls,
		ClassFactor:  5,
		StateTable:   Big5St,
		CharLenTable: Big5CharLenTable,
	}
}

func JohabSmModel() StateMachineModel {
	JohabCls := []byte{
		4, 4, 4, 4, 4, 4, 4, 4, // 00 - 07
		4, 4, 4, 4, 4, 4, 0, 0, // 08 - 0f
		4, 4, 4, 4, 4, 4, 4, 4, // 10 - 17
		4, 4, 4, 0, 4, 4, 4, 4, // 18 - 1f
		4, 4, 4, 4, 4, 4, 4, 4, // 20 - 27
		4, 4, 4, 4, 4, 4, 4, 4, // 28 - 2f
		4, 3, 3, 3, 3, 3, 3, 3, // 30 - 37
		3, 3, 3, 3, 3, 3, 3, 3, // 38 - 3f
		3, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
		1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
		1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
		1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
		1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
		1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
		1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
		1, 1, 1, 1, 1, 1, 1, 2, // 78 - 7f
		6, 6, 6, 6, 8, 8, 8, 8, // 80 - 87
		8, 8, 8, 8, 8, 8, 8, 8, // 88 - 8f
		8, 7, 7, 7, 7, 7, 7, 7, // 90 - 97
		7, 7, 7, 7, 7, 7, 7, 7, // 98 - 9f
		7, 7, 7, 7, 7, 7, 7, 7, // a0 - a7
		7, 7, 7, 7, 7, 7, 7, 7, // a8 - af
		7, 7, 7, 7, 7, 7, 7, 7, // b0 - b7
		7, 7, 7, 7, 7, 7, 7, 7, // b8 - bf
		7, 7, 7, 7, 7, 7, 7, 7, // c0 - c7
		7, 7, 7, 7, 7, 7, 7, 7, // c8 - cf
		7, 7, 7, 7, 5, 5, 5, 5, // d0 - d7
		5, 9, 9, 9, 9, 9, 9, 5, // d8 - df
		9, 9, 9, 9, 9, 9, 9, 9, // e0 - e7
		9, 9, 9, 9, 9, 9, 9, 9, // e8 - ef
		9, 9, 9, 9, 9, 9, 9, 9, // f0 - f7
		9, 9, 5, 5, 5, 5, 5, 0, // f8 - ff
	}

	JohabSt := []consts.MachineState{
		// cls = 0                   1                   2                   3                   4                   5                   6                   7                   8                   9
		consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, 3, 3, 4, // consts.StartMachineState
		consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, consts.ItsMeMachineState, // consts.ItsMeMachineState
		consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.ErrorMachineState, // consts.ErrorMachineState
		consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, consts.StartMachineState, // 3
		consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.StartMachineState, consts.ErrorMachineState, consts.StartMachineState, // 4
	}

	JohabCharLenTable := []byte{0, 1, 1, 1, 1, 0, 0, 2, 2, 2}

	return StateMachineModel{
		Name:         consts.Johab,
		Language:     "",
		ClassTable:   JohabCls,
		ClassFactor:  10,
		StateTable:   JohabSt,
		CharLenTable: JohabCharLenTable,
	}
}
