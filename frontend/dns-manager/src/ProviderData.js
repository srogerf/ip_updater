import React from "react";
import ReactDOM from "react-dom";
import GoDaddyDetails from "./providers/GoDaddy";
import ProviderSelector from "./providers/ProviderSelector";
import {
  Button
} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

global.data = { };

// render the details input panel
ReactDOM.render(
  <GoDaddyDetails data={global.data}/>,
  document.getElementById('provider_details')
);

ReactDOM.render(
  <ProviderSelector/>,
  document.getElementById('provider_select')
);


/*
  submit data
 */
class SendData extends React.Component {
  submit(e) {
    console.log("submitted");
    console.log(global.data);
  }
  render() {
    return (
      <Button
        variant="primary"
        onClick={ (e) => this.submit(e) }
      >Submit
      </Button>
    )

  }
}

ReactDOM.render(
  <SendData/>,
  document.getElementById('submit')
);
