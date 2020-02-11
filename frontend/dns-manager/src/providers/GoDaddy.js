import React from "react";
import ReButton from "../riff/ReButton";

class GoDaddyDetails extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      count: -1,
      key: '',
      secret: '',
      domain: '',
      host: '*'
    }
    // This binding is necessary to make `this` work in the callback
//    this.handleClick = this.handleClick.bind(this);
//    this.handleChange = this.handleChange.bind(this);
  }

  handleChange = (e) => {
    console.log("change ", e, " times ", this.state.count);

    this.setState({ count: this.state.count + 0 });
  }

  handleClick = () => {
  }

  componentDidUpdate(prevProps, prevState, snapshot) {
    console.log("up");
  }

  render() {
    return (
      <div>
        <ReButton labelText="API Key" defaultText="Remote key to access REST API"
                  onChange={(e) => this.handleChange(e)}/>
        <ReButton labelText="API Secret" defaultText="Remote secret to access REST API"
                  onChange={this.handleChange}/>
        <ReButton labelText="DNS Domain" defaultText="Domain name to update"
                  onChange={this.handleChange}/>
        <ReButton labelText="DNS Host" defaultText="Host name to update"
                  onChange={this.handleChange}/>
      </div>
    )
  }
}
export default GoDaddyDetails;
