import React, { useState, useEffect } from "react";
import SearchService from "../services/SearchService";
import HistorySidebar from "./HistorySidebar";
import { Search, ShoppingBag, Store, X, Plus, ShoppingCart, Filter } from "lucide-react";

const Dashboard = () => {
  const [query, setQuery] = useState("");
  const [store, setStore] = useState("");
  const [cheapest, setCheapest] = useState(null);
  const [allResults, setAllResults] = useState([]);
  const [shoppingList, setShoppingList] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [activeTab, setActiveTab] = useState("products");

  // History sidebar state
  const [historyOpen, setHistoryOpen] = useState(false);

  const handleGroceryListItemClick = (historyItem) => {
    // TODO: implement click on grocery list item
  };

  const handleSearchCheapest = async () => {
    if (!query.trim()) return;

    setLoading(true);
    setError("");
    setAllResults([]);
    try {
      const result = await SearchService.getCheapestProduct(query, store);
      setCheapest(result);
    } catch (err) {
      setError(err.message);
      setCheapest(null);
    } finally {
      setLoading(false);
    }
  };

  const handleShowAll = async () => {
    setLoading(true);
    setError("");
    try {
      const results = await SearchService.searchProducts(query, store);
      setAllResults(results);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const addToShoppingList = (product) => {
    setShoppingList((prev) => [...prev, product]);
  };

  const removeFromShoppingList = (indexToRemove) => {
    setShoppingList((prev) => prev.filter((_, index) => index !== indexToRemove));
  };

  const totalPrice = shoppingList.reduce((acc, item) => acc + parseFloat(item.price), 0).toFixed(2);
  const uniqueStores = [...new Set(shoppingList.map((item) => item.store))];

  return (
    <div className="min-h-screen bg-gradient-to-br from-indigo-100 to-blue-100 py-8">
      <div className="max-w-7xl mx-auto px-4 relative">
        {/* History Sidebar */}
        <HistorySidebar
          historyOpen={historyOpen}
          setHistoryOpen={setHistoryOpen}
          onHistoryListItemClick={handleGroceryListItemClick}
        />

        {/* Main Content */}
        <div className={`transition-all duration-300 ${historyOpen ? 'lg:ml-80' : 'ml-0'}`}>
          <h1 className="text-3xl font-bold text-gray-800 mb-2 text-center">Shopping Dashboard</h1>
          <p className="text-gray-600 text-center mb-8">Find the best deals across different stores</p>
          
          {/* Tabs for small screens */}
          <div className="lg:hidden mb-6 flex justify-center gap-4">
            <button
              className={`px-4 py-2 rounded-lg shadow flex items-center gap-2 transition ${activeTab === "products" ? "bg-blue-600 text-white" : "bg-white text-gray-700"}`}
              onClick={() => setActiveTab("products")}
            >
              <ShoppingBag size={18} />
              <span>Products</span>
            </button>
            <button
              className={`px-4 py-2 rounded-lg shadow flex items-center gap-2 transition ${activeTab === "shopping" ? "bg-blue-600 text-white" : "bg-white text-gray-700"}`}
              onClick={() => setActiveTab("shopping")}
            >
              <ShoppingCart size={18} />
              <span>Shopping List</span>
            </button>
          </div>

          <div className="flex flex-col lg:flex-row gap-8">
            {/* Left Side - Products */}
            <div className={`lg:w-2/3 ${activeTab !== "products" ? "hidden lg:block" : ""}`}>
              <div className="bg-white rounded-2xl shadow-xl overflow-hidden mb-8">
                <div className="bg-gradient-to-r from-blue-600 to-indigo-700 p-6">
                  <h2 className="text-2xl font-bold text-white">Find Products</h2>
                  <p className="text-blue-100">Search across multiple stores for the best prices</p>
                </div>
                
                <div className="p-6">
                  <div className="flex flex-col md:flex-row gap-4 mb-6">
                    <div className="relative flex-1">
                      <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <Search size={18} className="text-gray-400" />
                      </div>
                      <input
                        type="text"
                        placeholder="Search products..."
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      />
                    </div>
                    
                    <div className="relative md:w-1/3">
                      <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <Store size={18} className="text-gray-400" />
                      </div>
                      <select
                        value={store}
                        onChange={(e) => setStore(e.target.value)}
                        className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 appearance-none bg-white"
                      >
                        <option value="">All Stores</option>
                        <option value="Spar">Spar</option>
                        <option value="Lidl">Lidl</option>
                        <option value="Dm">Dm</option>
                      </select>
                      <div className="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                        <Filter size={18} className="text-gray-400" />
                      </div>
                    </div>
                    
                    <button
                      onClick={handleSearchCheapest}
                      className="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition shadow-md flex items-center justify-center gap-2"
                    >
                      <Search size={18} />
                      <span>Search</span>
                    </button>
                  </div>

                  {loading && (
                    <div className="flex justify-center py-8">
                      <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
                    </div>
                  )}
                  
                  {error && (
                    <div className="bg-red-50 text-red-600 px-4 py-3 rounded-lg mb-6 flex items-center">
                      <svg className="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 20 20">
                        <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                      </svg>
                      {error}
                    </div>
                  )}

                  {/* Cheapest option display */}
                  {cheapest && (
                    <div className="mb-6 border-l-4 border-green-500 bg-green-50 rounded-lg p-5 shadow-sm flex items-center justify-between">
                      <div className="flex-1 pr-4">
                        <h3 className="font-bold text-lg text-gray-800">Best Price Found:</h3>
                        <p className="text-gray-800 font-medium mb-1">{cheapest.name}</p>
                        <div className="flex items-center gap-2 mb-1">
                          <Store size={16} className="text-gray-600" />
                          <p className="text-gray-600">{cheapest.store}</p>
                        </div>
                        <p className="text-sm text-gray-600 mb-2">{cheapest.description}</p>
                        <div className="flex items-baseline gap-3">
                          <p className="text-lg font-bold text-green-600">{cheapest.price} €</p>
                          <p className="text-sm text-gray-500">
                            {cheapest.quantity} {cheapest.unit} — {cheapest.pricePerUnit} €/unit
                          </p>
                        </div>
                        <button
                          onClick={handleShowAll}
                          className="mt-3 bg-gray-800 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition text-sm shadow-sm flex items-center gap-1"
                        >
                          <Filter size={16} />
                          Show all options
                        </button>
                      </div>
                      <button
                        onClick={() => addToShoppingList(cheapest)}
                        className="bg-green-600 text-white h-10 w-10 rounded-full hover:bg-green-700 transition flex items-center justify-center shadow-md"
                      >
                        <Plus size={20} />
                      </button>
                    </div>
                  )}

                  {/* All results */}
                  {allResults.length > 0 && (
                    <div className="space-y-4">
                      {allResults.map((item, index) => (
                        <div
                          key={index}
                          className="border border-gray-200 rounded-lg p-4 bg-white shadow-sm hover:shadow-md transition flex justify-between items-center"
                        >
                          <div>
                            <h3 className="font-bold text-lg text-gray-800">{item.name}</h3>
                            <p className="text-sm text-gray-600 mb-2">{item.description}</p>
                            <div className="grid grid-cols-2 gap-x-8 gap-y-1 text-sm">
                              <div className="flex items-center gap-2">
                                <Store size={14} className="text-gray-500" />
                                <p>{item.store}</p>
                              </div>
                              <div className="flex items-center gap-2">
                                <ShoppingBag size={14} className="text-gray-500" />
                                <p>{item.quantity} {item.unit}</p>
                              </div>
                              <div className="flex items-center gap-2">
                                <span className="font-bold text-green-600 text-base">{item.price} €</span>
                              </div>
                              <div className="flex items-center gap-2">
                                <p className="text-gray-500">{item.pricePerUnit} €/unit</p>
                              </div>
                            </div>
                          </div>
                          <button
                            onClick={() => addToShoppingList(item)}
                            className="bg-green-600 text-white h-9 w-9 rounded-full hover:bg-green-700 transition flex items-center justify-center shadow-sm ml-4"
                          >
                            <Plus size={18} />
                          </button>
                        </div>
                      ))}
                    </div>
                  )}
                </div>
              </div>
            </div>

            {/* Right Side - Shopping List */}
            <div className={`lg:w-1/3 ${activeTab !== "shopping" ? "hidden lg:block" : ""}`}>
              <div className="bg-white rounded-2xl shadow-xl overflow-hidden sticky top-8">
                <div className="bg-gradient-to-r from-indigo-600 to-purple-700 p-6">
                  <h2 className="text-2xl font-bold text-white">Shopping List</h2>
                  <p className="text-indigo-100">Your selected items</p>
                </div>
                
                <div className="p-6">
                  {shoppingList.length === 0 ? (
                    <div className="text-center py-8">
                      <ShoppingCart size={48} className="mx-auto text-gray-300 mb-3" />
                      <p className="text-gray-500">Your shopping list is empty</p>
                      <p className="text-gray-400 text-sm mt-1">Add items from the search results</p>
                    </div>
                  ) : (
                    <div>
                      <ul className="space-y-3 mb-6">
                        {shoppingList.map((item, index) => (
                          <li key={index} className="border border-gray-200 p-3 rounded-lg flex justify-between items-center hover:bg-gray-50 transition">
                            <div>
                              <p className="font-medium text-gray-800">{item.name}</p>
                              <div className="flex items-center gap-2 text-sm text-gray-600">
                                <Store size={14} className="text-gray-500" />
                                <span>{item.store}</span>
                                <span className="font-semibold text-green-600">{item.price} €</span>
                              </div>
                            </div>
                            <button
                              onClick={() => removeFromShoppingList(index)}
                              className="text-red-600 hover:text-red-800 hover:bg-red-50 rounded-full h-8 w-8 flex items-center justify-center transition"
                            >
                              <X size={18} />
                            </button>
                          </li>
                        ))}
                      </ul>
                      
                      <div className="border-t border-gray-200 pt-4 space-y-2">
                        <div className="flex justify-between items-center">
                          <span className="text-gray-600">Total Items:</span>
                          <span className="font-medium">{shoppingList.length}</span>
                        </div>
                        <div className="flex justify-between items-center">
                          <span className="text-gray-600">Stores:</span>
                          <span className="font-medium">{uniqueStores.join(", ") || "None"}</span>
                        </div>
                        <div className="flex justify-between items-center text-lg font-bold">
                          <span>Total:</span>
                          <span className="text-green-600">{totalPrice} €</span>
                        </div>
                        
                        <button className="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition shadow-md mt-4 font-medium">
                          Save Shopping List
                        </button>
                      </div>
                    </div>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
