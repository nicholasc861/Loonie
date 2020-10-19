import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom'

import GlobalStyle from './components/globalStyles'
import NavigationBar from './components/NavigationBar'
import Login from './components/Login'
import Register from './components/Register'


const App = () => {
  return(
      <Router>
        <GlobalStyle />
        <div>
          <NavigationBar />
          <Switch>
            <Route path="/login">
              <Login />
            </Route>
            <Route path="/register">
              <Register />
            </Route>
          </Switch>
        </div>
      </Router>
  )
};

ReactDOM.render(
    <App />,
    document.getElementById('root')
);
