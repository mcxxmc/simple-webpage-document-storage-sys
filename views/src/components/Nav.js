import React from "react";
import DefaultView from "./DefaultView";
import ReactDOM from "react-dom";

class Nav extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div>
                <button onClick = {initializeDefaultView}>DefaultView</button>
            </div>
        )
    }
}

function initializeDefaultView() {
    ReactDOM.render(
        <DefaultView/>,
        document.getElementById("div_content")
    )
}

export default Nav;