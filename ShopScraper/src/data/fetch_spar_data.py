import requests
from datetime import datetime
import json
import os

def fetch_spar_data():
    try:
        url = "https://search-spar.spar-ics.com/fact-finder/rest/v4/search/products_lmos_si?q=*&query=*&hitsPerPage=17073&substringFilter=pos-visible:81701&q1=&x1=product-lifestyleInf"

        response = requests.get(url)
        response.raise_for_status()  # Raise an exception for HTTP errors

        data = response.json()

        # Ensure directory exists
        os.makedirs("data/raw/spar", exist_ok=True)

        # Save the JSON data to a file
        file_path = "data/raw/spar/spar_data.json"
        with open(file_path, "w", encoding="utf-8") as file:
            json.dump(data, file, ensure_ascii=False, indent=2)

        print(f"Fetching successful. Data saved to {file_path} at {datetime.now()}")

    except requests.RequestException as e:
        # Print error message if there is a problem fetching the file
        print(f"Error fetching data: {e}")

if __name__ == "__main__":
    fetch_spar_data()