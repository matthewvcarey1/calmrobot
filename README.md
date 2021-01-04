# Calmrobot

This program calculates the coordinates that are visitable by a robot starting on (0,0) that can move
horizontally or vertically one position at a time. This is in a world where every coodinate pair whose combined individual absolute digits adds up to a sum greater than 23 has a mine on it.

## To build
go build -o robot.exe ./cmd/calmrobot/

## To run

./robot.exe

or

./robot.exe -images

The later writes two png image files mines.png and robot.png showing the world as a picture.
The file mines.png shows the mines and robot.png shows the accessable areas for the robot. (see picture)

./robot.exe -safe X

Where X is a number. Sets a new value for the safe number other than the standard 23.

./robot.exe -help

Prints a help text and exits


![Robot](robot.png)

## Improvements

The code probably could be made to go faster. We use the symmetry in the 4 quarters to only calculate 1 quadrant and multiply the result by 4. We have to be careful to count the axes only 4 times and the origin only once. 

The big O value is poor as the performance of the code degrades sharply as the safe number increments.

The lack of Unit tests is more of an issue. 

There is no error checking for input conditions.
The out of bounds errors will be caught by overflow and 
I, otherwise would have to cause a Panic and terminate 
at that point anyway.

