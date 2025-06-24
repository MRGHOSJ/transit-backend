<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>router Package</title>
    <link rel="stylesheet" href="../styles.css">
</head>
<body>
<div class="container">

<div class="breadcrumbs">
    <a href="../README.md">Home</a> &raquo; <span>router</span>
</div>

<div class="package-header">
    <h1>📦 router</h1>
    <p class="filepath"><i>C:\Users\DELL\Desktop\transit-backend-main\internal\router\router.go</i></p>
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
                <span class="func-name">Setup</span>
                <span class="func-sig">(data *model.Data)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>
    </div>
</section>

</div>
</body>
</html>
