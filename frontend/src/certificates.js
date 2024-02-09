import { Link } from "react-router-dom";
import React, { useState } from "react";
import { Sidebar, Menu, MenuItem} from "react-pro-sidebar";
import { Button, TableBody, TableCell, TableContainer, TableRow} from '@mui/material';

import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import Diversity1OutlinedIcon from '@mui/icons-material/Diversity1Outlined';
import CardGiftcardOutlinedIcon from '@mui/icons-material/CardGiftcardOutlined';
import SettingsOutlinedIcon from "@mui/icons-material/HelpOutlineOutlined";

const Certificates = (props) => {

    const username = useState("Jeffrey"); //Change to username

    const handleRequest = (taskId) => { //Link to backend for approval

        const updatedTasks = tasks.map((task) => {
          if (task.id === taskId) {

            task.requestable = 'no'; 
          }
          return task;
        });
        setTasks(updatedTasks);
      };

      const [tasks, setTasks] = useState([
        {id: "0", name: "Volunteering @ NUS", requestable: "yes", downloadable: "no"},
        {id: "1", name: "Card Making Workshop", requestable: "no", downloadable: "no"},
        {id: "2", name: "Donation Drive @ AMK", requestable: "no", downloadable: "yes"}
    ]);

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
            <div className = "CertDescription">
                <h1>Eligible Activities</h1>
                <p>Here you can request certificates for your activities. Once approved, you can download it for your personal use!</p>

                <TableContainer>
                    <TableBody>
                        {tasks.map((task) => (
                            <TableRow
                                key={task.name}
                                sx={{ '&:last-child td, &:last-child th':
                                    { border: 0 } }}
                            >
                                <TableCell component="th" scope="row">
                                    {task.name}
                                </TableCell>
                                <TableCell align="right">
                                    {task.requestable === 'yes' ? (
                                        <Button variant="contained" color="primary" onClick={() => handleRequest(task.id)}>
                                        Request
                                        </Button>
                                    ) : (
                                        <Button variant="contained" disabled>
                                        Requested
                                        </Button>
                                    )}
                                </TableCell>
                                <TableCell align="right">
                                {task.downloadable === 'yes' ? (
                                    <Button variant="contained" color="success">
                                    Download
                                    </Button>
                                ) : (
                                    <Button variant="contained" disabled>
                                    Download
                                    </Button>
                                )}
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </TableContainer>
            </div>
        </div>
    </div>
}
export default Certificates