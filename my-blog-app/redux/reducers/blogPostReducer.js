import { FETCH_POSTS, CREATE_POST } from '../actions/blogPostActions';

const initialState = {
  posts: [],
};

const blogPostReducer = (state = initialState, action) => {
  switch (action.type) {
    case FETCH_POSTS:
      return {
        ...state,
        posts: action.payload,
      };
    case CREATE_POST:
      return {
        ...state,
        posts: [action.payload, ...state.posts],
      };
    default:
      return state;
  }
};

export default blogPostReducer;
