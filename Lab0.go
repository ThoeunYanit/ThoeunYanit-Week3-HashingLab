package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/sha3"
)

func main(){
	// data := []byte("hello")
	// fmt.Printf("%x\n", md5.Sum(data))
	// fmt.Printf("%x\n", sha1.Sum(data))
	// fmt.Printf("%x\n", sha256.Sum256(data))
	// fmt.Printf("%x\n", sha512.Sum512(data))
	// fmt.Printf("%x\n", sha3.Sum256(data))

	var value1 string
	var value2 string

	fmt.Println("=====Name + Hashing Program=====")
	fmt.Print("Please input value 1: ")
	fmt.Scanln(&value1)

	fmt.Print("Please input value 2: ")
	fmt.Scanln(&value2)

	// convert value into byte, md5 return 16 bytes
	fmt.Printf("\nMD5:\nOutput A = %x\nOutput B = %x\n=> %s\n\n", md5.Sum([]byte(value1)), md5.Sum([]byte(value2)), compareMD5(md5.Sum([]byte(value1)), md5.Sum([]byte(value2))))
	
	fmt.Printf("SHA1:\nOutput A = %x\nOutput B = %x\n=> %s\n\n", sha1.Sum([]byte(value1)), sha1.Sum([]byte(value2)), compareSHA1(sha1.Sum([]byte(value1)), sha1.Sum([]byte(value2))))
	
	fmt.Printf("SHA256:\nOutput A = %x\nOutput B = %x\n=> %s\n\n", sha256.Sum256([]byte(value1)), sha256.Sum256([]byte(value2)), compareSHA256(sha256.Sum256([]byte(value1)), sha256.Sum256([]byte(value2))))
	
	fmt.Printf("SHA512:\nOutput A = %x\nOutput B = %x\n=> %s\n\n", sha512.Sum512([]byte(value1)), sha512.Sum512([]byte(value2)), compareSHA512(sha512.Sum512([]byte(value1)), sha512.Sum512([]byte(value2))))
	
	fmt.Printf("SHA3-256:\nOutput A = %x\nOutput B = %x\n=> %s\n\n", sha3.Sum256([]byte(value1)), sha3.Sum256([]byte(value2)), compareSHA3_256(sha3.Sum256([]byte(value1)), sha3.Sum256([]byte(value2))))
	
}

func compareMD5(value1, value2 [16]byte) string{
	if (value1 == value2){
		return "Match!"
	}
	return "No Match!"
}

func compareSHA1(value1, value2 [20]byte) string{
	if (value1 == value2){
		return "Match!"
	}
	return "No Match!"
}

func compareSHA256(value1, value2 [32]byte) string{
	if (value1 == value2){
		return "Match!"
	}
	return "No Match!"
}

func compareSHA512(value1, value2 [64]byte) string{
	if (value1 == value2){
		return "Match!"
	}
	return "No Match!"
}

func compareSHA3_256(value1, value2 [32]byte) string{
	if (value1 == value2){
		return "Match!"
	}
	return "No Match!"
}

