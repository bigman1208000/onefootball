package main

import (
	"onefootball/business"
    "fmt"
    "onefootball/concurrency"
    "onefootball/model"
)

func main() {
    var outPuts []model.Output
    
    noOfJobs := 10000
    go concurrency.Allocate(noOfJobs)
    done := make(chan bool)
    go concurrency.GetResult(done)
    concurrency.CreateWorkerPool(&outPuts)
    <-done

    fmt.Println("-----------------------------\nResult:\n")
    business.Print(outPuts)

}
