import { FETCH_HP_CATEGORIES_SUCCESS } from '../actions/actionTypes';


//Initialize starting state
let initialState = {
    isLoggedIn: false,
    homepageCategories: []
};

const reducer = (state = initialState, action ) => {
    switch(action.type) {
        case FETCH_HP_CATEGORIES_SUCCESS: 
            return [...state, ...action.items];
        default:
            return state;
    }
}



export default reducer;