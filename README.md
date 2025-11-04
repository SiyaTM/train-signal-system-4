# train-signal-system-4
A train signaling system that uses Go programming. It focuses on how concurrency can be used to run the system. 

## Concept
GPS goroutine
It periodically emits the trains coordinates. 
Sensor system 
Listens, find the segments that are nearby and a SensorEvent is sent. 
Signal system
Listens to SensorEvents and updates signal states. 

## Centralization 
Central broadcasting is provisioned, using the GPS. It helps with movement updates. Sensors can listen to the updates and provide the required signal changes. 

## GPS
Model 
In go the model plays a crucial role in how the fields are managed. There are 3 fields to consider. The ID, coordinates, and timestamp. 
System 
Simulation file 
Simulation of a train that emits position data over time via a channel. 

## Sensor system 
It listens for GPS updates, checks coordinates for each segment, and emit SensorEvents to the SignalCh. 

## Signal system
Signal updates are made for each segment. Additionally, listen for sensor events is used to check when a train enters and exits a segment. 
Signal
The signal ahead of a red one would be yellow. It warns the next train that the upcoming segment is occupied.

## Output 
Snippet
```
[GPS] Train Train-1 position: -25.755, 28.231
[SENSOR] Train Train-1 ENTERED Segment-A
[SENSOR] Train Train-1 ENTERED Segment-B
[SIGNAL] Signal-0 for Segment-A → RED
[SIGNAL] Signal-1 for Segment-B → RED
[SIGNAL] Signal-0 for Segment-A → YELLOW
```
