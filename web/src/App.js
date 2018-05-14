import React, { Component } from 'react';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="header bg-black text-white relative">
        <div className="nav-bar container mx-auto flex py-4">
            <nav className="w-2/5">
            </nav>           
            <div className="w-1/5 text-center"></div>
        
            <div className="w-2/5 flex justify-end">
              <a className="px-4">Login</a>
              <a className="bg-white text-black px-2 py-1 rounded-sm">Sign Up</a>
            </div>
        </div>

        <div className="search-box absolute">
            <div className="container mx-auto py-8 px-8 text-white">
              <div className="tagline_for_search">Taglines goes here</div>
              <div className="search_form flex">
                <div className="w-1/2 relative">
                  <input type="text" className="w-1/2 h-8 px-4 py-2 text-xs" placeholder="Search Business..."/>
                  <span className="flex item-center absolute pin-r pin-y mb-5"><i className="fa text-grey"></i></span>
                </div>
                
                <div className="w-1/2 relative">
                  <input type="text" className="inherit important h-8 px-4 py-2 text-xs" placeholder="Location for business" />
                  <span className="flex item-center absolute pin-r pin-y mb-5"><i className="fa text-grey"></i></span>
                </div>
                
                <button className="bg-black text-white px-2">Search</button>
              </div>
              The links will be aligned here
            </div>
        </div>
      </div>
    );
  }
}

export default App;
