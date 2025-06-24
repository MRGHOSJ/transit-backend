# 📦 Package: `router`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\router\router.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Setup(data *model.Data)</code></b></summary>

**Summary:** Configures HTTP router with application data

**Parameters:**
- `data` (*model.Data): Application data structure

**Returns:** Configured *mux.Router instance

**Complexity:**
- Time: O(n) where n is route count
- Space: O(n) for route storage

**Example:**
```go
router := Setup(appData)
```

**Edge Cases:**
- Nil data parameter handling
- Route collision detection


</details>

