const API_URL = process.env.REACT_APP_USER_SERVICE;

class UserService {
    token = null;
    tokenExpiration = null;

    constructor() {
        const savedToken = localStorage.getItem("auth_token");
        const savedExpiration = localStorage.getItem("auth_token_expiration");
    
        if (savedToken && savedExpiration && new Date(savedExpiration) > new Date()) {
          this.token = savedToken;
          this.tokenExpiration = savedExpiration;
        }
    }

    async register({ username, email, password }) {
        const response = await fetch(`${API_URL}/register`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ username, email, password }),
        });
    
        const data = await response.json();
    
        if (!response.ok) {
          throw new Error(data.message || "Registration failed");
        }
    
        return data; 
    }
  
    async login({ username, password }) {
        const response = await fetch(`${API_URL}/login`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ username, password }),
        });
    
        const data = await response.json();
    
        if (!response.ok) {
          throw new Error(data.message || "Login failed");
        }
    
        this.token = data.token;
        this.tokenExpiration = data.expiration;
        localStorage.setItem("auth_token", data.token);
        localStorage.setItem("auth_token_expiration", data.expiration);
    
        return { username };
    }

    logout() {
        this.token = null;
        this.tokenExpiration = null;
        localStorage.removeItem("auth_token");
        localStorage.removeItem("auth_token_expiration");
    }

    setToken(token) {
      this.token = token;
    }

    clearToken() {
      this.token = null;
    }

    getToken() {
      return this.token;
    }
}

export default new UserService();