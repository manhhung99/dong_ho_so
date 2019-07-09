package main
import "fmt"
import "time"
import "strings"
import "os"
import "os/exec"
//create font of characters :
var font=[][]string{
	[]string{" _ _  ",
	         "|   | ",
			 "|_ _| "},
	[]string{"   ",
	         "/| ",
			 " | "},
	[]string{" _ _  ",
			 " _ _| ",
			 "|_ _  "},
	[]string{"_ _  ",
	         "_ _| ",
			 "_ _| "},
	[]string{"      ",
	         "|_ _| ",
			 "    | "},
	[]string{" _ _  ",
	         "|_ _  ",
			 " _ _| "},
	[]string{" _ _  ",
	         "|_ _  ",
			 "|_ _| "},
	[]string{"_ _ ",
	         "  / ",
			 " /  "},
	[]string{" _ _  ",
	         "|_ _| ",
			 "|_ _| "},
	[]string{" _ _  ",
	         "|_ _| ",
			 " _ _| "},
	[]string{"    ",
	         " *  ",
			 " *  "},	
}

// in: datetime of OS
// out: string "07:35:43"
func read_system_clock()string{
	full_time := time.Now()
	fmt.Println(full_time)
	hms:=full_time.String()
	hhmmss:=string(hms[11:19])
	fmt.Println(hhmmss)
	return hhmmss
}
// int: character & font
// out: [] tuong ung voi 1 font
func convert_1_character_to_1_font(character string,font [][]string)[]string{
	var str="0123456789:"
	s:=strings.Split(str,"")
	var f []string
	if character==s[0]{
		f=font[0]
	}
	if character==s[1]{
		f=font[1]
	}
	if character==s[2]{
		f=font[2]
	}
	if character==s[3]{
		f=font[3]
	}
	if character==s[4]{
		f=font[4]
	}
	if character==s[5]{
		f=font[5]
	}
	if character==s[6]{
		f=font[6]
	}
	if character==s[7]{
		f=font[7]
	}
	if character==s[8]{
		f=font[8]
	}
	if character==s[9]{
		f=font[9]
	}
	if character==s[10]{
		f=font[10]
	}
	return f
}
//in: string "07:35:43"
//out: ["0","7",":","3",....}
func string_strings_of_array(str string)[]string{
	arr:=strings.Split(str,"")
	return arr
}
// in : string_hhmmss
// out: fonts_array
func string_hhmmss_to_fonts_in_array(str_arr []string)[][]string{
	var font_string [][]string
	for i:=0;i<8;i++{
		str_ar:=str_arr[i]
		font_str:=convert_1_character_to_1_font(str_ar,font)
		font_string=append(font_string,font_str)
	}
	return font_string
}
//in:font1 []string ,font2 []string
//out: font []string
func merge_2_font(font1 []string,font2 []string)[]string{
	var fontnew=[]string {"","",""}
	for i:=0;i<(len(font1));i++{
		fontnew[i]=font1[i]+font2[i]
	}
	return fontnew
}
func merge_array_of_fonts(array_of_fonts [][]string)[]string{
	var merged_font=[]string{"","",""}
	for i:=0;i<(len(array_of_fonts));i++{
		merged_font=merge_2_font(merged_font,array_of_fonts[i])
	}
	return merged_font
} 
func clear_screen(){
	cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
func display(array []string){
	for i:=0;i<(len(array));i++{
		row_i:=array[i]
		fmt.Println(row_i)
	}
}
//in :	
//
func main(){
	pre_sec:=0
	for true{
		t:=time.Now()
		s:=t.Second()
		if s!=pre_sec{
			clear_screen()
			str:=read_system_clock()
			ar:=string_strings_of_array(str)
			font_hhmmss:=string_hhmmss_to_fonts_in_array(ar)
			a:=merge_array_of_fonts(font_hhmmss)
			display(a)
		}
		pre_sec=s
		time.Sleep((1/1000)*time.Second)
	}
}
