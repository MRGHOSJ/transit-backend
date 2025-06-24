<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>components Package</title>
    <link rel="stylesheet" href="../styles.css">
</head>
<body>
<div class="container">

<div class="breadcrumbs">
    <a href="../README.md">Home</a> &raquo; <span>components</span>
</div>

<div class="package-header">
    <h1>📦 components</h1>
    <p class="filepath"><i>C:\Users\DELL\Desktop\transit-backend-main\internal\transport\components\component.go</i></p>
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
                <span class="func-name">HaversineDistance</span>
                <span class="func-sig">(lat1 float64, lon1 float64, lat2 float64, lon2 float64)</span>
            </button>
            <div class="accordion-content">
                **Summary:** Calculates Haversine distance between two geographic coordinates

**Parameters:**
- `lat1` (float64): Latitude of first point
- `lon1` (float64): Longitude of first point
- `lat2` (float64): Latitude of second point
- `lon2` (float64): Longitude of second point

**Returns:** Distance in kilometers as float64

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
dist := HaversineDistance(52.5200, 13.4050, 48.8566, 2.3522) // Berlin to Paris
```

**Edge Cases:**
- Antipodal points (exact opposite sides of Earth)
- Identical coordinates returning zero
- Invalid coordinate ranges (lat outside [-90,90])


            </div>
        </div>
    </div>
</section>

</div>
</body>
</html>
