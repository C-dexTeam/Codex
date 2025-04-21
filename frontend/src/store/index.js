// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import coursesSlice from './courses/coursesSlice'
import planguagesSlice from './planguages/planguagesSlice'
import profileSlice from './profile/profileSlice'
import chaptersSlice from './chapters/chaptersSlice'
import admin from './admin'

export const store = configureStore({
  reducer: {
    courses : coursesSlice,
    planguages : planguagesSlice,
    profile : profileSlice,
    chapters : chaptersSlice,
    admin,
  },
  
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
