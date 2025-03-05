// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import admin from './admin'

export const store = configureStore({
  reducer: {
    admin,
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
