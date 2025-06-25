# 📦 Package: `handler`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\transport\handler\route_handler.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `RouteHandler`

```go
type RouteHandler struct {
	routePlanner *service.RoutePlanner 
}
```

**Summary:** Struct that holds a reference to a route planner service.

**Parameters:**
- `routePlanner` (*service.RoutePlanner): Reference to the route planner service

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
handler := RouteHandler{routePlanner: svc}
```

**Edge Cases:**
- Nil routePlanner reference
- Uninitialized routePlanner


---

## 🔧 Functions

<details>
<summary><b><code>NewRouteHandler(planner *service.RoutePlanner)</code></b></summary>

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


</details>

<details>
<summary><b><code>GetRoutes(w http.ResponseWriter, r *http.Request)</code></b></summary>

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


</details>

<details>
<summary><b><code>response(w http.ResponseWriter, status int, data )</code></b></summary>

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


</details>

<details>
<summary><b><code>errorResponse(w http.ResponseWriter, status int, message string)</code></b></summary>

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


</details>

