import { Link } from "react-router-dom";
import React, { useState } from "react";
import { Sidebar, Menu, MenuItem} from "react-pro-sidebar";
import { Card, CardContent, Typography, Button } from '@mui/material';

import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import Diversity1OutlinedIcon from '@mui/icons-material/Diversity1Outlined';
import CardGiftcardOutlinedIcon from '@mui/icons-material/CardGiftcardOutlined';
import SettingsOutlinedIcon from "@mui/icons-material/HelpOutlineOutlined";

const Activities = (props) => {
    const username = useState("Jeffrey"); //Change to username
    
    const details = [
      { id: "0", title: "Volunteering with NUS", date: "27 December", time: "3:30pm - 6:30pm", venue: "Kent Ridge Hall"},
      { id: "1", title: "Donation Drive @ City Square", date: "28 December", time: "2:30pm - 4:30pm", venue: "City Square Mall"},
      { id: "2", title: "Volunteering with NUS", date: "1 January", time: "12:30pm - 6:30pm", venue: "University Town"}
    ];

    return (
      <div>
        <div className='Banner'>
          <h1>Hi, {username}!</h1>
        </div>
        <div className="MainBody">
          <Sidebar>
            <Menu>
              <MenuItem icon={<HomeOutlinedIcon />} component={<Link to="/dashboard" />}>Dashboard</MenuItem>
              <MenuItem icon={<Diversity1OutlinedIcon />} component={<Link to="/activities" />}>Opportunities</MenuItem>
              <MenuItem icon={<CardGiftcardOutlinedIcon />} component={<Link to="/certificates" />}>Certificates</MenuItem>
              <MenuItem icon={<SettingsOutlinedIcon />} component={<Link to="/settings" />}>Profile</MenuItem>
            </Menu>
          </Sidebar>
          <div className="Modules">
            {details.map((detail, index) => (
              <Card key={index}>
                <CardContent>
                  <Typography className = "ModuleTitle">{detail.title}</Typography>
                  <Typography>{detail.date} | {detail.time}</Typography>
                  <Typography>{detail.venue}</Typography>
                </CardContent>
                <Link to={`/activities/${detail.id}`}>
                  <Button variant="contained" color="primary"> Enrol </Button>
                </Link>
              </Card>
            ))}
          </div>
        </div>
      </div>
    );
}
export default Activities;
