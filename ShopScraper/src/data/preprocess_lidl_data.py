import pandas as pd
from bs4 import BeautifulSoup
import json


def preprocess_lidl_data():
    with open("data/raw/lidl/lidl_data.json", "r", encoding="utf-8") as file:
        data = json.load(file)

    products = []
    for item in data:
        try:
            gridbox = item.get("gridbox", {})
            details = gridbox.get("data", {})
            price_info = details.get("price", {})
            packaging = price_info.get("packaging", {}).get("text", "")
            
            # Discount logic
            old_price = price_info.get("oldPrice", 0.0)
            current_price = price_info.get("price", None)

            if old_price == 0.0:
                price = current_price
                discounted_price = None
            else:
                price = old_price
                discounted_price = current_price

            product = {
                "name": details.get("fullTitle", ""),
                "discounted_price": discounted_price,
                "price": price,
                #"amount": 1,
                "unit": packaging,
                "price_per_unit": price_info.get("basePrice", {}).get("text", ""),
                # "description": keyfacts.get("description", "")
                #     .replace("<ul>", "")
                #     .replace("</ul>", "")
                #     .replace("<li>", "- ")
                #     .replace("</li>", "\n")
                #     .strip()
            }
            products.append(product)
        except Exception as e:
            print(f"Error processing item {item.get('code')}: {e}")
    
    df = pd.DataFrame(products)     
    csv_path = f"data/preprocessed/lidl/lidl_data.csv"
    df.to_csv(csv_path, index=False)
    print(f"Saved: {csv_path}")

if __name__ == "__main__":
    preprocess_lidl_data()