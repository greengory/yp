import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import './index.css';

import {Provider} from 'react-redux';
import store from './redux/store/configureStore';
import App from './App';
import Listings from './components/Listings';
import BusinessDetails from './components/BusinessDetails';
import NewBusiness from './components/NewBusiness';
import WriteReview from './components/WriteReview';
//import registerServiceWorker from './registerServiceWorker';


ReactDOM.render(
    <Provider store={store}>
        <Router>
            <Switch>
                <Route path="/" component={App} />
                <Route path="/search/:city/:type" component={Listings} />
                <Route path="listings/new" component={NewBusiness} />
                <Route path="/business/:city/:business_name" component={BusinessDetails} />
                <Route path="/writereview" component={WriteReview} />
            </Switch>
        </Router>
    </Provider>,
    document.getElementById('root')
);