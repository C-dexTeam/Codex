import { combineReducers } from '@reduxjs/toolkit';
import adminCoursesReducer from './courses';
import rewardsReducer from './rewards';
import chaptersReducer from './chapters';
import languagesReducer from './languages';

const adminReducer = combineReducers({
    adminCourses: adminCoursesReducer,
    rewards: rewardsReducer,
    chapters: chaptersReducer,
    languages: languagesReducer,
});

export default adminReducer;
