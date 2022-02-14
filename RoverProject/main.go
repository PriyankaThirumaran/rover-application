package main
import (
    "fmt"
    "strconv"
    // "strings"
    // "errors"
    "log"
    "net/http"

    sw "./go"
)

  
// main function
func main() {
    var option int
    fmt.Println("Select any option")
    fmt.Println("1. Get input through CLI")
    fmt.Println("2. Get input through REST API")
    fmt.Scanf("%d", &option)
    if (option == 1){
        xyCoordOfPlateau := [2]int{0, 0}
        
        fmt.Println("Enter the upper-right coordinates of the plateau(space separated): ")
        fmt.Scan(&xyCoordOfPlateau[0])
        fmt.Scan(&xyCoordOfPlateau[1])
        
        if(xyCoordOfPlateau[0] >= 0 && xyCoordOfPlateau[1] >= 0){
            var noOfRovers int
            fmt.Println("Enter the number of rovers: ")
            fmt.Scan(&noOfRovers)
            var rovers []sw.Rover
            for i:=0; i<noOfRovers; i++{
                var rover sw.Rover
                fmt.Println("Enter the rover position(space separated x-cordinate, y-coordinate and direction): ")
                fmt.Scan(&rover.XCoordinate)
                fmt.Scan(&rover.YCoordinate)
                fmt.Scan(&rover.Direction)
                fmt.Println("Enter the rover movements to be moved: ")
                fmt.Scan(&rover.Moves)
                rovers = append(rovers, rover)
            }

            for _, rover := range rovers {
                roverDest, err := sw.FindDestination(rover, xyCoordOfPlateau)
                if err!=nil{
                    fmt.Println(err)
                }else{
                    fmt.Println(strconv.Itoa(roverDest.XCoordinate) + " " + strconv.Itoa(roverDest.YCoordinate) + " " + roverDest.Direction)
                }
            }
        }else{
            fmt.Println("Enter the valid coordinates for plateau")
        }    
    }else if (option == 2 ){
        log.Printf("Server started")

        router := sw.NewRouter()

        log.Fatal(http.ListenAndServe(":8080", router))
    }else{
        fmt.Println("Invalid input from user")
    }  
}