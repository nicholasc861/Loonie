import { useState, useEffect } from 'react'
import axios from 'axios'

const AuthCallback = () => {
    const [accessToken, setAccessToken] = useState()
    const [questradeAPIURL, setQuestradeAPIURL] = useState()

    useEffect(() => {
        const questrackCallback = window.location.search;
        const urlParams = new URLSearchParams(questrackCallback);
        const code = urlParams.get('code')

        const tokenResponse = getToken(code);
        const refreshToken = tokenResponse.refresh_token;

        storeRefreshToken(refreshToken);
    })


    const getToken = async (code) => {
        try {
            const data = await axios.post(`https://login.questrade.com/oauth2/token?client_id=${process.env.REACT_APP_QUESTRADE_CLIENT_ID}&code=${code}&grant_type=authorization_code`)
            return data.data;
        } catch (err) {
            console.error(err);
        }
    }

    const storeRefreshToken = async (refreshToken) => {
        try {
            axios.post(`${process.env.REACT_APP_API_URL}/addtoken`, {
                RefreshToken: refreshToken,
            })
        } catch (err) {
            console.error(err);
        }
    }
}

export default AuthCallback