import React from "react";
import {checkValidInput} from "../constants/validation";

class Login extends React.Component {

    constructor(props) {
        super(props);
        this.state={
            name: "",
            pwd: ""
        }
        this.handleName = this.handleName.bind(this)
        this.handlePwd = this.handlePwd.bind(this)
    }

    handleName(event) {
        this.setState({name: event.target.value})
    }

    handlePwd(event) {
        this.setState({pwd: event.target.value})
    }

    login() {
        if (checkValidInput(this.state.name) && checkValidInput(this.state.pwd)) {
            this.props.callbackLogin({"name": this.state.name, "pwd": this.state.pwd})
        }
    }

    register() {
        if (checkValidInput(this.state.name) && checkValidInput(this.state.pwd)) {
            this.props.callbackRegister({"name": this.state.name, "pwd": this.state.pwd})
        }
    }

    render() {
        return (
            <div className={"div-login"}>
                <h1>Login</h1>
                <br/>
                <p>Please note that your input can only contain numbers, dash, underscore
                    and letters a-z (uppercase and lowercase).</p>
                <br/>
                <p>Name</p><input onChange={this.handleName}/>
                <br/>
                <p>Password</p><input onChange={this.handlePwd}/>
                <br/>
                <button className={"basic-btn button-confirm"} onClick={() => this.login()}>Confirm</button>
                <button className={"basic-btn button-create"} onClick={() => this.register()}>Register</button>
            </div>
        )
    }

}

export default Login