import pandas as pd
import json
import os

def preprocess_dm_data():
    with open("data/raw/dm/dm_data.json", "r", encoding="utf-8") as file:
        data = json.load(file)

    products = []
  
    for item in data:
        try:
            gtin = item.get("gtin")
            name = item.get("name")
            description = ""
            discounted_price = ""
            regular_price = item.get('price', {}).get('value')
            amount = item.get('netQuantityContent')
            unit = item.get("contentUnit")
            price_per_unit = item.get('basePrice', {}).get('formattedValue')
            base_price_unit = item.get("basePriceUnit")
            base_price_amount = item.get("basePriceQuantity")
            
            products.append({
                "gtin": gtin,
                "name": name,
                "description": description,
                "discounted_price": discounted_price,
                "price": regular_price,
                "amount": amount,
                "unit": unit,
                "price_per_unit": price_per_unit,
                "base_price_unit": base_price_unit,
                "base_price_amount": base_price_amount
            })
            
        except Exception as e:
            print(f"Error processing item {item.get('gtin')}: {e}")
    
    os.makedirs("data/preprocessed/dm", exist_ok=True)
    
    df = pd.DataFrame(products)    
    df.drop_duplicates(subset=["gtin"], inplace=True)
    df.reset_index(drop=True, inplace=True)
    df.drop(columns=["gtin"], inplace=True)
    
    file_path = "data/preprocessed/dm/dm_data.csv" 
    df.to_csv(file_path, index=False)
    print(f"Saved: {file_path}")

if __name__ == "__main__":
    preprocess_dm_data()