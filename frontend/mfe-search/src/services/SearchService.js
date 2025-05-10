const API_URL = process.env.REACT_APP_SEARCH_SERVICE;

class SearchService {
  /**
   * Searches for products matching a query and optional store.
   * @param {string} q - The search term (required).
   * @param {string} [store] - Optional store filter.
   * @returns {Promise<Array>} List of matching products.
   */
  async searchProducts(q, store = "") {
    if (!q) throw new Error("Query parameter 'q' is required");

    const params = new URLSearchParams({ q });
    if (store) params.append("store", store);

    const response = await fetch(`${API_URL}/search-products?${params.toString()}`, {
      method: "GET",
    });

    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.message || "Failed to search products");
    }

    return data;
  }

  /**
   * Gets the cheapest product for a query and optional store.
   * @param {string} q - The search term (required).
   * @param {string} [store] - Optional store filter.
   * @returns {Promise<Object>} Cheapest matching product.
   */
  async getCheapestProduct(q, store = "") {
    if (!q) throw new Error("Query parameter 'q' is required");

    const params = new URLSearchParams({ q });
    if (store) params.append("store", store);

    const response = await fetch(`${API_URL}/cheapest-product?${params.toString()}`, {
      method: "GET",
    });

    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.message || "Failed to fetch cheapest product");
    }

    return data;
  }
}

export default new SearchService();