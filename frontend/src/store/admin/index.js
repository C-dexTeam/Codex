import { combineReducers } from '@reduxjs/toolkit';
import coursesReducer from './courses';
import rewardsReducer from './rewards';
import chaptersReducer from './chapters';
import languagesReducer from './languages';
import planguagesReducer from './planguages';
import compilerReducer from './compiler';

const adminReducer = combineReducers({
    courses: coursesReducer,
    rewards: rewardsReducer,
    chapters: chaptersReducer,
    languages: languagesReducer,
    planguages : planguagesReducer,
    compiler : compilerReducer
});

export default adminReducer;
