package main

import "testing"

// тестим подходят ли словари, берем за пример input.txt
// чтобы что то сломать меняй значения что я добавляю в словарь и будет ошибка
func Test_fill_in_dbl_dic(t *testing.T) {
	var dict [][]byte
	var res [][]byte
	var n int
	n = 7
	var path = "input.txt"

	res = append(res, []byte{97})
	res = append(res, []byte{98})
	res = append(res, []byte{99})
	res = append(res, []byte{100}) // ломать здесь !
	res = append(res, []byte{101})
	res = append(res, []byte{13})
	res = append(res, []byte{10})
	dict = fill_in_dbl_dic(dict, path)
	// fmt.Println(" res= ", res)
	// fmt.Println(" dict= ", dict)

	for i := 0; i < n; i++ {
		for j := 0; j < 1; j++ {
			// fmt.Println(" res el= ", res[i][j])
			// fmt.Println(" dict els= ", dict[i][j])
			if dict[i][j] != res[i][j] {
				t.Errorf("%d != %d", res, dict)
			}
		}
	}
}
