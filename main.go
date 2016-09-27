package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)
//это для обработки ошибок
//она будет юзаться после каждого раза, где будет заполняться err
//логической нагрузки не несет
func check(e error) {
    if e != nil {
			log.Fatal(e)
			}
}
//подготавливает файл
//и еще возвращает инфу о файле в переменной stat (мне нужна была для того что бы узнать длинну файла)
func work_with_files(path string) (file *os.File, stat os.FileInfo){
	file, err := os.Open(path);
	check(err)
	// defer file.Close()
	stat, err = file.Stat()
	check(err)
	return file, stat
}

func main() {
	var dict [][]byte
	//заполняем словарь единичными элементами
	dict = fill_in_dbl_dic(dict)
	fmt.Println("dictionary=  ", dict)
	//вызываем compress
	//тут будет собственно закодированый текст
	message := compress(dict)
	fmt.Println("message=  ", message)
	//это вывод в файл, он безобразный, я потом поправлю
	fout, err := os.Create("output.txt");
	check(err)
	w := bufio.NewWriter(fout)
	fmt.Fprintln(w, message)
	w.Flush()
}

//Тут проыеряется есть ли в словаре arr "символ" ch, и если есть, то возвращается еще позиция этого символа в массиве
func byte_in_dbl_slice(dict [][]byte, char []byte) (bool, int){
	var result = false
	var hlp = false
	var id int
	for i := range dict {
		if !hlp {
			if len(dict[i]) == len(char){
			hlp = true
			id = i
				for j := range char {
					if dict[i][j] != char[j]{
						hlp = false
						break
					}
				}
			}
		} else {result = true}
	}
	return result, id
}

//заполняем массив уникальных байтов
func fill_in_dbl_dic(dict [][]byte) [][]byte {
	char := make([]byte, 1)
	fin, stat := work_with_files("input.txt")
	 for i := 0; i < int(stat.Size()); i++ {
		_, err := fin.Read(char)
		check(err)
		bl, _ := byte_in_dbl_slice(dict, char)
		if !bl{
			line := make([]byte, len(char))
			copy(line, char)
			dict = append(dict, line)
		}
	}
	return dict
}
//ГЛАВНЫЙ ДВИЖ
func compress(dict [][]byte) (message []int) {
	//искусственный char
	//нужен что бы считывать посимвольно
	char := make([]byte, 1)
	next_char := make([]byte, 1)
	//поточная строка
	var curent_line []byte
	var hlp_line []byte
	var id int
	var bl bool
	fin, stat := work_with_files("input.txt")
	//считываем первый символ в файле
	_, err := fin.Read(next_char)
	check(err)
	//забиваем в поточную строку первый символ
	curent_line = append(curent_line, next_char[0])
	//главный цыкл. пока не считался весь файл
	for i := 1; i < int(stat.Size()); i++ {
		fmt.Println("-----------------------------------------------------------")
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
