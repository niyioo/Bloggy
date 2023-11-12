import { combineReducers } from 'redux';
import userReducer from './userReducer';
import blogPostReducer from './blogPostReducer';

const rootReducer = combineReducers({
  user: userReducer,
  blogPosts: blogPostReducer,
  // Add more reducers here if needed
});

export default rootReducer;
