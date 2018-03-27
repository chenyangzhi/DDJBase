// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package plural

// CLDRVersion is the CLDR version from which the tables in this package are derived.
const CLDRVersion = "32"

var ordinalRules = []pluralCheck{ // 64 elements
	0:  {cat: 0x2f, setID: 0x4},
	1:  {cat: 0x3a, setID: 0x5},
	2:  {cat: 0x22, setID: 0x1},
	3:  {cat: 0x22, setID: 0x6},
	4:  {cat: 0x22, setID: 0x7},
	5:  {cat: 0x2f, setID: 0x8},
	6:  {cat: 0x3c, setID: 0x9},
	7:  {cat: 0x2f, setID: 0xa},
	8:  {cat: 0x3c, setID: 0xb},
	9:  {cat: 0x2c, setID: 0xc},
	10: {cat: 0x24, setID: 0xd},
	11: {cat: 0x2d, setID: 0xe},
	12: {cat: 0x2d, setID: 0xf},
	13: {cat: 0x2f, setID: 0x10},
	14: {cat: 0x35, setID: 0x3},
	15: {cat: 0xc5, setID: 0x11},
	16: {cat: 0x2, setID: 0x1},
	17: {cat: 0x5, setID: 0x3},
	18: {cat: 0xd, setID: 0x12},
	19: {cat: 0x22, setID: 0x1},
	20: {cat: 0x2f, setID: 0x13},
	21: {cat: 0x3d, setID: 0x14},
	22: {cat: 0x2f, setID: 0x15},
	23: {cat: 0x3a, setID: 0x16},
	24: {cat: 0x2f, setID: 0x17},
	25: {cat: 0x3b, setID: 0x18},
	26: {cat: 0x2f, setID: 0xa},
	27: {cat: 0x3c, setID: 0xb},
	28: {cat: 0x22, setID: 0x1},
	29: {cat: 0x23, setID: 0x19},
	30: {cat: 0x24, setID: 0x1a},
	31: {cat: 0x22, setID: 0x1b},
	32: {cat: 0x23, setID: 0x2},
	33: {cat: 0x24, setID: 0x1a},
	34: {cat: 0xf, setID: 0x15},
	35: {cat: 0x1a, setID: 0x16},
	36: {cat: 0xf, setID: 0x17},
	37: {cat: 0x1b, setID: 0x18},
	38: {cat: 0xf, setID: 0x1c},
	39: {cat: 0x1d, setID: 0x1d},
	40: {cat: 0xa, setID: 0x1e},
	41: {cat: 0xa, setID: 0x1f},
	42: {cat: 0xc, setID: 0x20},
	43: {cat: 0xe4, setID: 0x0},
	44: {cat: 0x5, setID: 0x3},
	45: {cat: 0xd, setID: 0xe},
	46: {cat: 0xd, setID: 0x21},
	47: {cat: 0x22, setID: 0x1},
	48: {cat: 0x23, setID: 0x19},
	49: {cat: 0x24, setID: 0x1a},
	50: {cat: 0x25, setID: 0x22},
	51: {cat: 0x22, setID: 0x23},
	52: {cat: 0x23, setID: 0x19},
	53: {cat: 0x24, setID: 0x1a},
	54: {cat: 0x25, setID: 0x22},
	55: {cat: 0x22, setID: 0x24},
	56: {cat: 0x23, setID: 0x19},
	57: {cat: 0x24, setID: 0x1a},
	58: {cat: 0x25, setID: 0x22},
	59: {cat: 0x21, setID: 0x25},
	60: {cat: 0x22, setID: 0x1},
	61: {cat: 0x23, setID: 0x2},
	62: {cat: 0x24, setID: 0x26},
	63: {cat: 0x25, setID: 0x27},
} // Size: 152 bytes

var ordinalIndex = []uint8{ // 22 elements
	0x00, 0x00, 0x02, 0x03, 0x04, 0x05, 0x07, 0x09,
	0x0b, 0x0f, 0x10, 0x13, 0x16, 0x1c, 0x1f, 0x22,
	0x28, 0x2f, 0x33, 0x37, 0x3b, 0x40,
} // Size: 46 bytes

var ordinalLangToIndex = []uint8{ // 775 elements
	// Entry 0 - 3F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x12, 0x12, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00,
	0x00, 0x10, 0x10, 0x00, 0x00, 0x05, 0x05, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 40 - 7F
	0x12, 0x12, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e,
	0x0e, 0x0e, 0x0e, 0x0e, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x14, 0x14, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 80 - BF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	// Entry C0 - FF
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c,
	0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 100 - 13F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02,
	0x00, 0x00, 0x00, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	// Entry 140 - 17F
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x11, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11,
	0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x03,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 180 - 1BF
	0x00, 0x00, 0x00, 0x00, 0x09, 0x09, 0x09, 0x09,
	0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x0a, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 1C0 - 1FF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x0f, 0x0f, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x0d, 0x0d, 0x02, 0x02, 0x02,
	0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 200 - 23F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x04, 0x04, 0x04, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x13, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 240 - 27F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 280 - 2BF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x0b, 0x0b, 0x0b, 0x0b, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01,
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x07, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 2C0 - 2FF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x06, 0x06, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 300 - 33F
	0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x0c,
} // Size: 799 bytes

var ordinalInclusionMasks = []uint64{ // 100 elements
	// Entry 0 - 1F
	0x0000002000010009, 0x00000018482000d3, 0x0000000042840195, 0x000000410a040581,
	0x00000041040c0081, 0x0000009840040041, 0x0000008400045001, 0x0000003850040001,
	0x0000003850060001, 0x0000003800049001, 0x0000000800052001, 0x0000000040660031,
	0x0000000041840331, 0x0000000100040f01, 0x00000001001c0001, 0x0000000040040001,
	0x0000000000045001, 0x0000000070040001, 0x0000000070040001, 0x0000000000049001,
	0x0000000080050001, 0x0000000040200011, 0x0000000040800111, 0x0000000100000501,
	0x0000000100080001, 0x0000000040000001, 0x0000000000005001, 0x0000000050000001,
	0x0000000050000001, 0x0000000000009001, 0x0000000000010001, 0x0000000040200011,
	// Entry 20 - 3F
	0x0000000040800111, 0x0000000100000501, 0x0000000100080001, 0x0000000040000001,
	0x0000000000005001, 0x0000000050000001, 0x0000000050000001, 0x0000000000009001,
	0x0000000200050001, 0x0000000040200011, 0x0000000040800111, 0x0000000100000501,
	0x0000000100080001, 0x0000000040000001, 0x0000000000005001, 0x0000000050000001,
	0x0000000050000001, 0x0000000000009001, 0x0000000080010001, 0x0000000040200011,
	0x0000000040800111, 0x0000000100000501, 0x0000000100080001, 0x0000000040000001,
	0x0000000000005001, 0x0000000050000001, 0x0000000050000001, 0x0000000000009001,
	0x0000000200050001, 0x0000000040200011, 0x0000000040800111, 0x0000000100000501,
	// Entry 40 - 5F
	0x0000000100080001, 0x0000000040000001, 0x0000000000005001, 0x0000000050000001,
	0x0000000050000001, 0x0000000000009001, 0x0000000080010001, 0x0000000040200011,
	0x0000000040800111, 0x0000000100000501, 0x0000000100080001, 0x0000000040000001,
	0x0000000000005001, 0x0000000050000001, 0x0000000050000001, 0x0000000000009001,
	0x0000000080070001, 0x0000000040200011, 0x0000000040800111, 0x0000000100000501,
	0x0000000100080001, 0x0000000040000001, 0x0000000000005001, 0x0000000050000001,
	0x0000000050000001, 0x0000000000009001, 0x0000000200010001, 0x0000000040200011,
	0x0000000040800111, 0x0000000100000501, 0x0000000100080001, 0x0000000040000001,
	// Entry 60 - 7F
	0x0000000000005001, 0x0000000050000001, 0x0000000050000001, 0x0000000000009001,
} // Size: 824 bytes

// Slots used for ordinal: 40 of 0xFF rules; 16 of 0xFF indexes; 40 of 64 sets

var cardinalRules = []pluralCheck{ // 166 elements
	0:   {cat: 0x2, setID: 0x3},
	1:   {cat: 0x22, setID: 0x1},
	2:   {cat: 0x2, setID: 0x4},
	3:   {cat: 0x2, setID: 0x4},
	4:   {cat: 0x7, setID: 0x1},
	5:   {cat: 0x62, setID: 0x3},
	6:   {cat: 0x22, setID: 0x4},
	7:   {cat: 0x7, setID: 0x3},
	8:   {cat: 0x42, setID: 0x1},
	9:   {cat: 0x22, setID: 0x4},
	10:  {cat: 0x22, setID: 0x4},
	11:  {cat: 0x22, setID: 0x5},
	12:  {cat: 0x22, setID: 0x1},
	13:  {cat: 0x22, setID: 0x1},
	14:  {cat: 0x7, setID: 0x4},
	15:  {cat: 0x92, setID: 0x3},
	16:  {cat: 0xf, setID: 0x6},
	17:  {cat: 0x1f, setID: 0x7},
	18:  {cat: 0x82, setID: 0x3},
	19:  {cat: 0x92, setID: 0x3},
	20:  {cat: 0xf, setID: 0x6},
	21:  {cat: 0x62, setID: 0x3},
	22:  {cat: 0x4a, setID: 0x6},
	23:  {cat: 0x7, setID: 0x8},
	24:  {cat: 0x62, setID: 0x3},
	25:  {cat: 0x1f, setID: 0x9},
	26:  {cat: 0x62, setID: 0x3},
	27:  {cat: 0x5f, setID: 0x9},
	28:  {cat: 0x72, setID: 0x3},
	29:  {cat: 0x29, setID: 0xa},
	30:  {cat: 0x29, setID: 0xb},
	31:  {cat: 0x4f, setID: 0xb},
	32:  {cat: 0x61, setID: 0x2},
	33:  {cat: 0x2f, setID: 0x6},
	34:  {cat: 0x3a, setID: 0x7},
	35:  {cat: 0x4f, setID: 0x6},
	36:  {cat: 0x5f, setID: 0x7},
	37:  {cat: 0x62, setID: 0x2},
	38:  {cat: 0x4f, setID: 0x6},
	39:  {cat: 0x72, setID: 0x2},
	40:  {cat: 0x21, setID: 0x3},
	41:  {cat: 0x7, setID: 0x4},
	42:  {cat: 0x32, setID: 0x3},
	43:  {cat: 0x21, setID: 0x3},
	44:  {cat: 0x22, setID: 0x1},
	45:  {cat: 0x22, setID: 0x1},
	46:  {cat: 0x23, setID: 0x2},
	47:  {cat: 0x2, setID: 0x3},
	48:  {cat: 0x22, setID: 0x1},
	49:  {cat: 0x24, setID: 0xc},
	50:  {cat: 0x7, setID: 0x1},
	51:  {cat: 0x62, setID: 0x3},
	52:  {cat: 0x74, setID: 0x3},
	53:  {cat: 0x24, setID: 0x3},
	54:  {cat: 0x2f, setID: 0xd},
	55:  {cat: 0x34, setID: 0x1},
	56:  {cat: 0xf, setID: 0x6},
	57:  {cat: 0x1f, setID: 0x7},
	58:  {cat: 0x62, setID: 0x3},
	59:  {cat: 0x4f, setID: 0x6},
	60:  {cat: 0x5a, setID: 0x7},
	61:  {cat: 0xf, setID: 0xe},
	62:  {cat: 0x1f, setID: 0xf},
	63:  {cat: 0x64, setID: 0x3},
	64:  {cat: 0x4f, setID: 0xe},
	65:  {cat: 0x5c, setID: 0xf},
	66:  {cat: 0x22, setID: 0x10},
	67:  {cat: 0x23, setID: 0x11},
	68:  {cat: 0x24, setID: 0x12},
	69:  {cat: 0xf, setID: 0x1},
	70:  {cat: 0x62, setID: 0x3},
	71:  {cat: 0xf, setID: 0x2},
	72:  {cat: 0x63, setID: 0x3},
	73:  {cat: 0xf, setID: 0x13},
	74:  {cat: 0x64, setID: 0x3},
	75:  {cat: 0x74, setID: 0x3},
	76:  {cat: 0xf, setID: 0x1},
	77:  {cat: 0x62, setID: 0x3},
	78:  {cat: 0x4a, setID: 0x1},
	79:  {cat: 0xf, setID: 0x2},
	80:  {cat: 0x63, setID: 0x3},
	81:  {cat: 0x4b, setID: 0x2},
	82:  {cat: 0xf, setID: 0x13},
	83:  {cat: 0x64, setID: 0x3},
	84:  {cat: 0x4c, setID: 0x13},
	85:  {cat: 0x7, setID: 0x1},
	86:  {cat: 0x62, setID: 0x3},
	87:  {cat: 0x7, setID: 0x2},
	88:  {cat: 0x63, setID: 0x3},
	89:  {cat: 0x2f, setID: 0xa},
	90:  {cat: 0x37, setID: 0x14},
	91:  {cat: 0x65, setID: 0x3},
	92:  {cat: 0x7, setID: 0x1},
	93:  {cat: 0x62, setID: 0x3},
	94:  {cat: 0x7, setID: 0x15},
	95:  {cat: 0x64, setID: 0x3},
	96:  {cat: 0x75, setID: 0x3},
	97:  {cat: 0x7, setID: 0x1},
	98:  {cat: 0x62, setID: 0x3},
	99:  {cat: 0xf, setID: 0xe},
	100: {cat: 0x1f, setID: 0xf},
	101: {cat: 0x64, setID: 0x3},
	102: {cat: 0xf, setID: 0x16},
	103: {cat: 0x17, setID: 0x1},
	104: {cat: 0x65, setID: 0x3},
	105: {cat: 0xf, setID: 0x17},
	106: {cat: 0x65, setID: 0x3},
	107: {cat: 0xf, setID: 0xf},
	108: {cat: 0x65, setID: 0x3},
	109: {cat: 0x2f, setID: 0x6},
	110: {cat: 0x3a, setID: 0x7},
	111: {cat: 0x2f, setID: 0xe},
	112: {cat: 0x3c, setID: 0xf},
	113: {cat: 0x2d, setID: 0xa},
	114: {cat: 0x2d, setID: 0x17},
	115: {cat: 0x2d, setID: 0x18},
	116: {cat: 0x2f, setID: 0x6},
	117: {cat: 0x3a, setID: 0xb},
	118: {cat: 0x2f, setID: 0x19},
	119: {cat: 0x3c, setID: 0xb},
	120: {cat: 0x55, setID: 0x3},
	121: {cat: 0x22, setID: 0x1},
	122: {cat: 0x24, setID: 0x3},
	123: {cat: 0x2c, setID: 0xc},
	124: {cat: 0x2d, setID: 0xb},
	125: {cat: 0xf, setID: 0x6},
	126: {cat: 0x1f, setID: 0x7},
	127: {cat: 0x62, setID: 0x3},
	128: {cat: 0xf, setID: 0xe},
	129: {cat: 0x1f, setID: 0xf},
	130: {cat: 0x64, setID: 0x3},
	131: {cat: 0xf, setID: 0xa},
	132: {cat: 0x65, setID: 0x3},
	133: {cat: 0xf, setID: 0x17},
	134: {cat: 0x65, setID: 0x3},
	135: {cat: 0xf, setID: 0x18},
	136: {cat: 0x65, setID: 0x3},
	137: {cat: 0x2f, setID: 0x6},
	138: {cat: 0x3a, setID: 0x1a},
	139: {cat: 0x2f, setID: 0x1b},
	140: {cat: 0x3b, setID: 0x1c},
	141: {cat: 0x2f, setID: 0x1d},
	142: {cat: 0x3c, setID: 0x1e},
	143: {cat: 0x37, setID: 0x3},
	144: {cat: 0xa5, setID: 0x0},
	145: {cat: 0x22, setID: 0x1},
	146: {cat: 0x23, setID: 0x2},
	147: {cat: 0x24, setID: 0x1f},
	148: {cat: 0x25, setID: 0x20},
	149: {cat: 0xf, setID: 0x6},
	150: {cat: 0x62, setID: 0x3},
	151: {cat: 0xf, setID: 0x1b},
	152: {cat: 0x63, setID: 0x3},
	153: {cat: 0xf, setID: 0x21},
	154: {cat: 0x64, setID: 0x3},
	155: {cat: 0x75, setID: 0x3},
	156: {cat: 0x21, setID: 0x3},
	157: {cat: 0x22, setID: 0x1},
	158: {cat: 0x23, setID: 0x2},
	159: {cat: 0x2c, setID: 0x22},
	160: {cat: 0x2d, setID: 0x5},
	161: {cat: 0x21, setID: 0x3},
	162: {cat: 0x22, setID: 0x1},
	163: {cat: 0x23, setID: 0x2},
	164: {cat: 0x24, setID: 0x23},
	165: {cat: 0x25, setID: 0x24},
} // Size: 356 bytes

var cardinalIndex = []uint8{ // 36 elements
	0x00, 0x00, 0x02, 0x03, 0x04, 0x06, 0x09, 0x0a,
	0x0c, 0x0d, 0x10, 0x14, 0x17, 0x1d, 0x28, 0x2b,
	0x2d, 0x2f, 0x32, 0x38, 0x42, 0x45, 0x4c, 0x55,
	0x5c, 0x61, 0x6d, 0x74, 0x79, 0x7d, 0x89, 0x91,
	0x95, 0x9c, 0xa1, 0xa6,
} // Size: 60 bytes

var cardinalLangToIndex = []uint8{ // 775 elements
	// Entry 0 - 3F
	0x00, 0x08, 0x08, 0x08, 0x00, 0x00, 0x06, 0x06,
	0x01, 0x01, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21, 0x21,
	0x01, 0x01, 0x08, 0x08, 0x04, 0x04, 0x08, 0x00,
	0x00, 0x08, 0x08, 0x00, 0x00, 0x1a, 0x1a, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x06, 0x00, 0x00,
	// Entry 40 - 7F
	0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x1e, 0x1e,
	0x08, 0x08, 0x13, 0x00, 0x00, 0x13, 0x13, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x00, 0x00, 0x00, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x18, 0x18, 0x00, 0x00, 0x22, 0x22, 0x09, 0x09,
	0x09, 0x00, 0x00, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x00, 0x00, 0x16, 0x16, 0x00,
	0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 80 - BF
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	// Entry C0 - FF
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	// Entry 100 - 13F
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x04, 0x04,
	0x08, 0x08, 0x00, 0x00, 0x01, 0x01, 0x01, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x04, 0x04, 0x0c, 0x0c,
	0x08, 0x08, 0x08, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	// Entry 140 - 17F
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x08, 0x08, 0x04, 0x04, 0x1f, 0x1f,
	0x14, 0x14, 0x04, 0x04, 0x08, 0x08, 0x08, 0x08,
	0x01, 0x01, 0x06, 0x00, 0x00, 0x20, 0x20, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x17, 0x17, 0x01,
	0x01, 0x13, 0x13, 0x13, 0x16, 0x16, 0x08, 0x08,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 180 - 1BF
	0x00, 0x04, 0x0a, 0x0a, 0x04, 0x04, 0x04, 0x04,
	0x04, 0x10, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08,
	0x00, 0x08, 0x08, 0x00, 0x00, 0x08, 0x08, 0x02,
	0x02, 0x08, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x00, 0x00, 0x00, 0x00, 0x01,
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08,
	0x08, 0x08, 0x00, 0x00, 0x0f, 0x0f, 0x08, 0x10,
	// Entry 1C0 - 1FF
	0x10, 0x08, 0x08, 0x0e, 0x0e, 0x08, 0x08, 0x08,
	0x08, 0x00, 0x00, 0x06, 0x06, 0x06, 0x06, 0x06,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x1b, 0x1b, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x0d, 0x08,
	0x08, 0x08, 0x00, 0x00, 0x00, 0x00, 0x06, 0x06,
	0x00, 0x00, 0x08, 0x08, 0x0b, 0x0b, 0x08, 0x08,
	0x08, 0x08, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00,
	0x00, 0x1c, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 200 - 23F
	0x00, 0x08, 0x10, 0x10, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x00,
	0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x00, 0x08,
	0x06, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x06, 0x00,
	0x00, 0x06, 0x06, 0x08, 0x19, 0x19, 0x0d, 0x0d,
	0x08, 0x08, 0x03, 0x04, 0x03, 0x04, 0x04, 0x04,
	// Entry 240 - 27F
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x00,
	0x00, 0x00, 0x00, 0x08, 0x08, 0x00, 0x00, 0x12,
	0x12, 0x12, 0x08, 0x08, 0x1d, 0x1d, 0x1d, 0x1d,
	0x1d, 0x1d, 0x1d, 0x00, 0x00, 0x08, 0x08, 0x00,
	0x00, 0x08, 0x08, 0x00, 0x00, 0x08, 0x08, 0x08,
	0x10, 0x10, 0x10, 0x10, 0x08, 0x08, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x11, 0x11,
	0x05, 0x05, 0x18, 0x18, 0x15, 0x15, 0x10, 0x10,
	// Entry 280 - 2BF
	0x10, 0x10, 0x10, 0x10, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x13,
	0x13, 0x13, 0x13, 0x13, 0x13, 0x13, 0x13, 0x13,
	0x13, 0x13, 0x08, 0x08, 0x08, 0x04, 0x04, 0x04,
	0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x00, 0x00, 0x00, 0x00, 0x06, 0x06, 0x06,
	0x08, 0x08, 0x08, 0x00, 0x08, 0x00, 0x00, 0x08,
	// Entry 2C0 - 2FF
	0x08, 0x08, 0x08, 0x00, 0x00, 0x00, 0x00, 0x07,
	0x07, 0x08, 0x08, 0x1d, 0x1d, 0x04, 0x04, 0x04,
	0x08, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x08,
	0x08, 0x08, 0x08, 0x06, 0x08, 0x08, 0x00, 0x00,
	0x08, 0x08, 0x08, 0x00, 0x00, 0x04, 0x04, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry 300 - 33F
	0x00, 0x00, 0x00, 0x01, 0x01, 0x04, 0x04,
} // Size: 799 bytes

var cardinalInclusionMasks = []uint64{ // 100 elements
	// Entry 0 - 1F
	0x0000000200500419, 0x0000000000512153, 0x000000000a327105, 0x0000000ca23c7101,
	0x00000004a23c7201, 0x0000000482943001, 0x0000001482943201, 0x0000000502943001,
	0x0000000502943001, 0x0000000522943201, 0x0000000540543401, 0x00000000454128e1,
	0x000000005b02e821, 0x000000006304e821, 0x000000006304ea21, 0x0000000042842821,
	0x0000000042842a21, 0x0000000042842821, 0x0000000042842821, 0x0000000062842a21,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000000400421, 0x0000000000400061,
	// Entry 20 - 3F
	0x000000000a004021, 0x0000000022004021, 0x0000000022004221, 0x0000000002800021,
	0x0000000002800221, 0x0000000002800021, 0x0000000002800021, 0x0000000022800221,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000000400421, 0x0000000000400061,
	0x000000000a004021, 0x0000000022004021, 0x0000000022004221, 0x0000000002800021,
	0x0000000002800221, 0x0000000002800021, 0x0000000002800021, 0x0000000022800221,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	// Entry 40 - 5F
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000040400421, 0x0000000044400061,
	0x000000005a004021, 0x0000000062004021, 0x0000000062004221, 0x0000000042800021,
	0x0000000042800221, 0x0000000042800021, 0x0000000042800021, 0x0000000062800221,
	0x0000000200400421, 0x0000000000400061, 0x000000000a004021, 0x0000000022004021,
	0x0000000022004221, 0x0000000002800021, 0x0000000002800221, 0x0000000002800021,
	0x0000000002800021, 0x0000000022800221, 0x0000000040400421, 0x0000000044400061,
	0x000000005a004021, 0x0000000062004021, 0x0000000062004221, 0x0000000042800021,
	// Entry 60 - 7F
	0x0000000042800221, 0x0000000042800021, 0x0000000042800021, 0x0000000062800221,
} // Size: 824 bytes

// Slots used for cardinal: A6 of 0xFF rules; 24 of 0xFF indexes; 37 of 64 sets

// Total table size 3860 bytes (3KiB); checksum: B96047F
