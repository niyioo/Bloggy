import axios from 'axios';

// Action Types
export const FETCH_BLOG_POSTS_REQUEST = 'FETCH_BLOG_POSTS_REQUEST';
export const FETCH_BLOG_POSTS_SUCCESS = 'FETCH_BLOG_POSTS_SUCCESS';
export const FETCH_BLOG_POSTS_FAILURE = 'FETCH_BLOG_POSTS_FAILURE';
export const CREATE_BLOG_POST_REQUEST = 'CREATE_BLOG_POST_REQUEST';
export const CREATE_BLOG_POST_SUCCESS = 'CREATE_BLOG_POST_SUCCESS';
export const CREATE_BLOG_POST_FAILURE = 'CREATE_BLOG_POST_FAILURE';

// Action Creators
export const fetchBlogPostsRequest = () => ({
  type: FETCH_BLOG_POSTS_REQUEST,
});

export const fetchBlogPostsSuccess = (posts) => ({
  type: FETCH_BLOG_POSTS_SUCCESS,
  payload: posts,
});

export const fetchBlogPostsFailure = (error) => ({
  type: FETCH_BLOG_POSTS_FAILURE,
  payload: error,
});

export const createBlogPostRequest = () => ({
  type: CREATE_BLOG_POST_REQUEST,
});

export const createBlogPostSuccess = (newPost) => ({
  type: CREATE_BLOG_POST_SUCCESS,
  payload: newPost,
});

export const createBlogPostFailure = (error) => ({
  type: CREATE_BLOG_POST_FAILURE,
  payload: error,
});

// Thunk Action Creator for fetching blog posts
export const fetchBlogPosts = () => {
  return (dispatch) => {
    dispatch(fetchBlogPostsRequest());

    // Make an API request to get blog posts
    axios.get('/api/posts')
      .then((response) => {
        dispatch(fetchBlogPostsSuccess(response.data));
      })
      .catch((error) => {
        dispatch(fetchBlogPostsFailure(error.message));
      });
  };
};

// Thunk Action Creator for creating a new blog post
export const createBlogPost = (post) => {
  return (dispatch) => {
    dispatch(createBlogPostRequest());

    // Make an API request to create a new blog post
    axios.post('/api/posts', post)
      .then((response) => {
        dispatch(createBlogPostSuccess(response.data));
      })
      .catch((error) => {
        dispatch(createBlogPostFailure(error.message));
      });
  };
};
