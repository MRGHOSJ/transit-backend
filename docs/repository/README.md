# 📦 Package: `repository`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\transport\repository\repository.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (5)](#-functions)

## 🧱 Structs

### `transportRepository`

```go
type transportRepository struct {
}
```

**Summary:** Empty struct for transport repository

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repo := transportRepository{}
```

**Edge Cases:**
- Lacks implementation details
- May be part of an interface pattern


---

## 🔧 Functions

<details>
<summary><b><code>NewTransportRepository(data *model.Data)</code></b></summary>

**Summary:** Constructor for TransportRepository

**Parameters:**
- `data` (*model.Data): Initialization data

**Returns:** Initialized TransportRepository instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repo := NewTransportRepository(dataSource)
```

**Edge Cases:**
- Nil data parameter
- Uninitialized model.Data fields


</details>

<details>
<summary><b><code>buildTypeIndex(lines []model.Line)</code></b></summary>

**Summary:** Builds type index from lines

**Parameters:**
- `lines` ([]model.Line): Input lines to index

**Returns:** Constructed TypeIndex mapping

**Complexity:**
- Time: O(n) where n is lines count
- Space: O(n) for storing index

**Example:**
```go
index := buildTypeIndex(documentLines)
```

**Edge Cases:**
- Empty lines array
- Duplicate line types
- Nil line elements


</details>

<details>
<summary><b><code>GetAllLines()</code></b></summary>

**Summary:** Retrieves all transport lines from repository

**Returns:** Slice of Line model objects

**Complexity:**
- Time: O(n) where n is number of lines
- Space: O(n) for returned slice

**Example:**
```go
lines := repo.GetAllLines()
```

**Edge Cases:**
- Empty repository returns empty slice
- Large dataset may impact performance


</details>

<details>
<summary><b><code>GetLine(t string, l string)</code></b></summary>

**Summary:** Finds specific transport line by type and ID

**Parameters:**
- `t` (string): Transport type
- `l` (string): Line identifier

**Returns:** Pointer to Line model or error if not found

**Complexity:**
- Time: O(1) average, O(n) worst case
- Space: O(1)

**Example:**
```go
line, err := repo.GetLine("bus", "25")
```

**Edge Cases:**
- Non-existent line returns error
- Empty input strings may return error


</details>

<details>
<summary><b><code>GetLinesByType(t string)</code></b></summary>

**Summary:** Gets all lines of specified transport type

**Parameters:**
- `t` (string): Transport type

**Returns:** Slice of Line models or error

**Complexity:**
- Time: O(n) where n is matching lines
- Space: O(n) for returned slice

**Example:**
```go
busLines, err := repo.GetLinesByType("bus")
```

**Edge Cases:**
- Unknown type may return empty slice or error
- Case sensitivity in type matching


</details>

