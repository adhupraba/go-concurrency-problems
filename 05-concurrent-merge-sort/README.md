# Concurrent Merge Sort

### Problem Statement:

Implement merge sort in a concurrent fashion. When sorting a large slice of integers, split the slice into two halves and sort each half concurrently using goroutines. Then merge the sorted halves. Limit the number of concurrent goroutines to avoid excessive resource usage.

### Key Concepts:

- Divide and conquer algorithms.
- Spawning goroutines for independent recursive calls.
- Synchronizing using channels or WaitGroups.

### Hints:

- Define a threshold for the slice size under which you perform the sort sequentially.
- Use a WaitGroup or channels to synchronize the results of the recursive sort calls.
- Implement the merge function carefully so that it handles merging two sorted slices.
