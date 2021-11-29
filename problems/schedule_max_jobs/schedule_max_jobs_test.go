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

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			`
ScheduleMaxJobs()
given:    a set of jobs with start and end
should:   return max number of non-overlapping jobs
input:    %v
result:   %v
expected: %v
			`,
			input, result, expected)
	} else {
		fmt.Printf(
			`
ScheduleMaxJobs()
given:    a set of jobs with start and end
should:   return max number of non-overlapping jobs
input:    %v
result:   %v
expected: %v
		`,
			input, result, expected)
	}
}
