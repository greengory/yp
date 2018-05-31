import React, { Component } from 'react';
import './App.css';
import Navigation from './components/Navigation';
import SearchArea from './components/SearchArea';
import { NearbyBusinesses } from './components/NearbyBusinesses';
import { BrowseByCategories } from './components/BrowseByCategories';

class App extends Component {
  render() {
    return (
      <div className="mainApp">
        <div className="header bg-black text-white relative">
          <Navigation />
          <SearchArea />
        </div>

        <div className="advertise-with-us bg-white text-black">
          <div className="container mx-auto py-8">
            <h2 className="text-center">Advertise with Us</h2>
            <div className="adv-wrapper w-3/4 container mx-auto flex py-4">
              <div className="w-1/2">
               <p>Lorem ipsum dolor sit amet, et pro iuvaret perfecto prodesset.</p>
              </div>
              <div className="w-1/2 flex flex-wrap">
               <p className="adv-tagline">Lorem ipsum dolor sit amet, et pro iuvaret 
               perfecto prodesset, ad detracto nominavi indoctum est. 
               Qui ad option aeterno sapientem, in doming impedit dolores his, 
               an eum dicam graeci signiferumque.</p>
               <button className="bg-yellow-dark text-black mt-8 py-3 px-8 p-2 rounded-sm justify-center items-center">Get Started</button>
              </div>
            </div>
          </div>
        </div>
        <NearbyBusinesses />
      </div>
    );
  }
}

export default App;
