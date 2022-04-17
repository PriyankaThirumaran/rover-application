# rover-application
This is a golang application that finds the final position of a rover on a rectangular plateau.
-> Takes user input from CLI and can be accessed through REST API.

Application Workflow:
A squad of robotic rovers are to be landed by NASA on a plateau on Mars.
This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the surrounding terrain to send back to Earth.
A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points.
The plateau is divided up into a grid to simplify navigation. An example position might be '0, 0, N' which means the rover is in the bottom left corner and facing North.
In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. ' L' and ' R' makes the rover spin 90 degrees left or right respectively, without moving from its current spot.
'M' means move forward one grid point, and maintain the same heading.

Working: (Sample interaction thro' CLI)
Select any option
1. Get input through CLI
2. Get input through REST API

->  1

Enter the upper-right coordinates of the plateau(space separated):(the lower-left coordinates of the plateau are assumed to be 0,0.)

->  5 5

Enter the number of rovers:

->  2

Enter the rover position(space separated x-cordinate, y-coordinate and direction):

->  1 2 N

Enter the rover movements to be moved:

->  LMLMLMLMM

Enter the rover position(space separated x-cordinate, y-coordinate and direction):

->  3 3 E

Enter the rover movements to be moved:

-> MMRMMRMRRM

1 3 N

5 1 E

Working of REST API :
Select any option
1. Get input through CLI
2. Get input through REST API

-> 2

2022/02/14 09:22:01 Server started

POST API:
URL : http://localhost:8080/mars/rover/move/{xOfPlateau}/{yOfPlateau}
Sample URL : http://localhost:8080/mars/rover/move/5/5
Payload: 
{
    "x_coordinate" : 1,
    "y_coordinate" : 2,
    "direction" : "N",
    "moves" : "LMLMLMLMM"
}
Response : 
{
    "x_coordinate": 1,
    "y_coordinate": 3,
    "direction": "N"
}
