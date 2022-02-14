# rover-application
This is a golang application that finds the final position of a rover on a rectangular plateau.
-> Takes user input from CLI and can be accessed through REST API.

Working: (Sample interaction thro' CLI)
Select any option
1. Get input through CLI
2. Get input through REST API
1
Enter the upper-right coordinates of the plateau(space separated):
5 5
Enter the number of rovers:
2
Enter the rover position(space separated x-cordinate, y-coordinate and direction):
1 2 N
Enter the rover movements to be moved:
LMLMLMLMM
Enter the rover position(space separated x-cordinate, y-coordinate and direction):
3 3 E
Enter the rover movements to be moved:
MMRMMRMRRM
1 3 N
5 1 E

Working of REST API :
Select any option
1. Get input through CLI
2. Get input through REST API
2
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
