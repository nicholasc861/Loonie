import React, { useState, useEffect } from 'react'
import Chart from 'react-apexcharts'
import axios from 'axios'
import { Form, Row, Container, Col, Table} from 'react-bootstrap'

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
    const [investmentData, updateInvestmentData] = useState([2, 3, 4])
    const [selectedAccount, setSelectedAccount] = useState()
    const [accountNumbers, updateAccountNumbers] = useState([])
    const [currentPositions, updateCurrentPositions] = useState([1, 2])

    useEffect(() => {
        // setInterval(async () => {
        //     await axios.post('https://api01.iq.questrade.com/v1/markets/quotes/options')
        // }, 30000) 

        getAccounts()
        getCurrentPositions()
    }, [])

    const getAccounts = async () => {
        try {
            const data = await axios.get(`http://localhost:8080/user/accounts`, {withCredentials: true, origin: 'http://localhost:8080'})

            if (data){
                for (let i = 0; i < data.data.length; i ++){
                    console.log(data.data[i].AccountID)
                    updateAccountNumbers((accountNumbers) => [...accountNumbers, data.data[i]])
                }
                console.log(accountNumbers)
            }
        } catch (err){
            console.error(err)
        }
    }
    
    const getCurrentPositions = async () => {
        try {
            const currentAccount = selectedAccount
            const questradePositions = await axios.get(`https://api01.iq.questrade.com/v1/markets/${currentAccount}/options`, {
                headers: { Authorization: `Bearer ${process.env.QUESTRADE_API_ACCESS_TOKEN}` }
            })

            if (questradePositions){
                const positions = questradePositions.data.positions

                for (let i = 0; i < positions.length; i++) {
                    const selectedPosition = positions[i]
                    const addPosition = await axios.post(`http://localhost:8080/user/account/positions`, 
                        JSON.stringify({
                            AccountID: currentAccount,
                            QuestradeID: selectedPosition.symbolId,
                            Symbol: selectedPosition.symbol,
                            OpenQuantity: selectedPosition.openQuantity,
                            AverageEntryPrice: selectedPosition.averageEntryPrice,
                            IsOption: /\d/.test(selectedPosition.symbol),
                            status: true
                        }), {withCredentials: true})
                }
            }

            const data = await axios.get(`http://localhost:8080/user/account/positions`, {withCredentials: true})
        } catch (err) {
            console.error(err)
        }

        console.log(selectedAccount)
        // accountNumbers.forEach(async (AccountNumber) => {
        //     try {
        //         const accountData = await axios.get(`https://api01.iq.questrade.com/v1/accounts/${AccountNumber}/postions`,
        //             {headers: {
        //                 Authorization: "test",
        //             }})
        //         updateCurrentPositions([...currentPositions, accountData.positions])
        //     } catch (err) {
        //         console.error(err)
        //     }
        // })
    }

    const updateChart = () => {
        updateInvestmentData(investmentData => [...investmentData, ...test])

    }

    return (
        <>
            <Chart
                options={chartTheme}
                series={[{name: "Desktops", data: investmentData}]}
                type="line"
                width="700"
            />

            <Container>
                <Row md={4}>
                    <Col>
                        <h3>Account: </h3>
                    </Col>
                    <Col>
                        <Form>
                            <Form.Group controlId="exampleForm.SelectCustom">
                                <Form.Control as="select" custom onChange={(e) => setSelectedAccount(e.target.value)}>
                                    {accountNumbers &&
                                        accountNumbers.map((account, id) => {
                                            return (<option key={id} value={account.AccountID}>{account.AccountType} - {account.AccountID}</option>)
                                        })}
                                </Form.Control>
                            </Form.Group>
                        </Form>
                    </Col>
                </Row>
            </Container>


            <button onClick={getCurrentPositions}>Update Positions</button>
            <Table borderless>
                <thead>
                    <tr>
                        <th>Position</th>
                        <th>Open Quantity</th>
                        <th>Closed Quantity</th>
                        <th>Average Cost</th>
                        <th>Live P&L</th>
                    </tr>
                </thead>
                {currentPositions ? 
                    currentPositions.map((position, index) => {
                        return (
                        <tr key={index}>
                            <th>{position}</th>
                        </tr>
                    )
                    })
                : null
                }
            </Table>
        </>

    )
}

export default Dashboard