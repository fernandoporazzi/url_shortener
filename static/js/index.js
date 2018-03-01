import React, {Component, Fragment} from "react";
import ReactDOM from "react-dom";

import Header from './components/header';

import '../scss/main.scss';

class App extends Component {
  render() {
    return (
      <Fragment>
        <Header />
        <div>hello world!</div>
      </Fragment>
    );
  }
}

const mountNode = document.getElementById("app");
ReactDOM.render(<App />, mountNode);