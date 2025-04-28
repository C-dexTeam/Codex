import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: false,
  errors: null,
  data: null,
};

export const fethcCompiler = createAsyncThunk(
  'admin/adminCompiler/fethcCompiler',
  async (params = {}, { rejectWithValue }) => {
    try {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_BASE_URL}/private/chapters/compilerNames`,
        { params }
      );
      return response.data?.data;
    } catch (error) {
      return rejectWithValue(error.response.data);
    }
  }
);



export const adminCompilerSlice = createSlice({
  name: 'admin/adminCompiler',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fethcCompiler.pending, (state) => {
        state.loading = true;
      })
      .addCase(fethcCompiler.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
      })
      .addCase(fethcCompiler.rejected, (state) => {
        state.loading = false;
      })
  },
});

export const getCompiler = (state) => state.admin.adminCompiler.data;
export const getCompilerLoading = (state) => state.admin.adminCompiler.loading;
export const getCompilerErrors = (state) => state.admin.adminCompiler.errors;

export default adminCompilerSlice.reducer;
