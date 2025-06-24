<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>middleware Package</title>
    <link rel="stylesheet" href="../styles.css">
</head>
<body>
<div class="container">

<div class="breadcrumbs">
    <a href="../README.md">Home</a> &raquo; <span>middleware</span>
</div>

<div class="package-header">
    <h1>📦 middleware</h1>
    <p class="filepath"><i>C:\Users\DELL\Desktop\transit-backend-main\internal\middleware\middleware.go</i></p>
</div>

<div class="toc">
    <h2>📑 Package Contents</h2>
    <ul>
        <li><a href="#functions">Functions (3)</a></li>    </ul>
</div>

<section id="functions" class="card">
    <h2>🔧 Functions</h2>
    <div class="func-accordion">

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">Logging</span>
                <span class="func-sig">(next http.Handler)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">CORS</span>
                <span class="func-sig">(next http.Handler)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>

        <div class="accordion-item">
            <button class="accordion-btn" onclick="this.classList.toggle('active');
                this.nextElementSibling.classList.toggle('show')">
                <span class="func-name">Recoverer</span>
                <span class="func-sig">(next http.Handler)</span>
            </button>
            <div class="accordion-content">
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


            </div>
        </div>
    </div>
</section>

</div>
</body>
</html>
