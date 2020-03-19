import React from "react";
import ReButton from "../riff/ReButton";

class GoDaddyDetails extends React.Component {
  constructor(props) {
    super(props);
    console.log(props);
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

  handleChange = (e, name) => {
    // useful: the [] notation allows variables as property names
    this.setState({ [ name ]: e.target.value });
    this.props.data[name] = e.target.value;
    this.setState({ count: this.state.count + 0 });
  };

  componentDidUpdate(prevProps, prevState, snapshot) {
    //console.log("states ", this.state);
  };

  render() {
    const inputs = [
      { label: "API Key", default: "Remote key to access REST API", symbol: "key"},
      { label: "API Secret", default: "Remote secret to access REST API", symbol: "secret"},
      { label: "DNS Domain", default: "Domain name to update", symbol: "domain"},
      { label: "DNS Host", default: "Host name to update", symbol: "host"},

    ];
    return (
      <div>
        { inputs.map((btn, k) => {
          return <ReButton key={k} labelText={btn.label} defaultText={btn.default} onChange={(e) => this.handleChange(e, btn.symbol)} />
        })}
      </div>
    )
  }
}
export default GoDaddyDetails;
