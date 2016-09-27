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

//это для обработки ошибок
//она будет юзаться после каждого раза, где будет заполняться err
//логической нагрузки не несет
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

//заполняем массив уникальных байтов (первое считывание)
	dic = fill_in_dic(dic)
	fmt.Println(dic)
//по сути это dic_2 = dic
//dic_2 - двумерный массив, и приходится юзать спец функцию чтобы заполнить его
//dic_2 - будет основным словарем где будут накопляться сочетания байтов
	dic_2 = add_onechar_lines(dic)
	fmt.Println("dictionary=  ", dic_2)
//тут будет собственно закодированый текст
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
//Тут проыеряется есть ли в словаре arr "символ" ch, и если есть, то возвращается еще позиция этого символа в массиве
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

//тоже самое что и byte_in_dbl_slice только для одномерного массива
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
//заполняем массив уникальных байтов
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
//ГЛАВНЫЙ ДВИЖ
func compress(dict [][]byte) (message []int) {
	//искусственный Cchar
	//нужен что бы считывать посимвольно
	char := make([]byte, 1)
	next_char := make([]byte, 1)
	//поточная строка
	var curent_line []byte
	var hlp_line []byte
	var id int
	var bl bool
	//открываем файл
	fin, err := os.Open("input.txt");
	check(err)
	//закрываем файл. в го можно его закрыть заранее)))
	defer fin.Close()
	//узнаем длинну файла - это вместо EOF
	stat, err := fin.Stat()
	check(err)
	//считываем первый символ в файле
	_, err = fin.Read(next_char)
	check(err)
	//забиваем в поточную строку первый символ
	curent_line = append(curent_line, next_char[0])
	//главный цыкл. пока не считался весь файл
	for i := 1; i < int(stat.Size()); i++ {
		fmt.Println("------------------------------")
		_, err := fin.Read(char)
		check(err)
		hlp_line = append(curent_line, char[0])
		//проверка есть ли в словаре char. если да то в bl = true а id содержит тот самый код который мы выводим
		//по сути набор из idшников комбинаци1 из словаря и есть выходящим сообщением
		bl, id = byte_in_dbl_slice(dict, hlp_line)
		fmt.Println("id ", id, " bl ", bl, " cl ", string(curent_line), " ch ", string(char[0]), " mess ", message)
		if bl {
			curent_line = append (curent_line, char[0])
		} else {
			bl, id = byte_in_dbl_slice(dict, curent_line)
			//добавлем код в строку вывода
			message = append(message, id)
			//тут происходила настолько черная магия, что мне пришлось пойти на этот костыль и искусственно создать line
			//скопировать в line строку и добавить непосредственно line в словарь
			line := make([]byte, len(hlp_line))
			copy(line, hlp_line)
			dict = append(dict, line)
			//обнуляем поточную строку
			curent_line = curent_line[:0]
			//доюавляем в пустую поточную строку последний считаный симол
			curent_line = append(curent_line, char[0])
		}
		bl, id = byte_in_dbl_slice(dict, curent_line)
	}
	message = append(message, id)
	return message
}
//позволяет присвоить двумерному массиву одномерный
//например есть dict [1 2 3 4 5], тогда res заполниться так [[1] [2] [3] [4] [5]]
func add_onechar_lines (dict []byte) (res [][]byte) {
	for i := range dict {
		res = append(res, []byte{})
		res[i] = append(res[i], dict[i])
	}
	return res
}
