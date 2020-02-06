import React, { Component } from 'react';
import { FormControl, InputGroup} from 'react-bootstrap';

class ReButton extends Component {

  render() {
    return (
      <InputGroup size="sm" className="mb-3">
        <InputGroup.Prepend>
          <InputGroup.Text id="basic-addon1">{this.props.labelText}</InputGroup.Text>
        </InputGroup.Prepend>
        <FormControl
          onChange={(e) => this.props.onChange(e)}
          placeholder={this.props.defaultText}
          aria-label="key"
          aria-describedby="basic-addon1"
        />
      </InputGroup>
    )
  }

}

export default ReButton;
