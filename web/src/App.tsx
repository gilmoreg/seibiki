import * as React from 'react';
import Answer from './components/Answer';
import Form from './components/Form';
import { WordData } from './types';
import './App.css';

import logo from './logo.svg';

export interface AppState {
  results: WordData[];
}

class App extends React.Component {
  constructor(props: any) {
    super(props);
    this.state = { results: [] };
    this.update = this.update.bind(this);
  }

  update(results: WordData[]) {
    this.setState({ results });
  }

  public render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <div className="App-intro">
          <Form update={this.update} />
          <Answer words={undefined} />
        </div>
      </div>
    );
  }
}

export default App;
