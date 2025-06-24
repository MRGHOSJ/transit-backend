<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>handler Package</title>
    <link rel="stylesheet" href="../styles.css">
</head>
<body>
<div class="container">

<div class="breadcrumbs">
    <a href="../README.md">Home</a> &raquo; <span>handler</span>
</div>

<div class="package-header">
    <h1>📦 handler</h1>
    <p class="filepath"><i>C:\Users\DELL\Desktop\transit-backend-main\internal\transport\handler\route_handler.go</i></p>
</div>

<div class="toc">
    <h2>📑 Package Contents</h2>
    <ul>
        <li><a href="#structs">Structs (1)</a></li>        <li><a href="#functions">Functions (4)</a></li>    </ul>
</div>

<section id="structs" class="card">
    <h2>🧱 Structs</h2>
    <div class="struct-grid">

        <div class="struct-card">
            <h3>RouteHandler</h3>
            <div class="struct-details">
                <pre><code>type RouteHandler struct {
}</code></pre>
                <div class="doc-section">
                    **Summary:** Empty struct defining a route handler type

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
handler := RouteHandler{}
```

**Edge Cases:**
- No functionality defined yet
- Requires method implementations


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
                <span class="func-name">NewRouteHandler</span>
                <span class="func-sig">(planner *service.RoutePlanner)</span>
            </button>
            <div class="accordion-content">
                **Summary:** Constructor for RouteHandler with dependency injection

**Parameters:**
- `planner` (*service.RoutePlanner): Route planning service dependency

**Returns:** Pointer to initialized RouteHandler

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
handler := NewRouteHandler(myRoutePlanner)
```

**Edge Cases:**
- Nil planner input
- Memory allocation failure


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">GetRoutes</span>
                <span class="func-sig">(w http.ResponseWriter, r *http.Request)</span>
            </button>
            <div class="accordion-content">
                **Summary:** HTTP handler method for retrieving routes

**Parameters:**
- `w` (http.ResponseWriter): Response writer interface
- `r` (*http.Request): Incoming HTTP request

**Returns:** None (writes directly to response)

**Complexity:**
- Time: O(n) where n is route computation complexity
- Space: O(1) for handler, O(n) for response data

**Example:**
```go
router.HandleFunc("/routes", handler.GetRoutes)
```

**Edge Cases:**
- Invalid request parameters
- Route computation failures
- Response writer errors


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">response</span>
                <span class="func-sig">(w http.ResponseWriter, status int, data )</span>
            </button>
            <div class="accordion-content">
                **Summary:** Helper method for sending HTTP responses

**Parameters:**
- `w` (http.ResponseWriter): Response writer interface
- `status` (int): HTTP status code
- `data` (interface{}): Response payload (type varies)

**Returns:** None (writes directly to response)

**Complexity:**
- Time: O(1) for simple responses
- Space: O(1) for handler, O(n) for response data

**Example:**
```go
handler.response(w, http.StatusOK, routes)
```

**Edge Cases:**
- Invalid status codes
- Unserializable data
- Response writer errors


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">errorResponse</span>
                <span class="func-sig">(w http.ResponseWriter, status int, message string)</span>
            </button>
            <div class="accordion-content">
                **Summary:** Sends an HTTP error response with status and message

**Parameters:**
- `w` (http.ResponseWriter): HTTP response writer
- `status` (int): HTTP status code
- `message` (string): Error message to send

**Returns:** None (writes directly to response)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
errorResponse(w, 404, "Page not found")
```

**Edge Cases:**
- Invalid status codes
- Large messages causing buffer issues


            </div>
        </div>
    </div>
</section>

</div>
</body>
</html>
