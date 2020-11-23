import React, { Component, useContext } from 'react'
import { Route, Redirect } from 'react-router-dom'
import { AuthContext, useAuth } from './context/auth'


const PrivateRoute = ({component: Component, ...rest}) => {
    const context = useContext(AuthContext)

    return (
        <Route {...rest} render={props => 
            context.isLoggedIn ?
            (<Component {...props} />)
            : (<Redirect to="/login" />)
        } />
    )
}

export default PrivateRoute