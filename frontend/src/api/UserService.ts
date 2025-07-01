import axios, { AxiosInstance, AxiosResponse } from "axios";

const API_URL = import.meta.env.VITE_USER_SERVICE_URL;

if (!API_URL) {
  throw new Error("VITE_USER_SERVICE environment variable is not set");
}

export interface GroceryListSummary {
  id: string;
  name: string;
  createdAt: string;
}

export interface RegisterPayload {
  username: string;
  email: string;
  password: string;
}

export interface LoginPayload {
  username: string;
  password: string;
}

export interface AuthResponse {
  token: string;
  expiration: string; // ISO date string
}

class UserService {
  private api: AxiosInstance;
  private token: string | null = null;
  private tokenExpiration: string | null = null;

  constructor() {
    this.api = axios.create({
      baseURL: API_URL,
      headers: {
        "Content-Type": "application/json",
      },
    });

    const savedToken = localStorage.getItem("auth_token");
    const savedExpiration = localStorage.getItem("auth_token_expiration");

    if (savedToken && savedExpiration && new Date(savedExpiration) > new Date()) {
      this.token = savedToken;
      this.tokenExpiration = savedExpiration;
      this.setAuthHeader(savedToken);
    }
  }

  /**
   * Get the grocery list history for the current user.
   * @param token - The JWT token for authorization.
   * @returns Array of grocery list history items.
   */
  async getGroceryListHistory(token: string): Promise<GroceryListSummary[]> {
    if (!token) throw new Error("Authorization token is required");

    const response: AxiosResponse<GroceryListSummary[]> = await this.api.get("/groceryList", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    return response.data;
  }

  private setAuthHeader(token: string) {
    this.api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  }

  private clearAuthHeader() {
    delete this.api.defaults.headers.common["Authorization"];
  }

  async register(payload: RegisterPayload): Promise<AuthResponse> {
    const response: AxiosResponse<AuthResponse> = await this.api.post("/register", payload);
    return response.data;
  }

  async login(payload: LoginPayload): Promise<{ username: string }> {
    const response: AxiosResponse<AuthResponse & { username: string }> = await this.api.post("/login", payload);
    const { token, expiration, username } = response.data;

    this.token = token;
    this.tokenExpiration = expiration;

    localStorage.setItem("auth_token", token);
    localStorage.setItem("auth_token_expiration", expiration);

    this.setAuthHeader(token);

    return { username };
  }

  logout(): void {
    this.token = null;
    this.tokenExpiration = null;

    localStorage.removeItem("auth_token");
    localStorage.removeItem("auth_token_expiration");

    this.clearAuthHeader();
  }

  setToken(token: string): void {
    this.token = token;
    this.setAuthHeader(token);
  }

  clearToken(): void {
    this.token = null;
    this.clearAuthHeader();
  }

  getToken(): string | null {
    return this.token;
  }

  getTokenExpiration(): string | null {
    return this.tokenExpiration;
  }

  isTokenValid(): boolean {
    return !!(this.token && this.tokenExpiration && new Date(this.tokenExpiration) > new Date());
  }
}

export default new UserService();