import React, { useEffect, useState } from 'react';
import { History, X, Calendar, Menu } from 'lucide-react';
import UserService from '../services/UserService';

const groupHistoryByDate = (history) => {
  const today = new Date();
  const startOfToday = new Date(today.setHours(0, 0, 0, 0));
  const startOfYesterday = new Date(startOfToday);
  startOfYesterday.setDate(startOfYesterday.getDate() - 1);
  const startOfLast7Days = new Date(startOfToday);
  startOfLast7Days.setDate(startOfLast7Days.getDate() - 6);
  const startOfLastMonth = new Date(startOfToday);
  startOfLastMonth.setDate(startOfLastMonth.getDate() - 30);

  const groups = {
    Today: [],
    Yesterday: [],
    'Last 7 Days': [],
    'Last Month': [],
  };

  for (const item of history) {
    const createdAt = new Date(item.createdAt);

    if (createdAt >= startOfToday) {
      groups.Today.push(item);
    } else if (createdAt >= startOfYesterday && createdAt < startOfToday) {
      groups.Yesterday.push(item);
    } else if (createdAt >= startOfLast7Days && createdAt < startOfYesterday) {
      groups['Last 7 Days'].push(item);
    } else if (createdAt >= startOfLastMonth && createdAt < startOfLast7Days) {
      groups['Last Month'].push(item);
    }
  }

  return groups;
};

const HistorySidebar = ({ historyOpen, setHistoryOpen, onHistoryListItemClick }) => {
  const [groceryHistory, setGroceryHistory] = useState([]);
  const [historyLoading, setHistoryLoading] = useState(false);

  const fetchGroceryHistory = async () => {
    setHistoryLoading(true);
    try {
      const history = await UserService.getGroceryListHistory();
      const sortedHistory = history.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt));
      setGroceryHistory(sortedHistory);
    } catch (err) {
      console.error("Failed to fetch grocery history:", err);
    } finally {
      setHistoryLoading(false);
    }
  };

  useEffect(() => {
    if (historyOpen && groceryHistory.length === 0) {
      fetchGroceryHistory();
    }
  }, [historyOpen]);

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  return (
    <>
      {/* Sidebar */}
      <div className={`fixed left-0 top-0 h-full bg-white shadow-2xl transition-transform duration-300 ease-in-out z-50 ${historyOpen ? 'translate-x-0' : '-translate-x-full'}`} style={{ width: '300px' }}>
        <div className="h-full flex flex-col">
          <div className="bg-gradient-to-r from-purple-600 to-indigo-700 p-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2">
                <History size={24} className="text-white" />
                <h2 className="text-xl font-bold text-white">History</h2>
              </div>
              <button
                onClick={() => setHistoryOpen(false)}
                className="text-white hover:bg-white/20 rounded-full p-1 transition"
              >
                <X size={20} />
              </button>
            </div>
            <p className="text-purple-100 text-sm mt-1">Your saved grocery lists</p>
          </div>

          <div className="flex-1 overflow-y-auto p-4">
            {historyLoading ? (
              <div className="flex justify-center py-8">
                <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-purple-600"></div>
              </div>
            ) : groceryHistory.length === 0 ? (
              <div className="text-center py-8">
                <History size={48} className="mx-auto text-gray-300 mb-3" />
                <p className="text-gray-500">No saved lists yet</p>
                <p className="text-gray-400 text-sm mt-1">Your grocery lists will appear here</p>
              </div>
            ) : (
              <>
                {Object.entries(groupHistoryByDate(groceryHistory)).map(([label, items]) =>
                  items.length > 0 ? (
                    <div key={label} className="mb-6">
                      <h3 className="text-lg font-semibold text-gray-700 mb-2">{label}</h3>
                      <div className="space-y-3">
                        {items.map((item) => (
                          <div
                            key={item.id}
                            className="border border-gray-200 rounded-lg p-3 hover:bg-gray-50 transition cursor-pointer"
                            onClick={() => onHistoryListItemClick(item)}
                          >
                            <h3 className="font-medium text-gray-800 mb-1">{item.name}</h3>
                            <div className="flex items-center gap-2 text-sm text-gray-500">
                              <Calendar size={14} />
                              <span>{formatDate(item.createdAt)}</span>
                            </div>
                          </div>
                        ))}
                      </div>
                    </div>
                  ) : null
                )}
              </>
            )}
          </div>
        </div>
      </div>

      {/* Toggle Button */}
      {!historyOpen && (
        <button
          onClick={() => setHistoryOpen(true)}
          className="fixed left-4 top-4 bg-white shadow-lg rounded-lg p-3 hover:bg-gray-50 transition z-40"
        >
          <Menu size={20} className="text-gray-600" />
        </button>
      )}

      {/* Mobile Overlay */}
      {historyOpen && (
        <div
          className="fixed inset-0 backdrop-blur-sm bg-white/20 z-40 md:hidden"
          onClick={() => setHistoryOpen(false)}
        />
      )}
    </>
  );
};

export default HistorySidebar;