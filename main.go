package main

import (
	"fmt"
	"log"
//	"io/ioutil"
	"os"
//	"strings"
)
var dic []byte
var dic_2 [][]byte
func check(e error) {
    if e != nil {
			log.Fatal(e)
			}
}

func main() {
	// fin, err := os.Open("input.txt");
	// check(err)
	// defer fin.Close()

//	fout, err := os.Open("output.txt");
//	check(err)
//	defer fout.Close()

	dic = fill_in_dic(dic)
	fmt.Println(dic)
	dic_2 = add_onechar_lines(dic)
	fmt.Println("dictionary=  ", dic_2)
	var resultatik []int
	resultatik = compress(dic_2)
	fmt.Println("message=  ", resultatik)

//	str := string(dic)
// twoD := make([][]int, 3)
//var twoD [][]int



}

func byte_in_slice(arr []byte, ch byte) bool{
	var result = false
//	var id = 0
	for i := range arr {
		if arr[i] == ch {
			result = true
			break
		}
	}
	return result
}
func byte_in_dbl_slice(arr [][]byte, ch []byte) (bool, int){
	var result = false
	var hlp = false
	var id int
	for i := range arr {
		if !hlp {
			if len(arr[i]) == len(ch){
			hlp = true
			id = i
				for j := range ch {
					if arr[i][j] != ch[j]{
						hlp = false
						break
					}
				}
			}
		} else {result = true}
	}
	return result, id
}


func byte_id_slice(arr []byte, ch byte) int{
	var result = 0
//	var id = 0
	for i := range arr {
		if arr[i] == ch {
			result = i
			break
		}
	}
	return result
}

func fill_in_dic(dict []byte) []byte {
	char := make([]byte, 1)
	fin, err := os.Open("input.txt");
	check(err)
	defer fin.Close()
	stat, err := fin.Stat()
	check(err)
	for i := 0; i < int(stat.Size()); i++ {
		_, err := fin.Read(char)
		check(err)
		if !byte_in_slice(dict, char[0]){
			dict = append(dict, char[0])
		}
	}
	return dict
}
//главный движ
func compress(dict [][]byte) (message []int) {
	char := make([]byte, 1)
	next_char := make([]byte, 1)
	var curent_line []byte
	var hlp_line []byte
	var id int
	var bl bool
	fin, err := os.Open("input.txt");
	check(err)
	defer fin.Close()
	stat, err := fin.Stat()
	check(err)

	_, err = fin.Read(next_char)
	check(err)
	curent_line = append(curent_line, next_char[0])

	for i := 1; i < int(stat.Size()); i++ {
		fmt.Println("------------------------------")
		_, err := fin.Read(char)
		check(err)
		hlp_line = append(curent_line, char[0])

		bl, id = byte_in_dbl_slice(dict, hlp_line)
		fmt.Println("id ", id, " bl ", bl, " cl ", string(curent_line), " ch ", string(char[0]), " mess ", message)
		if bl {
			curent_line = append (curent_line, char[0])
		} else {
			bl, id = byte_in_dbl_slice(dict, curent_line)
			message = append(message, id)
//			dict = append(dict, []byte{})
//			dict[len(dict)-1] = curent_line
			line := make([]byte, len(hlp_line))
			copy(line, hlp_line)
			//fmt.Println("line  ",line)
			//fmt.Println("Cline  ",curent_line)
			dict = append(dict, line)
			curent_line = curent_line[:0]
			curent_line = append(curent_line, char[0])
//			bl, id = byte_in_dbl_slice(dict, curent_line)
//			message = append(message, id)
		}
		bl, id = byte_in_dbl_slice(dict, curent_line)
	}
	message = append(message, id)

	return message
}

func add_onechar_lines (dict []byte) (res [][]byte) {
	for i := range dict {
		res = append(res, []byte{})
		res[i] = append(res[i], dict[i])
	}
	return res
}
