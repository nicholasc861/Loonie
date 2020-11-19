import React, {useState} from 'react';

import styled from 'styled-components'
import axios from 'axios'
import { Button, Form } from 'react-bootstrap'

const RegisterHeader = styled.div`
    font-size: 20px;
    font-weight: 600;
    text-align: center;
`

const RegisterWrapper = styled.div`
    font-family: "Karla", sans-serif;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 500px;
`

const FormGroup = styled.div`
    margin: 20px 0px;
`

const RegisterInput = styled.input`
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

const RegisterForm = styled(Form)`
    display: block;
    margin: auto;

`

const RegisterButton = styled(Button)`
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

const Register = () => {
    const [firstName, setFirstName] = useState('')
    const [lastName, setLastName] = useState('')
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [errorMessage, setErrorMessage] = useState('')

    const registerUser = async() => {
        try {
            const data = await axios.post('http://localhost:8080/register', 
                JSON.stringify ({
                    firstname: firstName,
                    lastname: lastName,
                    email: email,
                    password: password,
                })
            )

            if (data.status === 401) {
                setErrorMessage("Invalid email or password.")
            }
        } catch (err) {
            console.log(err)
        }
    }

    return (
        <RegisterWrapper>
            <RegisterForm>
                <RegisterHeader>Questrack allows you visualize your investments on Questrade.</RegisterHeader>
                <div class="row">
                    <FormGroup>
                        <FormLabel>First Name</FormLabel>
                        <div>
                            <RegisterInput onChange={(e) => setFirstName(e.target.value)} style={{"margin-right": '10px'}}></RegisterInput>
                        </div>
                    </FormGroup>
                    <FormGroup>
                        <FormLabel>Last Name</FormLabel>
                        <div>
                            <RegisterInput onChange={(e) => setLastName(e.target.value)}></RegisterInput>
                        </div>
                    </FormGroup>
                </div>
                <div class="row">
                    <FormGroup>
                        <FormLabel>Email</FormLabel>
                        <div>
                            <RegisterInput onChange={(e) => setEmail(e.target.value)} style={{width: '678px'}}></RegisterInput>
                        </div>
                    </FormGroup>
                </div>
                <div class="row">
                    <FormGroup>
                    <FormLabel>Password (min 8. characters)</FormLabel>
                        <div>
                            <RegisterInput type="password" onChange={(e) => setPassword(e.target.value)} style={{width: '678px'}}></RegisterInput>
                        </div>
                    </FormGroup>
                </div>
                <div class="row" style={{"align-items": "center"}}>
                    <RegisterButton onClick={registerUser}>Register!</RegisterButton>
                </div>
                {errorMessage &&
                    <div>{errorMessage}</div>
                }
            </RegisterForm>
        </RegisterWrapper>
    )
}

export default Register