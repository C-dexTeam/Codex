import { combineReducers } from '@reduxjs/toolkit';
import adminCoursesReducer from './courses';
import adminRewardsReducer from './rewards';
import adminChaptersReducer from './chapters';
import adminLanguagesReducer from './languages';
import adminAttributesReducer from './attributes';
import adminPlanguagesReducer from './planguages';
import adminCompilerReducer from './compiler';

const adminReducer = combineReducers({
    adminCourses: adminCoursesReducer,
    adminRewards: adminRewardsReducer,
    adminChapters: adminChaptersReducer,
    adminLanguages: adminLanguagesReducer,
    adminAttributes: adminAttributesReducer,
    adminPlanguages: adminPlanguagesReducer,
    adminCompiler: adminCompilerReducer
});

export default adminReducer;
