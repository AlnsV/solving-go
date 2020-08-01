package solver

import (
	"container/list"
)

/*
Data Structures: Stacks/Queues

School has decided to supply packed lunch boxes to students.
- There are N lunch boxes that are placed on each other. The lunch boxes are either circular or
rectangular in shape.

- Each student has his/her own preference for the type of lunchbox.

If the student finds that the tiffin at the top of the stack is not as per his/her preference (of shape)
he/she will go back and rejoin the queue and the process will continue.

Estimate the number of students who will not be able to get lunch.

Input Format
Input1: N, denoting the number of Children and lunchboxes.
Input2: An array of N elements, each element can be either 0(Rectangle) or 1(circle) denoting the
type of lunch box from top to bottom.
Input3: An array of N elements, each element denoting the preference 0(Rectangle) or 1(Circle) of a
student from the start till the end of the queue.

Output Format
For the given input your code should output the number of students who will not be able to eat lunch.

*/

func LunchBoxes(numChildren int, availableBoxes *list.List, boxesPreference []string) int {
	var result int
	for children := 0; children < numChildren; children++ {
		currentWantedBox := boxesPreference[children]
		for e := availableBoxes.Front(); e != nil; e = e.Next() {
			availableBox := e.Value.(string)
			if currentWantedBox == availableBox {
				availableBoxes.Remove(e)
				result++
				break
			}
		}
	}
	return result
}
