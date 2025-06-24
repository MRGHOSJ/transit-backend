# 📦 Package: `utils`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\utils\loader.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>LoadTransportData(path string)</code></b></summary>

**Summary:** Loads transport data from a file

**Parameters:**
- `path` (string): File path to load data from

**Returns:** Loaded data as *model.Data or error

**Complexity:**
- Time: O(n) where n is file size
- Space: O(n) where n is data size

**Example:**
```go
data, err := LoadTransportData('routes.json')
```

**Edge Cases:**
- Nonexistent file returns error
- Malformed data returns parsing error


</details>

