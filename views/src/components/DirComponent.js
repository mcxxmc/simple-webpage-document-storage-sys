import React from "react";

const tab = "  "

class DirComponent extends React.Component {

    render() {
        return (
            <div className={"div_dir"}>
                <pre>
                    <p>{tab.repeat(this.props.level) + this.props.name}</p>
                </pre>
            </div>
        )
    }

}

export default DirComponent