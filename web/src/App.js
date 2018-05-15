import React, { Component } from 'react';
import './App.css';
import Navigation from './components/Navigation';
import SearchArea from './components/SearchArea';

class App extends Component {
  render() {
    return (
      <div className="header bg-black text-white relative">
        <Navigation />
        <SearchArea />
      </div>
    );
  }
}

export default App;
