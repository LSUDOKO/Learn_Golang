package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 	f, err := os.Open("example.txt")
	// 	if err != nil {
	// 		//log the error
	// 		panic(err)
	// 	}
	// 	fileinfo, err := f.Stat()
	// 	if err != nil {
	// 		//log the error
	// 		panic(err)
	// 	}
	// 	fmt.Println("file name is:", fileinfo.Name())
	// 	fmt.Println("", fileinfo.Size())
	// 	fmt.Println("check file or folder",fileinfo.IsDir())
	// 	fmt.Println("mode of file",fileinfo.Mode())
	// 	fmt.Println("last modified time",fileinfo.ModTime())

	// f,err:=os.Open("example.txt")
	// if err!=nil{
	// 	panic(err)
	// }

	// defer f.Close()
	// buff:=make([]byte,12)
	// d,err:=f.Read(buff)
	// if err!=nil{
	// 	panic(err)
	// }
	// for i:=0;i<len(buff);i++{
	// 	println(string(buff[i]))
	// }
	// fmt.Println("data",d,buff)
	//
	//read folders
	//
	//write a file
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	er:=os.Remove("example2.txt")
	if er !=nil{
		panic(er)
	}
	fmt.Println("file deleted")
	// // f.WriteString("hello")
	// // f.WriteString("world")
	// bytes:=[]byte("hello world")
	// f.Write(bytes)
	// //read and write to another file(stream fashion)
	// sou, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sou.Close()
	// destfile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destfile.Close()
	// reader := bufio.NewReader(sou)
	// writer := bufio.NewWriter(destfile)
	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("Retured to new file sucessfully")
	//delete file
	

}
