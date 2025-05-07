import requests
from datetime import datetime
import time
import json
import os

def fetch_lidl_data():
    try:
        all_products = []
        offset = 0
        fetchsize = 108  # Actual working limit

        while True:
            params = {
                "q": "",
                "offset": offset,
                "version": "v2.0.0",
                "assortment": "SI",
                "locale": "sl_SI",
                "fetchsize": fetchsize
            }

            response = requests.get("https://www.lidl.si/q/api/search", params=params)
            response.raise_for_status() 
            data = response.json()

            items = data.get("items", [])
            if not items:
                break

            all_products.extend(items)
            print(f"Fetched {len(items)} items (Total: {len(all_products)})")
            time.sleep(0.5)

            offset += fetchsize
            if offset >= data.get("numFound", 0):
                break
            
        print(f"Total products fetched: {len(all_products)}") 
        
        os.makedirs("data/raw/lidl", exist_ok=True)
        
        file_path = "data/raw/lidl/lidl_data.json"
        with open(file_path, "w", encoding="utf-8") as file:
            json.dump(all_products, file, ensure_ascii=False, indent=2)

        print(f"Fetching successful. Data saved to {file_path} at {datetime.now()}")
        
    except requests.exceptions.RequestException as e:
        print(f"Error fetching data: {e}")
    
if __name__ == "__main__":
    fetch_lidl_data()