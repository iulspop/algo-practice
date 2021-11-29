package schedule_max_jobs

func ScheduleMaxJobs(jobs [][2]int) (res [][2]int) {
	jobSchedules := [][][2]int{}
	jobSchedule := [][2]int{}

	createJobSets := func(jobs [][2]int) {
		if len(jobs) == 0 {
			append(jobSchedules, jobSchedule)
			jobSchedule = [][2]int{}
		} else {
			min := 0
			refIndex := 0
			for i := 0; i < len(jobs); i++ {
				job := jobs[i]
				start :=job[0]

				if start < min {
					refIndex = i
					min = start
				}
			}

			append(jobSchedule, jobs[refIndex])
			createJobSets(append(jobs[:refIndex], jobs[refIndex+1:]...))
		}
	}


	return jobs
}

/*

Input: [[0, 10], [5, 10]] // set of projects with start and end times
Output: [[]...] // largest possible set of non-overlapping projects.

Approach: Full Verification through recursive branch check

Algo:
- Create array of possible job sets
- Create array for temp job set

call create possible job sets(all jobs array)
  - If no more jobs
      - add job set to possible job sets
      - reset job set to empty
  - Pick the project that is most recent
      - add to temp job set
      - call create possible job sets without it

- array with all possible job sets
  - map/reduce to number of jobs
  - select first job set with maximum number




*/
