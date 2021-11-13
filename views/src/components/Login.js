import React from "react";

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
        this.props.callbackLogin({"name": this.state.name, "pwd": this.state.pwd})
    }

    register() {
        this.props.callbackRegister({"name": this.state.name, "pwd": this.state.pwd})
    }

    render() {
        return (
            <div className={"div-login"}>
                <h1>Login</h1>
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