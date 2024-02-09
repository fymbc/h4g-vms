import { useNavigate } from "react-router-dom";
import React, { useState } from "react";

const Dashboard = (props) => {

    const username = useState("Admin"); //Change to username
    return <div>
        <div className = 'Banner'>
            <h1>Welcome, {username}!</h1>
        </div>
        <div className = "MainBody">
            <div>
                <table>
                    <tbody className = "Sidebar">
                        <tr>
                            <td>
                                <img src = "logo192.png"></img>
                                <br />Dashboard
                            </td>
                        </tr>
                        <tr>
                            <td><img src = "logo192.png"></img>
                                <br />Opportunities
                            </td>   
                        </tr>
                        <tr>
                            <td><img src = "logo192.png"></img>
                                <br />Certificates
                            </td>
                        </tr>
                        <tr>
                            <td><img src = "logo192.png"></img>
                                <br />Settings
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div className = "Timeline">
                <h1>My <br />Timeline</h1>
            </div>
        </div>
    </div>
}

export default Dashboard