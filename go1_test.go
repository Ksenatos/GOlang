package main 
import "testing"



func BenchmarkCompress(b *testing.B) {
 var dict [][]byte
 path := "input.txt"
 //b.SetBytes(2)
 for m:= 0; m < 5; m++{  //set max value "m" bigger for decrease amount of itteration
 for n:= 0; n< b.N; n++{
    b.ReportAllocs()
    compress(fill_in_dbl_dic(dict,path), path)
 } }
}

func BenchmarkDecompress(b *testing.B) {
 var dict [][]byte
 path := "input.txt"
 //b.SetBytes(2)
 for m:= 0; m < 5; m++{  //set max value "m" bigger for decrease amount of itteration
 for n:= 0; n< b.N; n++{
    b.ReportAllocs()
    decompress(fill_in_dbl_dic(dict,path), compress(fill_in_dbl_dic(dict,path), path))
 }}
}
