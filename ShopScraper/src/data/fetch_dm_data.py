import requests
import time
import pandas as pd
import os
import json

# Define the base URL for the product search API
BASE_URL = "https://product-search.services.dmtech.com/si/search/crawl"

# Function to retrieve the available categories
def fetch_categories():
    url = f"{BASE_URL}?query=*" 
    response = requests.get(url)
    if response.status_code == 200:
        data = response.json()
        facets = data.get("facets", [])
        for facet in facets:
            if facet.get("key") == "categoryNames":
                categories = [value["name"] for value in facet.get("values", [])]
                return categories
    else:
        print("Error fetching categories")
    return []

# Function to fetch products for a specific category
def fetch_products_by_category(category, page_size=1000):
    print(f"Fetching products for category: {category}")
    time.sleep(3)  # To avoid hitting rate limits
    
    products = []
    url = f"{BASE_URL}?query=*&categoryName={category}&pageSize={page_size}"
    
    retries = 0
    while retries < 5:
        try:
            response = requests.get(url)
            if response.status_code == 200:
                data = response.json()
                category_products = data.get("products", [])
                products.extend(category_products)
                break  # success
            elif response.status_code == 429:
                wait_time = 10 * (retries + 2) 
                print(f"Rate limited. Waiting {wait_time:.2f} seconds before retrying...")
                time.sleep(wait_time)
                retries += 1
            else:
                print(f"Error {response.status_code} fetching category: {category}")
                break
        except Exception as e:
            print(f"Exception fetching category {category}: {e}")
            break

    return products

# Function to fetch all products for all categories
def fetch_all_products():
    categories = fetch_categories()  # Fetch all category names
    print("Categories fetched:", categories)
    all_products = []

    for category in categories:
        category_products = fetch_products_by_category(category)
        all_products.extend(category_products)
        print(f"Fetched {len(category_products)} products for {category}")
    
    return all_products

def save_to_json(products, filename="all_products.csv"):
    os.makedirs("data/raw/dm", exist_ok=True)

    file_path = "data/raw/dm/dm_data.json"
    with open(file_path, "w", encoding="utf-8") as file:
        json.dump(products, file, ensure_ascii=False, indent=2)
        
    print(f"Saved {len(products)} products to {filename}")

def main():
    all_products = fetch_all_products()
    save_to_json(all_products)  

if __name__ == "__main__":
    main()