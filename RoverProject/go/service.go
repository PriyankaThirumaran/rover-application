package mars

import (
    // "strconv"
    // "strings"
    "errors"
)

func FindDestination(rover Rover, xyCoordOfPlateau [2]int) (Rover, error) {
    directions := [4]string{"N", "E", "S", "W"}
    leftDirections := [4]string{"W", "N", "E", "S"}
    rightDirections := [4]string{"E", "S", "W", "N"}
    xPosition := [4]int{0, 1, 0, -1}
    yPosition := [4]int{1, 0, -1, 0}
    for _, move:= range rover.Moves{
        switch move {
        case 'L':
            for index, dir := range directions{
                if rover.Direction == dir{
                    rover.Direction = leftDirections[index]
                    break
                }
            }

        case 'R':
            for index, dir := range directions{
                if rover.Direction == dir{
                    rover.Direction = rightDirections[index]
                    break
                }
            }
        case 'M':
            for index, dir := range directions{
                if rover.Direction == dir{
                    rover.XCoordinate += xPosition[index]
                    rover.YCoordinate += yPosition[index]
                    break
                }
            }
            if !isPositionValid(rover, xyCoordOfPlateau){
                return rover, errors.New("Rover moves out of the plateau")
            }
        }
   }
   rover.Moves = ""
   return rover, nil
}

func isPositionValid(rover Rover, xyCoordOfPlateau [2]int) bool {
    if rover.XCoordinate>=0 && rover.XCoordinate<= xyCoordOfPlateau[0] && rover.YCoordinate>=0 && rover.YCoordinate<= xyCoordOfPlateau[1]{
        return true
    }else{
        return false
    }
}