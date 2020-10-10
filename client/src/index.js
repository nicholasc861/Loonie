import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom'
import { createStore, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import thunk from 'redux-thunk'

import rootReducer from './reducers'
import NavigationBar from './components/NavigationBar'

const store = createStore(rootReducer, applyMiddlware(thunk))

const App = () => {
  return(
      <Router>
        <div>
          <NavigationBar />
        </div>
      </Router>
  )
};

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>, 
  document.getElementById('root')
);
