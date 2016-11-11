package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"runtime"
	"sync"
	"time"
//	"net/http"
//	"github.com/wblackecaldwell/profiler"
//	"strconv"
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
//вывод в файл *.lzw
func work_with_out_files(path string, message []byte) {
	path = path + ".lzw"
	fout, err := os.Create(path);
	check(err)
	w := bufio.NewWriter(fout)
  for _, line := range message {
    fmt.Fprint(w, line)
  }
  w.Flush()
}

func work_with_out_files_decompress(path string, message string) {
	fout, err := os.Create(path);
	check(err)
	w := bufio.NewWriter(fout)
    fmt.Fprint(w, message)
  w.Flush()
}

func read_the_path() string{
	fmt.Println("enter files path: ")
	in := bufio.NewScanner(os.Stdin)
  in.Scan()
  if err := in.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
  }
  return in.Text()
}
var threads_numbers int = 6
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	var dict [][][]byte
	var dict_artificial [][]byte
  //var dict_o [][]byte
	var message [][]byte
	var message_artificial []byte
	var result_message []byte
	//получаем путь к input файлу
	//path_in := "files/" +read_the_path()
	//path_out := path_in + "2"
	path_in := "files/Hamlet.txt"
	//заполняем словарь единичными элементами
//	t0 := time.Now();
	//dict_o = fill_in_dbl_dic_old(dict_o, path_in)
	//t1 := time.Now();
//	fmt.Println(" time= ", t1.Sub(t0))
//	 fmt.Println("dictionary=  ", dict_o)
	//вызываем compress
	//тут будет собственно закодированый текст
	// t0 = time.Now();
	 //_ = compress_old(dict_o, path_in)
	// t1 := time.Now();
	 //fmt.Println("compress time old= ", t1.Sub(t0))
	//это вывод в файл
	//work_with_out_files(path_in, message_old)

	// var message2 string
	// t0 = time.Now();
	// message2 = decompress(dict, message)
	// t1 = time.Now();
	// fmt.Println("decompress time= ", t1.Sub(t0))
	// //fmt.Println("m2=", message2)
	// work_with_out_files_decompress(path_out, message2)
		//c := make(chan [][]byte)

	//	t0 := time.Now();

		for i:=1; i<=threads_numbers; i++{

			 go func(){

				dict = append(dict, dict_artificial)
				message = append(message, message_artificial)
				dict[i-1] = fill_in_dbl_dic(dict[i-1], path_in, i)
				message[i-1] = compress(dict[i-1], path_in, i)
				fmt.Println("res", message)
			}()
			time.Sleep(100 * time.Millisecond)

		}
		time.Sleep(100 * time.Millisecond)
		for j:=1; j<=threads_numbers; j++{
		for _, value:=range message[j-1]{
			result_message = append(result_message, value)
			}}
		work_with_out_files(path_in, result_message)
		//t1:= time.Now();

		//t0=t0.Add(time.Millisecond*10*time.Duration(threads_numbers))
		//fmt.Println("compress time= ", t1.Sub(t0) - (time.Millisecond*200*time.Duration(threads_numbers)))
		//fmt.Println("real", compress(dict_o, path_in))
}

//Тут проыеряется есть ли в словаре dict "символ" char, и если есть, то возвращается еще позиция этого символа в массиве
func byte_in_dbl_slice(dict [][]byte, char []byte) (bool, byte){
	var (result = false
	 	hlp = false
	 	id byte
 )
	for i := range dict {
		if !hlp {
			if len(dict[i]) == len(char){
			hlp = true
			id = byte(i)
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
func fill_in_dbl_dic(dict [][]byte, path string, t_num int) ([][]byte){
	char := make([]byte, 1)
	fin, stat := work_with_files(path)
	for j:=0; j<t_num-1; j++{
		for i := 0; i < (int(stat.Size())/threads_numbers); i++ {
			_, _ = fin.Read(char)
		}
	}
	 for i := 0 ; i < (int(stat.Size())/threads_numbers); i++ {
		 //fmt.Println("i=", i)
	// for i := 0; i < int(stat.Size()); i++ {
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

func fill_in_dbl_dic_old(dict [][]byte, path string) [][]byte {
	char := make([]byte, 1)
	fin, stat := work_with_files(path)
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
func compress(dict [][]byte, path string, t_num int) (message []byte) {
	//искусственный char
	//нужен что бы считывать посимвольно
	char := make([]byte, 1)
	next_char := make([]byte, 1)
	//поточная строка
	var (curent_line []byte
	 hlp_line []byte
	 id byte
	 bl bool
 )
	fin, stat := work_with_files(path)
	//считываем первый символ в файле
	_, err := fin.Read(next_char)
	check(err)
	//забиваем в поточную строку первый символ
	curent_line = append(curent_line, next_char[0])
	//главный цыкл. пока не считался весь файл
	for j:=0; j<t_num-1; j++{
		for i := 0; i < (int(stat.Size())/threads_numbers); i++ {
			_, _ = fin.Read(char)
		}
	}
	for i := 1; i < (int(stat.Size())/threads_numbers); i++ {
		// fmt.Println("-----------------------------------------------------------")
		_, err := fin.Read(char)
		check(err)
		hlp_line = append(curent_line, char[0])
		//проверка есть ли в словаре char. если да то в bl = true а id содержит тот самый код который мы выводим
		//по сути набор из idшников комбинаци1 из словаря и есть выходящим сообщением
		bl, id = byte_in_dbl_slice(dict, hlp_line)
		// fmt.Println("id ", id, " bl ", bl, " cl ", string(curent_line), " ch ", string(char[0]), " mess ", message)
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

func compress_old(dict [][]byte, path string) (message []byte) {
	char := make([]byte, 1)
	next_char := make([]byte, 1)
	var (curent_line []byte
	 hlp_line []byte
	 id byte
	 bl bool
 )
	fin, stat := work_with_files(path)
	_, err := fin.Read(next_char)
	check(err)
	curent_line = append(curent_line, next_char[0])
	for i := 1; i < int(stat.Size()); i++ {
		_, err := fin.Read(char)
		check(err)
		hlp_line = append(curent_line, char[0])
		bl, id = byte_in_dbl_slice(dict, hlp_line)
		if bl {
			curent_line = append (curent_line, char[0])
		} else {
			bl, id = byte_in_dbl_slice(dict, curent_line)
			message = append(message, id)
			line := make([]byte, len(hlp_line))
			copy(line, hlp_line)
			dict = append(dict, line)
			curent_line = curent_line[:0]
			curent_line = append(curent_line, char[0])
		}
		bl, id = byte_in_dbl_slice(dict, curent_line)
	}
	message = append(message, id)
	return message
}

func decompress(dict [][]byte, message []byte) (message2 string){
	var char byte
	message2 += string(dict[message[0]])
	for i := 1; i < int(len(message)-1); i++ {
		curent_line := make([]byte, len(dict[message[i]]))
		copy(curent_line, dict[message[i]])
		message2 = message2 + string(curent_line)
		dict = append(dict, dict[char])
		dict[len(dict)-1] = append(dict[len(dict)-1], curent_line[0])
		char = message[i]
	}
	return message2
}

// func byte_to_str_to_int_to_str(num []byte, dict [][]byte) string {
// 	i, err := strconv.Atoi(string(num))
// 	check(err)
// 	return string(dict[i])
// }
//
// func byte_to_str_to_int(num []byte, dict [][]byte) []byte {
// 	i, err := strconv.Atoi(string(num))
// 	check(err)
// 	return dict[i]
// }
