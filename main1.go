package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"time"
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
func work_with__out_files(path string, message []byte) {
	path = path + ".lzw"
	fout, err := os.Create(path);
	check(err)
	w := bufio.NewWriter(fout)
  for _, line := range message {
    fmt.Fprint(w, line)
  }
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

func main() {

	var dict [][]byte
	//получаем путь к input файлу
	path := read_the_path()
	//заполняем словарь единичными элементами
	t0 := time.Now();
	dict = fill_in_dbl_dic(dict, path)
	t1 := time.Now();
	fmt.Println(" time= ", t1.Sub(t0))
	// fmt.Println("dictionary=  ", dict, " time= ", t1.Sub(t0))
	//вызываем compress
	//тут будет собственно закодированый текст
	t0 = time.Now();
	message := compress(dict, path)
	t1 = time.Now();
	fmt.Println(" time= ", t1.Sub(t0))

	decompress(message, dict)
 // fmt.Println("message=  ", message)
	//это вывод в файл
	work_with__out_files(path, message)
}

//Тут проыеряется есть ли в словаре dict "символ" char, и если есть, то возвращается еще позиция этого символа в массиве
func byte_in_dbl_slice(dict [][]byte, char []byte) (bool, byte){
	var result = false
	var hlp = false
	var id byte
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
func fill_in_dbl_dic(dict [][]byte, path string) [][]byte {
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
func compress(dict [][]byte, path string) (message []byte) {
	//искусственный char
	//нужен что бы считывать посимвольно
	char := make([]byte, 1)
	next_char := make([]byte, 1)
	//поточная строка
	var curent_line []byte
	var hlp_line []byte
	var id byte
	var bl bool
	fin, stat := work_with_files(path)
	//считываем первый символ в файле
	_, err := fin.Read(next_char)
	check(err)
	//забиваем в поточную строку первый символ
	curent_line = append(curent_line, next_char[0])
	//главный цыкл. пока не считался весь файл
	for i := 1; i < int(stat.Size()); i++ {
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
//-------------------------------------------------

func Fill_dic_decomp (dict [][]byte, line []byte, position int )  [][]byte{
     dict[position] = append(dict, line)
} // или dict[position] := line

func decompress (massage []byte, dict [][]byte) (word []byte){
  var code []byte
	var hlp_l []byte
	var buf []byte
	var id byte
	var flag bool
	var j byte
	var current_line []byte

 //m := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
	file, err := os.Open("input.txt.lzw")
	    if err != nil {
	        // handle the error here
	        return
	    }
	    defer file.Close()

	    // get the file size
	    stat, err := file.Stat()
	    if err != nil {
	        return
	    }


	 _, err = file.Read(code)  // '0'
	 check(err)
	hlp_l = append(hlp_l, dict[code])
	current_line = append(current_line, hlp_l)    // 'a'
	 j = 0
	 for i:= 1; i< int(stat.Size()); i++ {
	    _, err := file.Read(code)   // first itteration code = '1'
	    // hlp_l := append(hlp_l, dict[code][j]) // first itteration hlp_l = 'ab'  âîçëå àïïåíäà ñòîèò íå ïðèñâîèòü à ïðîñòî ðàâíî, ïðîâåðèòü
	    // buf := append(buf, dict[code])
	    current_line := append(current_line, dict[code])


	    if len(hlp_l)-1 == len(buf){   // -1 îòíèìàåì òîò ñèìâîë, êîòîðûé ïðèøåë ñ ïðîøëîãî ðàçà
	       flag, id := byte_in_dbl_slice(dict, hlp_l)
	       if !flag {
	         Fill_dic_decomp(dict, hlp_l, i+4)
	         hlp_l := hlp_l[:-1]
	       }  //else buf := buf[:0] // åñëè óæå åñòü â ìàññèâå, íî ïî ìîåìó íå îáÿçàòåëüíî îáíóëÿòü. åñëè íå áóäåò ðàáîòàòü óáðàòü)

	    } else {
	        flag, id := byte_in_dbl_slice(dict, hlp_l)
	        if !flag {
	          Fill_dic_decomp(dict, hlp_l, i+4)
	          hlp_l := buf
	          buf := buf[:0]
	        } //else buf := buf[:0] // åñëè óæå åñòü â ìàññèâå, íî ïî ìîåìó íå îáÿçàòåëüíî îáíóëÿòü åñëè íå áóäåò ðàáîòàòü óáðàòü)
	    }


	 }
	fmt.Println(current_line)
	return current_line
	}
