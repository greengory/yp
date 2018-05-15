import React from 'react'

const SearchArea = () => {
    return (
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
                
                <button className="bg-teal text-white px-2">Search</button>
              </div>
              <div className="links mt-8 text-white text-center">
                  <a href="/restaurants">Restaurants</a>
                  <a href="/plumbers">Plumbers</a>
                  <a href="/locksmiths">Locksmiths</a>
                  <a href="/dentists">Dentists</a>
                  <a href="/auto-insurance">Auto Insurance</a>
              </div>
            </div>
        </div>
    );
};

export default SearchArea;