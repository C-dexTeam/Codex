import { showToast } from "@/utils/showToast";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: false,
  errors: null,
  data: null,
  test: null,
  filters: {
    page: 1,
    limit: 10,
}
};

export const fetchTest = createAsyncThunk(
  "admin/adminTest/fetchTest",
  async (params = {}) => {
    try {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_BASE_URL}/private/tests`,
        { params }
      );
      return response.data?.data;
    } catch (error) {
      return error.response.data;
    }
  }
);

export const fetchTestOne = createAsyncThunk(
  "admin/adminTest/fetchTestOne",
  async (params = {}) => {
    try {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_BASE_URL}/private/tests/${params.id}`,
        { params }
      );
      return response.data?.data;
    } catch (error) {
      return error.response.data;
    }
  }
);



export const deleteTest = createAsyncThunk(
  "admin/adminTest/deleteTest",
  async (id, { rejectWithValue, dispatch }) => {
    try {
      const response = await axios.delete(
        `${process.env.NEXT_PUBLIC_BASE_URL}/admin/tests/${id}`
      );

      dispatch(fetchTest());
      return { id: id };
    } catch (error) {
      return rejectWithValue(error.response);
    }
  }
);

export const createTest = createAsyncThunk(
  "admin/adminTest/createTest",
  async ({ data, callback }, { dispatch, rejectWithValue }) => {
    console.log("data", data);
    try {
      const response = await axios.post(
        `${process.env.NEXT_PUBLIC_BASE_URL}/admin/tests`,
        data
      );

      dispatch(fetchTest());
      callback();

      return response.data?.data;
    } catch (error) {
      return rejectWithValue(error.response);
    }
  }
);

export const updateTest = createAsyncThunk(
  "admin/adminTest/updateTest",
  async (data, { rejectWithValue, dispatch }) => {
    try {
      const response = await axios.patch(
        `${process.env.NEXT_PUBLIC_BASE_URL}/admin/tests`,
        data
      );

      dispatch(fetchTest());

      return response.data?.data;
    } catch (error) {
      console.log("error", error);

      return rejectWithValue(error.response);
    }
  }
);

const adminTestSlice = createSlice({
  name: "admin/adminTest",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchTest.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchTest.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
      })
      .addCase(fetchTest.rejected, (state) => {
        state.loading = false;
      })
      .addCase(createTest.pending, (state) => {
        state.loading = true;
      }
      )
      .addCase(createTest.fulfilled, (state, action) => {
        state.loading = false;
        state.test = action.payload;
        showToast("dismiss")
        showToast("success", "Test created successfully");
      })
      .addCase(createTest.rejected, (state, action) => {
        state.loading = false;
        showToast("dismiss")
        showToast("error", action.payload?.data?.message);
      })
      .addCase(updateTest.pending, (state) => {
        state.loading = true;
      })
      .addCase(updateTest.fulfilled, (state, action) => {
        state.loading = false;
        state.test = action.payload;
        showToast("dismiss")
        showToast("success", "Test updated successfully");
      })
      .addCase(updateTest.rejected, (state, action) => {
        state.loading = false;
        showToast("dismiss")
        showToast("error", action.payload?.data?.message);
      })
      .addCase(deleteTest.pending, (state) => {
        state.loading = true;
      })
      .addCase(deleteTest.fulfilled, (state, action) => {
        state.loading = false;
        state.data = state.data.filter((item) => item.id !== action.payload.id);
        showToast("dismiss")
        showToast("success", "Test deleted successfully");
      })
      .addCase(deleteTest.rejected, (state, action) => {
        state.loading = false;
        showToast("dismiss")
        showToast("error", action.payload?.data?.message);
      });
    builder
      .addCase(fetchTestOne.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchTestOne.fulfilled, (state, action) => {
        state.loading = false;
        state.test = action.payload;
      })
      .addCase(fetchTestOne.rejected, (state) => {
        state.loading = false;
      });
  },
});

export const getLoading = (state) => state.admin.adminTest.loading;
export const getTest = (state) => state.admin.adminTest.data;
export const getFilters = (state) => state.admin.adminTest.filters;
export const getCurrentTest = (state) => state.admin.adminTest.test;
export const getTotalCount = (state) => state.admin.adminTest.totalCount;
export const getTestById = (state) => state.admin.adminTest.test;

export default adminTestSlice.reducer;
