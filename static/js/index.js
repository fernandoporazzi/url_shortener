import React, {Component, Fragment} from "react";
import ReactDOM from "react-dom";

import Header from './components/header';
import Form from './components/form';
import Response from './components/response';

import '../scss/main.scss';

class App extends Component {
  constructor() {
    super();

    this.state = {
      response: null
    };

    this._encode = this._encode.bind(this); 
  }

  _encode(url) {
    const data = JSON.stringify({url});
    fetch('/short/', 
      {
        method: 'POST',
        body: data,
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then((res) => res.json())
      .then((data) => {
        this.setState({response: data})
      });
  }

  render() {
    return (
      <Fragment>
        <Header />
        <Form encode={this._encode} />

        {this.state.response != null ? <Response response={this.state.response} /> : ''}
      </Fragment>
    );
  }
}

const mountNode = document.getElementById('app');
ReactDOM.render(<App />, mountNode);