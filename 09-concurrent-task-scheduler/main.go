package main

import (
	"container/heap"
	"fmt"
	"time"
)

// Task represents a unit of work.
type Task struct {
	Id             string
	ProcessingTime int      // in seconds
	Dependencies   []string // list of task IDs that must complete first
	Priority       int      // lower number means higher priority
}

// TaskItem is the element stored in our priority queue.
type TaskItem struct {
	task  *Task
	index int // required by heap.Interface methods
}

// PriorityQueue implements heap.Interface and holds TaskItems.
type PriorityQueue []*TaskItem

// Len returns the number of items in the queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less returns true if the task at index i has higher priority than the task at index j.
func (pq PriorityQueue) Less(i, j int) bool {
	// Lower Priority value means higher priority.
	if pq[i].task.Priority == pq[j].task.Priority {
		return pq[i].task.Id < pq[j].task.Id
	}
	return pq[i].task.Priority < pq[j].task.Priority
}

// Swap swaps two items in the queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds a new item to the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*TaskItem)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority item.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// TaskScheduler schedules tasks with dependencies and a maximum concurrency.
type TaskScheduler struct {
	tasks         map[string]*Task    // maps task id to Task
	inDegree      map[string]int      // counts the number of unmet dependencies for each task
	reverseDeps   map[string][]string // maps a task id to the tasks that depend on it
	readyQueue    *PriorityQueue      // tasks that are ready to run (in-degree == 0)
	maxConcurrent int                 // maximum tasks running concurrently
	totalTasks    int                 // total number of tasks scheduled
}

// NewTaskScheduler creates a new TaskScheduler.
func NewTaskScheduler(maxConcurrent int) *TaskScheduler {
	return &TaskScheduler{
		tasks:         make(map[string]*Task),
		inDegree:      make(map[string]int),
		reverseDeps:   make(map[string][]string),
		readyQueue:    &PriorityQueue{},
		maxConcurrent: maxConcurrent,
	}
}

// AddTasks adds the provided tasks and builds the dependency graph.
func (s *TaskScheduler) AddTasks(taskList []Task) {
	s.totalTasks = len(taskList)
	// Populate tasks and inDegree.
	for i := range taskList {
		task := taskList[i]
		s.tasks[task.Id] = &Task{
			Id:             task.Id,
			ProcessingTime: task.ProcessingTime,
			Dependencies:   task.Dependencies,
			Priority:       task.Priority,
		}
		s.inDegree[task.Id] = len(task.Dependencies)
	}
	// Build the reverse dependency graph.
	for _, task := range s.tasks {
		for _, dep := range task.Dependencies {
			s.reverseDeps[dep] = append(s.reverseDeps[dep], task.Id)
		}
	}
	// Initialize the priority queue with tasks that have no dependencies.
	heap.Init(s.readyQueue)
	for id, count := range s.inDegree {
		if count == 0 {
			heap.Push(s.readyQueue, &TaskItem{task: s.tasks[id]})
		}
	}
}

// Run executes the tasks following their dependencies and priority constraints.
func (s *TaskScheduler) Run() {
	completions := make(chan string)
	running := 0        // current number of running tasks
	completedCount := 0 // number of tasks completed

	// Continue until all tasks have completed.
	for completedCount < s.totalTasks {
		// Schedule ready tasks while we have capacity.
		for running < s.maxConcurrent && s.readyQueue.Len() > 0 {
			// Pop the highest priority task.
			item := heap.Pop(s.readyQueue).(*TaskItem)
			task := item.task
			// Print the "started" message here in order.
			fmt.Printf("Task %s started. (Priority: %d)\n", task.Id, task.Priority)
			running++
			go func(task *Task) {
				// Simulate processing time.
				time.Sleep(time.Duration(task.ProcessingTime) * time.Second)
				fmt.Printf("Task %s completed.\n", task.Id)
				completions <- task.Id
			}(task)
		}
		// Wait for a task to finish.
		finishedTaskId := <-completions
		running--
		completedCount++
		fmt.Printf("Tasks Completed: %d\n", completedCount)
		// Update dependencies for tasks that depend on the finished task.
		for _, dependentId := range s.reverseDeps[finishedTaskId] {
			s.inDegree[dependentId]--
			if s.inDegree[dependentId] == 0 {
				heap.Push(s.readyQueue, &TaskItem{task: s.tasks[dependentId]})
			}
		}
	}
}

func main() {
	// Define the tasks.
	tasks := []Task{
		{Id: "task1", ProcessingTime: 2, Dependencies: []string{}, Priority: 1},
		{Id: "task2", ProcessingTime: 1, Dependencies: []string{"task1"}, Priority: 2},
		{Id: "task3", ProcessingTime: 3, Dependencies: []string{"task1"}, Priority: 1},
		{Id: "task4", ProcessingTime: 1, Dependencies: []string{"task2", "task3"}, Priority: 3},
		{Id: "task5", ProcessingTime: 2, Dependencies: []string{"task4"}, Priority: 2},
		{Id: "task6", ProcessingTime: 2, Dependencies: []string{"task5"}, Priority: 1},
		{Id: "task7", ProcessingTime: 1, Dependencies: []string{"task5"}, Priority: 3},
		{Id: "task8", ProcessingTime: 2, Dependencies: []string{"task5"}, Priority: 2},
	}

	// Create a scheduler that allows up to 2 concurrent tasks.
	scheduler := NewTaskScheduler(2)
	scheduler.AddTasks(tasks)
	scheduler.Run()
}
