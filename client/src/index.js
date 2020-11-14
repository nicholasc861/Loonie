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
import Dashboard from './components/Dashboard'

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
            <Route path="/dashboard">
              <Dashboard />
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
