<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>repository Package</title>
    <link rel="stylesheet" href="../styles.css">
</head>
<body>
<div class="container">

<div class="breadcrumbs">
    <a href="../README.md">Home</a> &raquo; <span>repository</span>
</div>

<div class="package-header">
    <h1>📦 repository</h1>
    <p class="filepath"><i>C:\Users\DELL\Desktop\transit-backend-main\internal\transport\repository\repository.go</i></p>
</div>

<div class="toc">
    <h2>📑 Package Contents</h2>
    <ul>
        <li><a href="#structs">Structs (1)</a></li>        <li><a href="#functions">Functions (5)</a></li>    </ul>
</div>

<section id="structs" class="card">
    <h2>🧱 Structs</h2>
    <div class="struct-grid">

        <div class="struct-card">
            <h3>transportRepository</h3>
            <div class="struct-details">
                <pre><code>type transportRepository struct {
}</code></pre>
                <div class="doc-section">
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
                <span class="func-name">NewTransportRepository</span>
                <span class="func-sig">(data *model.Data)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">buildTypeIndex</span>
                <span class="func-sig">(lines []model.Line)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">GetAllLines</span>
                <span class="func-sig">()</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">GetLine</span>
                <span class="func-sig">(t string, l string)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">GetLinesByType</span>
                <span class="func-sig">(t string)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>
    </div>
</section>

</div>
</body>
</html>
