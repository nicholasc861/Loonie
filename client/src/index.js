import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom'

import NavigationBar from './components/NavigationBar'
import Login from './components/Login'


const App = () => {
  return(
      <Router>
        <div>
          <NavigationBar />
          <Switch>
            <Route path="/login">
              <Login />
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
