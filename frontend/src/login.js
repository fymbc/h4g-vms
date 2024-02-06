import { useNavigate } from "react-router-dom";
import React, { useState } from "react";

const Login = (props) => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [emailError, setEmailError] = useState("")
    const [passwordError, setPasswordError] = useState("")
    
    const navigate = useNavigate()

    const onButtonClick = () => {
    }



return <div className="App-header">
    <div className="LeftBox">
    <div className={"inputContainer"}>
        <input
            value={email}
            placeholder="Enter your email here"
            onChange={ev => setEmail(ev.target.value)}
            className={"inputBox"} />
        <label className="errorLabel">{emailError}</label>
    </div>
    <div className={"inputContainer"}>
        <input
            value={password}
            placeholder="Enter your password here"
            onChange={ev => setPassword(ev.target.value)}
            className={"inputBox"} />
        <label className="errorLabel">{passwordError}</label>
    </div>
    <br />
    <div className={"inputContainer"}>
      <input
          className={"inputButton"}
          type="button"
          onClick={onButtonClick}
          value={"Log in"} />
    </div>
    <div className={"inputContainer"}>
      <input
          className={"signupButton"}
          type="button"
          onClick={onButtonClick}
          value={"Sign up with Singpass"} />
    </div>
    </div>
    <div className= "LoginBox">
        <h1>Hello!</h1>
        <h2>Welcome to<br />VolunteerBAH</h2>
    </div>
</div>
}

export default Login
