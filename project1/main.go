package main
import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
)


type Product struct {
	ID          string
	Name        string
	Price       float64
	Description string
	ImageURL    string
	Category    string
}

var cartLock sync.Mutex
var cart = make(map[string]int)
// Handler for serving the homepage




func homePage(w http.ResponseWriter, r *http.Request) {
    // Parse home template
    tpl, err := template.New("home").Parse(homeTemplate)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }



 
	// Assume you have a list of products
var products = []Product{
	{ID: "1", Name: "Product 1", Price: 10.99, Description: "Description of Product 1", ImageURL: "/static/images/product1.jpg", Category: "clothing"},
    {ID: "2", Name: "Product 2", Price: 19.99, Description: "Description of Product 2", ImageURL: "/static/images/product2.jpg", Category: "clothing"},
	{ID: "3", Name: "Product 3", Price: 20.99, Description: "Description of Product 3", ImageURL: "/static/images/product3.jpg", Category: "clothing"},
	{ID: "4", Name: "Product 4", Price: 10.99, Description: "Description of Product 4", ImageURL: "/static/images/product1.jpg", Category: "clothing"},
	{ID: "5", Name: "Product 5", Price: 19.99, Description: "Description of Product 5", ImageURL: "/static/images/product8.jpg", Category: "clothing"},
	{ID: "6", Name: "Product 6", Price: 20.99, Description: "Description of Product 6", ImageURL: "/static/images/product9.jpg", Category: "clothing"},
	{ID: "7", Name: "Product 7", Price: 20.99, Description: "Description of Product 7", ImageURL: "/static/images/product10.jpg",Category: "clothing"},
	{ID: "8", Name: "Product 8", Price: 20.99, Description: "Description of Product 8", ImageURL: "/static/images/product7.jpg", Category: "clothing"},
	{ID: "9", Name: "Product 9", Price: 20.99, Description: "Description of Product 9", ImageURL: "/static/images/product5.jpg", Category: "clothing"},
	{ID: "10", Name: "Product 10", Price: 20.99, Description: "Description of Product 10", ImageURL: "/static/images/product11.jpg",Category: "electronics"},
	
}

    // Get the search query from the URL parameter "q"
    query := r.URL.Query().Get("q")


    // Get the selected category from the dropdown menu
    category := r.URL.Query().Get("category")

    // Filter products based on the search query and category
    filteredProducts := filterProductsByCategory(products, query, category)

    // Set content type to HTML
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    // Execute template of products
    if err := tpl.Execute(w, filteredProducts); err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}


// Function to filter products by category
func filterProductsByCategory(products []Product, query, category string) []Product {
    if category == "" {
        return filterProducts(products, query)
    }

    var filtered []Product
    for _, p := range products {
        // Check if the product matches the selected category
        if strings.ToLower(p.Category) == strings.ToLower(category) && strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) {
            filtered = append(filtered, p)
        }
    }
    return filtered
}

	



// Function to filter products based on search query
func filterProducts(products []Product, query string) []Product {
	if query == "" {
		return products
	}
	var filtered []Product
	for _, p := range products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
// Home template string
const homeTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>

    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to the Shop</title>
	<style>
	

    .Cart {
       
        color: purple;
        padding: 10px;
        text-align: center; /* Center align the text */
    }
    
    .product-container {
        display: flex;
        justify-content: center; /* Center align items horizontally */
        flex-wrap: wrap; /* Wrap items to the next row if needed */
    }
    
    .product {
        margin: 10px;
        padding: 10px;
        border: 1px solid #ccc;
        width: 200px;
        text-align: center; /* Center align the text */
		height: 60%;
    }

    .product img {
        max-width: 50%;
        height: auto;
        display: block; /* Ensures margin:auto works */
        margin: 0 auto; /* Center aligns the image horizontally */
    }

   
   
	
	.product-container .product[data-id="1"] img {
        max-width: 50%; /* Set a larger max-width for the image */
		margin-top:10%;
		
		
    }
	.product-container .product[data-id="2"] img {
        max-width: 70%; /* Set a larger max-width for the image */
		margin-top:12%;

    }
	.product-container .product[data-id="3"] img {
        max-width: 40%; /* Set a larger max-width for the image */
	}
	.product-container .product[data-id="4"] img {
        max-width: 50%; /* Set a larger max-width for the image */
		margin-top:10%;
	

    }
	.product-container .product[data-id="5"] img {
        max-width: 36%; /* Set a larger max-width for the image */
		
	
		
    }

	.product-container .product[data-id="6"] img {
		max-width: 35%; /* Set a larger max-width for the image */
		margin-top:10%;
			
		}
	.product-container .product[data-id="7"] img {
		max-width: 50%; /* Set a larger max-width for the image */
		margin-top:12%;
		}
	.product-container .product[data-id="8"] img {
		max-width: 45%; /* Set a larger max-width for the image */
		margin-top: 15%;
		}
	.product-container .product[data-id="9"] img {
		max-width: 50%; /* Set a larger max-width for the image */
		margin-top:10%;
		}
	.product-container .product[data-id="10"] img {
		max-width: 65%; /* Set a larger max-width for the image */
		margin-top:8%;
		}

		/* Center the search form */
		form {
			text-align: center;
			margin-top: 20px; /* Adjust the margin-top as needed */
		}
	
		/* Increase the width of the input field */
		input[type="text"] {
			width: 60%; /* Adjust the width as needed */
			padding: 10px; /* Add padding to make the input field taller */
			font-size: 16px; /* Increase font size */
		}
	
		/* Style the search button */
		button[type="submit"] {
			padding: 10px 20px; /* Add padding to make the button taller */
			font-size: 16px; /* Increase font size */
			background-color: purple; /* Change the background color */
			color: white; /* Change text color */
			border: none; /* Remove border */
			cursor: pointer; /* Add pointer cursor on hover */
		}
	
		button[type="submit"]:hover {
			background-color: red; /* Darken the background color on hover */
		}
		/* CSS styles for the search form */
        .search-container {
            background-color: yellow; /* Set background color to yellow */
            padding: 10px; /* Add padding for spacing */
        }

        .search-form {
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .search-form input[type="text"] {
            width: 60%;
            padding: 10px;
            font-size: 16px;
            border: 1px solid #ccc;
            border-radius: 4px;
        }

        .search-form button[type="submit"] {
            padding: 10px 20px;
            font-size: 16px;
            background-color: #f0c040;
            color: white;
            border: none;
            cursor: pointer;
        }

        .search-form button[type="submit"]:hover {
            background-color: #e0b030;
        }

        /* CSS styles for the cart icon */
        .cart-icon {
            position: absolute;
            top: 10px;
            right: 10px;
			color: black;
			width:50px;
        }
		
		/* Adjust the width of the search input */
    input[type="text"] {
        width: 100%; /* Change width to 100% to occupy the entire width of the search form */
        padding: 10px;
        font-size: 16px;
        border: 1px solid #ccc;
        border-radius: 4px;
    }
			
		
		
		.dropdown {
			margin-right: 10px; /* Adjust spacing between dropdown and search bar */
		}
		
		/* Style dropdown menu */
		.dropdown select {
			padding: 8px;
			font-size: 14px;
			border: 1px solid #ccc;
			border-radius: 4px;
		}
		
		/* Style search input */
		#search {
			flex: 1;
			width: 20px; /* Adjust the width as needed */
			padding: 8px;
			font-size: 14px;
			border: 1px solid #ccc;
			border-radius: 4px;
		}
		
		/* Style search button */
		button[type="submit"] {
			padding: 8px 16px;
			font-size: 14px;
			border: none;
			border-radius: 4px;
			background-color: green;
			color: #fff;
			cursor: pointer;
			
		}
		
		button[type="submit"]:hover {
			background-color: pink;
		}
		nav {
            background-color: #454545; /* Background color of the navbar */
            color: red; /* Text color of the navbar links */
            padding: 10px 20px; /* Padding inside the navbar */
        }

		

        nav ul {
            list-style-type: none; /* Remove bullet points from the list */
            margin: 0; /* Remove default margin */
            padding: 0; /* Remove default padding */
        }

        nav ul li {
            display: inline; /* Display list items horizontally */
            margin-right: 20px; /* Add spacing between navbar links */
        }

        nav ul li a {
            color: green; /* Text color of the navbar links */
            text-decoration: none; 
            font-weight: bold; 
        }

        nav ul li a:hover {
            text-decoration: underline; 
		}
		/* Your existing styles */
        .cart-icon {
            position: absolute;
            top: 10px;
            right: 10px;
            color: black; /* Set color of the icon */
        }
    </style>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css" integrity="sha384-Nz4hp2VBRr9IcKFHgkaDmTdi6L8rRVLPQ5Exmt9KZRFy9x8sXr6kNM6dKeQTKMT4" crossorigin="anonymous">
</head>
<body>




<!-- Your existing content -->
<h1 class="Cart">Welcome to the Shop</h1>

<!-- Add the cart icon -->
<div class="search-container">
	<a href="/cart" class="cart-icon">
		<i class="fas fa-shopping-cart fa-2x"></i> <!-- Font Awesome cart icon -->
	</a>

	<!-- Add the search form -->
	<form action="/" method="GET" class="search-form">
		<div class="dropdown">
			<select name="category">
				<option value="">All Categories</option>
				<option value="electronics">Electronics</option>
				<option value="clothing">Clothing</option>
				<option value="books">Books</option>
				<!-- Add more options as needed -->
			</select>
		</div>
		<input type="text" id="search" name="q">
		<button type="submit">Search</button>
	</form>
</div>


<nav>
<ul>
<li><a href="/">Home</a></li>
<li><a href="/cart">Cart</a></li>
<!-- Add more navigation links as needed -->
</ul>
</nav>

    <div class="product-container">
        {{ range . }}
		<div class="product" data-id="{{ .ID }}"> <!-- Add data-id attribute here -->
        <img src="{{ .ImageURL }}" alt="{{ .Name }}">
		
      
            <h2>{{ .Name }}</h2>
            <p>{{ .Description }}</p>
            <p>Price: ${{ .Price }}</p>
            <form action="/addToCart" method="post">
                <input type="hidden" name="productId" value="{{ .ID }}">
                <button type="submit">Add to Cart</button>
            </form>
        </div>
        {{ end }}
    </div>
    <a href="/cart">View Cart</a>
    <div id="cart"></div>

    <script>
        var ws = new WebSocket("ws://" + window.location.host + "/ws");
        ws.onmessage = function(event) {
            document.getElementById('cart').innerText = event.data;
        };
    </script>
</body>
</html>

`

// Handler for serving the cart page
func cartPage(w http.ResponseWriter, r *http.Request) {
	// Parse cart template
	tpl, err := template.New("cart").Parse(cartTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Lock the cart for reading
	cartLock.Lock()
	defer cartLock.Unlock()

	// Assume you have maps of products and their prices
	// This assumes that products and prices are stored elsewhere in your application
	products := map[string]string{
		"1": "Product 1",
		"2": "Product 2",
		"3": "Product 3",
		"4": "Product 4",
		"5": "Product 5",
		"6": "Product 6",
		"7": "Product 6",
		"8": "Product 6",
		"9": "Product 6",
		"10": "Product 6",
	}
	prices := map[string]float64{
		"1": 10.99,
		"2": 19.99,
		"3": 10.99,
		"4": 60.99,
		"5": 14.99,
		"6": 32.99,
		"7": 32.99,
		"8": 32.99,
		"9": 32.99,
		"10": 32.99,

	}

	// Prepare data for rendering the cart
type CartItem struct {
    ProductID string //stores the id of product
    Quantity  int    //stores quantity of product in cart
    ImageURL  string //stores the URL of the product image
}
var cartItems []CartItem //Declares a slice named cartItems to store multiple CartItem instances.
//This slice will hold the cart data in a format suitable for rendering in the template.

// Fixing the loop to populate cartItems with product details
for productId, quantity := range cart {
    product := getProductByID(productId)
    cartItems = append(cartItems, CartItem{
        ProductID: productId,
        Quantity:  quantity,
        ImageURL:  product.ImageURL,
    })
}


	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute template
	if err := tpl.Execute(w, struct {
		Products map[string]string
		Prices   map[string]float64
		Cart     []CartItem
	}{
		Products: products,
		Prices:   prices,
		Cart:     cartItems,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}




// Cart template string
const cartTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>View Cart</title>
</head>
<body>
	<h1>Cart</h1>
	<table border="1">
		<tr>
			<th>ID</th>
			<th>Name</th>
			<th>Price</th>
			<th>Quantity</th>
			<th>Action</th> <!-- New column for delete button -->
		</tr>
		{{ range .Cart }}
<tr>
    <td><img src="{{ .ImageURL }}" alt="{{ index $.Products .ProductID }}"></td>
    <td>{{ index $.Products .ProductID }}</td>
    <td>{{ index $.Prices .ProductID }}</td>
    <td>{{ .Quantity }}</td>
    <td>
        <form action="/removeFromCart" method="post">
            <input type="hidden" name="productId" value="{{ .ProductID }}">
            <button type="submit">Delete</button>
        </form>
    </td>
</tr>
{{ end }}

	</table>
	<a href="/">Continue Shopping</a>
</body>
</html>
`
func getProductByID(id string) Product {
    // Implement logic to fetch product details from your data source
    // For simplicity, I'm returning a hardcoded product
    switch id {
    case "1":
        return Product{
            ID:          "1",
            Name:        "Product 1",
            Price:       10.99,
            Description: "Description of Product 1",
            ImageURL:    "/static/images/product1.jpg",
        }
    case "2":
        return Product{
            ID:          "2",
           Price:       19.99,
            Description: "Description of Product 2",
            ImageURL:    "/static/images/product2.jpg",
        }
	case "3":
        return Product{
            ID:          "3",
            Name:        "Product 3",
            Price:       40.99,
            Description: "Description of Product 3",
            ImageURL:    "/static/images/product3.jpg",
        }
	case "4":
        return Product{
            ID:          "4",
            Name:        "Product 4",
            Price:       10.99,
            Description: "Description of Product 4",
            ImageURL:    "/static/images/product1.jpg",
        }
    case "5":
        return Product{
            ID:          "5",
			Name:        "Product 5",
           Price:       19.99,
            Description: "Description of Product 5",
            ImageURL:    "/static/images/product2.jpg",
        }
	case "6":
        return Product{
            ID:          "6",
            Name:        "Product 6",
            Price:       40.99,
            Description: "Description of Product 6",
            ImageURL:    "/static/images/product3.jpg",
        }
	case "7":
        return Product{
            ID:          "7",
            Name:        "Product 7",
            Price:       45.99,
            Description: "Description of Product 7",
            ImageURL:    "/static/images/product3.jpg",
        }
	case "8":
        return Product{
            ID:          "8",
            Name:        "Product 8",
            Price:       40.99,
            Description: "Description of Product 8",
            ImageURL:    "/static/images/product3.jpg",
        }
	case "9":
        return Product{
            ID:          "9",
            Name:        "Product 9",
            Price:       40.99,
            Description: "Description of Product 9",
            ImageURL:    "/static/images/product3.jpg",
        }
	case "10":
        return Product{
            ID:          "10",
            Name:        "Product 10",
            Price:       40.99,
            Description: "Description of Product 10",
            ImageURL:    "/static/images/product3.jpg",
        }
    default:
        // Handle unknown product ID
        return Product{}
    }
}

// Handler for adding a product to the cart
func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from the form submission
	productId := r.FormValue("productId")

	// Lock the cart for writing
	cartLock.Lock()
	defer cartLock.Unlock()

	// Increment the quantity of the product in the cart
	cart[productId]++

	// Redirect back to the product page or any other page
	http.Redirect(w, r, "/", http.StatusFound) // Change the redirect URL according to your application's logic
}


// Handler for serving the product page
func productPageHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}


// Handler for removing a product from the cart
func removeFromCartHandler(w http.ResponseWriter, r *http.Request) {
    // Get the product ID from the form submission
    productId := r.FormValue("productId")

    // Lock the cart for writing
    cartLock.Lock()
    defer cartLock.Unlock()

    // Remove the product from the cart
    delete(cart, productId)

    // Redirect back to the cart page
    http.Redirect(w, r, "/cart", http.StatusFound)
}

func main() {
	// Register static file server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Register handlers
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cart", cartPage)
	http.HandleFunc("/product", productPageHandler)
	http.HandleFunc("/addToCart", addToCartHandler)
	http.HandleFunc("/removeFromCart", removeFromCartHandler) // New handler for deleting from cart

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
