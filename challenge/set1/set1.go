package challenge

import (
	"crypto/aes"
	utils "cryptoGo/util"
	"fmt"
)

type fn func([]byte, []byte) []byte

func ConvertHexToBase64(hexstr string) (string, error) {
	_, dst := utils.HexToByte(hexstr)

	str, e := utils.ByteToBase64(dst)
	if e != nil {
		return str, e
	}
	return str, e
}

func FixedXOR(str string, key string) (string, error) {
	_, dsts := utils.HexToByte(str)
	_, dstk := utils.HexToByte(key)

	val := utils.XOR(dsts, dstk)

	rv := utils.ByteToHexString(val)
	return rv, nil
}

func GetMostLikelyKeySingleByte(dsts []byte, converterFunc fn) ([]byte, error) {
	weight := 0
	rv := make([]byte, 1)

	// find max byte val in map
	for k := 0; k < 256; k++ {
		b := []byte{byte(k)}
		res := converterFunc(dsts, b)
		if v := utils.GetWeights(res); v > weight {
			weight = v
			rv = b
		}
	}
	return rv, nil
}

func SingleByteXOR(dsts []byte) ([]byte, error) {
	weight := 0
	rv := make([]byte, len(dsts))

	// find max byte val in map
	for k := 0; k < 256; k++ {
		b := []byte{byte(k)}
		res := utils.XOR(dsts, b)
		// compare the weights and find the min
		if v := utils.GetWeights(res); v > weight {
			weight = v
			rv = res
		}
	}
	return rv, nil
}

func SingleByteXORHexStr(str string) (string, error) {
	_, dsts := utils.HexToByte(str)
	rv, e := SingleByteXOR(dsts)
	return string(rv), e
}

func FindSingleByteXOR(fp string) (string, error) {
	list := utils.ReadFile(fp)

	weight := 0
	rv := make([]byte, len(list))
	for _, v := range list {

		res, _ := SingleByteXORHexStr(v)
		b := []byte(res)
		if v := utils.GetWeights(b); v > weight {
			weight = v
			rv = b
		}
	}
	return string(rv), nil
}

func RepeatingKeyXOR(valueB []byte, keyB []byte) ([]byte, error) {
	rv := []byte{}
	offset := 0
	repeatKeyB := []byte{}
	// get repeated key
	for i := 0; i < len(valueB); i++ {
		repeatKeyB = append(repeatKeyB, keyB[(offset+i)%len(keyB)])
	}

	res := utils.XOR(valueB, repeatKeyB)

	rv = append(rv, res...)
	offset += len(valueB)
	return rv, nil
}

func RepeatingKeyXORString(list string, key string) (string, error) {
	keyB := []byte(key)
	valueB := []byte(list)
	res, _ := RepeatingKeyXOR(valueB, keyB)
	return utils.ByteToHexString(res), nil
}

func FindMinKeysize(value []byte, minKey int, maxKey int, sampleCount int) int {
	// get a big num by flipping bits of 0, since int is arch dep, use 30 to not get signed int
	min := 1 << 30
	minKeySize := minKey
	for keySize := minKey; keySize <= maxKey; keySize++ {
		// if we cant get an equal length amount of samples break, this way we always have equal length compares
		if len(value)/keySize < sampleCount {
			break
		}
		testCases := [][]byte{}
		for caseNum := 0; caseNum < sampleCount; caseNum++ {
			end := utils.GetEnd(keySize, caseNum*keySize, len(value))
			testCases = append(testCases, value[keySize*caseNum:end])
		}

		hammAgg := 0
		tested := 0
		for j := 0; j < len(testCases)-1; j++ {
			for k := j + 1; k < len(testCases); k++ {
				// normalize by keysize
				hammAgg += utils.HammingDistance(testCases[j], testCases[k]) / keySize
				tested++
			}
		}
		if newHamm := hammAgg / tested; tested > 0 && newHamm < min {
			min = newHamm
			minKeySize = keySize
		}
	}

	return minKeySize
}

func BreakIntoKeySize(value []byte, keysize int) [][]byte {
	rv := [][]byte{}
	for i := 0; i < len(value); i += keysize {
		end := utils.GetEnd(keysize, i, len(value))
		rv = append(rv, value[i:end])
	}
	return rv
}

func TransposeByteBlocks(value [][]byte) [][]byte {
	rv := make([][]byte, len(value[0]))
	for i := 0; i < len(value[0]); i++ {
		rv[i] = make([]byte, len(value))
		for j := 0; j < len(value); j++ {
			rv[i][j] = value[j][i]
		}
	}
	return rv
}

func BreakAndTransposeToBlocks(barr []byte, keySize int) [][]byte {
	barrs := BreakIntoKeySize(barr, keySize)
	// trim off extra values that aren't the same length
	if len(barrs[0]) != len(barrs[len(barrs)-1]) {
		barrs = barrs[:len(barrs)-1]
	}

	tarrs := TransposeByteBlocks(barrs)
	return tarrs
}

func BreakRepeatingXOR(byteArr []byte, minKey int, maxKey int, matches int) ([]byte, []byte) {
	keySize := FindMinKeysize(byteArr, minKey, maxKey, matches)
	transposedarr := BreakAndTransposeToBlocks(byteArr, keySize)
	keyByteArr := make([]byte, keySize)
	for i, transposed := range transposedarr {
		res, _ := GetMostLikelyKeySingleByte(transposed, utils.XOR)
		keyByteArr[i] = res[0]
	}

	rep, _ := RepeatingKeyXOR(byteArr, keyByteArr)

	return rep, keyByteArr
}

func IterateAndFindBestMatch(b []byte) ([]byte, []byte) {
	wordWeight := 0
	guesslen := 2
	maxGuesslen := 60
	key := []byte{}
	rv := []byte{}
	for i := guesslen; i < maxGuesslen; i++ {
		res, k := BreakRepeatingXOR(b, i, maxGuesslen, 3)
		if v := utils.GetWeights(res); v > wordWeight {
			wordWeight = v
			rv = res
			key = k
		}
	}
	return rv, key
}

func DecryptAES128ECB(input []byte, key []byte) []byte {
	// using ecb so use the key length as block size, should be around 16bytes
	block, e := aes.NewCipher(key)
	if e != nil {
		panic(e.Error())
	}
	// if len(input)%block.BlockSize() != 0 {
	// 	panic("block size and input are not the same length")
	// }
	keysize := block.BlockSize()
	decrypted := []byte{}
	for i := 0; i < len(input); i += keysize {
		b := make([]byte, block.BlockSize())
		block.Decrypt(b, input[i:i+keysize])
		decrypted = append(decrypted, b...)
	}
	return decrypted
}

func DecryptAES128Base64(input []byte, key []byte) []byte {
	bytes, _ := utils.Base64ByteToByte(input)
	rv := DecryptAES128ECB(bytes, key)
	return rv
}

func SingleByteAES128ECB(input []byte, key []byte) []byte {
	keyLong := make([]byte, 16)
	for i := 0; i < 16; i++ {
		keyLong[i] = key[0]
	}
	rv := DecryptAES128ECB(input, keyLong)
	return rv
}

func IdentifyAES128HexFromList(input []string, keylen int) []byte {
	rv := make([]byte, len(input[0]))
	for _, str := range input {
		_, bytes := utils.HexToByte(str)
		freqMap := utils.FindByteFrequency(bytes)
		// find the best fit to 
		// https://en.wikipedia.org/wiki/Letter_frequency
	}

	return rv
}
