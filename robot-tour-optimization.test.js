/*

Problem: Robot Tour Optimization
Input: A set `S` of `n` points in the plane.
[ [x, y], ... ]

Output: What is the shortest cycle tour that visits each point in the set S?
Sorted set such that traveling point to point is shortest distance.
[... points ordered by shortest cycle tour]

req:
  explicit
  - must take into account distance to return to first point
  - time is proportional to distance

  implicit
  - plane arbirarily large or small
  - plane is always in first quadrant (positive x, y)
  - assume no point is at same coordinates

mental model
- start with first point in set (arbitrary)
- check against every other point for the nearest
- the nearest is the second point
- check every every other point for the nearest
- third point is found
- so on until last point is found

algo
1. point

2. grab set of other points
3. check set of other points, that's next nearest point
4. grab set of other points ...

*/

function pointTravelOptimization(points) {
  const sortedPoints = [points[0]]
  points = points.slice(1, points.length)

  while (points.length > 0) {
    const lastPoint = sortedPoints[sortedPoints.length - 1]
    let nearestPoint = points[0]
    let nearestPointDistance = pointDistance(lastPoint, points[0])
    let nearestPointIndex = 0

    for (let i = 1; i < points.length; i++) {
      const point = points[i]

      const distance = pointDistance(lastPoint, point)
      if (distance < nearestPointDistance) {
        nearestPoint = point
        nearestPointDistance = distance
        nearestPointIndex = i
      }
    }

    sortedPoints.push(nearestPoint)
    points.splice(nearestPointIndex, 1)
  }

  return sortedPoints
}

const pointDistance = (pointA, pointB) => {
  const xDiff = Math.abs(pointA[0] - pointB[0])
  const yDiff = Math.abs(pointA[1] - pointB[1])
  return xDiff + yDiff
}

test('Sorts four points travel', () => {
  const sortedPoints = pointTravelOptimization([
    [0, 0],
    [1, 1],
    [0, 1],
    [1, 0],
  ])

  console.log(sortedPoints)

  expect(sortedPoints).toEqual([
    [0, 0],
    [0, 1],
    [1, 1],
    [1, 0],
  ])
})

test('Sorts height points travel', () => {
  const sortedPoints = pointTravelOptimization([
    [4, 2],
    [9, 9],
    [2, 4],
    [11, 7],
    [2, 7],
    [9, 2],
    [4, 9],
    [11, 4],
  ])

  console.log(sortedPoints)

  expect(sortedPoints).toEqual([
    [4, 2],
    [2, 4],
    [2, 7],
    [4, 9],
    [9, 9],
    [11, 7],
    [11, 4],
    [9, 2],
  ])
})
