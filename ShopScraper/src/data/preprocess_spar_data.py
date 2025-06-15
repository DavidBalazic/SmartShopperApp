import pandas as pd
import json
import os

def preprocess_spar_data():
    with open("data/raw/spar/spar_data.json", "r", encoding="utf-8") as file:
        data = json.load(file)

    hits = data.get("hits", [])
    products = []
    # TODO: unit is sales unit and not conent unit, also there is no amount data
    for item in hits:
        try:
            values = item.get("masterValues", {})

            name = values.get("title")
            description = values.get("description")
            discounted_price = values.get("best-price")
            regular_price = values.get("regular-price")
            unit = values.get("sales-unit")
            price_per_unit = values.get("price-per-unit")
            price_per_unit_number = values.get("price-per-unit-number")
            image_url = values.get("image-url")
            products.append({
                "name": name,
                "store": "Spar",
                #"description": description,
                #"discounted_price": discounted_price,
                "price": regular_price,
                "unit": unit,
                #"price_per_unit": price_per_unit,
                "price_per_unit": price_per_unit_number,
                "image_url": image_url
            })
            
        except Exception as e:
            print(f"Error processing item {item.get('code')}: {e}")
    
    os.makedirs("data/preprocessed/spar", exist_ok=True)
    df = pd.DataFrame(products)    
    
    # Drop duplicates
    df = df.drop_duplicates(subset=["name", "store", "price"])
    
    file_path = "data/preprocessed/spar/spar_data.csv" 
    df.to_csv(file_path, index=False)
    print(f"Saved: {file_path}")

if __name__ == "__main__":
    preprocess_spar_data()