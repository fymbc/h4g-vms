import { useNavigate } from "react-router-dom";
import React, { useState } from "react";

const Dashboard = (props) => {

    const username = useState("Admin"); //Change to username


    return <div>
        <div className = 'Banner'>
            <h1>Welcome, {username}!</h1>
        </div>
        <div>
            <table className = "Sidebar">
                <tbody>
                    <tr>
                        <td>Dashboard</td>
                    </tr>
                    <tr>
                        <td>Opportunities</td>
                    </tr>
                    <tr>
                        <td>Certificates</td>
                    </tr>
                    <tr>
                        <td>Settings</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
}

export default Dashboard