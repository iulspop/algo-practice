package schedule_max_jobs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScheduleMaxJobs(t *testing.T) {
	input := [][2]int{{4, 12}, {2, 9}, {10, 15}, {6, 15}, {14, 34}, {16, 20}, {21, 30}, {22, 30}, {28, 46}, {32, 48}}

	result := ScheduleMaxJobs(input)
	expected := [][2]int{{2, 9}, {10, 15}, {16, 20}, {22, 30}, {32, 48}}

	AssertEquals(t, result, expected)
}

func TestSortIntervalsByEnd(t *testing.T) {
	input := [][2]int{{32, 48}, {4, 12}, {2, 9}, {14, 34}, {10, 15}, {6, 15}}

	result := SortIntervalsByEnd(input)
	expected := [][2]int{{32, 48}, {4, 12}, {2, 9}, {14, 34}, {10, 15}, {6, 15}}

	AssertEquals(t, result, expected)
}

func AssertEquals(t *testing.T, result [][2]int, expected [][2]int) {
	message := CreateTestMessage(result, expected)
	if reflect.DeepEqual(result, expected) {
		fmt.Print(message)
	} else {
		t.Errorf(message)
	}
}

func CreateTestMessage(result [][2]int, expected [][2]int) string {
	return fmt.Sprintf(
		`
ScheduleMaxJobs()
given:    a set of jobs with start and end
should:   return max number of non-overlapping jobs
result:   %v
expected: %v
	`, result, expected)
}
