import React from 'react';

const Navigation = () => {
    return (
        <div className="nav-bar container mx-auto flex py-4">
            <nav className="w-2/5">
            </nav>           
            <div className="w-1/5 text-center"></div>
        
            <div className="w-2/5 flex justify-end">
              <a className="px-4" href="/writereview">Write a Review</a>
              <a className="px-4" href="/listing/new">Advertise with Us</a>
              <a className="px-4">Login</a>
              <a className="bg-white text-black px-2 py-1 rounded-sm">Sign Up</a>
            </div>
        </div>
    );
};

export default Navigation;