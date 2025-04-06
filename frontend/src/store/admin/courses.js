import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import axios from 'axios';

const initialState = {
    loading: false,
    data: null,
};

/**
 * Fetch courses with optional query parameters.
 * 
 * @param {Object} params - Query parameters for fetching courses.
 * @param {string} params.id - Course ID.
 * @param {string} params.languageID - Language ID.
 * @param {string} params.pLanguageID - Programming Language ID.
 * @param {string} params.title - Course Title.
 * @param {string} params.page - Page number.
 * @param {string} params.limit - Number of items per page.
 */
export const fetchCourses = createAsyncThunk('courses/fetchCourses', async (params = {}) => {
    try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_BASE_URL}/private/courses`, { params });
        return response.data?.data;
    } catch (error) {
        return error.response.data;
    }
});

/**
 * Create a new course.
 * 
 * @param {FormData} formData - Form data containing course details.
 * @param {File} formData.imageFile - Course Image File (required).
 * @param {string} [formData.description] - Course description (optional).
 * @param {string} [formData.languageID] - Language ID (optional).
 * @param {string} formData.programmingLanguageID - Programming Language ID (required).
 * @param {number} [formData.rewardAmount] - Reward Amount (optional).
 * @param {string} [formData.rewardID] - Reward ID (optional).
 * @param {string} formData.title - Course Title (required).
 */
export const createCourse = createAsyncThunk('courses/createCourse', async ({ formData, callback }, { dispatch, rejectWithValue }) => {
    try {
        const response = await axios.post(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/courses`, formData);

        dispatch(fetchCourses());
        callback()

        return response.data?.data;
    } catch (error) {
        return rejectWithValue(error.response);
    }
});

/**
 * Update an existing course.
 * 
 * @param {FormData} formData - Form data containing updated course details.
 * @param {File} formData.imageFile - Course Image File (required).
 * @param {string} [formData.description] - Course description (optional).
 * @param {string} formData.id - Course ID (required).
 * @param {string} [formData.languageID] - Language ID (optional).
 * @param {string} formData.programmingLanguageID - Programming Language ID (required).
 * @param {number} [formData.rewardAmount] - Reward Amount (optional).
 * @param {string} [formData.rewardID] - Reward ID (optional).
 * @param {string} formData.title - Course Title (required).
 */
export const updateCourse = createAsyncThunk('courses/updateCourse', async (formData, { rejectWithValue }) => {
    try {
        const response = await axios.patch(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/courses`, formData);
        return response.data?.data;
    } catch (error) {
        return rejectWithValue(error.response);
    }
});

/**
 * Delete a course by ID.
 * 
 * @param {string} id - Course ID (required).
 */
export const deleteCourse = createAsyncThunk('courses/deleteCourse', async (id, { rejectWithValue }) => {
    try {
        const response = await axios.delete(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/courses/${id}`);
        return response.data?.data;
    } catch (error) {
        return rejectWithValue(error.response);
    }
});

const coursesSlice = createSlice({
    name: 'courses',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchCourses.pending, (state) => {
                state.loading = true;
            })
            .addCase(fetchCourses.fulfilled, (state, action) => {
                state.loading = false;
                state.data = action.payload;
            })
            .addCase(fetchCourses.rejected, (state) => {
                state.loading = false;
            })
            .addCase(createCourse.pending, (state) => {
                state.loading = true;
            })
            .addCase(createCourse.fulfilled, (state, action) => {
                state.loading = false;
                state.data = [...state.data || [], action.payload];
            })
            .addCase(createCourse.rejected, (state) => {
                state.loading = false;
            })
            .addCase(updateCourse.pending, (state) => {
                state.loading = true;
            })
            .addCase(updateCourse.fulfilled, (state, action) => {
                state.loading = false;
                const index = state.data.findIndex(course => course.id === action.payload.id);
                if (index !== -1) {
                    state.data[index] = action.payload;
                }
            })
            .addCase(updateCourse.rejected, (state) => {
                state.loading = false;
            })
            .addCase(deleteCourse.pending, (state) => {
                state.loading = true;
            })
            .addCase(deleteCourse.fulfilled, (state, action) => {
                state.loading = false;
                state.data = state.data.filter(course => course.id !== action.payload.id);
            })
            .addCase(deleteCourse.rejected, (state) => {
                state.loading = false;
            });
    },
});

export const getLoading = (state) => state.admin.courses.loading;
export const getCourses = (state) => state.admin.courses.data;

export default coursesSlice.reducer;
