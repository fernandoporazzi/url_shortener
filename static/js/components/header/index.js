import React, {Component} from "react";

export default class Header extends Component {
  render() {
    return (
      <header>
        <div className="header-container">
          <h1>
            <a href="./">Url Shortener</a>
          </h1>
        </div>
      </header>
    )
  }
}