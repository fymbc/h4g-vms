import { Link } from "react-router-dom";
import React, { useState } from "react";
import { Sidebar, Menu, MenuItem} from "react-pro-sidebar";
import { Chrono } from "react-chrono";

import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import Diversity1OutlinedIcon from '@mui/icons-material/Diversity1Outlined';
import CardGiftcardOutlinedIcon from '@mui/icons-material/CardGiftcardOutlined';
import SettingsOutlinedIcon from "@mui/icons-material/HelpOutlineOutlined";

const Dashboard = (props) => {

    const username = useState("Jeffrey"); //Change to username

    const items = [{            // Take data from backend
        title: "22 Dec",
        cardTitle: "CPR-AED Workshop",
        cardSubtitle:"3:30pm - 6:30pm",
        cardDetailedText: "Woodlands MRT"
      },
      {
        title: "03 Jan",
        cardTitle: "Volunteering @ THK",
        cardSubtitle:"9:00am - 2:00pm",
        cardDetailedText: "AMK-THK Hospital"
      },
      {
        title: "14 Jan",
        cardTitle: "Donation Drive @ Hougang",
        cardSubtitle:"1:00pm - 3:00pm",
        cardDetailedText: "Hougang Community Club"
      }
        ];

    return <div>
        <div className = 'Banner'>
            <h1>Hi, {username}!</h1>
        </div> 
        <div className = "MainBody">
            <Sidebar>
                <Menu>
                <MenuItem icon={<HomeOutlinedIcon />} component={<Link to="/dashboard" />}>Dashboard</MenuItem>
                <MenuItem icon={<Diversity1OutlinedIcon />} component={<Link to="/activities" />}>Opportunities</MenuItem>
                <MenuItem icon={<CardGiftcardOutlinedIcon />} component={<Link to="/certificates" />}>Certificates</MenuItem>
                <MenuItem icon={<SettingsOutlinedIcon />} component={<Link to="/settings" />}>Profile</MenuItem>
                </Menu>
            </Sidebar>  
            <div className = "Timeline">
                <h1>My Timeline</h1>
                <Chrono 
                    items={items} 
                    mode="VERTICAL" 
                    hideControls = "true"
                    cardWidth = "900"
                    fontSizes={{
                        cardSubtitle: '3vh',
                        cardText: '50px',
                        cardTitle: '5vh',
                        title: '3vh'}}
                    >
                </Chrono>
            </div>
            <div className = "PastEvents">
                <h1>My Past Events</h1>
                <h3> You currently do not have any past events.</h3>
            </div>
        </div>
    </div>
}
export default Dashboard