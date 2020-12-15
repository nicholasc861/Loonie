import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom'

import GlobalStyle from './components/GlobalStyles'
import PrivateRoute from './PrivateRoute'
import NavigationBar from './components/NavigationBar'
import Login from './pages/Login'
import Register from './pages/Register'
import Dashboard from './pages/Dashboard'
import AuthCallback from './components/Questrade'
import useSessionStorage from './utils/SessionStorageHook'
import { AuthContext } from './context/auth';

const App = () => {
  const [isLoggedIn, setLoggedIn] = useSessionStorage("isAutheticated", false)

  const login = () => {
    setLoggedIn(true)
  }

  const logout = () => {
    setLoggedIn(false)
  }

  return(
    <AuthContext.Provider value={{isLoggedIn, login, logout}}>
      <Router>
        <GlobalStyle />
          <NavigationBar />
          <Switch>
            <Route path="/login">
              <Login />
            </Route>
            <Route path="/register">
              <Register />
            </Route>
            <Route path="/auth">
              <AuthCallback />
            </Route>
            <PrivateRoute path="/dashboard" component={Dashboard} />
          </Switch>
      </Router>
    </AuthContext.Provider>
  )
};

ReactDOM.render(
    <App />,
    document.getElementById('root')
);
