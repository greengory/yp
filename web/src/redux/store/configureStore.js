import { createStore, applyMiddleware } from 'redux';
import reducer from '../reducers/rootReducers';
import thunkMiddleware from 'redux-thunk';
import { createLogger} from 'redux-logger';


const store = createStore(
    reducer, /* preloadedState, */
    window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__(),
    applyMiddleware(
        createLogger(),
        thunkMiddleware
    ),
);

export default store;