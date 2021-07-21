import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

export interface UserData {
  id: string;
  token: string;
}

export const login = createAsyncThunk("user/login", async () => {
  return {} as UserData;
});

const userSlice = createSlice({
  name: "user",
  initialState: {
    loading: false,
    user: {} as UserData,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(login.pending, (state) => {
      state.loading = true;
    });
    builder.addCase(login.fulfilled, (state, action) => {
      state.user = action.payload;
      state.loading = false;
    });
    builder.addCase(login.rejected, (state) => {
      state.loading = false;
    });
  },
});

export default userSlice.reducer;
