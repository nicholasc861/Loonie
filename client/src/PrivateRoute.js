import React, { Component } from 'react'
import { Route, Redirect } from 'react-router-dom'
import { useAuth } from './context/auth'


const PrivateRoute = ({component: Component, ...rest}) => {

    return (
        <Route {...rest} render={props => 
            true ?
            (<Component {...props} />)
            : (<Redirect to="/login" />)
        } />
    )
}

export default PrivateRoute