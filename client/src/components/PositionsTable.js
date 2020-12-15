import React, { useEffect, useState } from 'react'
import {Table} from 'react-bootstrap'
import styled from 'styled-components'
import axios from 'axios'

const TableHeader = styled.th`
    color: #2F2F2F;
    font-size: 14px;
    width: ${props => props.size ? props.size : '13%'};
    text-align: ${props => props.align ? 'left' : 'center'};
`

const PositionHeader = styled.div`
    font-size: 14px;
    font-weight: 700;
`

const OptionText = styled.div`
    font-weight: 400;
    font-size: 12px;
`

const PosTable = styled(Table)`
    padding: 50px;
    font-size: 14px;
    vertical-align: middle;

    tbody td {
        text-align: center;
        vertical-align: middle;
    }
`
const PosRow = styled.tr`
    cursor: pointer;
`

const PLWrapper = styled.div`
    color: ${props => props.isNegative ? 'red' : 'green'};
`

const PositionsTable = ({onclick, positions, p_l}) => {
    return (
        <PosTable>
            <thead>
                <tr>
                    <TableHeader align={true} size={'35%'}>Position</TableHeader>
                    <TableHeader>Open Quantity</TableHeader>
                    <TableHeader>Closed Quantity</TableHeader>
                    <TableHeader>Average Cost</TableHeader>
                    <TableHeader>Total Cost</TableHeader>
                    <TableHeader>Live P&L</TableHeader>
                </tr>
            </thead>
            <tbody>
            {positions.map((position, index) => {
                return(
                    <PosRow onClick={() => onclick(position.QuestradeID)} key={index}>
                        <th>
                            {position.IsOption ?
                                <PositionHeader>
                                    {position.Symbol.split(/[0-9]/)[0]}
                                    <OptionText>
                                        {position.Symbol.split(/(?<=\D)(?=\d)|(?<=\d)(?=\D)/g).slice(1,4).join('-') + " "}
                                        {position.Symbol.split(/(?<=\D)(?=\d)|(?<=\d)(?=\D)/g).slice(4,)}
                                    </OptionText>
                                </PositionHeader>
                            :
                                <PositionHeader>
                                    {position.Symbol}
                                </PositionHeader>
                            }
                        </th>
                        <td>
                            {position.OpenQuantity}
                        </td>
                        <td>
                            {position.ClosedQuantity}
                        </td>
                        <td>
                            {position.AverageEntryPrice.toFixed(2)}
                        </td>
                        <td>
                            {position.TotalEntry.toFixed(2)}
                        </td>
                        <td>
                            {p_l && p_l[index] && <PLWrapper isNegative={p_l[index].isNegative}>{p_l[index].P_L.toFixed(2)}</PLWrapper>}
                        </td>
                    </PosRow>
                    )
                })
            }   
            </tbody>    
        </PosTable>
    )
}

export default PositionsTable