import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { useEffect, useState } from 'react';
import Login from './login'
import Dashboard from './dashboard'
import Activities from './activities';
import Certificates from './certificates';


function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  const [email, setEmail] = useState("")

  return (
    <div className = "App">
      <BrowserRouter>
      <Routes>
          <Route path="/" element={<Login/>} />
          <Route path="/dashboard" element={<Dashboard/>}/>
          <Route path="/activities" element = {<Activities/>}/>
          <Route path="/certificates" element = {<Certificates/>}/>
      </Routes>
      </BrowserRouter>
    </div>
    );
  }

export default App;
