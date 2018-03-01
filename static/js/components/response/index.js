import React, {Component} from "react";

export default class Response extends Component {
  constructor(props) {
    super(props);

    this.state = {
      isCopied: false
    };

    this._copyToClipboard = this._copyToClipboard.bind(this);
  }

  _copyToClipboard(e) {
    const textField = document.createElement('textarea');
    textField.innerText = window.location.href + this.props.response.encoded;
    document.body.appendChild(textField)
    textField.select();
    document.execCommand('copy');
    textField.remove();

    this.setState({isCopied: true});

    setTimeout(() => {
      this.setState({isCopied: false});
    }, 2000);
  }

  render() {
    const isCopied = this.state.isCopied ? <span className="copy-tooltip">copiado</span> : '';

    return (
      <div className="response">
        <span className="short-url">{window.location.href + this.props.response.encoded}</span>
        <button className="copy-short-url" onClick={this._copyToClipboard}>
          {isCopied}
          <span>Copy short URL</span>
        </button>
      </div>
    )
  }
}