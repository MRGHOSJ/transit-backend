# 📦 Package: `middleware`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\middleware\middleware.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Logging(next http.Handler)</code></b></summary>

**Summary:** HTTP middleware for request logging

**Parameters:**
- `next` (http.Handler): Next handler in the chain

**Returns:** http.Handler with logging middleware

**Complexity:**
- Time: O(1) per middleware call
- Space: O(1)

**Example:**
```go
router.Use(Logging)
```

**Edge Cases:**
- May log sensitive data if not filtered
- Performance impact with high traffic


</details>

<details>
<summary><b><code>CORS(next http.Handler)</code></b></summary>

**Summary:** HTTP middleware for CORS headers

**Parameters:**
- `next` (http.Handler): Next handler in the chain

**Returns:** http.Handler with CORS support

**Complexity:**
- Time: O(1) per middleware call
- Space: O(1)

**Example:**
```go
router.Use(CORS)
```

**Edge Cases:**
- Overly permissive CORS headers security risk
- Preflight request handling


</details>

<details>
<summary><b><code>Recoverer(next http.Handler)</code></b></summary>

**Summary:** HTTP middleware for panic recovery

**Parameters:**
- `next` (http.Handler): Next handler in the chain

**Returns:** http.Handler with panic recovery

**Complexity:**
- Time: O(1) per middleware call
- Space: O(1)

**Example:**
```go
router.Use(Recoverer)
```

**Edge Cases:**
- May mask underlying issues if overused
- Recovery from specific panic types


</details>

