package challenge

import (
	challenge "cryptoGo/challenge/set1"
	utils "cryptoGo/util"
	"fmt"
	"strings"
	"testing"
)

func TestChallenge1(t *testing.T) {
	testValue := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	res, _ := challenge.ConvertHexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if res != testValue {
		t.Error("Expected ", testValue, " but got", res)
	}
}

func TestChallenge2(t *testing.T) {
	str := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	exp := "746865206b696420646f6e277420706c6179"
	res, _ := challenge.FixedXOR(str, key)
	if res != exp {
		t.Error("Expected ", exp, " but got", res)
	}
}

func TestChallenge3(t *testing.T) {
	str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	exp := "Cooking MC's like a pound of bacon"
	res, _ := challenge.SingleByteXORHexStr(str)
	if res != exp {
		t.Error("Expected ", exp, " but got", res)
	}
	fmt.Printf("out-> %s\n", res)
}

func TestChallenge4(t *testing.T) {
	exp := "Now that the party is jumping\n"
	res, _ := challenge.FindSingleByteXOR("../data/challenge4.txt")
	if res != exp {
		t.Error("Expected ", exp, " but got", res)
	}
	fmt.Printf("out-> %s", res)
}

func TestChallenge5(t *testing.T) {
	str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	exp := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	key := "ICE"
	res, _ := challenge.RepeatingKeyXORString(str, key)
	if res != exp {
		t.Error("Expected \n", exp, "\n but got\n", res)
	}
	fmt.Printf("out-> %s\n", res)
}

func TestHammingDist(t *testing.T) {
	exp := 37
	str1 := "this is a test"
	str2 := "wokka wokka!!!"
	res := utils.HammingDistanceStr(str1, str2)
	if res != exp {
		t.Error("Expected \n", exp, "\n but got\n", res)
	}
}

func TestFindMinKeysize(t *testing.T) {
	barr := []byte("zxcvbnmasdfzxcvbnmasdfxxcvbnmasdf")
	exp := 11
	res := challenge.FindMinKeysize(barr, 2, 40, 3)
	if res != exp {
		t.Error("Expected ", exp, "\n but got ", res)
	}
}

func TestBreakIntoKeySize(t *testing.T) {
	barr := []byte("123123123")
	exp := [][]byte{[]byte("123"), []byte("123"), []byte("123")}
	res := challenge.BreakIntoKeySize(barr, 3)

	for k, varr := range res {
		for i, v := range varr {
			if v != exp[k][i] {
				t.Error("Expected ", exp, "\n but got ", res)
			}
		}
	}
}

func TestTransposeByteBlocks(t *testing.T) {
	barr := [][]byte{[]byte("1234"), []byte("1234"), []byte("1234")}
	exp := [][]byte{[]byte("111"), []byte("222"), []byte("333"), []byte("444")}
	res := challenge.TransposeByteBlocks(barr)

	for k, varr := range res {
		for i, v := range varr {
			if v != exp[k][i] {
				t.Error("Expected ", exp, "\n but got ", res)
			}
		}
	}
	fmt.Printf("out-> %d\n", res)
}

func TestBreakAndTranspose(t *testing.T) {
	barr := []byte("123412341234")
	exp := [][]byte{[]byte("111"), []byte("222"), []byte("333"), []byte("444")}
	res := challenge.BreakAndTransposeToBlocks(barr, 4)

	for k, varr := range res {
		for i, v := range varr {
			if v != exp[k][i] {
				t.Error("Expected ", exp, "\n but got ", res)
			}
		}
	}
	fmt.Printf("out-> %d\n", res)
}

func TestChallenge6Simple(t *testing.T) {
	str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	_, exp := utils.HexToByte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	res, _ := challenge.BreakRepeatingXOR(exp, 2, 40, 3)
	if string(res) != str {
		t.Error("\n\tExpected\t", str, "\n\tbut got\t", string(res))
	}
}

func TestChallenge6(t *testing.T) {
	loadedFile := utils.ReadFile("../data/challenge6.txt")
	joinStr := strings.Join(loadedFile, "")
	byteArr, _ := utils.Base64ToByte(joinStr)
	res, key := challenge.IterateAndFindBestMatch(byteArr)
	fmt.Printf("out-> %s\t%s\n", res, key)
}

func TestChallenge7(t *testing.T) {
	loadedFile := utils.ReadFile("../data/challenge7.txt")
	joinStr := strings.Join(loadedFile, "")
	challenge.DecryptAES128Base64([]byte(joinStr), []byte("YELLOW SUBMARINE"))
}
