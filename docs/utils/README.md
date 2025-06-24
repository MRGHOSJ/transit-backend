<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>utils Package</title>
    <link rel="stylesheet" href="../styles.css">
</head>
<body>
<div class="container">

<div class="breadcrumbs">
    <a href="../README.md">Home</a> &raquo; <span>utils</span>
</div>

<div class="package-header">
    <h1>📦 utils</h1>
    <p class="filepath"><i>C:\Users\DELL\Desktop\transit-backend-main\internal\utils\loader.go</i></p>
</div>

<div class="toc">
    <h2>📑 Package Contents</h2>
    <ul>
        <li><a href="#functions">Functions (1)</a></li>    </ul>
</div>

<section id="functions" class="card">
    <h2>🔧 Functions</h2>
    <div class="func-accordion">

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">LoadTransportData</span>
                <span class="func-sig">(path string)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>
    </div>
</section>

</div>
</body>
</html>
