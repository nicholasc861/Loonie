import React, { useContext, useEffect } from "react";

import { ReactComponent as Logo } from "../assets/Logo.svg";
import { ReactComponent as DashboardIcon } from "../assets/dashboard.svg";
import { ReactComponent as CalendarIcon } from "../assets/calendar.svg";
import { ReactComponent as ReportIcon } from "../assets/report.svg";
import { ReactComponent as CoinIcon } from "../assets/coins.svg";
import { ReactComponent as BriefcaseIcon } from "../assets/briefcase.svg";
import { ReactComponent as GearIcon } from "../assets/gear.svg";

import {
  NavWrapper,
  NavBody,
  NavBrand,
  NavSection,
  NavTitle,
  NavItem,
} from "./NavigationBar.styles";

const contents: Record<string, any> = {
  Analytics: [
    { icon: <DashboardIcon fill={"#B2B7C2"} />, name: "Dashboard" },
    { icon: <CalendarIcon fill={"#B2B7C2"} />, name: "Calendar" },
    { icon: <ReportIcon stroke={"#B2B7C2"} />, name: "Reports" },
  ],
  Trades: [
    { icon: <CoinIcon stroke={"#B2B7C2"} />, name: "Trades"},
  ],
  Management: [
    {icon: <BriefcaseIcon stroke={"#B2B7C2"} />, name: "Accounts"},
    {icon: <GearIcon fill={"#B2B7C2"} />, name: "Settings"}
  ]
};

const NavigationBar = () => {
  return (
    <NavWrapper>
      <NavBrand>
        <Logo width={36} height={36} />
      </NavBrand>
      <NavBody>
        {Object.keys(contents).map((key) => (
          <NavSection>
            <NavTitle>{key}</NavTitle>
            <ul>
                {contents[key].map(({ icon, name }: any) => (
                <NavItem>
                    <div>
                        {icon}
                        {name}
                    </div>
                </NavItem>
                ))}
            </ul>
          </NavSection>
        ))}
      </NavBody>
    </NavWrapper>
  );
};

export default NavigationBar;
