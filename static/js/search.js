document.addEventListener("DOMContentLoaded", () => {
    const searchBar = document.getElementById("search-bar");
    const suggestionsList = document.getElementById("suggestions-list");

    // Listen for every time the user types a key
    searchBar.addEventListener("input", function() {
        const query = searchBar.value.trim();

        // If the box is empty, hide the dropdown and stop
        if (query.length === 0) {
            suggestionsList.innerHTML = "";
            suggestionsList.style.display = "none";
            return;
        }
        let debounceTimer;
        searchBar.addEventListener("input", function() {
            clearTimeout(debounceTimer);
            debounceTimer = setTimeout(() => {
                // your existing fetch(...) logic goes here
            }, 250);
        });


        // 1. Send the GET request to your Go backend
        fetch(`/search?q=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => {
                // Clear out the old suggestions
                suggestionsList.innerHTML = "";

                // If no results, hide the box
                if (!data || data.length === 0) {
                    suggestionsList.style.display = "none";
                    return;
                }

                // 2. Loop through the JSON data and build HTML for each item
                data.forEach(item => {
                    const li = document.createElement("li");
                    
                    // Notice we use item.artistid because of your struct tag `json:"artistid"`
                    li.innerHTML = `
                        <a href="/artist?id=${item.artistID}">
                            <span class="suggestion-text">${item.text}</span>
                            <span class="suggestion-type">-> ${item.type}</span>
                        </a>
                    `;
                    suggestionsList.appendChild(li);
                });

                // Show the dropdown now that it has data
                suggestionsList.style.display = "block";
            })
            .catch(error => console.error("Error fetching search data:", error));
    });

    // Optional: Hide dropdown if user clicks outside of it
    document.addEventListener("click", function(event) {
        if (!searchBar.contains(event.target) && !suggestionsList.contains(event.target)) {
            suggestionsList.style.display = "none";
        }
    });
}); 