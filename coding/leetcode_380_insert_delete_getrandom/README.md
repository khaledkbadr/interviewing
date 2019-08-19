# 380. Insert Delete GetRandom O(1)

Design a data structure that supports all following operations in average O(1) time.

* insert(val): Inserts an item val to the set if not already present.
* remove(val): Removes an item val from the set if present.
* getRandom: Returns a random element from current set of elements. Each element must have the same probability of being returned.

## Implementation

we need to find out how to to do the following

* store values with find complexity of O(1)

The only data structures that support find complexity of one is `map`, so we are going to use `map` as a `set` by putting zero-value as values

* store values in a way that can be accessed randomly

since `map` keys can't be accessed randomly since we can't know which keys are there in O(1) we have to use another data structure to store the values. `slice` seems to be the right data structure since it can be accessed randomly by index.

* deletion of values should be in O(1)

We need to obtain the index of the deleted value in the vals slice. Therefore, we will add the index value as the value in the `map` and we no longer use it as a `set`. With that whenever we want to delete a value we can delete it from the vals slice using the index stored in the map

## Varation

Designing a HashMap with the same characteristics as the discussed data structure. You can find the solution [here](../randomizedmap/randomizedmap.go)
