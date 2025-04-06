import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import axios from 'axios';

const initialState = {
    loading: false,
    data: null,
};

/**
 * Fetch rewards with optional query parameters.
 * 
 * @param {Object} params - Query parameters for fetching rewards.
 * @param {string} params.rewardID - Reward id.
 * @param {string} params.name - Reward name.
 * @param {string} params.symbol - Reward symbol.
 * @param {string} params.rewardType - Reward Title.
 * @param {string} params.page - Page number.
 * @param {string} params.limit - Number of items per page.
 */
export const fetchRewards = createAsyncThunk('rewards/fetchRewards', async (params = {}) => {
    try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_BASE_URL}/private/rewards`, { params });
        return response.data?.data;
    } catch (error) {
        return error.response.data;
    }
});

/**
 * Create a new reward.
 * 
 * @param {FormData} formData - Form data containing reward details.
 * @param {File} formData.imageFile - Reward Image File (required).
 * @param {string} formData.name - Reward name (required).
 * @param {string} formData.description - Reward description (optional).
 * @param {string} formData.rewardType - Reward Type (required).
 * @param {string} formData.sellerFee - Seller Fee (optional).
 * @param {string} formData.symbol - Reward Symbol (required).
 */
export const createReward = createAsyncThunk('rewards/createReward', async (formData, { dispatch, rejectWithValue }) => {
    try {
        const response = await axios.post(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/rewards`, formData);
        dispatch(fetchRewards());
        return response.data?.data;
    } catch (error) {
        return rejectWithValue(error.response);
    }
});

/**
 * Update an existing reward.
 * 
 * @param {FormData} formData - Form data containing reward details.
 * @param {File} formData.id - Reward id (required).
 * @param {File} formData.imageFile - Reward Image File (optional).
 * @param {string} formData.name - Reward name (optional).
 * @param {string} formData.description - Reward description (optional).
 * @param {string} formData.rewardType - Reward Type (optional).
 * @param {string} formData.sellerFee - Seller Fee (optional).
 * @param {string} formData.symbol - Reward Symbol (optional).
 */
export const updateReward = createAsyncThunk('rewards/updateReward', async (formData, { rejectWithValue }) => {
    try {
        const response = await axios.patch(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/rewards`, formData);
        return response.data?.data;
    } catch (error) {
        return rejectWithValue(error.response);
    }
});

/**
 * Delete a reward by ID.
 * 
 * @param {string} id - Reward ID (required).
 */
export const deleteReward = createAsyncThunk('rewards/deleteReward', async (id, { rejectWithValue }) => {
    try {
        const response = await axios.delete(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/rewards/${id}`);
        return response.data?.data;
    } catch (error) {
        return rejectWithValue(error.response);
    }
});

const rewardsSlice = createSlice({
    name: 'rewards',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchRewards.pending, (state) => {
                state.loading = true;
            })
            .addCase(fetchRewards.fulfilled, (state, action) => {
                state.loading = false;
                state.data = action.payload;
            })
            .addCase(fetchRewards.rejected, (state) => {
                state.loading = false;
            })
            .addCase(createReward.pending, (state) => {
                state.loading = true;
            })
            .addCase(createReward.fulfilled, (state, action) => {
                state.loading = false;
                state.data = [...state.data || [], action.payload];
            })
            .addCase(createReward.rejected, (state) => {
                state.loading = false;
            })
            .addCase(updateReward.pending, (state) => {
                state.loading = true;
            })
            .addCase(updateReward.fulfilled, (state, action) => {
                state.loading = false;
                const index = state.data.findIndex(course => course.id === action.payload.id);
                if (index !== -1) {
                    state.data[index] = action.payload;
                }
            })
            .addCase(updateReward.rejected, (state) => {
                state.loading = false;
            })
            .addCase(deleteReward.pending, (state) => {
                state.loading = true;
            })
            .addCase(deleteReward.fulfilled, (state, action) => {
                state.loading = false;
                state.data = state.data.filter(course => course.id !== action.payload.id);
            })
            .addCase(deleteReward.rejected, (state) => {
                state.loading = false;
            });
    },
});

export const getLoading = (state) => state.admin.rewards.loading;
export const getRewards = (state) => state.admin.rewards.data;

export default rewardsSlice.reducer;
