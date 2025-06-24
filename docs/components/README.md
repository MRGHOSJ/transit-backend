# 📦 Package: `components`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\transport\components\component.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>HaversineDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64)</code></b></summary>

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


</details>

