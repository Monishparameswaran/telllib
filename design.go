package use

import ("fmt")

func Pattern(num int){
	for i:=0;i<num;i++{
		for j:=0;j<5;j++{
			fmt.Print("*");
		}
		fmt.Println();
	}
}
