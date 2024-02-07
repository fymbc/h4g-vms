import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { useEffect, useState } from 'react';
import Login from './login'
import Dashboard from './dashboard'


function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  const [email, setEmail] = useState("")

  return (
    <div className = "App">
      <BrowserRouter>
      <Routes>
          <Route path="/" element={<Login setLoggedIn={setLoggedIn} setEmail={setEmail} />} />
          <Route path="/dashboard" element={<Dashboard setLoggedIn={setLoggedIn} setEmail={setEmail} />} />
      </Routes>
      </BrowserRouter>
    </div>
    );
  }

export default App;
