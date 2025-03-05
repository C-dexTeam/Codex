// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import coursesSlice from './courses/coursesSlice'
import planguagesSlice from './planguages/planguagesSlice'
import profileSlice from './profile/profileSlice'


export const store = configureStore({
  reducer: {
    courses : coursesSlice,
    planguages : planguagesSlice,
    profile : profileSlice,
  },
  
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
