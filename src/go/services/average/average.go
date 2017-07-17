package average

import (
	"errors"
)


// Numbers will return the average for the array passed in
func Numbers(allNumbers []int ) (average float64, err error){
if len(allNumbers) == 0{
	err = errors.New("Require at least one number calculate an average")
	return
}
sum :=0
for _,num := range allNumbers{
	sum += num
}
average = float64(sum) / float64(len(allNumbers))
return
}