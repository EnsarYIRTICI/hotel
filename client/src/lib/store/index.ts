import { configureStore, createSlice } from "@reduxjs/toolkit";

// Create a slice of state for authentication
const authSlice = createSlice({
  name: "auth",
  initialState: {
    isLoggedIn: !!localStorage.getItem("token"), // check if token exists in localStorage
  },
  reducers: {
    login: (state) => {
      state.isLoggedIn = true;
      localStorage.setItem("token", "your-token"); // Save token in localStorage
    },
    logout: (state) => {
      state.isLoggedIn = false;
      localStorage.removeItem("token"); // Remove token from localStorage
    },
  },
});

export const { login, logout } = authSlice.actions;

// Create Redux store with the authSlice
const store = configureStore({
  reducer: {
    auth: authSlice.reducer,
  },
});

export default store;
