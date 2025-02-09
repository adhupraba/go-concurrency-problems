# Producer-Consumer Pattern

### Problem Statement:

Implement a producer-consumer system using channels. Create multiple producer goroutines that generate data (e.g., random numbers or strings) and send them on a channel. Create several consumer goroutines that receive data from the channel and process it (for example, by printing or accumulating values).

### Key Concepts:

- Use channels for communication.
- Use sync.WaitGroup to wait for goroutines to finish.
- Consider a graceful shutdown (e.g., using a done channel or closing the channel).

### Hints:

- Decide on a buffer size for your channel.
- Producers can use a loop to send data; consumers should exit gracefully when the channel is closed.
- Use a WaitGroup to ensure that all producers and consumers finish their work.
