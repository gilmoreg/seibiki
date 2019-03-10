import * as React from 'react';
import Answer from './components/Answer';
import Form from './components/Form';
import { WordData } from './types';
import './App.css';

export interface AppState {
  results: WordData[];
}

class App extends React.Component<any, AppState> {
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
          <h1 className="App-title">Seibiki (正引き)</h1>
        </header>
        <div className="App-intro">
          <Form update={this.update} />
          <Answer words={this.state.results} />
        </div>
      </div>
    );
  }
}

export default App;
