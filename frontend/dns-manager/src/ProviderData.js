import React from "react";
import ReactDOM from "react-dom";
import ReButton from './riff/ReButton';
import {
  Button,
  Dropdown,
  InputGroup,
  FormControl,
  DropdownButton
} from 'react-bootstrap';

class GoDaddyDetails extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      count: 0,
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
    
    this.setState({ count: this.state.count + 1 });
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

ReactDOM.render(
  <GoDaddyDetails/>,
  document.getElementById('provider_details')
);

class Provider extends React.Component {
  render() {
    return (
      <div>
        <InputGroup className="mb-3">
          <DropdownButton
            as={InputGroup.Append}
            variant="outline-secondary"
            title="Provider"
            id="input-group-dropdown-2"
          >
            <Dropdown.Item href="#">Go Daddy</Dropdown.Item>
          </DropdownButton>
          <FormControl
            placeholder="Provider Name"
            aria-label="Provider Name"
            aria-describedby="basic-addon2"
          />

        </InputGroup>
      </div>
    )
  }
}

ReactDOM.render(
  <Provider/>,
  document.getElementById('provider_select')
);


/*
  submit data
 */
class SendData extends React.Component {
  render() {
    return (
      <Button
        variant="primary"
      >Submit
      </Button>
    )

  }
}

ReactDOM.render(
  <SendData/>,
  document.getElementById('submit')
);
