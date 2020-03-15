import React, { Component } from 'react';
import { FormControl, InputGroup} from 'react-bootstrap';

// basic button
class ReButton extends Component {
  render() {
    return (
      <InputGroup size="sm" className="re-input-group">
        <InputGroup.Prepend className="re-input-prepend">
          <InputGroup.Text className="re-label" id="basic-addon1">{this.props.labelText}</InputGroup.Text>
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
