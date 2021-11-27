package schedule_max_jobs

func ScheduleMaxJobs(jobs [][2]int) (res [][2]int) {
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
