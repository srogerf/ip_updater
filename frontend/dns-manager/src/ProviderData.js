import React from "react";
import ReactDOM from "react-dom";

class Key extends React.Component {
    render() {
        return (
            <div>
                <div className="label">API Key</div>
                <div>
                    <input type="text" />
                </div>

            </div>
        )
    }
}
ReactDOM.render(
    <Key />,
    document.getElementById('api_key')
);

class Provider extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: ''};

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.setState({value: event.target.value});
    }

    handleSubmit(event) {
        alert('A name was submitted: ' + this.state.value);
        event.preventDefault();
    }
    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Name:
                    <input type="text" value={this.state.value} onChange={this.handleChange} />
                </label>
                <input type="submit" value="Submit" />
            </form>
        )
    }
}
ReactDOM.render(
    <Provider />,
    document.getElementById('provider')
);
