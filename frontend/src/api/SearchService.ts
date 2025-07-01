import axios, { AxiosInstance, AxiosResponse } from "axios";

const API_URL = import.meta.env.VITE_SEARCH_SERVICE_URL;

if (!API_URL) {
  throw new Error("VITE_APP_SEARCH_SERVICE environment variable is not set");
}

// Define types for product
export interface Product {
  name: string;
  price: number;
  unit: string;
  store: string;
  pricePerUnit: number;
  imageUrl: string;
}

class SearchService {
  private api: AxiosInstance;

  constructor() {
    this.api = axios.create({
      baseURL: API_URL,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  /**
   * Searches for products matching a query and optional store.
   * @param q - The search term (required).
   * @param store - Optional store filter.
   * @returns List of matching products.
   */
  async searchProducts(q: string, store: string = ""): Promise<Product[]> {
    if (!q) throw new Error("Query parameter 'q' is required");

    const params: Record<string, string> = { q };
    if (store) params.store = store;

    const response: AxiosResponse<Product[]> = await this.api.get("/search-products", { params });

    return response.data;
  }

  /**
   * Gets the cheapest product for a query and optional store.
   * @param q - The search term (required).
   * @param store - Optional store filter.
   * @returns Cheapest matching product.
   */
  async getCheapestProduct(q: string, store: string = ""): Promise<Product> {
    if (!q) throw new Error("Query parameter 'q' is required");

    const params: Record<string, string> = { q };
    if (store) params.store = store;

    const response: AxiosResponse<Product> = await this.api.get("/cheapest-product", { params });

    return response.data;
  }
}

export default new SearchService();
