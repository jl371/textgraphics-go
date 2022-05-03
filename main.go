package main
import (
	"log"
	"fmt"
)
type Screen struct {
	 sc [][]bool
}
var layer1 Screen

var x int
var y int 
//this is where the cursor is
var pen bool

func InitScreen() [][]bool{
	matrix := make([][]bool, 80)
    rows := make([]bool, 80*31)
    for i, startRow := 0, 0; i < 80; i, startRow = i+1, startRow+31 {
        endRow := startRow + 31
        matrix[i] = rows[startRow:endRow:endRow]
    }
	setPen(true)
    return matrix
}
func ResetCursor() {
	x = 0
	y = 0
}
func DisplayMe(){
	fmt.Print("\033[H\033[2J")
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
func setPen(b bool) {
	if b == true {
		pen = true
	} else {
		pen = false
	}
}
func removePixel(w int, z int) {
	layer1.sc[w][z] = false;
}
func movecursor(w int, z int) {
	x = w
	y = z
}
func drawLine(length int, orientation int) {
	plotPixel(x,y)
	switch (orientation) {
	case 0:
		//draw line upwards
		for i:=1; i < length; i++ {
			h := y - i
			if (h < 0 || h > 30) {
				break
			} else {
				if pen == true {
					layer1.sc[x][h] = true
				} else {
					layer1.sc[x][h] = false
				}
				
			}
		}
		break
	case 1:
		//diagonal line to top right
		for i := 1; i < length; i++ {
                h := y - i
                j := x + i
                if (h < 0 || h > 30) || (j < 0 || j > 80) {
					break 
				} else {
					if pen == true {
						layer1.sc[j][h] = true
					} else {
						layer1.sc[j][h] = false
					}
				}
            }
            break
	case 2:
		for i := 1; i < length; i++ {
                j := x + i;
                if ((j < 0 || j > 80)) {
					break
				} else { 
					if pen == true {
						layer1.sc[j][y] = true
					} else {
						layer1.sc[j][y] = false
					}
				}
            }
            break
	case 3:
		//diagonal line to the bottom right
		for i := 1; i < length; i++ {
			h := y + i
			j := x + i
			if ((h < 0 || h > 30) || (j < 0 || j > 80)) {
				break
			} else {
				if pen == true {
					layer1.sc[j][h] = true
				} else {
					layer1.sc[j][h] = false
				}
			}
		}
		break
	case 4:
		//draw line downwards, starting at x, y
		for i := 1; i < length; i++ {
			h := y + i
			if (h < 0 || h > 30) { 
				break 
			} else { 
				if pen == true {
					layer1.sc[x][h] = true
				} else {
					layer1.sc[x][h] = false
				}
			}
		}
		break;
	case 5:
		//diagonal line to the bottom left
		for i := 1; i < length; i++ {
			h := y + i
			j := x - i
			if ((h < 0 || h > 30) || (j < 0 || j > 80)) { 
				break; 
			} else {
				if pen == true {
					layer1.sc[j][h] = true
				} else {
					layer1.sc[j][h] = false
				}
			} 
		}
		break;
	case 6:
		for i := 1; i < length; i++ {
                j := x - i
                if ((j < 0 || j > 80)) {
					break
				} else { 
					if pen == true {
						layer1.sc[j][y] = true
					} else {
						layer1.sc[j][y] = false
					}
				}
            }
        break
	case 7:
		for i := 1; i < length; i++ {
                h := y - i
                j := x - i
                if ((h < 0 || h > 30) || (j < 0 || j > 80)) {
					break
				} else { 
					if pen == true {
						layer1.sc[j][h] = true
					} else {
						layer1.sc[j][h] = false
					}
				}
				
            }
            break
	}


}

func main(){
	log.Println("things")
	layer1.sc = InitScreen()
	plotPixel(0,20)
	movecursor(10,10)
	drawLine(5,0)
	drawLine(5,1)
	drawLine(5,2)
	drawLine(5,3)
	drawLine(5,4)
	drawLine(5,5)
	drawLine(5,6)
	drawLine(5,7)
	DisplayMe()

}
