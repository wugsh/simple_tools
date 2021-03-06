package main

/*
* Using Go binary file and encryption method to prevent Python code leakage.
* Encrypting Python files --> Decrypting Python files --> Runing Python files --> Delete Decrypted Python files
* Only Go binary and Encrypted Python files are left at run time and stop time.
*/

import (
    "fmt"
    "io/ioutil"
	"os"
	"os/exec"
    "crypto/cipher"
    "crypto/des"
    "crypto/sha256"
	"encoding/base64"
	"time"
)

//File encryption
func EncryptFile(fileName string, writfileName string, key []byte) (int, error) {
	file, err:=os.Open(fileName)
	writfile, err1:=os.Create(writfileName)
    if err!=nil {
        fmt.Println("File not found")
        return -1, err
	}
	if err1!=nil {
		fmt.Println("File not found")
		return -1, err1
	}
	defer file.Close()
	defer writfile.Close()
    //Read file contents
    plain,_:=ioutil.ReadAll(file)
    //Create block
	block,_:=des.NewTripleDESCipher(key)
	
	EncryptMode:=cipher.NewCBCEncrypter(block,key[:8])  
    //Plaintext complement
    plain=PKCS5append(plain)
    EncryptMode.CryptBlocks(plain,plain)
    err2:= ioutil.WriteFile(writfileName, []byte(base64.StdEncoding.EncodeToString(plain)), 0777)
    if err2!=nil{
		fmt.Println("Failed to save encrypted file")
		return -1, err2
    }else{
		fmt.Println("The file is encrypted. Remember to encrypt the key")
		return 0, nil
    }
}


//Decrypt files
func DecryptFile(fileName string, writfileName string, key []byte) (int, error) {
	file,err:=os.Open(fileName)
	writfile, err1:=os.Create(writfileName)
    if err!=nil {
		fmt.Println("File not found")
        return -1, err
	}
	if err1!=nil {
		return -1, err1
	}
	defer file.Close()
	defer writfile.Close()
    //Read file contents
    plain,_:=ioutil.ReadAll(file)
    //Create block
	block,_:=des.NewTripleDESCipher(key)
	DecryptMode:=cipher.NewCBCDecrypter(block,key[:8])
    plain,_=base64.StdEncoding.DecodeString(string(plain))
    DecryptMode.CryptBlocks(plain,plain)
    plain=PKCS5remove(plain)    
    err2 := ioutil.WriteFile(writfileName,plain,0777)
    if err2!=nil{
		fmt.Println("Failed to save decrypted file")
		return -1, err2
    }else{
		fmt.Println("File decrypted")
		return 0, nil
    }
}

 
func PKCS5append(plaintext []byte) []byte {
    num := 8 - len(plaintext)%8
    for i:=0;i<num;i++{
        plaintext=append(plaintext,byte(num))
    }
    return plaintext
}
 

func PKCS5remove(plaintext []byte) []byte {
    length := len(plaintext)
    num := int(plaintext[length-1])
    return plaintext[:(length - num)]
}


func main() {
    var PassKey string = "***********************"
    //Generate key
    PassKeyByte :=sha256.Sum224([]byte(PassKey))
	key :=PassKeyByte[:24]
	num1, _:= EncryptFile("PrimaryFile.py", "EncryptFile.py", key)
	num2, _ := DecryptFile("EncryptFile.py", "DecryptFile.py", key)
	//Using a goroutine for task processing
	go RunPython()
	//Wait for the task to finish processing
	time.Sleep(time.Second * 1)
	os.Remove("DecryptFile.py")
	fmt.Println(num1, num2)
}

//Execute Python program
func RunPython() {
	cmd := exec.Command("python3", "DecryptFile.py")
	err := cmd.Start()
	if err!=nil{
        fmt.Println(err)
        os.Remove("DecryptFile.py")
    } 
	err3 := cmd.Wait()
	if err3 !=nil {
		fmt.Println(err3)
	} else {
		fmt.Println("succeeded")
	}
    os.Remove("DecryptFile.py")
}
