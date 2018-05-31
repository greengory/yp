import * as types from '../actions/actionTypes';


//Initialize starting state
let initialState = {
    isLoggedIn: false,
    homepageCategories: []
};


const reducer = (state = [], action ) => {
    switch(action.type) {
        case types.FETCH_HP_CATEGORIES_SUCCESS: 
            return [...state, ...action.items];
        case types.HP_CAT_RETRIEVED:
            return ''
        case types.HP_CAT_IS_LOADING:
            return action.isLoading;
        case types.HP_LOADING_HAS_ERRORED:
            return action.hasErrored;
        default:
            return state;
    }
}



export default reducer;