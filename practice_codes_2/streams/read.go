package main

import (
	"fmt"
	"io"
)

//create data holder and counter struct
type DataObject struct {
	str       string // holds the actual string data
	readIndex int    // counter to read next byte
}

func (dataObj *DataObject) Read(p []byte) (n int, err error) {
	//first convert string to bytes array

	strByte := []byte(dataObj.str)

	// check if readIndex is GTE to total size of string, if so return 0 and io.EOF
	if dataObj.readIndex >= len(strByte) {
		return 0, io.EOF
	}

	//get next readable limit
	newReadLimit := dataObj.readIndex + len(p)

	// check if readIndex+new data sixe is GTE size of string, is so set new index to new size and return io.EOF
	if newReadLimit >= len(strByte) {
		newReadLimit = len(strByte)
		err = io.EOF
	}

	// get the next bytes data
	nextByte := strByte[dataObj.readIndex:newReadLimit]
	n = len(nextByte)

	//copy the data to p
	copy(p, nextByte)
	//set the correct counter
	dataObj.readIndex = newReadLimit

	return
}

func main() {
	rawData := DataObject{
		str: "My Name is Rajesh Kumar Yadav ",
		//readIndex: 0,
	}
	//cretae a packet
	p := make([]byte, 3)

	//call the read method
	for {
		n, err := rawData.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occured!", err)
			break
		}
	}
}
