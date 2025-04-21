import { showToast } from '@/utils/showToast'
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import axios from 'axios'

/**
 * Async thunks for chapter operations
 * 
 * @param {Object} params - Query parameters for fetching chapters
 * @param {string} params.id - Chapter ID
 * @param {string} params.languageID - Language ID
 * @param {string} params.courseID - Course ID
 * @param {string} params.rewardID - Reward ID
 * @param {string} params.title - Chapter Title
 * @param {string} params.grantsExp - Whether chapter grants experience
 * @param {string} params.page - Page number
 * @param {string} params.limit - Number of items per page
 */
export const fetchChapters = createAsyncThunk(
    'admin/adminChapters/fetchChapters',
    async (_, { rejectWithValue, getState }) => {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_BASE_URL}/private/chapters`, { params: { ...getState().admin.adminChapters.filters } })

            return response.data.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error fetching chapters')
        }
    }
)

/**
 * Delete a chapter by ID
 * 
 * @param {string} chapterId - The ID of the chapter to delete
 * @param {string} params.page - Page number
 * @param {string} params.limit - Number of items per page
 * @returns {Promise} A promise that resolves with the deleted chapter ID or rejects with an error
 */

export const fetchChapter = createAsyncThunk(
    'admin/adminChapters/fetchChapter',
    async ({ id, page = 1, limit = 10 }, { rejectWithValue }) => {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_BASE_URL}/private/chapters/${id}?page=${page}&limit=${limit}`)
            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error fetching chapter')
        }
    }
)

/**
 * Chapter data structure
 * 
 * @typedef {Object} ChapterData
 * @property {boolean} active - Whether the chapter is active
 * @property {string} content - Chapter content
 * @property {string} courseID - Course ID
 * @property {string} description - Chapter description
 * @property {string} dockerTemplate - Docker template
 * @property {string} frontendTemplate - Frontend template
 * @property {string} funcName - Function name
 * @property {boolean} grantsExperience - Whether chapter grants experience
 * @property {string} languageID - Language ID
 * @property {number} order - Chapter order
 * @property {number} rewardAmount - Reward amount
 * @property {string} rewardID - Reward ID
 * @property {string} title - Chapter title
 */

export const createChapter = createAsyncThunk(
    'admin/adminChapters/createChapter',
    async ({ data, callback }, { rejectWithValue, dispatch }) => {
        try {
            const response = await axios.post(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/chapters`, {
                ...data,
                rewardAmount: parseInt(data?.rewardAmount) ? parseInt(data?.rewardAmount) : 0,
                order: parseInt(data?.order) ? parseInt(data?.order) : 0,
                checkTemplate: data?.checkTemplate || '',
            })

            dispatch(fetchChapters({ params: { page: 1, limit: 10 } }))
            callback?.()

            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error creating chapter')
        }
    }
)
/**
 * Chapter data structure for update operation
 * 
 * @typedef {Object} UpdateChapterData
 * @property {boolean} active - Whether the chapter is active
 * @property {string} content - Chapter content
 * @property {string} courseID - Course ID
 * @property {string} description - Chapter description
 * @property {string} dockerTemplate - Docker template
 * @property {string} frontendTemplate - Frontend template
 * @property {string} funcName - Function name
 * @property {boolean} grantsExperience - Whether chapter grants experience
 * @property {string} languageID - Language ID
 * @property {number} rewardAmount - Reward amount
 * @property {string} rewardID - Reward ID
 * @property {string} title - Chapter title
 */

export const updateChapter = createAsyncThunk(
    'admin/adminChapters/updateChapter',
    async ({ data, callback }, { rejectWithValue, dispatch, getState }) => {
        try {
            const response = await axios.patch(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/chapters`, data || getState().admin.adminChapters.currentChapter)

            dispatch(fetchChapters({ params: { page: 1, limit: 10 } }))
            callback?.()

            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error updating chapter')
        }
    }
)

/**
 * Chapter data structure for delete operation
 * 
 * @property {string} id - Chapter ID
 */

export const deleteChapter = createAsyncThunk(
    'admin/adminChapters/deleteChapter',
    async (id, { rejectWithValue, dispatch }) => {
        try {
            await axios.delete(`${process.env.NEXT_PUBLIC_BASE_URL}/admin/chapters/${id}`)

            dispatch(fetchChapters({ params: { page: 1, limit: 10 } }))
            return id
        } catch (error) {
            return rejectWithValue(error.response?.data || 'Error deleting chapter')
        }
    }
)


// Initial state
const initialState = {
    chapters: [],
    currentChapter: null,
    filters: {
        page: 1,
        limit: 10,
        title: "",
        courseID: "",
        languageID: "",
    },
    loading: false,
}

// Slice
const adminChaptersSlice = createSlice({
    name: 'admin/adminChapters',
    initialState,
    reducers: {
        clearError: (state) => {
            state.error = null
        },
        setCurrentChapter: (state, action) => {
            state.currentChapter = action.payload
        },
        setFilters: (state, action) => {
            state.filters = action.payload
        }
    },
    extraReducers: (builder) => {
        // Fetch Chapters
        builder
            .addCase(fetchChapters.pending, (state) => {
                state.loading = true
            })
            .addCase(fetchChapters.fulfilled, (state, action) => {
                state.loading = false
                state.chapters = action.payload;
            })
            .addCase(fetchChapters.rejected, (state, action) => {
                state.loading = false
            })

        // Fetch Single Chapter
        builder
            .addCase(fetchChapter.pending, (state) => {
                state.loading = true
            })
            .addCase(fetchChapter.fulfilled, (state, action) => {
                state.loading = false
                state.currentChapter = action.payload.data
            })
            .addCase(fetchChapter.rejected, (state, action) => {
                state.loading = false
            })

        // Create Chapter
        builder
            .addCase(createChapter.pending, (state) => {
                state.loading = true
            })
            .addCase(createChapter.fulfilled, (state, action) => {
                state.loading = false
                showToast("dismiss")
                showToast("success", "Chapter created successfully")
            })
            .addCase(createChapter.rejected, (state, action) => {
                state.loading = false
                showToast("dismiss")
                showToast("error", "Error creating chapter")
            })

        // Update Chapter
        builder
            .addCase(updateChapter.pending, (state) => {
                state.loading = true
            })
            .addCase(updateChapter.fulfilled, (state, action) => {
                state.loading = false
                const index = state.chapters.findIndex(chapter => chapter.id === action.payload.id)
                if (index !== -1) {
                    state.chapters[index] = action.payload
                }
                showToast("dismiss")
                showToast("success", "Chapter updated successfully")
            })
            .addCase(updateChapter.rejected, (state, action) => {
                state.loading = false
                showToast("dismiss")
                showToast("error", "Error updating chapter")
            })

        // Delete Chapter
        builder
            .addCase(deleteChapter.pending, (state) => {
                state.loading = true
            })
            .addCase(deleteChapter.fulfilled, (state, action) => {
                state.loading = false
                showToast("dismiss")
                showToast("success", "Chapter deleted successfully")
            })
            .addCase(deleteChapter.rejected, (state, action) => {
                state.loading = false
                showToast("dismiss")
                showToast("error", "Error deleting chapter")
            })
    }
})

// Export actions
export const { clearError, setCurrentChapter } = adminChaptersSlice.actions

export const getChapters = (state) => state.admin.adminChapters.chapters
export const getCurrentChapter = (state) => state.admin.adminChapters.currentChapter
export const getLoading = (state) => state.admin.adminChapters.loading
export const getFilters = (state) => state.admin.adminChapters.filters
// Export reducer
export default adminChaptersSlice.reducer 