import { useNavigate } from "react-router-dom";
import React, { useState } from "react";

const Login = (props) => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [emailError, setEmailError] = useState("")
    const [passwordError, setPasswordError] = useState("")
    
    const navigate = useNavigate()

    const onButtonClick = () => {
        setEmailError("")
        setPasswordError("")

        if ("" === email) {
            setEmailError("Please enter your email")
            return
        }

        if (!/^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/.test(email)) {
            setEmailError("Please enter a valid email")
            return
        }

        if ("" === password) {
            setPasswordError("Please enter a password")
            return
        }

        if (password.length < 7) {
            setPasswordError("The password must be 8 characters or longer")
            return
        }

        //TODO: Authentication logic and link to database
    }



return <div className="App-header">
    <div className="LeftBox">
        <p className="LoginTitle">Login</p>
        <p className="EmailTitle">Email</p>
    <div className={"inputContainer"}>
        <input
            value={email}
            placeholder="Enter your email here"
            onChange={ev => setEmail(ev.target.value)}
            className={"inputBox"} />
        <label className="errorLabel">{emailError}</label>
    </div>
    <p className="EmailTitle">Password</p>
    <div className={"inputContainer"}>
        <input
            value={password}
            placeholder="Enter your password here"
            onChange={ev => setPassword(ev.target.value)}
            className={"inputBox"} />
        <label className="errorLabel">{passwordError}</label>
    </div>
    <div className={"inputContainer"}>
      <input
          className={"LoginButton"}
          type="button"
          onClick={onButtonClick}
          value={"Log in"} />
    </div>
    <div className={"inputContainer"}>
      <input
          className={"SignupButton"}
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
