import React from 'react'
import { Nav, Navbar, } from 'react-bootstrap'
import styled from 'styled-components'

import Logo from '../assets/Logo.png'

const Navigation = styled(Navbar)`
    font-family: "Open Sans", sans-serif;
    font-weight: 600;
    top: 0;
    left: 0;
    background-color: #2F2F2F;
    padding: 10px 30px;
    width: 100%;
    margin: 0 auto;
    height: 3 rem;
    display: flex;
    flex-direction: row;
`

const NavBrandWrapper = styled.div`
    vertical-align: center;
    
    .navbar-brand {
        color: #FFFCFF;
        font-size: 22px;
    }

`

const LogoImg = styled.img`
    width: 22px;
    height: 22px;
    margin: 0px 12px 3px;
    align-content: centre; 
`

const NavLinks = styled(Nav)`
    width: 100%;
    font-weight: 400;
    font-family: "Karla", sans-serif;
`

const RegisterButton = styled(Nav.Link)`
    margin: 0px 5px;
    border: 2px solid #4AAD52;
    border-radius: 3em;
    background: #4AAD52;
`

const NavText = styled.span`
    color: #FFFCFF;
    font-weight: 600;
    padding: 10px;
    font-size: 17px;
`

const NavigationBar = () => {
    const user = false;


    return (
        <Navigation>
            <NavBrandWrapper>
                <Navigation.Brand href="/">
                    Questrack
                    <LogoImg src={Logo} />
                </Navigation.Brand>
            </NavBrandWrapper>
            <Navigation.Toggle aria-controls="responsive-navbar-nav" />
            <Navbar.Collapse id="responsive-navbar-nav">
                <NavLinks className="justify-content-end">
                    {user ? (
                    <>
                    </>
                    ) : (
                    <>
                        <Nav.Link href="/login">
                            <NavText>
                                Log In
                            </NavText>
                        </Nav.Link>
                        <RegisterButton href="/register">
                            <NavText>
                                Sign Up
                            </NavText>
                        </RegisterButton>
                    </>
                    )}
                </NavLinks>
            </Navbar.Collapse>
        </Navigation>
    )
}

export default NavigationBar