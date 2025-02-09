# Pipeline Pattern

### Problem Statement:

Implement a multi-stage pipeline. For example, create a pipeline that generates numbers, filters even numbers, squares them, and then collects the results. Each stage should run in its own goroutine, and channels should be used to pass data from one stage to the next.

### Key Concepts:

- Chaining channels and goroutines.
- Decoupling stages of processing.
- Handling channel closing in a pipeline.

### Hints:

- Each stage is a function that receives an input channel and returns an output channel.
- Ensure that when a stage finishes processing, it closes its output channel.
- Consider using for-range loops to receive data until a channel is closed.
