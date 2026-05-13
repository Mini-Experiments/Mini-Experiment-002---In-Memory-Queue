# In-Memory System Entities

## Entities Involved
- Producer
- Consumer
- In-Memory Queue 
- Events

## 1. Producers

### 1.1 What job does this do?
Responsible for creating, detecting, capturing or forwarding state changes.

### 1.2 What breaks if it doesn't exist?
No ingress of information.

### 1.3 What does it talk to?
Talks to the in-memory queue and the source from where the data is coming or the previous process in the chain of communication/execution flow.

### 1.4 Different types

- **Data producer** — A fitness tracker producing data.
- **Log producer** — A web server generating logs.
- **Batch producer** — Nightly financial process, record transactions during the day and then combine them into a batch.

## 2. Consumers

### 2.1 What job does this do?
Operates on information received from producer. Might aggregate, persist, transform or forward information downstream.

### 2.2 What breaks if it doesn't exist?
No processing of data.

### 2.3 What does it talk to?
Talks to the in-memory queue and the next process in the chain of communication/execution flow.

### 2.4 Different types

- **Event processor** — Stock trading system reacting in real-time.
- **Storage consumer** — A DB logging actions.
- **Analytics consumer** — Dashboard that aggregates user statistics.

## 3. In-Memory Queue

### 3.1 What job does this do?
Acts like a buffer between producer and consumer. Technically, temporally decouples both producers and consumers (meaning they do not have to operate at identical speeds or at the same instant). Simply, lets the producer continue with their flow of execution without being blocked simply because of response time taken by the consumer.

### 3.2 What breaks if it doesn't exist?
Efficiency of the system plummets. Producers wait before they can push their data to an already occupied consumer being blocked in the process and thus wasting CPU cycles, throughput, responsiveness or concurrency.

### 3.3 What does it talk to?
Operationally it is bi-directional interacting with both producer and consumer. But data flow is uni-directional from producer to consumer.

### 3.4 Different types

- **In-memory queue** — local task queue inside a desktop or a smartphone, buffering tasks in RAM.
- **Persistent queue** — logging systems, logging each result in DB ensuring durability even if crashes (ACID in DB).
- **Distributed queue** — events flow across nodes giving fault tolerance at scale. Example — Apache Kafka.

### 3.5 Events
They are just bits of information and lack any intrinsic logic. They simply exist to convey state changes. Without them the entire system collapses, and they don’t actually talk to any specific entities. They are just data in transit. Can be thought of as packet of information representing state change. Can include — payload, metadata, timestamp, event type/schema, identifier etc.

**Type of events** —

- **User events** — click or form submission.
- **System events** — file update or server restart.
- **Business events** — purchase/transaction complete.
- **External events** — API call