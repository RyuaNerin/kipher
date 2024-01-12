package lea

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

type funcBlock func(ctx *leaContext, dst, src []byte)

var (
	leaEnc1 funcBlock = leaEnc1Go
	leaEnc4 funcBlock = leaEnc4Using1
	leaEnc8 funcBlock = leaEnc8Using4

	leaDec1 funcBlock = leaDec1Go
	leaDec4 funcBlock = leaDec4Using1
	leaDec8 funcBlock = leaDec8Using4
)

func leaEnc4Using1(ctx *leaContext, dst, src []byte) {
	leaEnc1(ctx, dst[BlockSize*0:], src[BlockSize*0:])
	leaEnc1(ctx, dst[BlockSize*1:], src[BlockSize*1:])
	leaEnc1(ctx, dst[BlockSize*2:], src[BlockSize*2:])
	leaEnc1(ctx, dst[BlockSize*3:], src[BlockSize*3:])
}
func leaDec4Using1(ctx *leaContext, dst, src []byte) {
	leaDec1(ctx, dst[BlockSize*0:], src[BlockSize*0:])
	leaDec1(ctx, dst[BlockSize*1:], src[BlockSize*1:])
	leaDec1(ctx, dst[BlockSize*2:], src[BlockSize*2:])
	leaDec1(ctx, dst[BlockSize*3:], src[BlockSize*3:])
}

func leaEnc8Using4(ctx *leaContext, dst, src []byte) {
	leaEnc4(ctx, dst[BlockSize*0:], src[BlockSize*0:])
	leaEnc4(ctx, dst[BlockSize*4:], src[BlockSize*4:])
}
func leaDec8Using4(ctx *leaContext, dst, src []byte) {
	leaDec4(ctx, dst[BlockSize*0:], src[BlockSize*0:])
	leaDec4(ctx, dst[BlockSize*4:], src[BlockSize*4:])
}

func (ctx *leaContext) initContext(key []byte) error {
	l := len(key)
	switch l {
	case 16:
	case 24:
	case 32:
	default:
		return KeySizeError(l)
	}

	ctx.round = uint8(leaSetKeyGo(ctx.rk[:], key))

	return nil
}

func (ctx *leaContext) BlockSize() int {
	return BlockSize
}

func (ctx *leaContext) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic(fmt.Sprintf("krypto/lea: invalid block size %d (src)", len(src)))
	}
	if len(dst) < BlockSize {
		panic(fmt.Sprintf("krypto/lea: invalid block size %d (dst)", len(dst)))
	}

	if !ctx.ecb {
		leaEnc1Go(ctx, dst, src)
	} else {
		if len(src)%BlockSize != 0 {
			panic("krypto/lea: input not full blocks")
		}

		remainBlock := len(src) / ctx.BlockSize()

		for remainBlock >= 8 {
			remainBlock -= 8
			leaEnc8(ctx, dst, src)

			dst, src = dst[0x80:], src[0x80:]
		}

		for remainBlock >= 4 {
			remainBlock -= 4
			leaEnc4(ctx, dst, src)

			dst, src = dst[0x40:], src[0x40:]
		}

		for remainBlock > 0 {
			remainBlock -= 1
			leaEnc1(ctx, dst, src)

			dst, src = dst[0x10:], src[0x10:]
		}
	}
}

func (ctx *leaContext) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic(fmt.Sprintf("krypto/lea: invalid block size %d (src)", len(src)))
	}
	if len(dst) < BlockSize {
		panic(fmt.Sprintf("krypto/lea: invalid block size %d (dst)", len(dst)))
	}

	if !ctx.ecb {
		leaDec1Go(ctx, dst, src)
	} else {
		if len(src)%BlockSize != 0 {
			panic("krypto/lea: input not full blocks")
		}

		remainBlock := len(src) / ctx.BlockSize()

		for remainBlock >= 8 {
			remainBlock -= 8
			leaDec8(ctx, dst, src)

			dst, src = dst[0x80:], src[0x80:]
		}

		for remainBlock >= 4 {
			remainBlock -= 4
			leaDec4(ctx, dst, src)

			dst, src = dst[0x40:], src[0x40:]
		}

		for remainBlock > 0 {
			remainBlock -= 1
			leaDec1(ctx, dst, src)

			dst, src = dst[0x10:], src[0x10:]
		}
	}
}

func leaSetKeyGo(rk []uint32, key []byte) int {
	keyLen := len(key)

	switch keyLen {
	case 16:
		rk[0] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*0:])+delta[0][0], 1)
		rk[6] = bits.RotateLeft32(rk[0]+delta[1][1], 1)
		rk[12] = bits.RotateLeft32(rk[6]+delta[2][2], 1)
		rk[18] = bits.RotateLeft32(rk[12]+delta[3][3], 1)
		rk[24] = bits.RotateLeft32(rk[18]+delta[0][4], 1)
		rk[30] = bits.RotateLeft32(rk[24]+delta[1][5], 1)
		rk[36] = bits.RotateLeft32(rk[30]+delta[2][6], 1)
		rk[42] = bits.RotateLeft32(rk[36]+delta[3][7], 1)
		rk[48] = bits.RotateLeft32(rk[42]+delta[0][8], 1)
		rk[54] = bits.RotateLeft32(rk[48]+delta[1][9], 1)
		rk[60] = bits.RotateLeft32(rk[54]+delta[2][10], 1)
		rk[66] = bits.RotateLeft32(rk[60]+delta[3][11], 1)
		rk[72] = bits.RotateLeft32(rk[66]+delta[0][12], 1)
		rk[78] = bits.RotateLeft32(rk[72]+delta[1][13], 1)
		rk[84] = bits.RotateLeft32(rk[78]+delta[2][14], 1)
		rk[90] = bits.RotateLeft32(rk[84]+delta[3][15], 1)
		rk[96] = bits.RotateLeft32(rk[90]+delta[0][16], 1)
		rk[102] = bits.RotateLeft32(rk[96]+delta[1][17], 1)
		rk[108] = bits.RotateLeft32(rk[102]+delta[2][18], 1)
		rk[114] = bits.RotateLeft32(rk[108]+delta[3][19], 1)
		rk[120] = bits.RotateLeft32(rk[114]+delta[0][20], 1)
		rk[126] = bits.RotateLeft32(rk[120]+delta[1][21], 1)
		rk[132] = bits.RotateLeft32(rk[126]+delta[2][22], 1)
		rk[138] = bits.RotateLeft32(rk[132]+delta[3][23], 1)

		/**
		rk[  1] = rk[  3] = rk[  5] = bits.RotateLeft32(loadU32(mk,1) + delta[0][ 1], 3);
		rk[  7] = rk[  9] = rk[ 11] = bits.RotateLeft32(rk[  1] + delta[1][ 2], 3);
		rk[ 13] = rk[ 15] = rk[ 17] = bits.RotateLeft32(rk[  7] + delta[2][ 3], 3);
		rk[ 19] = rk[ 21] = rk[ 23] = bits.RotateLeft32(rk[ 13] + delta[3][ 4], 3);
		rk[ 25] = rk[ 27] = rk[ 29] = bits.RotateLeft32(rk[ 19] + delta[0][ 5], 3);
		rk[ 31] = rk[ 33] = rk[ 35] = bits.RotateLeft32(rk[ 25] + delta[1][ 6], 3);
		rk[ 37] = rk[ 39] = rk[ 41] = bits.RotateLeft32(rk[ 31] + delta[2][ 7], 3);
		rk[ 43] = rk[ 45] = rk[ 47] = bits.RotateLeft32(rk[ 37] + delta[3][ 8], 3);
		rk[ 49] = rk[ 51] = rk[ 53] = bits.RotateLeft32(rk[ 43] + delta[0][ 9], 3);
		rk[ 55] = rk[ 57] = rk[ 59] = bits.RotateLeft32(rk[ 49] + delta[1][10], 3);
		rk[ 61] = rk[ 63] = rk[ 65] = bits.RotateLeft32(rk[ 55] + delta[2][11], 3);
		rk[ 67] = rk[ 69] = rk[ 71] = bits.RotateLeft32(rk[ 61] + delta[3][12], 3);
		rk[ 73] = rk[ 75] = rk[ 77] = bits.RotateLeft32(rk[ 67] + delta[0][13], 3);
		rk[ 79] = rk[ 81] = rk[ 83] = bits.RotateLeft32(rk[ 73] + delta[1][14], 3);
		rk[ 85] = rk[ 87] = rk[ 89] = bits.RotateLeft32(rk[ 79] + delta[2][15], 3);
		rk[ 91] = rk[ 93] = rk[ 95] = bits.RotateLeft32(rk[ 85] + delta[3][16], 3);
		rk[ 97] = rk[ 99] = rk[101] = bits.RotateLeft32(rk[ 91] + delta[0][17], 3);
		rk[103] = rk[105] = rk[107] = bits.RotateLeft32(rk[ 97] + delta[1][18], 3);
		rk[109] = rk[111] = rk[113] = bits.RotateLeft32(rk[103] + delta[2][19], 3);
		rk[115] = rk[117] = rk[119] = bits.RotateLeft32(rk[109] + delta[3][20], 3);
		rk[121] = rk[123] = rk[125] = bits.RotateLeft32(rk[115] + delta[0][21], 3);
		rk[127] = rk[129] = rk[131] = bits.RotateLeft32(rk[121] + delta[1][22], 3);
		rk[133] = rk[135] = rk[137] = bits.RotateLeft32(rk[127] + delta[2][23], 3);
		rk[139] = rk[141] = rk[143] = bits.RotateLeft32(rk[133] + delta[3][24], 3);
		*/
		tmp := bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*1:])+delta[0][1], 3)
		rk[1] = tmp
		rk[3] = tmp
		rk[5] = tmp

		for i := 1; i <= 23; i++ {
			tmp = bits.RotateLeft32(rk[(i-1)*6+1]+delta[i%4][i+1], 3)

			rk[i*6+1] = tmp
			rk[i*6+3] = tmp
			rk[i*6+5] = tmp
		}

		rk[2] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*2:])+delta[0][2], 6)
		rk[8] = bits.RotateLeft32(rk[2]+delta[1][3], 6)
		rk[14] = bits.RotateLeft32(rk[8]+delta[2][4], 6)
		rk[20] = bits.RotateLeft32(rk[14]+delta[3][5], 6)
		rk[26] = bits.RotateLeft32(rk[20]+delta[0][6], 6)
		rk[32] = bits.RotateLeft32(rk[26]+delta[1][7], 6)
		rk[38] = bits.RotateLeft32(rk[32]+delta[2][8], 6)
		rk[44] = bits.RotateLeft32(rk[38]+delta[3][9], 6)
		rk[50] = bits.RotateLeft32(rk[44]+delta[0][10], 6)
		rk[56] = bits.RotateLeft32(rk[50]+delta[1][11], 6)
		rk[62] = bits.RotateLeft32(rk[56]+delta[2][12], 6)
		rk[68] = bits.RotateLeft32(rk[62]+delta[3][13], 6)
		rk[74] = bits.RotateLeft32(rk[68]+delta[0][14], 6)
		rk[80] = bits.RotateLeft32(rk[74]+delta[1][15], 6)
		rk[86] = bits.RotateLeft32(rk[80]+delta[2][16], 6)
		rk[92] = bits.RotateLeft32(rk[86]+delta[3][17], 6)
		rk[98] = bits.RotateLeft32(rk[92]+delta[0][18], 6)
		rk[104] = bits.RotateLeft32(rk[98]+delta[1][19], 6)
		rk[110] = bits.RotateLeft32(rk[104]+delta[2][20], 6)
		rk[116] = bits.RotateLeft32(rk[110]+delta[3][21], 6)
		rk[122] = bits.RotateLeft32(rk[116]+delta[0][22], 6)
		rk[128] = bits.RotateLeft32(rk[122]+delta[1][23], 6)
		rk[134] = bits.RotateLeft32(rk[128]+delta[2][24], 6)
		rk[140] = bits.RotateLeft32(rk[134]+delta[3][25], 6)

		rk[4] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*3:])+delta[0][3], 11)
		rk[10] = bits.RotateLeft32(rk[4]+delta[1][4], 11)
		rk[16] = bits.RotateLeft32(rk[10]+delta[2][5], 11)
		rk[22] = bits.RotateLeft32(rk[16]+delta[3][6], 11)
		rk[28] = bits.RotateLeft32(rk[22]+delta[0][7], 11)
		rk[34] = bits.RotateLeft32(rk[28]+delta[1][8], 11)
		rk[40] = bits.RotateLeft32(rk[34]+delta[2][9], 11)
		rk[46] = bits.RotateLeft32(rk[40]+delta[3][10], 11)
		rk[52] = bits.RotateLeft32(rk[46]+delta[0][11], 11)
		rk[58] = bits.RotateLeft32(rk[52]+delta[1][12], 11)
		rk[64] = bits.RotateLeft32(rk[58]+delta[2][13], 11)
		rk[70] = bits.RotateLeft32(rk[64]+delta[3][14], 11)
		rk[76] = bits.RotateLeft32(rk[70]+delta[0][15], 11)
		rk[82] = bits.RotateLeft32(rk[76]+delta[1][16], 11)
		rk[88] = bits.RotateLeft32(rk[82]+delta[2][17], 11)
		rk[94] = bits.RotateLeft32(rk[88]+delta[3][18], 11)
		rk[100] = bits.RotateLeft32(rk[94]+delta[0][19], 11)
		rk[106] = bits.RotateLeft32(rk[100]+delta[1][20], 11)
		rk[112] = bits.RotateLeft32(rk[106]+delta[2][21], 11)
		rk[118] = bits.RotateLeft32(rk[112]+delta[3][22], 11)
		rk[124] = bits.RotateLeft32(rk[118]+delta[0][23], 11)
		rk[130] = bits.RotateLeft32(rk[124]+delta[1][24], 11)
		rk[136] = bits.RotateLeft32(rk[130]+delta[2][25], 11)
		rk[142] = bits.RotateLeft32(rk[136]+delta[3][26], 11)
		break

	case 24:
		rk[0] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*0:])+delta[0][0], 1)
		rk[6] = bits.RotateLeft32(rk[0]+delta[1][1], 1)
		rk[12] = bits.RotateLeft32(rk[6]+delta[2][2], 1)
		rk[18] = bits.RotateLeft32(rk[12]+delta[3][3], 1)
		rk[24] = bits.RotateLeft32(rk[18]+delta[4][4], 1)
		rk[30] = bits.RotateLeft32(rk[24]+delta[5][5], 1)
		rk[36] = bits.RotateLeft32(rk[30]+delta[0][6], 1)
		rk[42] = bits.RotateLeft32(rk[36]+delta[1][7], 1)
		rk[48] = bits.RotateLeft32(rk[42]+delta[2][8], 1)
		rk[54] = bits.RotateLeft32(rk[48]+delta[3][9], 1)
		rk[60] = bits.RotateLeft32(rk[54]+delta[4][10], 1)
		rk[66] = bits.RotateLeft32(rk[60]+delta[5][11], 1)
		rk[72] = bits.RotateLeft32(rk[66]+delta[0][12], 1)
		rk[78] = bits.RotateLeft32(rk[72]+delta[1][13], 1)
		rk[84] = bits.RotateLeft32(rk[78]+delta[2][14], 1)
		rk[90] = bits.RotateLeft32(rk[84]+delta[3][15], 1)
		rk[96] = bits.RotateLeft32(rk[90]+delta[4][16], 1)
		rk[102] = bits.RotateLeft32(rk[96]+delta[5][17], 1)
		rk[108] = bits.RotateLeft32(rk[102]+delta[0][18], 1)
		rk[114] = bits.RotateLeft32(rk[108]+delta[1][19], 1)
		rk[120] = bits.RotateLeft32(rk[114]+delta[2][20], 1)
		rk[126] = bits.RotateLeft32(rk[120]+delta[3][21], 1)
		rk[132] = bits.RotateLeft32(rk[126]+delta[4][22], 1)
		rk[138] = bits.RotateLeft32(rk[132]+delta[5][23], 1)
		rk[144] = bits.RotateLeft32(rk[138]+delta[0][24], 1)
		rk[150] = bits.RotateLeft32(rk[144]+delta[1][25], 1)
		rk[156] = bits.RotateLeft32(rk[150]+delta[2][26], 1)
		rk[162] = bits.RotateLeft32(rk[156]+delta[3][27], 1)

		rk[1] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*1:])+delta[0][1], 3)
		rk[7] = bits.RotateLeft32(rk[1]+delta[1][2], 3)
		rk[13] = bits.RotateLeft32(rk[7]+delta[2][3], 3)
		rk[19] = bits.RotateLeft32(rk[13]+delta[3][4], 3)
		rk[25] = bits.RotateLeft32(rk[19]+delta[4][5], 3)
		rk[31] = bits.RotateLeft32(rk[25]+delta[5][6], 3)
		rk[37] = bits.RotateLeft32(rk[31]+delta[0][7], 3)
		rk[43] = bits.RotateLeft32(rk[37]+delta[1][8], 3)
		rk[49] = bits.RotateLeft32(rk[43]+delta[2][9], 3)
		rk[55] = bits.RotateLeft32(rk[49]+delta[3][10], 3)
		rk[61] = bits.RotateLeft32(rk[55]+delta[4][11], 3)
		rk[67] = bits.RotateLeft32(rk[61]+delta[5][12], 3)
		rk[73] = bits.RotateLeft32(rk[67]+delta[0][13], 3)
		rk[79] = bits.RotateLeft32(rk[73]+delta[1][14], 3)
		rk[85] = bits.RotateLeft32(rk[79]+delta[2][15], 3)
		rk[91] = bits.RotateLeft32(rk[85]+delta[3][16], 3)
		rk[97] = bits.RotateLeft32(rk[91]+delta[4][17], 3)
		rk[103] = bits.RotateLeft32(rk[97]+delta[5][18], 3)
		rk[109] = bits.RotateLeft32(rk[103]+delta[0][19], 3)
		rk[115] = bits.RotateLeft32(rk[109]+delta[1][20], 3)
		rk[121] = bits.RotateLeft32(rk[115]+delta[2][21], 3)
		rk[127] = bits.RotateLeft32(rk[121]+delta[3][22], 3)
		rk[133] = bits.RotateLeft32(rk[127]+delta[4][23], 3)
		rk[139] = bits.RotateLeft32(rk[133]+delta[5][24], 3)
		rk[145] = bits.RotateLeft32(rk[139]+delta[0][25], 3)
		rk[151] = bits.RotateLeft32(rk[145]+delta[1][26], 3)
		rk[157] = bits.RotateLeft32(rk[151]+delta[2][27], 3)
		rk[163] = bits.RotateLeft32(rk[157]+delta[3][28], 3)

		rk[2] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*2:])+delta[0][2], 6)
		rk[8] = bits.RotateLeft32(rk[2]+delta[1][3], 6)
		rk[14] = bits.RotateLeft32(rk[8]+delta[2][4], 6)
		rk[20] = bits.RotateLeft32(rk[14]+delta[3][5], 6)
		rk[26] = bits.RotateLeft32(rk[20]+delta[4][6], 6)
		rk[32] = bits.RotateLeft32(rk[26]+delta[5][7], 6)
		rk[38] = bits.RotateLeft32(rk[32]+delta[0][8], 6)
		rk[44] = bits.RotateLeft32(rk[38]+delta[1][9], 6)
		rk[50] = bits.RotateLeft32(rk[44]+delta[2][10], 6)
		rk[56] = bits.RotateLeft32(rk[50]+delta[3][11], 6)
		rk[62] = bits.RotateLeft32(rk[56]+delta[4][12], 6)
		rk[68] = bits.RotateLeft32(rk[62]+delta[5][13], 6)
		rk[74] = bits.RotateLeft32(rk[68]+delta[0][14], 6)
		rk[80] = bits.RotateLeft32(rk[74]+delta[1][15], 6)
		rk[86] = bits.RotateLeft32(rk[80]+delta[2][16], 6)
		rk[92] = bits.RotateLeft32(rk[86]+delta[3][17], 6)
		rk[98] = bits.RotateLeft32(rk[92]+delta[4][18], 6)
		rk[104] = bits.RotateLeft32(rk[98]+delta[5][19], 6)
		rk[110] = bits.RotateLeft32(rk[104]+delta[0][20], 6)
		rk[116] = bits.RotateLeft32(rk[110]+delta[1][21], 6)
		rk[122] = bits.RotateLeft32(rk[116]+delta[2][22], 6)
		rk[128] = bits.RotateLeft32(rk[122]+delta[3][23], 6)
		rk[134] = bits.RotateLeft32(rk[128]+delta[4][24], 6)
		rk[140] = bits.RotateLeft32(rk[134]+delta[5][25], 6)
		rk[146] = bits.RotateLeft32(rk[140]+delta[0][26], 6)
		rk[152] = bits.RotateLeft32(rk[146]+delta[1][27], 6)
		rk[158] = bits.RotateLeft32(rk[152]+delta[2][28], 6)
		rk[164] = bits.RotateLeft32(rk[158]+delta[3][29], 6)

		rk[3] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*3:])+delta[0][3], 11)
		rk[9] = bits.RotateLeft32(rk[3]+delta[1][4], 11)
		rk[15] = bits.RotateLeft32(rk[9]+delta[2][5], 11)
		rk[21] = bits.RotateLeft32(rk[15]+delta[3][6], 11)
		rk[27] = bits.RotateLeft32(rk[21]+delta[4][7], 11)
		rk[33] = bits.RotateLeft32(rk[27]+delta[5][8], 11)
		rk[39] = bits.RotateLeft32(rk[33]+delta[0][9], 11)
		rk[45] = bits.RotateLeft32(rk[39]+delta[1][10], 11)
		rk[51] = bits.RotateLeft32(rk[45]+delta[2][11], 11)
		rk[57] = bits.RotateLeft32(rk[51]+delta[3][12], 11)
		rk[63] = bits.RotateLeft32(rk[57]+delta[4][13], 11)
		rk[69] = bits.RotateLeft32(rk[63]+delta[5][14], 11)
		rk[75] = bits.RotateLeft32(rk[69]+delta[0][15], 11)
		rk[81] = bits.RotateLeft32(rk[75]+delta[1][16], 11)
		rk[87] = bits.RotateLeft32(rk[81]+delta[2][17], 11)
		rk[93] = bits.RotateLeft32(rk[87]+delta[3][18], 11)
		rk[99] = bits.RotateLeft32(rk[93]+delta[4][19], 11)
		rk[105] = bits.RotateLeft32(rk[99]+delta[5][20], 11)
		rk[111] = bits.RotateLeft32(rk[105]+delta[0][21], 11)
		rk[117] = bits.RotateLeft32(rk[111]+delta[1][22], 11)
		rk[123] = bits.RotateLeft32(rk[117]+delta[2][23], 11)
		rk[129] = bits.RotateLeft32(rk[123]+delta[3][24], 11)
		rk[135] = bits.RotateLeft32(rk[129]+delta[4][25], 11)
		rk[141] = bits.RotateLeft32(rk[135]+delta[5][26], 11)
		rk[147] = bits.RotateLeft32(rk[141]+delta[0][27], 11)
		rk[153] = bits.RotateLeft32(rk[147]+delta[1][28], 11)
		rk[159] = bits.RotateLeft32(rk[153]+delta[2][29], 11)
		rk[165] = bits.RotateLeft32(rk[159]+delta[3][30], 11)

		rk[4] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*4:])+delta[0][4], 13)
		rk[10] = bits.RotateLeft32(rk[4]+delta[1][5], 13)
		rk[16] = bits.RotateLeft32(rk[10]+delta[2][6], 13)
		rk[22] = bits.RotateLeft32(rk[16]+delta[3][7], 13)
		rk[28] = bits.RotateLeft32(rk[22]+delta[4][8], 13)
		rk[34] = bits.RotateLeft32(rk[28]+delta[5][9], 13)
		rk[40] = bits.RotateLeft32(rk[34]+delta[0][10], 13)
		rk[46] = bits.RotateLeft32(rk[40]+delta[1][11], 13)
		rk[52] = bits.RotateLeft32(rk[46]+delta[2][12], 13)
		rk[58] = bits.RotateLeft32(rk[52]+delta[3][13], 13)
		rk[64] = bits.RotateLeft32(rk[58]+delta[4][14], 13)
		rk[70] = bits.RotateLeft32(rk[64]+delta[5][15], 13)
		rk[76] = bits.RotateLeft32(rk[70]+delta[0][16], 13)
		rk[82] = bits.RotateLeft32(rk[76]+delta[1][17], 13)
		rk[88] = bits.RotateLeft32(rk[82]+delta[2][18], 13)
		rk[94] = bits.RotateLeft32(rk[88]+delta[3][19], 13)
		rk[100] = bits.RotateLeft32(rk[94]+delta[4][20], 13)
		rk[106] = bits.RotateLeft32(rk[100]+delta[5][21], 13)
		rk[112] = bits.RotateLeft32(rk[106]+delta[0][22], 13)
		rk[118] = bits.RotateLeft32(rk[112]+delta[1][23], 13)
		rk[124] = bits.RotateLeft32(rk[118]+delta[2][24], 13)
		rk[130] = bits.RotateLeft32(rk[124]+delta[3][25], 13)
		rk[136] = bits.RotateLeft32(rk[130]+delta[4][26], 13)
		rk[142] = bits.RotateLeft32(rk[136]+delta[5][27], 13)
		rk[148] = bits.RotateLeft32(rk[142]+delta[0][28], 13)
		rk[154] = bits.RotateLeft32(rk[148]+delta[1][29], 13)
		rk[160] = bits.RotateLeft32(rk[154]+delta[2][30], 13)
		rk[166] = bits.RotateLeft32(rk[160]+delta[3][31], 13)

		rk[5] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*5:])+delta[0][5], 17)
		rk[11] = bits.RotateLeft32(rk[5]+delta[1][6], 17)
		rk[17] = bits.RotateLeft32(rk[11]+delta[2][7], 17)
		rk[23] = bits.RotateLeft32(rk[17]+delta[3][8], 17)
		rk[29] = bits.RotateLeft32(rk[23]+delta[4][9], 17)
		rk[35] = bits.RotateLeft32(rk[29]+delta[5][10], 17)
		rk[41] = bits.RotateLeft32(rk[35]+delta[0][11], 17)
		rk[47] = bits.RotateLeft32(rk[41]+delta[1][12], 17)
		rk[53] = bits.RotateLeft32(rk[47]+delta[2][13], 17)
		rk[59] = bits.RotateLeft32(rk[53]+delta[3][14], 17)
		rk[65] = bits.RotateLeft32(rk[59]+delta[4][15], 17)
		rk[71] = bits.RotateLeft32(rk[65]+delta[5][16], 17)
		rk[77] = bits.RotateLeft32(rk[71]+delta[0][17], 17)
		rk[83] = bits.RotateLeft32(rk[77]+delta[1][18], 17)
		rk[89] = bits.RotateLeft32(rk[83]+delta[2][19], 17)
		rk[95] = bits.RotateLeft32(rk[89]+delta[3][20], 17)
		rk[101] = bits.RotateLeft32(rk[95]+delta[4][21], 17)
		rk[107] = bits.RotateLeft32(rk[101]+delta[5][22], 17)
		rk[113] = bits.RotateLeft32(rk[107]+delta[0][23], 17)
		rk[119] = bits.RotateLeft32(rk[113]+delta[1][24], 17)
		rk[125] = bits.RotateLeft32(rk[119]+delta[2][25], 17)
		rk[131] = bits.RotateLeft32(rk[125]+delta[3][26], 17)
		rk[137] = bits.RotateLeft32(rk[131]+delta[4][27], 17)
		rk[143] = bits.RotateLeft32(rk[137]+delta[5][28], 17)
		rk[149] = bits.RotateLeft32(rk[143]+delta[0][29], 17)
		rk[155] = bits.RotateLeft32(rk[149]+delta[1][30], 17)
		rk[161] = bits.RotateLeft32(rk[155]+delta[2][31], 17)
		rk[167] = bits.RotateLeft32(rk[161]+delta[3][0], 17)
		break

	case 32:
		rk[0] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*0:])+delta[0][0], 1)
		rk[8] = bits.RotateLeft32(rk[0]+delta[1][3], 6)
		rk[16] = bits.RotateLeft32(rk[8]+delta[2][6], 13)
		rk[24] = bits.RotateLeft32(rk[16]+delta[4][4], 1)
		rk[32] = bits.RotateLeft32(rk[24]+delta[5][7], 6)
		rk[40] = bits.RotateLeft32(rk[32]+delta[6][10], 13)
		rk[48] = bits.RotateLeft32(rk[40]+delta[0][8], 1)
		rk[56] = bits.RotateLeft32(rk[48]+delta[1][11], 6)
		rk[64] = bits.RotateLeft32(rk[56]+delta[2][14], 13)
		rk[72] = bits.RotateLeft32(rk[64]+delta[4][12], 1)
		rk[80] = bits.RotateLeft32(rk[72]+delta[5][15], 6)
		rk[88] = bits.RotateLeft32(rk[80]+delta[6][18], 13)
		rk[96] = bits.RotateLeft32(rk[88]+delta[0][16], 1)
		rk[104] = bits.RotateLeft32(rk[96]+delta[1][19], 6)
		rk[112] = bits.RotateLeft32(rk[104]+delta[2][22], 13)
		rk[120] = bits.RotateLeft32(rk[112]+delta[4][20], 1)
		rk[128] = bits.RotateLeft32(rk[120]+delta[5][23], 6)
		rk[136] = bits.RotateLeft32(rk[128]+delta[6][26], 13)
		rk[144] = bits.RotateLeft32(rk[136]+delta[0][24], 1)
		rk[152] = bits.RotateLeft32(rk[144]+delta[1][27], 6)
		rk[160] = bits.RotateLeft32(rk[152]+delta[2][30], 13)
		rk[168] = bits.RotateLeft32(rk[160]+delta[4][28], 1)
		rk[176] = bits.RotateLeft32(rk[168]+delta[5][31], 6)
		rk[184] = bits.RotateLeft32(rk[176]+delta[6][2], 13)

		rk[1] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*1:])+delta[0][1], 3)
		rk[9] = bits.RotateLeft32(rk[1]+delta[1][4], 11)
		rk[17] = bits.RotateLeft32(rk[9]+delta[2][7], 17)
		rk[25] = bits.RotateLeft32(rk[17]+delta[4][5], 3)
		rk[33] = bits.RotateLeft32(rk[25]+delta[5][8], 11)
		rk[41] = bits.RotateLeft32(rk[33]+delta[6][11], 17)
		rk[49] = bits.RotateLeft32(rk[41]+delta[0][9], 3)
		rk[57] = bits.RotateLeft32(rk[49]+delta[1][12], 11)
		rk[65] = bits.RotateLeft32(rk[57]+delta[2][15], 17)
		rk[73] = bits.RotateLeft32(rk[65]+delta[4][13], 3)
		rk[81] = bits.RotateLeft32(rk[73]+delta[5][16], 11)
		rk[89] = bits.RotateLeft32(rk[81]+delta[6][19], 17)
		rk[97] = bits.RotateLeft32(rk[89]+delta[0][17], 3)
		rk[105] = bits.RotateLeft32(rk[97]+delta[1][20], 11)
		rk[113] = bits.RotateLeft32(rk[105]+delta[2][23], 17)
		rk[121] = bits.RotateLeft32(rk[113]+delta[4][21], 3)
		rk[129] = bits.RotateLeft32(rk[121]+delta[5][24], 11)
		rk[137] = bits.RotateLeft32(rk[129]+delta[6][27], 17)
		rk[145] = bits.RotateLeft32(rk[137]+delta[0][25], 3)
		rk[153] = bits.RotateLeft32(rk[145]+delta[1][28], 11)
		rk[161] = bits.RotateLeft32(rk[153]+delta[2][31], 17)
		rk[169] = bits.RotateLeft32(rk[161]+delta[4][29], 3)
		rk[177] = bits.RotateLeft32(rk[169]+delta[5][0], 11)
		rk[185] = bits.RotateLeft32(rk[177]+delta[6][3], 17)

		rk[2] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*2:])+delta[0][2], 6)
		rk[10] = bits.RotateLeft32(rk[2]+delta[1][5], 13)
		rk[18] = bits.RotateLeft32(rk[10]+delta[3][3], 1)
		rk[26] = bits.RotateLeft32(rk[18]+delta[4][6], 6)
		rk[34] = bits.RotateLeft32(rk[26]+delta[5][9], 13)
		rk[42] = bits.RotateLeft32(rk[34]+delta[7][7], 1)
		rk[50] = bits.RotateLeft32(rk[42]+delta[0][10], 6)
		rk[58] = bits.RotateLeft32(rk[50]+delta[1][13], 13)
		rk[66] = bits.RotateLeft32(rk[58]+delta[3][11], 1)
		rk[74] = bits.RotateLeft32(rk[66]+delta[4][14], 6)
		rk[82] = bits.RotateLeft32(rk[74]+delta[5][17], 13)
		rk[90] = bits.RotateLeft32(rk[82]+delta[7][15], 1)
		rk[98] = bits.RotateLeft32(rk[90]+delta[0][18], 6)
		rk[106] = bits.RotateLeft32(rk[98]+delta[1][21], 13)
		rk[114] = bits.RotateLeft32(rk[106]+delta[3][19], 1)
		rk[122] = bits.RotateLeft32(rk[114]+delta[4][22], 6)
		rk[130] = bits.RotateLeft32(rk[122]+delta[5][25], 13)
		rk[138] = bits.RotateLeft32(rk[130]+delta[7][23], 1)
		rk[146] = bits.RotateLeft32(rk[138]+delta[0][26], 6)
		rk[154] = bits.RotateLeft32(rk[146]+delta[1][29], 13)
		rk[162] = bits.RotateLeft32(rk[154]+delta[3][27], 1)
		rk[170] = bits.RotateLeft32(rk[162]+delta[4][30], 6)
		rk[178] = bits.RotateLeft32(rk[170]+delta[5][1], 13)
		rk[186] = bits.RotateLeft32(rk[178]+delta[7][31], 1)

		rk[3] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*3:])+delta[0][3], 11)
		rk[11] = bits.RotateLeft32(rk[3]+delta[1][6], 17)
		rk[19] = bits.RotateLeft32(rk[11]+delta[3][4], 3)
		rk[27] = bits.RotateLeft32(rk[19]+delta[4][7], 11)
		rk[35] = bits.RotateLeft32(rk[27]+delta[5][10], 17)
		rk[43] = bits.RotateLeft32(rk[35]+delta[7][8], 3)
		rk[51] = bits.RotateLeft32(rk[43]+delta[0][11], 11)
		rk[59] = bits.RotateLeft32(rk[51]+delta[1][14], 17)
		rk[67] = bits.RotateLeft32(rk[59]+delta[3][12], 3)
		rk[75] = bits.RotateLeft32(rk[67]+delta[4][15], 11)
		rk[83] = bits.RotateLeft32(rk[75]+delta[5][18], 17)
		rk[91] = bits.RotateLeft32(rk[83]+delta[7][16], 3)
		rk[99] = bits.RotateLeft32(rk[91]+delta[0][19], 11)
		rk[107] = bits.RotateLeft32(rk[99]+delta[1][22], 17)
		rk[115] = bits.RotateLeft32(rk[107]+delta[3][20], 3)
		rk[123] = bits.RotateLeft32(rk[115]+delta[4][23], 11)
		rk[131] = bits.RotateLeft32(rk[123]+delta[5][26], 17)
		rk[139] = bits.RotateLeft32(rk[131]+delta[7][24], 3)
		rk[147] = bits.RotateLeft32(rk[139]+delta[0][27], 11)
		rk[155] = bits.RotateLeft32(rk[147]+delta[1][30], 17)
		rk[163] = bits.RotateLeft32(rk[155]+delta[3][28], 3)
		rk[171] = bits.RotateLeft32(rk[163]+delta[4][31], 11)
		rk[179] = bits.RotateLeft32(rk[171]+delta[5][2], 17)
		rk[187] = bits.RotateLeft32(rk[179]+delta[7][0], 3)

		rk[4] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*4:])+delta[0][4], 13)
		rk[12] = bits.RotateLeft32(rk[4]+delta[2][2], 1)
		rk[20] = bits.RotateLeft32(rk[12]+delta[3][5], 6)
		rk[28] = bits.RotateLeft32(rk[20]+delta[4][8], 13)
		rk[36] = bits.RotateLeft32(rk[28]+delta[6][6], 1)
		rk[44] = bits.RotateLeft32(rk[36]+delta[7][9], 6)
		rk[52] = bits.RotateLeft32(rk[44]+delta[0][12], 13)
		rk[60] = bits.RotateLeft32(rk[52]+delta[2][10], 1)
		rk[68] = bits.RotateLeft32(rk[60]+delta[3][13], 6)
		rk[76] = bits.RotateLeft32(rk[68]+delta[4][16], 13)
		rk[84] = bits.RotateLeft32(rk[76]+delta[6][14], 1)
		rk[92] = bits.RotateLeft32(rk[84]+delta[7][17], 6)
		rk[100] = bits.RotateLeft32(rk[92]+delta[0][20], 13)
		rk[108] = bits.RotateLeft32(rk[100]+delta[2][18], 1)
		rk[116] = bits.RotateLeft32(rk[108]+delta[3][21], 6)
		rk[124] = bits.RotateLeft32(rk[116]+delta[4][24], 13)
		rk[132] = bits.RotateLeft32(rk[124]+delta[6][22], 1)
		rk[140] = bits.RotateLeft32(rk[132]+delta[7][25], 6)
		rk[148] = bits.RotateLeft32(rk[140]+delta[0][28], 13)
		rk[156] = bits.RotateLeft32(rk[148]+delta[2][26], 1)
		rk[164] = bits.RotateLeft32(rk[156]+delta[3][29], 6)
		rk[172] = bits.RotateLeft32(rk[164]+delta[4][0], 13)
		rk[180] = bits.RotateLeft32(rk[172]+delta[6][30], 1)
		rk[188] = bits.RotateLeft32(rk[180]+delta[7][1], 6)

		rk[5] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*5:])+delta[0][5], 17)
		rk[13] = bits.RotateLeft32(rk[5]+delta[2][3], 3)
		rk[21] = bits.RotateLeft32(rk[13]+delta[3][6], 11)
		rk[29] = bits.RotateLeft32(rk[21]+delta[4][9], 17)
		rk[37] = bits.RotateLeft32(rk[29]+delta[6][7], 3)
		rk[45] = bits.RotateLeft32(rk[37]+delta[7][10], 11)
		rk[53] = bits.RotateLeft32(rk[45]+delta[0][13], 17)
		rk[61] = bits.RotateLeft32(rk[53]+delta[2][11], 3)
		rk[69] = bits.RotateLeft32(rk[61]+delta[3][14], 11)
		rk[77] = bits.RotateLeft32(rk[69]+delta[4][17], 17)
		rk[85] = bits.RotateLeft32(rk[77]+delta[6][15], 3)
		rk[93] = bits.RotateLeft32(rk[85]+delta[7][18], 11)
		rk[101] = bits.RotateLeft32(rk[93]+delta[0][21], 17)
		rk[109] = bits.RotateLeft32(rk[101]+delta[2][19], 3)
		rk[117] = bits.RotateLeft32(rk[109]+delta[3][22], 11)
		rk[125] = bits.RotateLeft32(rk[117]+delta[4][25], 17)
		rk[133] = bits.RotateLeft32(rk[125]+delta[6][23], 3)
		rk[141] = bits.RotateLeft32(rk[133]+delta[7][26], 11)
		rk[149] = bits.RotateLeft32(rk[141]+delta[0][29], 17)
		rk[157] = bits.RotateLeft32(rk[149]+delta[2][27], 3)
		rk[165] = bits.RotateLeft32(rk[157]+delta[3][30], 11)
		rk[173] = bits.RotateLeft32(rk[165]+delta[4][1], 17)
		rk[181] = bits.RotateLeft32(rk[173]+delta[6][31], 3)
		rk[189] = bits.RotateLeft32(rk[181]+delta[7][2], 11)

		rk[6] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*6:])+delta[1][1], 1)
		rk[14] = bits.RotateLeft32(rk[6]+delta[2][4], 6)
		rk[22] = bits.RotateLeft32(rk[14]+delta[3][7], 13)
		rk[30] = bits.RotateLeft32(rk[22]+delta[5][5], 1)
		rk[38] = bits.RotateLeft32(rk[30]+delta[6][8], 6)
		rk[46] = bits.RotateLeft32(rk[38]+delta[7][11], 13)
		rk[54] = bits.RotateLeft32(rk[46]+delta[1][9], 1)
		rk[62] = bits.RotateLeft32(rk[54]+delta[2][12], 6)
		rk[70] = bits.RotateLeft32(rk[62]+delta[3][15], 13)
		rk[78] = bits.RotateLeft32(rk[70]+delta[5][13], 1)
		rk[86] = bits.RotateLeft32(rk[78]+delta[6][16], 6)
		rk[94] = bits.RotateLeft32(rk[86]+delta[7][19], 13)
		rk[102] = bits.RotateLeft32(rk[94]+delta[1][17], 1)
		rk[110] = bits.RotateLeft32(rk[102]+delta[2][20], 6)
		rk[118] = bits.RotateLeft32(rk[110]+delta[3][23], 13)
		rk[126] = bits.RotateLeft32(rk[118]+delta[5][21], 1)
		rk[134] = bits.RotateLeft32(rk[126]+delta[6][24], 6)
		rk[142] = bits.RotateLeft32(rk[134]+delta[7][27], 13)
		rk[150] = bits.RotateLeft32(rk[142]+delta[1][25], 1)
		rk[158] = bits.RotateLeft32(rk[150]+delta[2][28], 6)
		rk[166] = bits.RotateLeft32(rk[158]+delta[3][31], 13)
		rk[174] = bits.RotateLeft32(rk[166]+delta[5][29], 1)
		rk[182] = bits.RotateLeft32(rk[174]+delta[6][0], 6)
		rk[190] = bits.RotateLeft32(rk[182]+delta[7][3], 13)

		rk[7] = bits.RotateLeft32(binary.LittleEndian.Uint32(key[4*7:])+delta[1][2], 3)
		rk[15] = bits.RotateLeft32(rk[7]+delta[2][5], 11)
		rk[23] = bits.RotateLeft32(rk[15]+delta[3][8], 17)
		rk[31] = bits.RotateLeft32(rk[23]+delta[5][6], 3)
		rk[39] = bits.RotateLeft32(rk[31]+delta[6][9], 11)
		rk[47] = bits.RotateLeft32(rk[39]+delta[7][12], 17)
		rk[55] = bits.RotateLeft32(rk[47]+delta[1][10], 3)
		rk[63] = bits.RotateLeft32(rk[55]+delta[2][13], 11)
		rk[71] = bits.RotateLeft32(rk[63]+delta[3][16], 17)
		rk[79] = bits.RotateLeft32(rk[71]+delta[5][14], 3)
		rk[87] = bits.RotateLeft32(rk[79]+delta[6][17], 11)
		rk[95] = bits.RotateLeft32(rk[87]+delta[7][20], 17)
		rk[103] = bits.RotateLeft32(rk[95]+delta[1][18], 3)
		rk[111] = bits.RotateLeft32(rk[103]+delta[2][21], 11)
		rk[119] = bits.RotateLeft32(rk[111]+delta[3][24], 17)
		rk[127] = bits.RotateLeft32(rk[119]+delta[5][22], 3)
		rk[135] = bits.RotateLeft32(rk[127]+delta[6][25], 11)
		rk[143] = bits.RotateLeft32(rk[135]+delta[7][28], 17)
		rk[151] = bits.RotateLeft32(rk[143]+delta[1][26], 3)
		rk[159] = bits.RotateLeft32(rk[151]+delta[2][29], 11)
		rk[167] = bits.RotateLeft32(rk[159]+delta[3][0], 17)
		rk[175] = bits.RotateLeft32(rk[167]+delta[5][30], 3)
		rk[183] = bits.RotateLeft32(rk[175]+delta[6][1], 11)
		rk[191] = bits.RotateLeft32(rk[183]+delta[7][4], 17)
		break
	}

	return (keyLen >> 1) + 16
}

func leaEnc1Go(ctx *leaContext, dst, src []byte) {
	X0 := binary.LittleEndian.Uint32(src[4*0:])
	X1 := binary.LittleEndian.Uint32(src[4*1:])
	X2 := binary.LittleEndian.Uint32(src[4*2:])
	X3 := binary.LittleEndian.Uint32(src[4*3:])

	X0, X1, X2, X3 = leaEnc1GoRoundUnder24(ctx, dst, src, X0, X1, X2, X3)

	if ctx.round > 24 {
		X0, X1, X2, X3 = leaEnc1GoRoundOver24(ctx, dst, src, X0, X1, X2, X3)
		if ctx.round > 28 {
			X0, X1, X2, X3 = leaEnc1GoRoundOver28(ctx, dst, src, X0, X1, X2, X3)
		}
	}

	binary.LittleEndian.PutUint32(dst[4*0:], X0)
	binary.LittleEndian.PutUint32(dst[4*1:], X1)
	binary.LittleEndian.PutUint32(dst[4*2:], X2)
	binary.LittleEndian.PutUint32(dst[4*3:], X3)
}
func leaEnc1GoRoundUnder24(ctx *leaContext, dst, src []byte, X0, X1, X2, X3 uint32) (uint32, uint32, uint32, uint32) {
	X3 = bitsRotateRight32((X2^ctx.rk[4])+(X3^ctx.rk[5]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[2])+(X2^ctx.rk[3]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[0])+(X1^ctx.rk[1]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[10])+(X0^ctx.rk[11]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[8])+(X3^ctx.rk[9]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[6])+(X2^ctx.rk[7]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[16])+(X1^ctx.rk[17]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[14])+(X0^ctx.rk[15]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[12])+(X3^ctx.rk[13]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[22])+(X2^ctx.rk[23]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[20])+(X1^ctx.rk[21]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[18])+(X0^ctx.rk[19]), 9)

	X3 = bitsRotateRight32((X2^ctx.rk[28])+(X3^ctx.rk[29]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[26])+(X2^ctx.rk[27]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[24])+(X1^ctx.rk[25]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[34])+(X0^ctx.rk[35]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[32])+(X3^ctx.rk[33]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[30])+(X2^ctx.rk[31]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[40])+(X1^ctx.rk[41]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[38])+(X0^ctx.rk[39]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[36])+(X3^ctx.rk[37]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[46])+(X2^ctx.rk[47]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[44])+(X1^ctx.rk[45]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[42])+(X0^ctx.rk[43]), 9)

	X3 = bitsRotateRight32((X2^ctx.rk[52])+(X3^ctx.rk[53]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[50])+(X2^ctx.rk[51]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[48])+(X1^ctx.rk[49]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[58])+(X0^ctx.rk[59]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[56])+(X3^ctx.rk[57]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[54])+(X2^ctx.rk[55]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[64])+(X1^ctx.rk[65]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[62])+(X0^ctx.rk[63]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[60])+(X3^ctx.rk[61]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[70])+(X2^ctx.rk[71]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[68])+(X1^ctx.rk[69]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[66])+(X0^ctx.rk[67]), 9)

	X3 = bitsRotateRight32((X2^ctx.rk[76])+(X3^ctx.rk[77]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[74])+(X2^ctx.rk[75]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[72])+(X1^ctx.rk[73]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[82])+(X0^ctx.rk[83]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[80])+(X3^ctx.rk[81]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[78])+(X2^ctx.rk[79]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[88])+(X1^ctx.rk[89]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[86])+(X0^ctx.rk[87]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[84])+(X3^ctx.rk[85]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[94])+(X2^ctx.rk[95]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[92])+(X1^ctx.rk[93]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[90])+(X0^ctx.rk[91]), 9)

	X3 = bitsRotateRight32((X2^ctx.rk[100])+(X3^ctx.rk[101]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[98])+(X2^ctx.rk[99]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[96])+(X1^ctx.rk[97]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[106])+(X0^ctx.rk[107]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[104])+(X3^ctx.rk[105]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[102])+(X2^ctx.rk[103]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[112])+(X1^ctx.rk[113]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[110])+(X0^ctx.rk[111]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[108])+(X3^ctx.rk[109]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[118])+(X2^ctx.rk[119]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[116])+(X1^ctx.rk[117]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[114])+(X0^ctx.rk[115]), 9)

	X3 = bitsRotateRight32((X2^ctx.rk[124])+(X3^ctx.rk[125]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[122])+(X2^ctx.rk[123]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[120])+(X1^ctx.rk[121]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[130])+(X0^ctx.rk[131]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[128])+(X3^ctx.rk[129]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[126])+(X2^ctx.rk[127]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[136])+(X1^ctx.rk[137]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[134])+(X0^ctx.rk[135]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[132])+(X3^ctx.rk[133]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[142])+(X2^ctx.rk[143]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[140])+(X1^ctx.rk[141]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[138])+(X0^ctx.rk[139]), 9)

	return X0, X1, X2, X3
}
func leaEnc1GoRoundOver24(ctx *leaContext, dst, src []byte, X0, X1, X2, X3 uint32) (uint32, uint32, uint32, uint32) {
	X3 = bitsRotateRight32((X2^ctx.rk[148])+(X3^ctx.rk[149]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[146])+(X2^ctx.rk[147]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[144])+(X1^ctx.rk[145]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[154])+(X0^ctx.rk[155]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[152])+(X3^ctx.rk[153]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[150])+(X2^ctx.rk[151]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[160])+(X1^ctx.rk[161]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[158])+(X0^ctx.rk[159]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[156])+(X3^ctx.rk[157]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[166])+(X2^ctx.rk[167]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[164])+(X1^ctx.rk[165]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[162])+(X0^ctx.rk[163]), 9)

	return X0, X1, X2, X3
}
func leaEnc1GoRoundOver28(ctx *leaContext, dst, src []byte, X0, X1, X2, X3 uint32) (uint32, uint32, uint32, uint32) {
	X3 = bitsRotateRight32((X2^ctx.rk[172])+(X3^ctx.rk[173]), 3)
	X2 = bitsRotateRight32((X1^ctx.rk[170])+(X2^ctx.rk[171]), 5)
	X1 = bits.RotateLeft32((X0^ctx.rk[168])+(X1^ctx.rk[169]), 9)
	X0 = bitsRotateRight32((X3^ctx.rk[178])+(X0^ctx.rk[179]), 3)
	X3 = bitsRotateRight32((X2^ctx.rk[176])+(X3^ctx.rk[177]), 5)
	X2 = bits.RotateLeft32((X1^ctx.rk[174])+(X2^ctx.rk[175]), 9)
	X1 = bitsRotateRight32((X0^ctx.rk[184])+(X1^ctx.rk[185]), 3)
	X0 = bitsRotateRight32((X3^ctx.rk[182])+(X0^ctx.rk[183]), 5)
	X3 = bits.RotateLeft32((X2^ctx.rk[180])+(X3^ctx.rk[181]), 9)
	X2 = bitsRotateRight32((X1^ctx.rk[190])+(X2^ctx.rk[191]), 3)
	X1 = bitsRotateRight32((X0^ctx.rk[188])+(X1^ctx.rk[189]), 5)
	X0 = bits.RotateLeft32((X3^ctx.rk[186])+(X0^ctx.rk[187]), 9)

	return X0, X1, X2, X3
}

func leaDec1Go(ctx *leaContext, dst, src []byte) {
	X0 := binary.LittleEndian.Uint32(src[4*0:])
	X1 := binary.LittleEndian.Uint32(src[4*1:])
	X2 := binary.LittleEndian.Uint32(src[4*2:])
	X3 := binary.LittleEndian.Uint32(src[4*3:])

	if ctx.round > 28 {
		X0, X1, X2, X3 = leaDec1GoRoundOver28(ctx, dst, src, X0, X1, X2, X3)
	}

	if ctx.round > 24 {
		X0, X1, X2, X3 = leaDec1GoRoundOver24(ctx, dst, src, X0, X1, X2, X3)
	}

	X0, X1, X2, X3 = leaDec1GoRoundUnder24(ctx, dst, src, X0, X1, X2, X3)

	binary.LittleEndian.PutUint32(dst[4*0:], X0)
	binary.LittleEndian.PutUint32(dst[4*1:], X1)
	binary.LittleEndian.PutUint32(dst[4*2:], X2)
	binary.LittleEndian.PutUint32(dst[4*3:], X3)
}
func leaDec1GoRoundOver28(ctx *leaContext, dst, src []byte, X0, X1, X2, X3 uint32) (uint32, uint32, uint32, uint32) {
	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[186])) ^ ctx.rk[187]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[188])) ^ ctx.rk[189]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[190])) ^ ctx.rk[191]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[180])) ^ ctx.rk[181]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[182])) ^ ctx.rk[183]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[184])) ^ ctx.rk[185]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[174])) ^ ctx.rk[175]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[176])) ^ ctx.rk[177]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[178])) ^ ctx.rk[179]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[168])) ^ ctx.rk[169]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[170])) ^ ctx.rk[171]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[172])) ^ ctx.rk[173]

	return X0, X1, X2, X3
}
func leaDec1GoRoundOver24(ctx *leaContext, dst, src []byte, X0, X1, X2, X3 uint32) (uint32, uint32, uint32, uint32) {
	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[162])) ^ ctx.rk[163]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[164])) ^ ctx.rk[165]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[166])) ^ ctx.rk[167]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[156])) ^ ctx.rk[157]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[158])) ^ ctx.rk[159]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[160])) ^ ctx.rk[161]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[150])) ^ ctx.rk[151]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[152])) ^ ctx.rk[153]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[154])) ^ ctx.rk[155]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[144])) ^ ctx.rk[145]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[146])) ^ ctx.rk[147]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[148])) ^ ctx.rk[149]

	return X0, X1, X2, X3
}
func leaDec1GoRoundUnder24(ctx *leaContext, dst, src []byte, X0, X1, X2, X3 uint32) (uint32, uint32, uint32, uint32) {
	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[138])) ^ ctx.rk[139]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[140])) ^ ctx.rk[141]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[142])) ^ ctx.rk[143]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[132])) ^ ctx.rk[133]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[134])) ^ ctx.rk[135]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[136])) ^ ctx.rk[137]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[126])) ^ ctx.rk[127]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[128])) ^ ctx.rk[129]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[130])) ^ ctx.rk[131]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[120])) ^ ctx.rk[121]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[122])) ^ ctx.rk[123]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[124])) ^ ctx.rk[125]

	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[114])) ^ ctx.rk[115]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[116])) ^ ctx.rk[117]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[118])) ^ ctx.rk[119]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[108])) ^ ctx.rk[109]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[110])) ^ ctx.rk[111]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[112])) ^ ctx.rk[113]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[102])) ^ ctx.rk[103]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[104])) ^ ctx.rk[105]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[106])) ^ ctx.rk[107]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[96])) ^ ctx.rk[97]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[98])) ^ ctx.rk[99]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[100])) ^ ctx.rk[101]

	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[90])) ^ ctx.rk[91]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[92])) ^ ctx.rk[93]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[94])) ^ ctx.rk[95]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[84])) ^ ctx.rk[85]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[86])) ^ ctx.rk[87]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[88])) ^ ctx.rk[89]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[78])) ^ ctx.rk[79]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[80])) ^ ctx.rk[81]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[82])) ^ ctx.rk[83]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[72])) ^ ctx.rk[73]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[74])) ^ ctx.rk[75]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[76])) ^ ctx.rk[77]

	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[66])) ^ ctx.rk[67]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[68])) ^ ctx.rk[69]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[70])) ^ ctx.rk[71]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[60])) ^ ctx.rk[61]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[62])) ^ ctx.rk[63]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[64])) ^ ctx.rk[65]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[54])) ^ ctx.rk[55]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[56])) ^ ctx.rk[57]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[58])) ^ ctx.rk[59]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[48])) ^ ctx.rk[49]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[50])) ^ ctx.rk[51]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[52])) ^ ctx.rk[53]

	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[42])) ^ ctx.rk[43]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[44])) ^ ctx.rk[45]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[46])) ^ ctx.rk[47]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[36])) ^ ctx.rk[37]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[38])) ^ ctx.rk[39]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[40])) ^ ctx.rk[41]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[30])) ^ ctx.rk[31]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[32])) ^ ctx.rk[33]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[34])) ^ ctx.rk[35]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[24])) ^ ctx.rk[25]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[26])) ^ ctx.rk[27]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[28])) ^ ctx.rk[29]

	X0 = (bitsRotateRight32(X0, 9) - (X3 ^ ctx.rk[18])) ^ ctx.rk[19]
	X1 = (bits.RotateLeft32(X1, 5) - (X0 ^ ctx.rk[20])) ^ ctx.rk[21]
	X2 = (bits.RotateLeft32(X2, 3) - (X1 ^ ctx.rk[22])) ^ ctx.rk[23]
	X3 = (bitsRotateRight32(X3, 9) - (X2 ^ ctx.rk[12])) ^ ctx.rk[13]
	X0 = (bits.RotateLeft32(X0, 5) - (X3 ^ ctx.rk[14])) ^ ctx.rk[15]
	X1 = (bits.RotateLeft32(X1, 3) - (X0 ^ ctx.rk[16])) ^ ctx.rk[17]
	X2 = (bitsRotateRight32(X2, 9) - (X1 ^ ctx.rk[6])) ^ ctx.rk[7]
	X3 = (bits.RotateLeft32(X3, 5) - (X2 ^ ctx.rk[8])) ^ ctx.rk[9]
	X0 = (bits.RotateLeft32(X0, 3) - (X3 ^ ctx.rk[10])) ^ ctx.rk[11]
	X1 = (bitsRotateRight32(X1, 9) - (X0 ^ ctx.rk[0])) ^ ctx.rk[1]
	X2 = (bits.RotateLeft32(X2, 5) - (X1 ^ ctx.rk[2])) ^ ctx.rk[3]
	X3 = (bits.RotateLeft32(X3, 3) - (X2 ^ ctx.rk[4])) ^ ctx.rk[5]

	return X0, X1, X2, X3
}
