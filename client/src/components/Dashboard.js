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
    const [investmentData, updateInvestmentData] = useState([{name: "Desktops",
    data: [10, 41, 35, 51, 49, 62, 69, 91, 148]}])

    useEffect(() => {

    })
    

    const updateChart = () => {
        const newData = []


        updateInvestmentData(newData)

    }

    return (
        <>
            <Chart
                options={chartTheme}
                series={investmentData}
                type="line"
                width="700"
            />
        </>

    )
}

export default Dashboard