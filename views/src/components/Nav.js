import React from "react";
import DefaultView from "./DefaultView";
import ReactDOM from "react-dom";

class Nav extends React.Component {

    render() {
        return (
            <div>
                <h1>Nav</h1>
                <button onClick = {initializeDefaultView}>DefaultView</button>
            </div>
        )
    }
}

function initializeDefaultView() {
    ReactDOM.render(
        <DefaultView/>,
        document.getElementById("root")
    )
}

export default Nav;