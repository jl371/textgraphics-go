package main
import (
	"log"
	"fmt"
	"time"
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
				currentline += "█"
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
	if (x > 80) {
		x = 0 + w 
	} else {
		x = w
	}
	if (y > 30) {
		y = 0 + z
	} else {
		y = z
	}

}
func drawLine(length int, orientation int) {
	if pen == true {
		plotPixel(x,y)
	} else {
		removePixel(x,y)
	}
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
	//DisplayMe()
	//time.Sleep(33 * time.Millisecond)
}
func AnimateStar(x int, y int, length int, repeats int) {
	movecursor(x,y)
	for i:=0; i < repeats; i++ {
		setPen(true)
		for j:=0; j < length; j++ {
			drawLine(j,0)
			drawLine(j,1)
			drawLine(j,2)
			drawLine(j,3)
			drawLine(j,4)
			drawLine(j,5)
			drawLine(j,6)
			drawLine(j,7)
			DisplayMe()
			time.Sleep(33 * time.Millisecond)
		}
		setPen(false)
		
		for k:=0; k < length; k++ {
			drawLine(k,0)
			drawLine(k,1)
			drawLine(k,2)
			drawLine(k,3)
			drawLine(k,4)
			drawLine(k,5)
			drawLine(k,6)
			drawLine(k,7)
			DisplayMe()
			time.Sleep(33 * time.Millisecond)
		}

	}
}
func AnimateLine(x int, y int, length int, orientation int, repeats int) {
	movecursor(x,y)
	for i:=0; i < repeats; i++ {
		setPen(true)
		
			setPen(true)
		for j:=0; j < length; j++ {
			drawLine(j,orientation)
			DisplayMe()
			time.Sleep(33 * time.Millisecond)
		}
		setPen(false)
		
		for k:=0; k < length; k++ {
			drawLine(k,orientation)
			DisplayMe()
			time.Sleep(33 * time.Millisecond)
		}
		
	}
}
func drawSquare(x int, y int, size int) {
	movecursor(x,y)
	drawLine(size,4)
	drawLine(size,2)
	movecursor(x+size-1, y+size-1)
	drawLine(size,0)
	drawLine(size,6)
	movecursor(x,y)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
}
func reset() {
	for i := 0; i < 30; i++ {
		for j := 0; j < 80; j++ {
			layer1.sc[j][i] = false
		}
	}
	ResetCursor()
	DisplayMe()
}
func main(){
	log.Println("things")
	layer1.sc = InitScreen()
	plotPixel(0,20)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	plotPixel(5,20)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	plotPixel(10,20)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	plotPixel(15,20)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	plotPixel(20,20)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	plotPixel(79,30)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	plotPixel(3,3)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	drawSquare(50, 0, 9)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	drawSquare(45,0,5)
	DisplayMe()
	time.Sleep(33 * time.Millisecond)
	AnimateStar(10,10,8,5)
	AnimateStar(20,20,6,8)
	AnimateStar(50,25,4,4)
	AnimateLine(20,10,10,0,5)
	for i:= 0; i < 8; i++ {
		AnimateLine(50,25,5,i,2)
	}
	movecursor(10,5)
	for i:=0; i < 7; i++ {
		setPen(true)
		movecursor(x+i,y+i)
		drawLine(6,1)
		DisplayMe()
		time.Sleep(33 * time.Millisecond)
		drawLine(6,5)
		drawLine(6,7)
		DisplayMe()
		time.Sleep(33 * time.Millisecond)

	}
}
