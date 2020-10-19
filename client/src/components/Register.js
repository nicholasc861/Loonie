import React from 'react';

import styled from 'styled-components'
import { Button, Form } from 'react-bootstrap'

const RegisterHeader = styled.div`
    font-size: 20px;
    font-weight: 600;
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

const Register = () => {
    return (
        <RegisterWrapper>
            <RegisterForm>
                <RegisterHeader>Questrack allows you visualize your investments on Questrade.</RegisterHeader>
                <FormGroup>
                    <RegisterInput></RegisterInput>
                    <RegisterInput></RegisterInput>
                </FormGroup>
            </RegisterForm>
        </RegisterWrapper>
    )
}

export default Register