package main 
import "fmt"
import "os" 
import "log"
import "bufio"


func letterize(ind *int, word string, index *int, font []string, res []string) []string{
	if (word[*ind]=='\\' && *ind!=len(word)-1 && word[*ind+1]=='n') {
		for i:=0; i<8; i++ {
			res=append(res,"")
		}
		*ind+=1;
		*index+=8;
		return res;
	}
	up:=(int(word[*ind])-32)*9+1;
	for i:=*index; i<*index+8; i++ {
		// if *ind!=0 {
		// 	res[i]+="";
		// }
		res[i]+=font[up+i-*index];
		//res[i]+="\n";
	}
	return res
} 
func Printres(res []string) {
	if (Output(res)) {
		return;
	}
	for i:=0; i<len(res); i++ {
		for j:=0; j<len(res[i]); j++ {
			fmt.Printf("%c", res[i][j])
		}
		fmt.Println();
	}
}
func Font() string{
	if (len(os.Args)>=3) {
		if (os.Args[2]=="standard") {
			return "fonts/standard.txt";
		}
		if (os.Args[2]=="shadow") {
			return "fonts/shadow.txt";
		}
		if (os.Args[2]=="thinkertoy") {
			return "fonts/thinkertoy.txt";
		}
		
	}
	return "fonts/standard.txt";
}
func Output(res[] string) bool {
	file:=""
	if (len(os.Args)>=4) {
		flag:="--output=";
		for i:=0; i<len(flag); i++ {
			if (flag[i]!=os.Args[3][i]) {
				return false;
			}
		}
		file=os.Args[3][len(flag):];
	}
	f,err:=os.Create(file);
	if err!=nil {
		log.Fatal(err);
	}
	defer f.Close();
	for i:=0; i<len(res); i++ {
		_,err=f.WriteString(res[i])
		_,err=f.WriteString("\n")
		if err!=nil {
			log.Fatal(err);
		}
	}
	return true;
}
func Conv(word string, fontName string) []string {
	var font []string;
	f,err:=os.Open("fonts/"+fontName	+".txt");
	defer f.Close();
	scanner:=bufio.NewScanner(f);
	for scanner.Scan(){
		font=append(font,scanner.Text());
	}
	if (err!=nil) {
		log.Fatal(err);
	} 
	index:=0
	res := make([]string, 8);
	for i:=0; i<len(word); i++ {
		res=letterize(&i,word,&index,font,res);
	}
	//singlestring:=""
	for i:=0; i<len(res); i++ {
		res[i]+="\n"
	}
	return res;
}
// func main() {
// 	if (len(os.Args)==1) {
// 		fmt.Println("Provide the string");
// 		return;
// 	}
// 	word := os.Args[1];
	
// 	res:=Conv(word);
// 	Printres(res);
	
// }	