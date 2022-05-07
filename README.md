# textgraphics-go
<p> This is a small and lightweight textmode graphics library, written in Go. This is used programatically. The screen is 80x30, monochrome</p>

### Features
* Can plot points, lines and squares
* Can animate stars and lines

### Functions
* movecursor(x,y) - moves cursor to specified point. This is where lines and squares originate.
* reset() - blanks screen, resets cursor to 0,0
* drawLine(length, orientation) - draws a line on the screen where the cursor is at, of length specified, in 8 different orientations. Going clockwise, 0 is line pointing up, 1 is diagonal going top right, etc.
* drawSquare(x,y,size) - draws a square of size specified, with the top left corner being at x,y on screen
* plotPixel(w,z) - plots a point at w,z, coordinates specified
* removePixel(w,z) - removes pixel at w,z, coordinates specified
* DisplayMe() - Display whatever is currently in the screen buffer
* setPen(b) - When true, it will draw lines and squares with a filled in character. If false, it will draw with a space (erase)
* AnimateStar(x,y,length,repeats) - Animates an 8 pointed star on screen, at position x,y. Length is length of points, repeats is how many times it will repeat the animation.
* AnimateLine(x,y,length,orientation,repeats) - Animates a line on screen at position x,y. Length is length of line, orientation is same as drawLine, repeats is how many times it will repeat the animation.
* ResetCursor() - moves cursor back to 0,0
* InitScreen() - Internal function for intializing screen.
