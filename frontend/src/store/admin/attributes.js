import { showToast } from '@/utils/showToast'
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from 'axios'

/**
 * Async thunks for attribute operations
 * 
 * @param {Object} params - Query parameters for fetching attributes
 * @param {string} params.page - Page number
 * @param {string} params.limit - Number of items per page
 */
export const fetchAttributes = createAsyncThunk(
    'admin/adminAttributes/fetchAttributes',
    async (_, { rejectWithValue, getState }) => {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_BASE_URL}/private/attributes`, { params: { ...getState().admin.adminAttributes.filters } })
            return response.data?.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error fetching attributes')
        }
    }
)

/**
 * Async thunk for fetching an attribute by ID
 * 
 * @param {string} id - ID of the attribute to fetch
 * @returns {Promise} - Promise with the response data
 */
export const fetchAttribute = createAsyncThunk(
    'admin/adminAttributes/fetchAttribute',
    async (id, { rejectWithValue }) => {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/attributes/${id}`)
            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error fetching attribute')
        }
    }
)


/**
 * Async thunk for creating an attribute
 * 
 * @param {Object} data - Data for creating the attribute
 * @returns {Promise} - Promise with the response data
 */
export const createAttribute = createAsyncThunk(
    'admin/adminAttributes/createAttribute',
    async ({ data, callback }, { rejectWithValue, dispatch }) => {
        try {
            const response = await axios.post(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/attributes`, data)

            dispatch(fetchAttributes({ params: { page: 1, limit: 10 } }))
            callback && callback()
            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error creating attribute')
        }
    }
)

/**
 * Async thunk for updating an attribute
 * 
 * @param {Object} data - Data for updating the attribute
 * @returns {Promise} - Promise with the response data
 */
export const updateAttribute = createAsyncThunk(
    'admin/adminAttributes/updateAttribute',
    async ({ id, data }, { rejectWithValue }) => {
        try {
            const response = await axios.patch(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/attributes/${id}`, data)
            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error updating attribute')
        }
    }
)

/**
 * Async thunk for deleting an attribute
 * 
 * @param {string} id - ID of the attribute to delete
 * @returns {Promise} - Promise with the response data
 */
export const deleteAttribute = createAsyncThunk(
    'admin/adminAttributes/deleteAttribute',
    async (id, { rejectWithValue, dispatch }) => {
        try {
            const response = await axios.delete(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/attributes/${id}`)
            dispatch(fetchAttributes({ page: 1, limit: 10 }))
            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error deleting attribute')
        }
    }
)

const initialState = {
    data: [],
    totalCount: 0,
    attribute: null,
    loading: false,
    filters: {
        page: 1,
        limit: 10,
        name: ""
    }
};

export const getAttributes = (state) => state.admin.adminAttributes.data
export const getLoading = (state) => state.admin.adminAttributes.loading
export const getFilters = (state) => state.admin.adminAttributes.filters
export const getCurrentAttribute = (state) => state.admin.adminAttributes.attribute
export const getTotalCount = (state) => state.admin.adminAttributes.totalCount

export const adminAttributesSlice = createSlice({
    name: 'admin/adminAttributes',
    initialState,
    reducers: {
        setCurrentAttribute: (state, action) => {
            state.attribute = action.payload
        },
        setFilters: (state, action) => {
            state.filters = action.payload
        }
    },
    extraReducers: (builder) => {
        builder
            // Get Attributes
            .addCase(fetchAttributes.pending, (state) => {
                state.loading = true
            })
            .addCase(fetchAttributes.fulfilled, (state, action) => {
                state.loading = false
                state.data = action.payload?.attributes || []
                state.totalCount = action.payload?.totalCount || 0
            })
            .addCase(fetchAttributes.rejected, (state, action) => {
                state.loading = false
            })
            // Get Attribute
            .addCase(fetchAttribute.pending, (state) => {
                state.loading = true
            })
            .addCase(fetchAttribute.fulfilled, (state, action) => {
                state.loading = false
                state.attribute = action.payload
            })
            .addCase(fetchAttribute.rejected, (state, action) => {
                state.loading = false
            })
            // Create Attribute
            .addCase(createAttribute.pending, (state) => {
                state.loading = true
                showToast("dismiss")
                showToast("loading", "Creating attribute...")
            })
            .addCase(createAttribute.fulfilled, (state) => {
                state.loading = false
                showToast("dismiss")
                showToast("success", "Attribute created successfully")
            })
            .addCase(createAttribute.rejected, (state, action) => {
                state.loading = false
                showToast("dismiss")
                showToast("error", "Error creating attribute")
            })
            // Update Attribute
            .addCase(updateAttribute.pending, (state) => {
                state.loading = true
                showToast("dismiss")
                showToast("loading", "Updating attribute...")
            })
            .addCase(updateAttribute.fulfilled, (state) => {
                state.loading = false
            })
            .addCase(updateAttribute.rejected, (state, action) => {
                state.loading = false
            })
            // Delete Attribute
            .addCase(deleteAttribute.pending, (state) => {
                state.loading = true
                showToast("dismiss")
                showToast("loading", "Deleting attribute...")
            })
            .addCase(deleteAttribute.fulfilled, (state) => {
                state.loading = false
                showToast("dismiss")
                showToast("success", "Attribute deleted successfully")
            })
            .addCase(deleteAttribute.rejected, (state, action) => {
                state.loading = false
                showToast("dismiss")
                showToast("error", "Error deleting attribute")
            })
    }
})

export const { setCurrentAttribute, setFilters } = adminAttributesSlice.actions

export default adminAttributesSlice.reducer 