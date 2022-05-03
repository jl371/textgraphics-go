package main
import (
	"log"
	
)
type Screen struct {
	 sc [][]bool
}
var layer1 Screen

var x int
var y int 
//this is where the cursor is

func InitScreen() [][]bool{
	matrix := make([][]bool, 80)
    rows := make([]bool, 80*31)
    for i, startRow := 0, 0; i < 80; i, startRow = i+1, startRow+31 {
        endRow := startRow + 31
        matrix[i] = rows[startRow:endRow:endRow]
    }
    return matrix
}
func ResetCursor() {
	x = 0
	y = 0
}
func DisplayMe(){
	var totalscreen string
	totalscreen+= "\n"
	var currentline string
	for i := 0; i < 30; i++ {
		for j := 0; j < 80; j++ {
			if layer1.sc[j][i] == true {
				currentline += "#"
			} else {
				currentline += " "
			}
		}
		totalscreen += currentline
		totalscreen += "\n"
		currentline = ""
	}
	log.Println(totalscreen)
}
func plotPixel(w int, z int) {
	layer1.sc[w][z] = true;
}

func main(){
	log.Println("things")
	layer1.sc = InitScreen()
	plotPixel(0,20)
	DisplayMe()

}git 
