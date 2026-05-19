# Decision Log

## Overall Decisions Made
1. A well defined statement of what "complete" looks like.
2. Standard for format of events to be used throughout the system.
3. Data structure/memory representation of events.
4. Type of producer or consumer for demonstration.
5. **Go** as the programming language.
6. Producers as goroutines of the same program.
7. A couple of consumers with different functionalities.
8. Queue implemented using Linked Lists.

## [DECISION 00]

### Context
Need a clearly scoped completion criterion for the first version of the in-memory queue system.

### Decision
The system is considered complete 
1. When multiple producers can concurrently generate events.
2. Offload them onto a shared in-memory queue with ordering preserved. 
3. A consumer processes them using defined logic while visibly demonstrating ordering.

### Rationale
Defines observable system behavior without prematurely introducing persistence, distribution, or advanced synchronization concerns.

### Consequence
The implementation must support concurrent producers, ordered enqueue/dequeue semantics, and observable event flow verification.

---
## [DECISION 01]

### Context
Standardizing data format for events, but the problem statement leaves the format open.

### Decision
Chose JSON as the data format for events.

### Alternatives
| Format      | Reason for Rejection / Acceptance |
|-------------|-------------------------------------|
| JSON        | Accepted: human-readable, flexible, widely supported |
| XML         | Rejected: more verbose, less flexible for quick iteration |
| Protobuf    | Rejected: more complex setup due to rigid schema definitions |
| CSV         | Rejected: too simple, not suitable for structured events |
| Plain Key-Value | Rejected: minimal structure, messire to scale |

### Rationale
1. JSON offers the right balance: it’s easy to read, flexible, and integrates well with most languages. 
2. Widely used in networking request and response cycles. Compounds on the previous build.

### Consequence
All events will be structured in JSON, allowing easy extension if the system scales.

---
## [DECISION 02]

### Context
In-memory representation needed for events.

### Decision
Decided to use strongly structured representation like custom structs/classes.

### Alternatives

| Alternative | Decision | Reason |
|---|---|---|
| Structs / Classes | Accepted | Strong structure, type safety, scalable, predictable under concurrency |
| Maps / Dictionaries | Rejected | Weakly structured, allows invalid fields/types, less predictable |
| Tuples | Rejected | Poor readability and maintainability at scale |
| DataFrames / Tabular Structures | Rejected | Better suited for analytics rather than real-time event flow |

### Rationale
1. Any system prefers having predictable representation for events. Ease in debugging.
2. The conept of classes and objects is scalable by design and in our case, events might increase exponentially with increase in number of producers or their speed.
3. Easy JSON serialization.

### Consequence
Each event will have a clearly defined representation in memory, reducing ambiguity and improving maintainability, debugging, and scalability.

---
## [DECISION 03]

### Context
The system requires multiple producers generating events concurrently and pushing them into a shared in-memory queue. A concrete producer type was needed for the first implementation.

### Decision
The producers will be implemented as log producers.

### Alternatives

| Alternative | Decision | Reason |
|---|---|---|
| Data Producers | Rejected | Would require simulating external sensors/devices unnecessarily |
| Batch Producers | Rejected | Less suitable for demonstrating continuous asynchronous flow |
| User Action Producers | Rejected | Difficult to automate and test repeatedly |
| Log Producers | Accepted | Structured, continuous, timestamp-friendly, easy to simulate, and compounds with previous WebSocket server build |

### Rationale
Log generation naturally fits the producer-consumer model and allows easy simulation of concurrent event generation while remaining simple enough for the first implementation.

### Consequence
Each producer will generate log events independently and push them into the shared in-memory queue for ordered consumption.

---
## [DECISION 04]

### Context
We need a consumer to process the queued events and produce a final output. The consumer’s task must be clear and fit well into the overall flow.

### Decision
The consumer will be a simple event printer, which receives events from the queue and prints them in the order they arrive.

### Alternatives

| Alternative | Decision | Reason |
|---|---|---|
| Complex Analytics Processor | Rejected | Adds unnecessary complexity; we want a simple demonstration first |
| Database Logger | Rejected | Persistent storage is outside the current scope; we focus on order demonstration |
| Real-Time Dashboard | Rejected | Requires external visualization; too complex for initial build |
| Simple Event Printer | Accepted | Keeps focus on order preservation and queue behavior; easy to verify via console output |

### Rationale
We need a minimal but clear demonstration: printing events in order keeps the focus on queue mechanics without introducing new complexities.

### Consequence
We will confirm that the events arrive in order and can be easily inspected, ensuring that the producer-to-queue-to-consumer pipeline works as expected.

---
## [DECISION 05]

### Context
A technology stack decision was needed for building the producer-consumer in-memory queue system, ensuring future scalability and skill compounding.

### Decision
Go will be the programming language chosen for this build.

### Alternatives

| Alternative | Decision | Reason |
|---|---|---|
| Python | Rejected | Faster iterations, but not as performant for scaling as Go; also diverges from prior Go-based build. |
| Java | Rejected | Mature, but verbose; slower iterations; heavier runtime. |
| Node.js | Rejected | Great for I/O, but less performant; less aligned with systems-level ambition. |
| Rust | Rejected | Powerful but steeper learning curve; Go better for quick, iterative development. |
| C# | Rejected | Cross-platform, but heavier runtime; not as nimble as Go in prototyping. |

### Rationale
Go was selected to compound with prior work, leverage its strong concurrency model, and align with industry demand for scalable, efficient systems.

### Consequence
The project will be implemented in Go, ensuring a lean, fast, and concurrency-ready foundation.

---
## [DECISION 06]

### Producer Implementation

1. Producers as goroutines/processes of one master program, not separate program.
    1. **Alternative** - Have different programs for each producer.
    2. **Rationale** - For experiment scope, assumed all 4 producers to be server logs, producing the same type of events.
2. Log event fields: producerID, consumerID, IPAddress, pushTimestamp, eventID, payload.
    1. **Alternative** - The complete standard structure of a log event.
    2. **Rationale** - Simplified enough for demo implementation. 
        1. producerID - To track parent producer of an event.
        2. consumerID - To decide consumer, assuming multiple consumers with different functionalities.
        3. pushTimestamp - To decide which event was pushed first onto the queue for lock allocation.
        4. eventID - For verifying order preservation.
        5. IPAddress and Payload - To draw some resemblance to the standard structure.

3. Push function prints to the terminal for slice 1.
4. Producers run for finite bursts with random delays, not indefinitely. 

---

## [DECISION 07]

### Consumer Implementation

1. Two consumers with different functions were implemented, in order to resemble multiple consumers playing different roles as in real systems.

2. ConsumerID was allocated randomly and simple filtering used in order to allocate consumers.

---

## [DECISION 08]

### Queue Implementation

1. Linked List chosen for queue implementation.
    1. **Alternatives** - Array (rejected - non-dynamic size) and Stack (unnecessary complexity)
    2. **Rationale** - 
        1. Since it's highly unpredictable as to how many events might be in the queue at any given instance of time. The best choice was to use Linked List. 
        2. Choosing an arbitrary absurdly big size for an array would have simply wasted storage space. 
        3. Furthermore linked lists also provide for a way to release resources when used.
2. Push/Pop output to terminal for isolated testing and simplicity.
---