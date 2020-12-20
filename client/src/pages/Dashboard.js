import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { Form, Row, Container, Col, Button} from 'react-bootstrap'
import styled from 'styled-components'

import AccountGraph from '../components/AccountGraph'
import PositionsTable from '../components/PositionsTable'
import SecuritySearch from '../components/SecuritySearch'

const AccountSelector = styled(Form)`
    width: 50%;
    float: right;
    padding: 0px 0px 10px;
    font-weight: 600;

    .custom-select {
        font-size: 0.8rem;
    }

    label {
        font-size: 0.8rem;
        margin-bottom: 0.2rem;
    }
`

const TotalAccountBalance = styled(Button)`
    position: relative;
    background-color: black;
    top: 25%;
`

const Dashboard = () => {
    const [loading, setLoading] = useState(true)

    const [investmentData, updateInvestmentData] = useState([])
    const [quoteInfo, updateQuoteInfo] = useState()
    const [selectedPositions, updateSelectedPositions] = useState([])
    const [selectedAccount, setSelectedAccount] = useState(0)

    const [selectedPositionForQuote, updatePositionForQuote] = useState()
    const [selectedPositionForQuery, updatePositionForQuery] = useState()

    const [optionsChain, updateOptionsChain] = useState([]);
    const [selectedOptionForQuote, updateOptionForQuote] = useState()

    const [queryType, updateQueryType] = useState('stock')
    const [searchResults, updateSearchResults] = useState([])

    const [accountNumbers, updateAccountNumbers] = useState([])
    const [currentPositions, updateCurrentPositions] = useState([])
    const [livePL, updateLivePL] = useState([])

    useEffect(() => {
        // Get all account numbers
        if (accountNumbers.length === 0){
            getAccounts()
        }
        //Wait for account numbers before retrieving positions
        if (selectedAccount && currentPositions.length === 0){
            getCurrentPositions()
        }
        // Only update positions displayed if there are positions and accounts returned
        if(currentPositions.length > 0 && selectedAccount){
            updateSelectedPositions(currentPositions.filter((position) => position.AccountID == selectedAccount))
        }
    }, [accountNumbers, selectedAccount, currentPositions])
 
    useEffect(() => {
        // Run once to retrieve P&L Live on initial load
        if(livePL.length !== selectedPositions.length){
            getQuote(selectedPositions, selectedAccount);
        }

        const interval = setInterval(() => getQuote(selectedPositions, selectedAccount), 30000)
        return () => clearInterval(interval)
    }, [selectedPositions, selectedAccount, livePL])

    // Check if Security Selected is Option or Stock and work accordingly
    useEffect(() => {
        if (queryType == "option" && selectedPositionForQuote){
            getOptionsChain(selectedPositionForQuote)
        } else if (selectedPositionForQuote){
            updateInvestmentData([])
            getHistoricalForChart(selectedPositionForQuote)
            const interval = setInterval(() => getQuoteForChart(selectedPositionForQuote), 30000)
            return () => clearInterval(interval)
        }
    }, [queryType, selectedPositionForQuote])

    // Get Quote for Option Selected
    useEffect(() => {
        if (selectedOptionForQuote){
            updateInvestmentData([])
            getHistoricalForChart(selectedOptionForQuote)
            const interval = setInterval(() => getQuoteForChart(selectedOptionForQuote), 30000)
            return () => clearInterval(interval)
        }
    }, [selectedOptionForQuote])


    // Search for Stocks matching Query Terms
    useEffect(() => {
        if(selectedPositionForQuery){
            getSearch(selectedPositionForQuery)
        } else {
            updateSearchResults()
        }
    }, [selectedPositionForQuery])


    const getAccounts = async () => {
        try {
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/user/accounts`, {withCredentials: true})
            if (res){
                for (let i = 0; i < res.data.data.length; i ++){
                    updateAccountNumbers((accountNumbers) => [...accountNumbers, res.data.data[i]])
                }
                setSelectedAccount(res.data.data[0].AccountID)
            }
        } catch (err){
            console.error(err)
        }
    }
    
    const getCurrentPositions = async () => {
        try {
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/user/positions`, {withCredentials: true})
            if (res.data.data) {
                updateCurrentPositions(res.data.data)
            }    
        } catch (err) {
            console.error(err)
        }
    }

    const getQuote = async (positions, account) => {
        const filteredPositions = positions.filter((position) => position.AccountID == account)
        try {
            const res = await axios.post(`${process.env.REACT_APP_API_URL}/user/livepl`, JSON.stringify({
                positions: filteredPositions
            }), {withCredentials: true})
            updateLivePL(res.data.data)
        } catch (err){
            console.error(err)
        }
    }

    const getSearch = async (query) => {
        try {
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/user/query/${query}`, {withCredentials: true})
            if (res.data.data) {
                updateSearchResults(res.data.data.slice(0,5))
            }
        } catch (err) {
            console.error(err)
        }
    }
    
    const getQuoteForChart = async (selectedPosition) => {
        try {
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/user/quote/${selectedPosition}`, {withCredentials: true});
            if (res.data.data) {
                updateQuoteInfo(res.data.data)
                updateInvestmentData((currentData) => [...currentData, [res.data.data.TimeQuoted *1000, res.data.data.BidPrice]]);
            }
        } catch (err) {
            console.error(err)
        }
    }

    const getHistoricalForChart = async (selectedPosition) => {
        try {
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/user/hisquote/${selectedPosition}`, {withCredentials: true})
            if (res.data.data) {
                updateQuoteInfo(res.data.data[0])
                updateInvestmentData(res.data.data.map(({TimeQuoted, BidPrice}) => [TimeQuoted * 1000, BidPrice]));
            }
        } catch (err) {
            console.error(err)
        }
    }

    const getOptionsChain = async (questradeID) => {
        try {
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/user/option/${questradeID}`, {withCredentials: true});
            if (res.data.data) {
                updateOptionsChain(res.data.data)
            }
        } catch (err) {
            console.error(err)
        }
    }

    const handlePositionChange = (newValue) => {
        updatePositionForQuery(newValue);
    }

    const handlePositionSelect = (newValue) => {
        updatePositionForQuote(newValue);
    }
    
    const handlePositionType = (newValue) => {
        updateQueryType(newValue);
    }

    const handleOptionQuote = (newValue) => {
        updateOptionForQuote(newValue);
    }

    return (
        <> 
            <Container>
                <Row>
                    <Col md={8}>
                        {investmentData &&
                            <AccountGraph quoteInfo={quoteInfo} position={selectedPositionForQuote} investmentData={investmentData} />
                        }
                    </Col>
                    <Col md={4}>                
                        <SecuritySearch optionChain={optionsChain} searchResults={searchResults} checked={queryType} onOptionChoose={handleOptionQuote} onChoosePosition={handlePositionSelect} onButton={handlePositionType} onChange={handlePositionChange}/>
                    </Col>
                </Row>
            </Container>
            <Container>
                <Row>
                    <Col md={4}>
                        <TotalAccountBalance onClick={() => {}}>See Live Balance</TotalAccountBalance>
                    </Col>
                    <Col md={4}></Col>
                    <Col md={4}>
                        <AccountSelector>
                            <AccountSelector.Label>Account:</AccountSelector.Label>
                            <AccountSelector.Control as="select" custom onChange={(e) => {setSelectedAccount(e.target.value); }}>
                                {accountNumbers &&
                                    accountNumbers.map((account, id) => {
                                        return (<option key={id} value={account.AccountID}>{account.AccountType} - {account.AccountID}</option>)
                                    })}
                            </AccountSelector.Control>
                        </AccountSelector>
                    </Col>
                </Row>
            </Container>
            <Container>
                <Row>
                    {selectedPositions && 
                        <PositionsTable onclick={handlePositionSelect} positions={selectedPositions} p_l={livePL} />
                    }
                </Row>
            </Container>
        </>
    )
}

export default Dashboard