import { showToast } from "@/utils/showToast";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";


const initialState = {
  loading: false,
  error: false,
  data: [],
};



export const getTest = createAsyncThunk(
  "test/getTest",
  async (data, { rejectWithValue,dispatch }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `${process.env.NEXT_PUBLIC_BASE_URL}/private/tests`,
        headers: {
          "Content-Type": "application/json",
        },
        params: {
            chapterID: data.id,
        },
        
      });
      if (response.status === 200) {
        return response.data;
      }
    } catch (error) {
      return rejectWithValue(response.message || error.message);
    }
  }
);


const testSlice = createSlice({
  name: "test",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
        .addCase(getTest.pending, (state) => {
            state.loading = true;
        })
        .addCase(getTest.fulfilled, (state, action) => {
            state.loading = false;
            state.data = action.payload;
        })
        .addCase(getTest.rejected, (state) => {
            state.loading = false;
            state.error = true;
        })
  },
});

export default testSlice.reducer;
