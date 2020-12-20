import React, { useEffect, useState } from 'react';
import { Row } from 'react-bootstrap';
import Chart from 'react-apexcharts'
import styled from 'styled-components';
import moment from 'moment';

const TimeframeWrapper = styled.div`
    margin: 30px;
`

const TimeframeButton = styled.div`
    background-color: green;
    display: inline-block;
    color: white;
    font-weight: 700;
    cursor: pointer;
    text-align: center;
    margin: 10px;
    width: 45px;
    border-radius: 25px;
`
const TitleWrapper = styled.div`
    margin: 30px;
`


const PositionChart = styled(Chart)`
    margin: 30px;
`

const Symbol = styled.div`
    font-weight: 700;
    font-size: 20px;
`

const BidPrice = styled.div`
    font-size: 20px;
    font-weight: 600;
`

const ChangeInPrice = styled.div`
    font-size: 16px;
    color: ${props => props.isNegative ? 'red' : 'green'};
`

const Option = styled.span`
    color: gray;
    font-weight: 500;
`

const SymbolDesc = styled.div`
    font-size: 10px;
`

const CustomTooltip = (series, seriesIndex, dataPointIndex, w) => {
    return (
        <div>
            
        </div>
    )
}

const AccountGraph = ({quoteInfo, investmentData}) => {
    const [timespan, changeTimespan] = useState(0)

    const chartTheme = {
        chart: {
            toolbar: {
                show: false
            },
            background: '#373F47',
        },
        zoom: {
            enabled: true,
        },
        colors: quoteInfo && quoteInfo.BidPrice >= quoteInfo.PrevDayClosePrice ? ['#08A045'] : ['#B22222'],
        grid: {
            show: false,
        },
        stroke: {
            curve: 'smooth',
            width: 2,
        },
        markers: {
            size: 0,
            hover: {
                size: 0,
            }
        },
        tooltip: {
            enabled: false,
            x: {
                format: 'dd MMM HH:mm',
            },
        },
        xaxis: {
            axisBorder: {
                show: false
            },
            axisTicks: {
                show: false
            },
            labels: {
                datetimeUTC: false,
            },
            type: 'datetime',
            min: timespan > 0 ? moment(moment().subtract(timespan, 'days').format('LL') + ' 9:30').unix() * 1000 : moment((moment().format('LL') + ' 09:30')).unix()*1000,
            max: moment((moment().format('LL') + ' 16:00')).unix()*1000,
        },
        yaxis: {
            axisBorder: {
                show: false
            },
            axisTicks: {
                show: false
            },
            labels: {
                show: false,
            },
            min: quoteInfo ? parseInt(quoteInfo.OpenPrice) - parseInt(quoteInfo.OpenPrice)*0.05: null,
            max: quoteInfo ? parseInt(quoteInfo.OpenPrice) + parseInt(quoteInfo.OpenPrice)*0.05: null,
        },
    }

    const computeChange = {
        change: quoteInfo && (quoteInfo.BidPrice - quoteInfo.PrevDayClosePrice).toFixed(2),
        percent: quoteInfo && ((Math.abs(quoteInfo.BidPrice - quoteInfo.PrevDayClosePrice) / quoteInfo.PrevDayClosePrice)*100).toFixed(2),
    }

    const PositionIsOption = quoteInfo && /[0-9]/g.test(quoteInfo.Symbol);

    return (
        <>  
            <TitleWrapper>
                {quoteInfo &&
                    (<>
                        
                            {PositionIsOption && quoteInfo
                            ? <Symbol>
                                {quoteInfo.Symbol.split(/[0-9]/)[0] + " "} 
                                <Option>({quoteInfo.Symbol.split(/(?<=\D)(?=\d)|(?<=\d)(?=\D)/g).slice(1,4).join('-') + " " + quoteInfo.Symbol.split(/(?<=\D)(?=\d)|(?<=\d)(?=\D)/g).slice(4,).join('')})</Option>
                              </Symbol>
                            : <Symbol>
                                {quoteInfo.Symbol}
                              </Symbol>
                            }
                        <SymbolDesc>{quoteInfo.Description}</SymbolDesc>
                        <BidPrice>${quoteInfo.BidPrice.toFixed(2)}</BidPrice>
                        <ChangeInPrice isNegative={quoteInfo.BidPrice < quoteInfo.PrevDayClosePrice}>
                            {computeChange.change && quoteInfo.BidPrice >= quoteInfo.PrevDayClosePrice
                            ?   <span>↗ </span>
                            :   <span>↘ </span>
                            }
                            {computeChange.change} ({computeChange.percent}%)
                        </ChangeInPrice>
                     </>
                    )
                }
            </TitleWrapper>
            <PositionChart
                options={chartTheme}
                series={[{data: investmentData}]}
                type="line"
                width="700"
            />
            <TimeframeWrapper>
                <TimeframeButton onClick={() => changeTimespan()}>1D</TimeframeButton>
                <TimeframeButton onClick={() => changeTimespan(2)}>2D</TimeframeButton>
                <TimeframeButton onClick={() => changeTimespan(7)}>1W</TimeframeButton>
                <TimeframeButton onClick={() => changeTimespan(30)}>1M</TimeframeButton>
            </TimeframeWrapper> 
        </>
    )

}

export default AccountGraph