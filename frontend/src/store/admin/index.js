import { combineReducers } from '@reduxjs/toolkit';
import coursesReducer from './courses';
import rewardsReducer from './rewards';
import chaptersReducer from './chapters';
import languagesReducer from './languages';

const adminReducer = combineReducers({
    courses: coursesReducer,
    rewards: rewardsReducer,
    chapters: chaptersReducer,
    languages: languagesReducer,
});

export default adminReducer;
