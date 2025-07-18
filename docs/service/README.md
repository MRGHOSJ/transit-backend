# 📦 Package: `service`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\transport\service\route_planner.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (6)](#-functions)

## 🧱 Structs

### `RoutePlanner`

```go
type RoutePlanner struct {
	graph *Graph 
}
```

**Summary:** Route planning component with graph dependency

**Parameters:**
- `graph` (*Graph): Reference to graph structure

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (type definition)
- Space: O(1)

**Example:**
```go
planner := RoutePlanner{graph: stationGraph}
```

**Edge Cases:**
- Nil graph reference
- Concurrent modification of underlying graph


---

## 🔧 Functions

<details>
<summary><b><code>NewRoutePlanner(graph *Graph)</code></b></summary>

**Summary:** Creates a new RoutePlanner instance with a given graph

**Parameters:**
- `graph` (*Graph): Graph structure representing the transport network

**Returns:** Pointer to a new RoutePlanner instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
planner := NewRoutePlanner(myGraph)
```

**Edge Cases:**
- Nil graph input may cause initialization issues


</details>

<details>
<summary><b><code>getPreferences(mode string)</code></b></summary>

**Summary:** Retrieves transport preferences for a given mode

**Parameters:**
- `mode` (string): Transport mode (e.g., 'walking', 'driving')

**Returns:** TransportPreferences struct containing mode-specific settings

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
prefs := planner.getPreferences('biking')
```

**Edge Cases:**
- Unsupported mode string may return default preferences


</details>

<details>
<summary><b><code>FindRoutes(fromLat float64, fromLng float64, toLat float64, toLng float64, modes []model.TransportType, prefsParam string)</code></b></summary>

**Summary:** Finds route options between two geographic points

**Parameters:**
- `fromLat` (float64): Origin latitude
- `fromLng` (float64): Origin longitude
- `toLat` (float64): Destination latitude
- `toLng` (float64): Destination longitude
- `modes` ([]model.TransportType): Allowed transport types
- `prefsParam` (string): Preference configuration identifier

**Returns:** Slice of RouteOption structs and potential error

**Complexity:**
- Time: O(E + V log V) where V=vertices, E=edges (Dijkstra's)
- Space: O(V) for priority queue storage

**Example:**
```go
routes, err := planner.FindRoutes(40.7, -74.0, 34.0, -118.0, []model.TransportType{model.BUS, model.TRAIN}, 'fastest')
```

**Edge Cases:**
- Invalid coordinates
- No available routes between points
- Empty transport modes slice


</details>

<details>
<summary><b><code>findNearestAccessPoints(lat float64, lng float64, modes []model.TransportType)</code></b></summary>

**Summary:** Finds nearest transport access points to a location

**Parameters:**
- `lat` (float64): Target latitude
- `lng` (float64): Target longitude
- `modes` ([]model.TransportType): Transport types to consider

**Returns:** Slice of Connection structs representing nearby access points

**Complexity:**
- Time: O(n) where n=number of candidate points
- Space: O(k) where k=number of returned connections

**Example:**
```go
connections := planner.findNearestAccessPoints(51.5, -0.1, []model.TransportType{model.TUBE})
```

**Edge Cases:**
- Location far from any transport nodes
- Empty modes list
- Invalid coordinates


</details>

<details>
<summary><b><code>countTransfers(segments []model.RouteSegment)</code></b></summary>

**Summary:** Counts the number of transfers in a route

**Parameters:**
- `segments` ([]model.RouteSegment): List of route segments to analyze

**Returns:** Number of transfers as int

**Complexity:**
- Time: O(n)
- Space: O(1)

**Example:**
```go
transfers := planner.countTransfers(segments) // returns 2
```

**Edge Cases:**
- Empty segments list returns 0
- Direct route with no transfers returns 0


</details>

<details>
<summary><b><code>getInstructions(conn model.Connection, from model.Station, to model.Station)</code></b></summary>

**Summary:** Generates navigation instructions for a connection

**Parameters:**
- `conn` (model.Connection): Connection details
- `from` (model.Station): Starting station
- `to` (model.Station): Destination station

**Returns:** Formatted instructions as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
instructions := planner.getInstructions(conn, from, to) // returns 'Take line A to X, transfer to line B'
```

**Edge Cases:**
- Invalid connection returns empty string
- Same from/to station returns 'Already at destination'


</details>

