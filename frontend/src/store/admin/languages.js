import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from 'axios'

// Async thunk
export const fetchLanguages = createAsyncThunk(
    'admin/adminLanguages/fetchLanguages',
    async (_, { rejectWithValue }) => {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_BASE_URL}/private/language/`)
            return response.data.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error fetching languages')
        }
    }
)

// Initial state
const initialState = {
    languages: [],
    loading: false,
    error: null
}

// Slice
const adminLanguagesSlice = createSlice({
    name: 'admin/adminLanguages',
    initialState,
    reducers: {
        clearError: (state) => {
            state.error = null
        }
    },
    extraReducers: (builder) => {
        // Fetch Languages
        builder
            .addCase(fetchLanguages.pending, (state) => {
                state.loading = true
                state.error = null
            })
            .addCase(fetchLanguages.fulfilled, (state, action) => {
                state.loading = false
                state.languages = action.payload
            })
            .addCase(fetchLanguages.rejected, (state, action) => {
                state.loading = false
                state.error = action.payload
            })
    }
})

// Selectors
export const getLanguages = (state) => state.admin.adminLanguages.languages
export const getLanguagesLoading = (state) => state.admin.adminLanguages.loading
export const getLanguagesError = (state) => state.admin.adminLanguages.error

// Export actions
export const { clearError } = adminLanguagesSlice.actions

// Export reducer
export default adminLanguagesSlice.reducer 