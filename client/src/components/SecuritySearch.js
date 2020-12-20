import React, { useState, useEffect } from 'react';
import moment from 'moment';
import styled from 'styled-components';
import { InputGroup,  Row, Form, Button } from 'react-bootstrap';

const InputWrapper = styled(InputGroup)`
    margin: 0px 30px;
`
const AssetSelector = styled.button`
    font-size: 14px;
    font-weight: 700;
    color: ${props => props.checked ? '#FFFFFF' : '#4AAD52'};
    box-shadow: none;
    border-radius: 2px;
    background-color: ${props => props.checked ? '#4AAD52' : '#FFFFFF'};
    border: 1px solid #4AAD52;
`

const SearchWrapper = styled.div`
    margin-top: 40px;
`

const SearchResultWrapper = styled.div`
    position: absolute;
    width: 82%;
    padding-bottom: 3px;
    margin-top: 34px;
    z-index: 999;
    text-align: left;
    background: #FFFFFF;
    box-shadow: 0 4px 6px 0 rgba(32,33,36,0.28);
    border-bottom-left-radius: 12px;
    border-bottom-right-radius: 12px;
    border: 1px;
`

const ResultWrapper = styled.div`
    &:hover {
        background-color: #d3d3d3;
    }
`

const SearchResult = styled.div`
    cursor: pointer;
    margin-left: 10px;
    padding: 5px 0px;
`

const PositionSymbol = styled.div`
    font-size: 14px;
`

const PositionDescription = styled.div`
    font-size: 10px;
    color: #80868b;
    text-overflow: ellipsis;

`

const PositionSearch = styled.input`
    width: 200px;
`

const PositionOptionsWrapper = styled.div`
    margin: 30px 0px 30px 30px;
`

const OptionDropdown = styled(Form)`
    font-weight: 600;
    margin-bottom: 10px;

    .custom-select {
        font-size: 0.8rem;
    }

    label {
        font-size: 0.8rem;
    }
`

const QuantityWrapper = styled(InputGroup)`
    height: 30px;

    button {
        background-color: #4AAD52;
        width: 35px;
        border: none;

        &:disabled {
            background-color: #4AAD52;
            border-color: #4AAD52;
        }
    }
`

const QuantitySelector = styled.input`
    text-align: center;
`

const GraphButton = styled(Button)`
    background-color: #4AAD52;
    border-color: #4AAD52;
    margin: 10px 0px;

    &:hover {
        background-color: #4AAD52;
        border-color: #4AAD52;
    }
`

const QuantityComponent = ({disabled, quantity, setQuantity}) => {
    return (
        <OptionDropdown>
            <OptionDropdown.Label>Quantity:</OptionDropdown.Label>
            <QuantityWrapper>  
                <Button disabled={disabled} onClick={() => {setQuantity(quantity - 5)}}>-</Button>
                <QuantitySelector disabled={disabled} value={quantity} onChange={(e) => setQuantity(parseInt(e.target.value), 10)} />
                <Button disabled={disabled} onClick={() => {setQuantity(quantity + 5)}}>+</Button>
            </QuantityWrapper>
        </OptionDropdown>
    )
}


const SecuritySearch = ({searchResults, checked, onChange, onOptionChoose, onButton, onChoosePosition, optionChain}) => {
    const [selected, setSelected] = useState();
    const [quantity, setQuantity] = useState(1);
    const [searchFocused, setSearchFocused] = useState(false);
    const [optionType, setOptionType] = useState('Call');

    const [optionSelected, setOptionSelected] = useState(0);
    const [strikePrice, setStrikePrice] = useState(0);

    const setGraph = () => {
        if (checked === 'option'){
            let id;
            if (optionType == 'Call'){
                id = optionChain[optionSelected].chainPerRoot[0].chainPerStrikePrice[strikePrice].callSymbolId;
                onOptionChoose(id)
            } else {
                id = optionChain[optionSelected].chainPerRoot[0].chainPerStrikePrice[strikePrice].putSymbolId;
                onOptionChoose(id)
            }
        }
    }

    useEffect(() => {

        console.log(optionSelected, strikePrice)

    }, [optionChain, optionSelected, strikePrice])
    


    return (
        <>  
            <Row>
                <SearchWrapper>
                    <InputWrapper>
                        <PositionSearch value={selected && selected} onBlur={() => setSearchFocused(false)} onFocus={() => setSearchFocused(true)} onChange={(e) => {onChange(e.target.value); setSelected()}}/>
                        <AssetSelector checked={checked === 'stock'} onClick={() => {onButton('stock')}}>STK</AssetSelector>
                        <AssetSelector checked={checked === 'option'} onClick={() => {onButton('option')}}>OPT</AssetSelector>

                        {searchFocused && !selected && searchResults && searchResults.length > 0 
                        && <SearchResultWrapper>
                            {
                                searchResults.map((result, index) => {
                                    return(
                                        <ResultWrapper key={index}>
                                            <SearchResult onMouseDown={() => {onChoosePosition(result.symbolId); setSelected(result.symbol)}}>
                                                <PositionSymbol>{result.symbol} {result.listingExchange ? '(' + result.listingExchange + ')' : ''}</PositionSymbol>
                                                <PositionDescription>
                                                    {result.description.length > 50 
                                                        ? result.description.slice(0, 40) + "..."
                                                        : result.description
                                                    }
                                                </PositionDescription>
                                            </SearchResult>
                                        </ResultWrapper>
                                    );
                                })
                            }
                            </SearchResultWrapper>
                        }
                    </InputWrapper>
                </SearchWrapper>
            </Row>
            <Row>
                {checked === 'option' 
                ?   (
                    <PositionOptionsWrapper>
                        <QuantityComponent disabled={!selected} quantity={quantity} setQuantity={setQuantity} />
                        <OptionDropdown>
                            <OptionDropdown.Label>Option Type:</OptionDropdown.Label>
                            <OptionDropdown.Control as="select" custom disabled={!selected} onChange={(e) => setOptionType(e.target.value)}> 
                                <option value="Call">Call</option>
                                <option value="Put">Put</option>
                            </OptionDropdown.Control>
                        </OptionDropdown>
                        <OptionDropdown>
                            <OptionDropdown.Label>Expiry Date:</OptionDropdown.Label>
                            <OptionDropdown.Control as="select" custom disabled={!selected} onChange={(e) => setOptionSelected(parseInt(e.target.value))}> 
                                {optionChain.length >0 &&
                                    optionChain.map((optionC, index) => {
                                        return (
                                            <option value={index}>{ moment(optionC.expiryDate).format('YYYY-MM-DD')}</option>
                                        )
                                    })
                                }
                            </OptionDropdown.Control>
                        </OptionDropdown>
                        <OptionDropdown>
                            <OptionDropdown.Label>Strike Price:</OptionDropdown.Label>
                            <OptionDropdown.Control as="select" custom disabled={!selected} onChange={(e) => {setStrikePrice(e.target.value);}}> 
                                {optionChain.length > 0 && optionSelected !== undefined && 
                                    optionChain[optionSelected].chainPerRoot[0].chainPerStrikePrice.map((strike, index) => {
                                        return (
                                            <option value={index}>{strike.strikePrice.toFixed(2)}</option>
                                        )
                                    })
                                }
                            </OptionDropdown.Control>
                        </OptionDropdown>

                        <GraphButton onClick={() => {setGraph()}}>Graph Option!</GraphButton>
                    </PositionOptionsWrapper> 
                 )
                :
                (
                    <PositionOptionsWrapper>
                        <QuantityComponent disabled={!selected} quantity={quantity} setQuantity={setQuantity} />
                        <GraphButton onClick={() => {setGraph()}}>Graph Option!</GraphButton>
                    </PositionOptionsWrapper>
                )
                }
            </Row>
        </>
    );

};

export default SecuritySearch;