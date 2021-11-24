package schedule_max_jobs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScheduleMaxJobs(t *testing.T) {
  got := ScheduleMaxJobs([][2]int{{1, 2},{1, 2}})
  want := [][2]int{{1, 2},{1, 2}}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %q, wanted %q", got, want)
  } else {
    fmt.Printf("got %q, wanted %q\n", got, want)
  }
}
