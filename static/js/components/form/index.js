import React, {Component} from "react";

export default class Form extends Component {
  constructor(props) {
    super(props);

    this.state = {
      url: ''
    };

    this._onChangeHandler = this._onChangeHandler.bind(this);
    this._onClickHandler = this._onClickHandler.bind(this);
  }

  _onChangeHandler(e) {
    const url =  e.target.value.trim();
    this.setState({url});
  }

  _onClickHandler(e) {
    if (this.state.url === '') return;

    this.props.encode(this.state.url);
    this.setState({url: ''});
  }

  render() {
    return (
      <main className="main">
        <div className="form-container">
          <h2>Simplify your links</h2>
          <div className="input-container">
            <input type="text" placeholder="Your original URL here" value={this.state.url} onChange={this._onChangeHandler} />
            <button onClick={this._onClickHandler}>Shorten URL</button>
          </div>
          <div className="tagline">Fill the input above to get a shorter url</div>
        </div>
      </main>
    )
  }
}