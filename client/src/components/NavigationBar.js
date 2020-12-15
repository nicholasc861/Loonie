import React, {useContext, useEffect} from 'react'
import { Nav, Navbar, } from 'react-bootstrap'
import styled from 'styled-components'
import { AuthContext } from '../context/auth'

import Logo from '../assets/Logo.png'

const Navigation = styled(Navbar)`
    font-weight: 700;
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

        :hover {
            color: #FFFCFF;
        }
    }
`

const Spacer = styled.div`
    width: 20px;
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
`

const RegisterButton = styled(Nav.Link)`
    heigth: 24px;
    width: 122px;
    align-items: center;
    border: 2px solid #4AAD52;
    border-radius: 24px;
    text-align: center;
    display: inline;
    background: #4AAD52;
`

const NavText = styled.span`
    color: #FFFCFF;
    line-height: 24px;
    font-weight: 700;
    font-size: 16px;
`

const NavigationBar = () => {
    const context = useContext(AuthContext)

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
                    {context.isLoggedIn ? (
                    <>
                        <Nav.Link href="/dashboard">
                            <NavText>
                                Dashboard
                            </NavText>
                        </Nav.Link>
                        <Nav.Link href="/" onClick={() => context.logout()}>
                            <NavText>
                                Sign Out
                            </NavText>
                        </Nav.Link>
                    </>
                    ) : (
                    <>
                        <Nav.Link href="/login">
                            <NavText>
                                Log In
                            </NavText>
                        </Nav.Link>
                        <Spacer />
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