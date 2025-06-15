import axios from "axios";

const API_BASE = process.env.REACT_APP_USER_SERVICE_URL;

const HistoryService = {
  async getGroceryListHistory(token) {
    const res = await axios.get(`${API_BASE}/groceryList`, {
      headers: {
         Authorization: `Bearer ${token}`,
      },
    });
    return res.data;
  },
};

export default HistoryService;
