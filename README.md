# Groupie Tracker - Geolocalization & Search

A high-performance web application built in Go that tracks band and artist information, visualizes concert locations on interactive maps, and provides real-time search capabilities. 

Built as a hands-on backend engineering project for the learn2earn Academy (01 Edu platform).

## 🚀 Features

* **Real-Time Debounced Search:** A custom-built, case-insensitive search engine that scans artists, members, locations, and dates. It uses JavaScript debouncing to optimize network requests and backend O(N) pre-computed indexing for lightning-fast JSON responses.
* **Geolocalization & Mapping:** Converts textual concert locations into geographic coordinates using the OpenStreetMap Nominatim API. Visualizes these coordinates seamlessly using Leaflet.js.
* **Background Cache Warming:** Utilizes Go concurrency (`goroutines` and `sync.RWMutex`) to preload and cache API coordinate data in the background, ensuring 0-second map load times for users without hitting API rate limits.
* **Concurrent API Consumption:** Fetches and aggregates relational data across four distinct REST endpoints simultaneously using Go's `sync.WaitGroup`.
* **Centralized Error Management:** A custom struct-driven error handler that safely catches internal server panics and invalid routes, displaying a stylized UI instead of raw crash logs.
* **Zero Dependencies:** The backend relies entirely on the Go Standard Library.

## 🛠️ Tech Stack

* **Backend:** Go (Standard Library: `net/http`, `html/template`, `sync`, `encoding/json`)
* **Frontend:** HTML5, CSS3 (Modular Architecture), Vanilla JavaScript
* **Mapping:** Leaflet.js API
* **Geocoding:** OpenStreetMap (Nominatim API)

## 📂 Project Architecture (MVC)

The codebase is structured for scalability and separation of concerns:
* `/model`: Contains Go structs representing the unified data schemas.
* `/controllers`: Manages HTTP routing, template rendering, and error handling.
* `/control_utils`: Handles the heavy algorithms (search indexing, background geocoding, API aggregation).
* `/static`: Houses modular CSS stylesheets and frontend JavaScript logic.
* `/templates`: HTML templates using Go's native templating engine.

## ⚙️ Installation & Usage

1. Clone the repository to your local machine.
2. Ensure you have [Go installed](https://go.dev/dl/).
3. Navigate to the root directory of the project in your terminal.
4. Run the application:
   ```bash
   go run .

## 🚧 Upcoming Features (Roadmap)
Advanced Filtering System: A comprehensive multi-criteria filtering engine to sort artists by creation date, first album, member count, and specific concert locations.

### 👨‍💻 Author
Oduh Emmanuel Aba (Maddox-Bayn)
Creator of SpiriTech