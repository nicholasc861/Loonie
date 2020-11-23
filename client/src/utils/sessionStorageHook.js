import React, { useState, useEffect } from 'react'

const useSessionStorage = (key, defaultValue) => {
    const storedValue = sessionStorage.getItem(key)
    const value = storedValue ? JSON.parse(storedValue) : defaultValue
    const [sessionValue, setSessionValue] = useState(value)

    useEffect(() => {
        sessionStorage.setItem(key, sessionValue)
    }, [key, sessionValue])

    return [sessionValue, setSessionValue]
}

export default useSessionStorage