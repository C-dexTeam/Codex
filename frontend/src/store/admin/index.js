import { combineReducers } from '@reduxjs/toolkit';
import adminCoursesReducer from './courses';
import rewardsReducer from './rewards';
import chaptersReducer from './chapters';
import languagesReducer from './languages';
import attributesReducer from './attributes';
const adminReducer = combineReducers({
    adminCourses: adminCoursesReducer,
    rewards: rewardsReducer,
    chapters: chaptersReducer,
    languages: languagesReducer,
    attributes: attributesReducer,
});

export default adminReducer;
