<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>service Package</title>
    <link rel="stylesheet" href="../styles.css">
</head>
<body>
<div class="container">

<div class="breadcrumbs">
    <a href="../README.md">Home</a> &raquo; <span>service</span>
</div>

<div class="package-header">
    <h1>📦 service</h1>
    <p class="filepath"><i>C:\Users\DELL\Desktop\transit-backend-main\internal\transport\service\route_planner.go</i></p>
</div>

<div class="toc">
    <h2>📑 Package Contents</h2>
    <ul>
        <li><a href="#structs">Structs (1)</a></li>        <li><a href="#functions">Functions (6)</a></li>    </ul>
</div>

<section id="structs" class="card">
    <h2>🧱 Structs</h2>
    <div class="struct-grid">

        <div class="struct-card">
            <h3>RoutePlanner</h3>
            <div class="struct-details">
                <pre><code>type RoutePlanner struct {
}</code></pre>
                <div class="doc-section">
                    **Summary:** Empty struct for route planning functionality

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) (zero-sized type)

**Example:**
```go
planner := RoutePlanner{}
```

**Edge Cases:**
- Methods may need nil receiver checks
- Uninitialized field dependencies


                </div>
            </div>
        </div>
    </div>
</section>

<section id="functions" class="card">
    <h2>🔧 Functions</h2>
    <div class="func-accordion">

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">NewRoutePlanner</span>
                <span class="func-sig">(graph *Graph)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">getPreferences</span>
                <span class="func-sig">(mode string)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">FindRoutes</span>
                <span class="func-sig">(fromLat float64, fromLng float64, toLat float64, toLng float64, modes []model.TransportType, prefsParam string)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">findNearestAccessPoints</span>
                <span class="func-sig">(lat float64, lng float64, modes []model.TransportType)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">countTransfers</span>
                <span class="func-sig">(segments []model.RouteSegment)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">getInstructions</span>
                <span class="func-sig">(conn model.Connection, from model.Station, to model.Station)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>
    </div>
</section>

</div>
</body>
</html>
