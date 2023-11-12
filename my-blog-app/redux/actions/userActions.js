import axios from 'axios';

// Action Types
export const LOGIN_REQUEST = 'LOGIN_REQUEST';
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS';
export const LOGIN_FAILURE = 'LOGIN_FAILURE';
export const REGISTER_REQUEST = 'REGISTER_REQUEST';
export const REGISTER_SUCCESS = 'REGISTER_SUCCESS';
export const REGISTER_FAILURE = 'REGISTER_FAILURE';
export const LOGOUT = 'LOGOUT';

// Action Creators
export const loginRequest = () => ({
  type: LOGIN_REQUEST,
});

export const loginSuccess = (user) => ({
  type: LOGIN_SUCCESS,
  payload: user,
});

export const loginFailure = (error) => ({
  type: LOGIN_FAILURE,
  payload: error,
});

export const registerRequest = () => ({
  type: REGISTER_REQUEST,
});

export const registerSuccess = () => ({
  type: REGISTER_SUCCESS,
});

export const registerFailure = (error) => ({
  type: REGISTER_FAILURE,
  payload: error,
});

export const logout = () => ({
  type: LOGOUT,
});

// Thunk Action Creator for user login
export const loginUser = (credentials) => {
  return (dispatch) => {
    dispatch(loginRequest());

    // Make an API request to authenticate the user
    axios.post('http://localhost:8080/login', credentials)
      .then((response) => {
        dispatch(loginSuccess(response.data));
      })
      .catch((error) => {
        dispatch(loginFailure(error.message));
      });
  };
};

// Thunk Action Creator for user registration
export const registerUser = (userData) => {
  return (dispatch) => {
    dispatch(registerRequest());

    // Make an API request to register a new user
    axios.post('http://localhost:8080/register', userData)
      .then(() => {
        dispatch(registerSuccess());
      })
      .catch((error) => {
        dispatch(registerFailure(error.message));
      });
  };
};
