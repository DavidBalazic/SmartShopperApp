import "./styles.css";
import "userApp/styles.css";
import "searchApp/styles.css";
import React, { Suspense, lazy } from "react";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";

// Load the remote component
const Login = lazy(() => import("userApp/Login"));
const Register = lazy(() => import("userApp/Register"));
const Dashboard = lazy(() => import("searchApp/Dashboard"));

const isAuthenticated = () => {
    const token = localStorage.getItem("auth_token");
    const expiration = new Date(localStorage.getItem("auth_token_expiration"));
    return token && expiration > new Date();
  };
  
  const App = () => (
    <BrowserRouter>
        <Suspense fallback={<div className="p-4">Loading...</div>}>
          <Routes>
            <Route path="/login/*" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/dashboard" element={<Dashboard />} />
            {/* <Route
              path="/history"
              element={isAuthenticated() ? <History /> : <Navigate to="/login" />}
            /> */}
            <Route path="*" element={<Navigate to="/login" />} />
          </Routes>
        </Suspense>
    </BrowserRouter>
  );
  
  export default App;