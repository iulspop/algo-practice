/*

stream of numbers

.........t

At any point 't', may get an API call asking for the highest product of 3 numbers

product as in result of multiplication

Assume:
  - biggest product is the multiplication of the 3 biggest numbers
  - numbers are given in array

Algo:
  - sort accending
  - grab the last three numbers (biggest)
  - multiply them
  - return product

*/

function getHighestProductOf3Numbers(numbers) {
  return [...numbers]
    .sort()
    .slice(numbers.length - 3, numbers.length)
    .reduce((product, number) => number * product)
}

console.log(getHighestProductOf3Numbers([1, 5, 5, 0, 0]))
