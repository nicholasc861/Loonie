import styled from "styled-components";

export const NavWrapper = styled.div`
  border-right: 1px solid #b2b7c2;
  max-width: 180px;
  height: 100%;
  margin: 0px;
  padding: 25px 0px;
  background-color: #161819;

  font-size: 14px;
`;

export const NavBrand = styled.div`
  display: inline-block;
  width: 100%;
  padding: 0 30px;
  margin-bottom: 50px;
`;

export const NavBody = styled.div`
  display: flex;
  flex-direction: column;
`;

export const NavSection = styled.div`
  display: block;
  color: #f8f8f9;
  margin-bottom: 20px;

  ul {
    padding: 0px;
  }
`;

export const NavTitle = styled.div`
  width: 100%;
  font-weight: 600;
  padding: 0 30px;
  margin-bottom: 10px;
`;

export const NavItem = styled.li`
  color: #b2b7c2;
  margin: 14px 0px;
  padding: 6px 0px;
  display: flex;
  align-items: center;

  div {
    display: flex;
    width: 100%;
    align-items: center;
    padding: 0px 30px;
  }

  svg {
    height: 18px;
    width: 18px;
    padding-right: 8px;
  }

  :hover {
    color: #FFD15C;

    svg {
        fill: #FFD15C;
    }
  }
`;
