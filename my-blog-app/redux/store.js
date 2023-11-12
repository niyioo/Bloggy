import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';

<<<<<<< HEAD
import rootReducer from './reducers';

=======
// Import your reducers and create a rootReducer if needed
import rootReducer from './reducers';

// Create the Redux store with the rootReducer and middleware
>>>>>>> 53834e38c4002610b63a29213c90384dab0149f8
const store = createStore(rootReducer, applyMiddleware(thunk));

function MyApp({ Component, pageProps }) {
  return (
    <Provider store={store}>
      <Component {...pageProps} />
    </Provider>
  );
}

export default MyApp;
