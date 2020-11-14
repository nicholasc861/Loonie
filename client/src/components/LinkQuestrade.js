import React from 'react'

exports.getInitialAuth = async () => {
    window.location = `https://login.questrade.com/oauth2/authorize?client_id=${process.env.APIkey}&response_type=code&redirect_uri=${placeholder}`
}