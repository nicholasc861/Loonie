import React, { useState, useContext } from 'react'
import styled from 'styled-components'
import { Container, Form, Button } from 'react-bootstrap'
import { useHistory } from 'react-router-dom'

import axios from "axios"
import { AuthContext } from '../context/auth'


const LoginHeader = styled.div`
    font-size: 25px;
    font-weight: 600;
`

const LoginWrapper = styled.div`
    font-family: "Karla", sans-serif;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 500px;
`

const FormGroup = styled.div`
    margin: 20px 0px;
`

const LoginInput = styled.input`
    border-width: 0px;
    border-radius: 4px;
    line-height: 1.5em;
    background-color: #E7E5DF;
    padding: 10px;
    height: 40px;
    width: 333px;

    :focus {
        outline-color: #4AAD52;
    }
`   

const FormLabel = styled.label`
    font-size: 15px;
    margin: 0px;
`

const LoginForm = styled(Form)`
    display: block;
    margin: auto;

`

const LoginButton = styled(Button)`
    background-color: #4AAD52;
    border: 0px;
    width: 90px;

    :hover {
        background-color: #488B49;
    }

    :focus {
        background-color: #4AAD52;
        border: 0px;
    }
`

const ErrorText = styled.div`
    color: red;
`

const Login = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [errorText, setErrorText] = useState('')
    const context = useContext(AuthContext)
    const history = useHistory()

    const tryLogin = async () => {
        try {
            const data = await axios.post(`${process.env.REACT_APP_API_URL}/login`, 
                JSON.stringify ({
                    "email": email,
                    "password": password,
                }), {withCredentials: true}
            )
            if (!data.data.status){
                setErrorText(data.data.message)
            } else {
                context.login()
                history.push('/dashboard')
            }
        } catch (err) {
            setErrorText(err.error)
            console.error(err)
        }
    }

    return (
        <LoginWrapper>
            <LoginForm>
                <LoginHeader>Welcome to Questrack!</LoginHeader>
                <FormGroup>
                    <FormLabel>Email or Username:</FormLabel>
                    <div>
                        <LoginInput
                            required
                            type="email" 
                            id="email" 
                            onChange={(e) => setEmail(e.target.value)} 
                        />
                    </div>
                </FormGroup>
                <FormGroup>
                    <FormLabel>Password:</FormLabel>
                    <div>
                        <LoginInput
                            required
                            type="password" 
                            id="password"
                            onChange={(e) => setPassword(e.target.value)} 
                        />
                    </div>
                </FormGroup>
                <LoginButton onClick={tryLogin}>
                    Sign In
                </LoginButton>
                <ErrorText>{errorText}</ErrorText>
            </LoginForm>
        </LoginWrapper>
    )

}

export default Login;