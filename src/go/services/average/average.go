package average

import (
	"errors"
)

func sum(numbers []int) (sum int){
	for _,num := range numbers{
		sum += num
	}
	return
}
// Numbers will return the average for the array passed in
func Numbers(allNumbers []int ) (average float64, err error){
	if len(allNumbers) == 0{
		err = errors.New("Require at least one number calculate an average")
		return
	}
	average = float64(sum(allNumbers)) / float64(len(allNumbers))
	return
}