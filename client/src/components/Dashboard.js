import React, { useState, useEffect } from 'react'
import Chart from 'react-apexcharts'
import axios from 'axios'

const chartTheme = {
    chart: {
        toolbar: {
            show: false
        },
        background: '#FFFFFF',
        fontFamily: 'Karla, sans-serif'
    },
    xaxis: {
        show: false,
        labels: {
            show: false
        }
    },
    yaxis: {
        show: false,
        labels: {
            show: false
        }
    }
}

const Dashboard = () => {
    const [investmentData, updateInvestmentData] = useState([10, 41, 35, 51, 49, 62, 69, 91, 148])

    useEffect(() => {
        setInterval(async () => {
            await axios.post('https://api01.iq.questrade.com/v1/markets/quotes/options')
        }, 30000) 
    })
    

    const updateChart = () => {
        const newData = [20,30,40,50]


        updateInvestmentData(investmentData => [...investmentData, ...newData])

    }

    return (
        <>
            <button onClick={updateChart}>test</button>
            <Chart
                options={chartTheme}
                series={[{name: "Desktops", data: investmentData}]}
                type="line"
                width="700"
            />
        </>

    )
}

export default Dashboard