// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Electronic Code Book (ECB) mode.

// ECB provides confidentiality by assigning a fixed ciphertext block to each
// plaintext block.

// See NIST SP 800-38A, pp 08-09

package ecb

import (
	"crypto/cipher"
	"crypto/aes"
	"encoding/base64"
)

func PaddingData(data []byte,blockSize int) []byte {
	size := len(data)

	nRemain := size % blockSize
    nBlocks := (size + blockSize - 1) / blockSize

    if (nRemain > 12 || nRemain == 0) {
        nBlocks += 1;
    }
    newSize := nBlocks * blockSize
    outData := make([]byte,newSize)
    end := newSize - 4
    outData[0 + end] = byte(size >> 24)
    outData[1 + end] = byte(size >> 16)
    outData[2 + end] = byte(size >> 8)
    outData[3 + end] = byte(size)

    for i:=0 ;i < size ; i++ {
    	outData[i] = data[i]
    }
    return outData
}

func RemoveEndPadding(data []byte) []byte {
	size := len(data)
	outData := make([]byte,0)
	for i := size-1 ; i >=0 ; i-- {
		if data[i] == 0 {
			continue
		} else {
			if len(outData) == 0 {
				outData = make([]byte,i+1)
			}
		}
		outData[i] = data[i]
	}
	return outData
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}


type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}


type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) []byte{
	dest,_ := base64.StdEncoding.DecodeString(data)
	return dest
}

func EncryptData(src,key []byte) []byte {
	aesCipher, _ := aes.NewCipher(key)
	blockSize := aesCipher.BlockSize()
	src = PaddingData(src, blockSize)
	dest := make([]byte, len(src))
	encrypter := NewECBEncrypter(aesCipher)
	encrypter.CryptBlocks(dest, src)
	return RemoveEndPadding(dest)
}

func DecryptData(src,key []byte) []byte {
	aesCipher, _ := aes.NewCipher(key)
	dest := make([]byte, len(src))
	encrypter := NewECBDecrypter(aesCipher)
	encrypter.CryptBlocks(dest, src)
	//dest = RemoveEndPadding(dest)
	return RemoveEndPadding(dest[0:len(dest) - 4])	
}
