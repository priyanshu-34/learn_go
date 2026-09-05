package main

import (
	"fmt"
	"sync"
	"time"
)

 func research(name string) (string, error){
	time.Sleep(1 * time.Second);
	fmt.Println("Done: ", name)
	return name, nil
 }

func goroutines(){
	start := time.Now()
	var wg sync.WaitGroup
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		wg.Add(1)
		go func(){
			defer wg.Done()
			research(task)
		}()
	}
	wg.Wait()
	fmt.Println("Total time taken: ", time.Since(start))
}