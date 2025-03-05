// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import coursesSlice from './courses/coursesSlice'
import planguagesSlice from './planguages/planguagesSlice'


export const store = configureStore({
  reducer: {
    courses : coursesSlice,
    planguages : planguagesSlice,
  },
  
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
