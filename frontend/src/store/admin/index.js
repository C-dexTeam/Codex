import { combineReducers } from '@reduxjs/toolkit';
import coursesReducer from './courses'; // Assuming courses.js is in the same directory
import rewardsReducer from './rewards'; // Assuming courses.js is in the same directory

const adminReducer = combineReducers({
    courses: coursesReducer,
    rewards: rewardsReducer,
    // Add other slices here as needed
});

export default adminReducer;
